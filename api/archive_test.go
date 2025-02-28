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
	"time"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/stretchr/testify/assert"
)

const testArchiveName = "libsacloud_test_archive"

//func TestGetCentOSArchiveByName(t *testing.T) {
//	archiveAPI := client.Archive
//
//	res, err := archiveAPI.Reset().WithNameLike("CentOS 7.2 64bit").Find()
//	assert.NoError(t, err)
//	assert.NotEmpty(t, res)
//	assert.Equal(t, len(res.Archives), 1)
//}

func TestGetArchiveWithLimitOffset(t *testing.T) {

	defer initArchive()()

	archiveAPI := client.Archive
	res, err := archiveAPI.Reset().Limit(2).Offset(1).Include("Name").Include("CreatedAt").Find()
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
	assert.Equal(t, len(res.Archives), 2)
	assert.Equal(t, res.From, 1)
	assert.Equal(t, res.Count, 2)
	assert.True(t, res.Total > 2)
}

func TestFindState(t *testing.T) {
	defer initArchive()()

	api := client.Archive

	api.Reset().WithNameLike("hoge").FilterBy("Fuga", "fuga").Limit(10).Offset(1).Include("inc").Exclude("enc")

	state := api.state

	assert.NotEmpty(t, state)
	assert.Equal(t, state.Filter["Name"], "hoge")
	assert.Equal(t, state.Filter["Fuga"], "fuga")
	assert.Equal(t, state.Count, 10)
	assert.Equal(t, state.From, 1)
	assert.Equal(t, state.Include[0], "inc")
	assert.Equal(t, state.Exclude[0], "enc")

	//clear state
	api.Reset()
	state = api.state
	assert.Empty(t, state.Filter)
	assert.Empty(t, state.Count)
	assert.Empty(t, state.From)
	assert.Empty(t, state.Include)
	assert.Empty(t, state.Exclude)

	res, err := api.withNameLike("CentOS").limit(1).Find()

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
	assert.Equal(t, res.Count, 1)
	assert.Contains(t, res.Archives[0].Name, "CentOS")
}

func TestFindStateWithSetter(t *testing.T) {
	defer initArchive()()

	api := client.Archive

	// set parameters by setter method
	api.SetEmpty()
	api.SetNameLike("hoge")
	api.SetFilterBy("Fuga", "fuga")
	api.SetLimit(10)
	api.SetOffset(1)
	api.SetInclude("inc")
	api.SetExclude("enc")

	state := api.state

	assert.NotEmpty(t, state)
	assert.Equal(t, state.Filter["Name"], "hoge")
	assert.Equal(t, state.Filter["Fuga"], "fuga")
	assert.Equal(t, state.Count, 10)
	assert.Equal(t, state.From, 1)
	assert.Equal(t, state.Include[0], "inc")
	assert.Equal(t, state.Exclude[0], "enc")

	//clear state
	api.SetEmpty()
	state = api.state
	assert.Empty(t, state.Filter)
	assert.Empty(t, state.Count)
	assert.Empty(t, state.From)
	assert.Empty(t, state.Include)
	assert.Empty(t, state.Exclude)

	api.SetNameLike("CentOS")
	api.SetLimit(1)

	res, err := api.Find()

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
	assert.Equal(t, res.Count, 1)
	assert.Contains(t, res.Archives[0].Name, "CentOS")
}

func TestArchiveCRUDAndFTP(t *testing.T) {
	defer initArchive()()

	api := client.Archive

	// get icon ID
	icons, err := client.Icon.Reset().WithSharedScope().Include("ID").Find()

	if !assert.NoError(t, err) || !assert.True(t, len(icons.Icons) > 0) {
		return
	}
	icon := icons.Icons[0]

	//CREATE
	newArchive := api.New()
	newArchive.Name = testArchiveName
	newArchive.SetDescription("hoge")
	newArchive.AppendTag("hoge")
	newArchive.SetIcon(&icon)
	newArchive.SetSizeGB(20)

	archive, err := api.Create(newArchive)

	assert.NoError(t, err)
	assert.NotEmpty(t, archive)

	assert.Equal(t, archive.Description, "hoge")
	assert.Len(t, archive.Tags, 1)
	assert.Equal(t, archive.Tags[0], "hoge")
	assert.Equal(t, archive.Icon.ID, icon.ID)

	archiveID := archive.ID

	//READ
	archive, err = api.Read(archiveID)
	assert.NoError(t, err)
	assert.NotEmpty(t, archive)

	// Update
	archive.SetDescription("")
	archive.ClearTags()
	archive.ClearIcon()

	archive, err = api.Update(archive.ID, archive)

	assert.Equal(t, archive.Description, "")
	assert.Len(t, archive.Tags, 0)
	assert.Nil(t, archive.Icon)

	//Open
	ftpServer, err := api.OpenFTP(archive.ID)
	assert.NoError(t, err)
	assert.NotEmpty(t, ftpServer.Password)

	password := ftpServer.Password

	////Close
	//res, err := api.CloseFTP(archiveID)
	//assert.NoError(t, err)
	//assert.True(t, res)

	//Re-Open(password not changed)
	//ftpServer, err = api.OpenFTP(archive.ID, false)
	//assert.NoError(t, err)
	//assert.Equal(t, ftpServer.Password, password)

	//Close
	api.CloseFTP(archiveID)

	//Re-Open(will password change)
	ftpServer, err = api.OpenFTP(archive.ID)
	assert.NoError(t, err)
	assert.NotEqual(t, ftpServer.Password, password)

	//Delete
	_, err = api.Delete(archiveID)
	assert.NoError(t, err)

}

func TestCreateAndWait(t *testing.T) {
	defer initArchive()()

	archiveAPI := client.Archive
	src, err := archiveAPI.FindLatestStableCentOS()

	assert.NoError(t, err)
	id := src.ID
	assert.NotEmpty(t, id)

	//CREATE
	newArchive := archiveAPI.New()
	newArchive.Name = testArchiveName
	newArchive.Description = "hoge"
	newArchive.SetSourceArchive(id)

	archive, err := archiveAPI.Create(newArchive)

	assert.NoError(t, err)
	assert.NotEmpty(t, archive)

	err = archiveAPI.SleepWhileCopying(archive.ID, client.DefaultTimeoutDuration)
	assert.NoError(t, err)

	archiveAPI.Delete(archive.ID)

}

func TestCreateAndAsyncWait(t *testing.T) {
	defer initArchive()()

	archiveAPI := client.Archive
	src, err := archiveAPI.FindLatestStableCentOS()

	assert.NoError(t, err)
	id := src.ID
	assert.NotEmpty(t, id)

	//CREATE
	newArchive := archiveAPI.New()
	newArchive.Name = testArchiveName
	newArchive.Description = "hoge"
	newArchive.SetSourceArchive(id)

	archive, err := archiveAPI.Create(newArchive)

	assert.NoError(t, err)
	assert.NotEmpty(t, archive)
	defer func() {
		archiveAPI.Delete(archive.ID)
	}()

	complete, progress, errChan := archiveAPI.AsyncSleepWhileCopying(archive.ID, client.DefaultTimeoutDuration)

	for {
		select {
		case a := <-progress:

			t.Logf("Copying...\t %d MB / %d MB", a.(*sacloud.Archive).GetMigratedMB(), a.(*sacloud.Archive).GetSizeMB())
		case a := <-complete:
			t.Logf("Done...\t %d MB / %d MB", a.(*sacloud.Archive).GetMigratedMB(), a.(*sacloud.Archive).GetSizeMB())
			//t.Logf("Trace:%#v", a)
			return
		case e := <-errChan:
			assert.Fail(t, e.Error(), nil)
			return
		case <-time.After(20 * time.Minute):
			assert.Fail(t, "Timeout: AsyncSleepWhileCopying: Disk -> %d", archive.ID)
			return
		}
	}
}

func TestArchiveAPI_FindStableOSs(t *testing.T) {
	defer initArchive()()

	api := client.Archive
	type target struct {
		label string
		f     func() (*sacloud.Archive, error)
	}

	targets := []target{
		{label: "CentOS", f: api.FindLatestStableCentOS},
		{label: "CentOS8", f: api.FindLatestStableCentOS8},
		{label: "CentOS7", f: api.FindLatestStableCentOS7},
		{label: "CentOS6", f: api.FindLatestStableCentOS6},
		{label: "Debian", f: api.FindLatestStableDebian},
		{label: "Debian10", f: api.FindLatestStableDebian10},
		{label: "Debian9", f: api.FindLatestStableDebian9},
		{label: "Ubuntu", f: api.FindLatestStableUbuntu},
		{label: "Ubuntu1804", f: api.FindLatestStableUbuntu1804},
		{label: "Ubuntu1604", f: api.FindLatestStableUbuntu1604},
		{label: "CoreOS", f: api.FindLatestStableCoreOS},
		{label: "RancherOS", f: api.FindLatestStableRancherOS},
		{label: "k3OS", f: api.FindLatestStableK3OS},
		{label: "Kusanagi", f: api.FindLatestStableKusanagi},
		{label: "FreeBSD", f: api.FindLatestStableFreeBSD},
		{label: "Netwiser", f: api.FindLatestStableNetwiser},
		{label: "OPNsense", f: api.FindLatestStableOPNsense},
		{label: "Windows2016", f: api.FindLatestStableWindows2016},
		{label: "Windows2016-RDS", f: api.FindLatestStableWindows2016RDS},
		{label: "Windows2016-RDS-Office", f: api.FindLatestStableWindows2016RDSOffice},
		{label: "Windows2016-SQLServer-Web", f: api.FindLatestStableWindows2016SQLServerWeb},
		{label: "Windows2016-SQLServer-Standard", f: api.FindLatestStableWindows2016SQLServerStandard},
		{label: "Windows2016-SQLServer2017-Standard", f: api.FindLatestStableWindows2016SQLServer2017Standard},
		{label: "Windows2016-SQLServer-Standard-All", f: api.FindLatestStableWindows2016SQLServerStandardAll},
		{label: "Windows2016-SQLServer2017-Standard-All", f: api.FindLatestStableWindows2016SQLServer2017StandardAll},
		{label: "Windows2019", f: api.FindLatestStableWindows2019},
	}

	for _, ts := range targets {
		res, err := ts.f()
		assert.NoError(t, err)
		assert.NotNil(t, res)
		t.Logf("Zone:%s / Current Stable %s: %#v", client.Zone, ts.label, res.Resource)

	}

}

func TestArchiveAPI_CanDiskEdit(t *testing.T) {
	defer initArchive()()

	api := client.Archive
	type target struct {
		label  string
		expect bool
		f      func() (*sacloud.Archive, error)
	}

	targets := []target{
		{label: "CentOS", expect: true, f: api.FindLatestStableCentOS},
		{label: "CentOS8", expect: true, f: api.FindLatestStableCentOS8},
		{label: "CentOS7", expect: true, f: api.FindLatestStableCentOS7},
		{label: "CentOS6", expect: true, f: api.FindLatestStableCentOS6},
		{label: "Debian", expect: true, f: api.FindLatestStableDebian},
		{label: "Debian10", expect: true, f: api.FindLatestStableDebian10},
		{label: "Debian9", expect: true, f: api.FindLatestStableDebian9},
		{label: "Ubuntu", expect: true, f: api.FindLatestStableUbuntu},
		{label: "Ubuntu1804", expect: true, f: api.FindLatestStableUbuntu1804},
		{label: "Ubuntu1604", expect: true, f: api.FindLatestStableUbuntu1604},
		{label: "CoreOS", expect: true, f: api.FindLatestStableCoreOS},
		{label: "RancherOS", expect: true, f: api.FindLatestStableRancherOS},
		{label: "k3OS", expect: true, f: api.FindLatestStableK3OS},
		{label: "Kusanagi", expect: true, f: api.FindLatestStableKusanagi},
		{label: "SophosUTM", expect: false, f: api.FindLatestStableSophosUTM}, // Sophos UTMはディスクの修正不可
		{label: "FreeBSD", expect: true, f: api.FindLatestStableFreeBSD},
		{label: "Netwiser", expect: false, f: api.FindLatestStableNetwiser}, // Netwiserはディスクの修正不可
		{label: "OPNsense", expect: false, f: api.FindLatestStableOPNsense}, // OPNsenseはディスクの修正不可
		{label: "Windows2016", expect: false, f: api.FindLatestStableWindows2016},
		{label: "Windows2016-RDS", expect: false, f: api.FindLatestStableWindows2016RDS},
		{label: "Windows2016-RDS-Office", expect: false, f: api.FindLatestStableWindows2016RDSOffice},
		{label: "Windows2016-SQLServer-Web", expect: false, f: api.FindLatestStableWindows2016SQLServerWeb},
		{label: "Windows2016-SQLServer-Standard", expect: false, f: api.FindLatestStableWindows2016SQLServerStandard},
		{label: "Windows2016-SQLServer2017-Standard", expect: false, f: api.FindLatestStableWindows2016SQLServer2017Standard},
		{label: "Windows2016-SQLServer-Standard-All", expect: false, f: api.FindLatestStableWindows2016SQLServerStandardAll},
		{label: "Windows2016-SQLServer2017-Standard-All", expect: false, f: api.FindLatestStableWindows2016SQLServer2017StandardAll},
		{label: "Windows2019", expect: false, f: api.FindLatestStableWindows2019},
	}

	for _, ts := range targets {
		archive, err := ts.f()
		assert.NoError(t, err)
		assert.NotNil(t, archive)

		res, err := client.Archive.CanEditDisk(archive.ID)
		assert.NoError(t, err)
		assert.Equal(t, res, ts.expect)

		t.Logf("Zone:%s / %s: CanDiskEdit: %t", client.Zone, ts.label, res)
	}

}

func initArchive() func() {
	cleanupArchive()
	return cleanupArchive
}

func cleanupArchive() {
	items, _ := client.Archive.Reset().WithNameLike(testArchiveName).Find()
	for _, item := range items.Archives {
		client.Archive.Delete(item.ID)
	}
}
