// Copyright 2016-2019 The Libsacloud Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const testIconName = "libsacloud_test_icon"

const testImageBody = `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`

func TestIconCRUD(t *testing.T) {
	defer initIcon()()

	api := client.Icon

	//CREATE
	newItem := api.New()
	newItem.Name = testIconName
	newItem.Image = testImageBody

	item, err := api.Create(newItem)

	assert.NoError(t, err)
	assert.NotEmpty(t, item)

	id := item.ID

	//READ
	item, err = api.Read(id) //normal
	assert.NoError(t, err)
	assert.NotEmpty(t, item)
	assert.NotEmpty(t, item.URL)

	image, err := api.GetImage(id, "small") //image base64
	assert.NoError(t, err)
	assert.NotEmpty(t, image)

	//Delete
	_, err = api.Delete(id)
	assert.NoError(t, err)
}

func initIcon() func() {
	cleanupIcon()
	return cleanupIcon
}

func cleanupIcon() {
	items, _ := client.Icon.Reset().WithNameLike(testIconName).Find()
	for _, item := range items.Icons {
		client.Icon.Delete(item.ID)
	}
}
