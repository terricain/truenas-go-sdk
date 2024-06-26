/*
TrueNAS RESTful API

Testing DatasetAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package truenas

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	openapiclient "github.com/terricain/truenas-go-sdk/pkg/truenas"
)

func Test_truenas_DatasetAPIService(t *testing.T) {
	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DatasetAPIService CreateDataset", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.DatasetAPI.CreateDataset(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test DatasetAPIService DeleteDataset", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var id string

		httpRes, err := apiClient.DatasetAPI.DeleteDataset(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test DatasetAPIService GetDataset", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.DatasetAPI.GetDataset(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test DatasetAPIService ListDatasets", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.DatasetAPI.ListDatasets(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test DatasetAPIService UpdateDataset", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.DatasetAPI.UpdateDataset(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})
}
