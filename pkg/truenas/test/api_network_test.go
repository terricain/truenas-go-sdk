/*
TrueNAS RESTful API

Testing NetworkAPIService

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

func Test_truenas_NetworkAPIService(t *testing.T) {
	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test NetworkAPIService GetNetworkConfiguration", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.NetworkAPI.GetNetworkConfiguration(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test NetworkAPIService GetNetworkSummary", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.NetworkAPI.GetNetworkSummary(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})
}
