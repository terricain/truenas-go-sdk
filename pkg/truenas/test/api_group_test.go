/*
TrueNAS RESTful API

Testing GroupAPIService

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

func Test_truenas_GroupAPIService(t *testing.T) {
	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test GroupAPIService CreateGroup", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.GroupAPI.CreateGroup(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test GroupAPIService DeleteGroup", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var id int32

		httpRes, err := apiClient.GroupAPI.DeleteGroup(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test GroupAPIService GetGroup", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var id int32

		resp, httpRes, err := apiClient.GroupAPI.GetGroup(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test GroupAPIService ListGroups", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.GroupAPI.ListGroups(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test GroupAPIService UpdateGroup", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var id int32

		resp, httpRes, err := apiClient.GroupAPI.UpdateGroup(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})
}