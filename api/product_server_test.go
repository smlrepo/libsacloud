package api

import (
	"testing"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/stretchr/testify/assert"
)

func TestGetProductServer(t *testing.T) {
	api := client.Product.Server

	oldZone := client.Zone
	defer func() {
		client.Zone = oldZone
	}()

	t.Run("valid server plan", func(t *testing.T) {
		//READ
		res, err := api.Read(100001001) // 1core 1GB
		assert.NoError(t, err)
		assert.NotEmpty(t, res)
		assert.NotEmpty(t, res.ID)
	})

	t.Run("valid server plan with default gen", func(t *testing.T) {
		client.Zone = "is1a"
		plan, err := api.GetBySpec(1, 1, sacloud.PlanDefault)
		assert.NoError(t, err)
		assert.NotNil(t, plan)
		assert.Equal(t, sacloud.PlanG2, plan.Generation)

		client.Zone = "is1b"
		plan, err = api.GetBySpec(1, 1, sacloud.PlanDefault)
		assert.NoError(t, err)
		assert.NotNil(t, plan)
		assert.Equal(t, sacloud.PlanG1, plan.Generation)
	})

	t.Run("invalid server plan", func(t *testing.T) {
		isValid, err := api.IsValidPlan(9999999, 99999999, sacloud.PlanG1)
		assert.Error(t, err)
		assert.False(t, isValid)
	})

	t.Run("server plan on is1a", func(t *testing.T) {
		client.Zone = "is1a"
		isValid, err := api.IsValidPlan(1, 1, sacloud.PlanG1)
		assert.NoError(t, err)
		assert.True(t, isValid)

		isValid, err = api.IsValidPlan(1, 1, sacloud.PlanG2)
		assert.NoError(t, err)
		assert.True(t, isValid)
	})

	t.Run("server plan on is1b", func(t *testing.T) {
		// isib
		client.Zone = "is1b"
		isValid, err := api.IsValidPlan(1, 1, sacloud.PlanG1)
		assert.NoError(t, err)
		assert.True(t, isValid)

		isValid, err = api.IsValidPlan(1, 1, sacloud.PlanG2)
		assert.Error(t, err)
		assert.False(t, isValid)
	})
}
