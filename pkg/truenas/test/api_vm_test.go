/*
TrueNAS RESTful API

Testing VmAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package truenas

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	openapiclient "github.com/terrycain/truenas-go-sdk/pkg/truenas"
)

func Test_truenas_VmAPIService(t *testing.T) {
	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test VmAPIService GetVM", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var id int32

		resp, httpRes, err := apiClient.VmAPI.GetVM(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test VmAPIService ListVMS", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.VmAPI.ListVMS(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})
}
