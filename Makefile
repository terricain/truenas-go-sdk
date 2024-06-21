SHELL:=bash
.SHELLFLAGS := -eu -o pipefail -c
.DELETE_ON_ERROR:
MAKEFLAGS += --warn-undefined-variables
MAKEFLAGS += --no-builtin-rules

GENERATOR_VERSION=7.1.0
GENERATOR_DOWNLOAD_URL=https://repo1.maven.org/maven2/org/openapitools/openapi-generator-cli/${GENERATOR_VERSION}/openapi-generator-cli-${GENERATOR_VERSION}.jar
GENERATOR_SHA1=7d9b3a162b5b5fddf28a995d330131cfff886333

# So regular `make` generates
gen: pkg/truenas
.PHONY: gen

.cache/openapi-generator-cli-${GENERATOR_VERSION}.jar:
	mkdir -p $(@D)
	curl -o $@ ${GENERATOR_DOWNLOAD_URL}
	echo -n "${GENERATOR_SHA1}  $@" | sha1sum --check


download_deps: .cache/openapi-generator-cli-${GENERATOR_VERSION}.jar
.PHONY: download_deps

clean:
	rm -rf .cache
.PHONY: clean

pkg/truenas: cfg/config.yaml cfg/openapi.yaml
	rm -rf $@ ./docs
	mkdir -p $@
	java -jar .cache/openapi-generator-cli-${GENERATOR_VERSION}.jar generate --input-spec cfg/openapi.yaml --config cfg/config.yaml --http-user-agent 'TruenasScaleSDK' --output $@ --generator-name go --git-user-id terricain --git-repo-id truenas-go-sdk
	# Clean up stuff as am too lazy to find the right generate options
	rm -f pkg/truenas/{git_push.sh,go.mod,go.sum,.travis.yml,.gitignore}
	sed -i'' 's@"github.com/terricain/truenas-go-sdk"@"github.com/terricain/truenas-go-sdk/pkg/truenas"@g' pkg/truenas/test/*.go pkg/truenas/README.md pkg/truenas/docs/*.md
	mv pkg/truenas/README.md .
	mv pkg/truenas/docs .
	# Apply some fixes
	patch -p1 < patches/client-jsonregex.patch
	gofumpt -l -w .
	go mod tidy
