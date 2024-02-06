package unit_test

import (
	"testing"

	"modular-monolithic/model"
	"modular-monolithic/module/v1/menu/helper"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestPrepareToMenusResponse(t *testing.T) {
	menus := []model.Menu{
		{
			ID:   uuid.New(),
			Name: &[]string{"Test Menu 1"}[0],
		},
		{
			ID:   uuid.New(),
			Name: &[]string{"Test Menu 2"}[0],
		},
	}

	response := helper.PrepareToMenusResponse(menus)

	assert.NotNil(t, response)
	assert.Len(t, response, 2)

	for i, menu := range response {
		assert.Equal(t, menus[i].ID, menu.ID)
		assert.Equal(t, *menus[i].Name, menu.Name)
	}
}

func TestPrepareToDetailMenuResponse(t *testing.T) {
	menu := &model.Menu{
		ID:   uuid.New(),
		Name: &[]string{"Test Menu"}[0],
	}

	response := helper.PrepareToDetailMenuResponse(menu)

	assert.NotNil(t, response)
	assert.Equal(t, menu.ID, response.ID)
	assert.Equal(t, *menu.Name, response.Name)
}
