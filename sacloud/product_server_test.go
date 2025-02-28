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

package sacloud

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testServerPlanJSON = `
{
	"ID": 1001,
	"Name": "\u30d7\u30e9\u30f3\/1Core-1GB",
	"CPU": 1,
	"MemoryMB": 1024,
	"ServiceClass": "cloud\/plan\/1core-1gb"
}
`

func TestMarshalProductServerJSON(t *testing.T) {
	var productServer ProductServer
	err := json.Unmarshal([]byte(testServerPlanJSON), &productServer)

	assert.NoError(t, err)
	assert.NotEmpty(t, productServer)

	assert.NotEmpty(t, productServer.ID)
}
