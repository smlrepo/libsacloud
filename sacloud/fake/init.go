package fake

import (
	"fmt"
	"sync"
	"time"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

var initOnce sync.Once

func ds() Store {
	InitDataStore()
	return DataStore
}

// DataStore fakeドライバが利用するデータストア
var DataStore Store

// InitDataStore データストアの初期化
func InitDataStore() {
	initOnce.Do(func() {
		if DataStore == nil {
			DataStore = NewInMemoryStore()
		}
		if err := DataStore.Init(); err != nil {
			panic(err)
		}

		// pool(id,ip,subnet,etc...)
		p := initValuePool(DataStore)

		if DataStore.NeedInitData() {
			initData(DataStore, p)
		}
	})
}

func initData(s Store, p *valuePool) error {
	initArchives(s, p)
	initBills(s, p)
	initCoupons(s, p)
	initCDROMs(s, p)
	initNotes(s, p)
	initSwitch(s, p)
	initZones(s, p)
	initRegions(s, p)
	initPrivateHostPlan(s, p)
	initDiskPlan(s, p)
	initLicensePlan(s, p)
	initInternetPlan(s, p)
	initServerPlan(s, p)
	initServiceClass(s, p)
	return nil
}

func initArchives(s Store, p *valuePool) {
	archives := []*sacloud.Archive{
		{
			ID:                   p.generateID(),
			Name:                 "CentOS 7.6 (1810) 64bit",
			Tags:                 []string{"@size-extendable", "arch-64bit", "current-stable", "distro-centos", "distro-ver-7.6", "os-linux"},
			DisplayOrder:         1,
			Scope:                types.Scopes.Shared,
			Availability:         types.Availabilities.Available,
			SizeMB:               20480,
			DiskPlanID:           types.DiskPlans.HDD,
			DiskPlanName:         "標準プラン",
			DiskPlanStorageClass: "iscsi9999",
		},
		{
			ID:                   p.generateID(),
			Name:                 "Ubuntu Server 18.04.2 LTS 64bit",
			Tags:                 []string{"@size-extendable", "arch-64bit", "current-stable", "distro-ubuntu", "distro-ver-18.04.2", "os-linux"},
			DisplayOrder:         2,
			Scope:                types.Scopes.Shared,
			Availability:         types.Availabilities.Available,
			SizeMB:               20480,
			DiskPlanID:           types.DiskPlans.HDD,
			DiskPlanName:         "標準プラン",
			DiskPlanStorageClass: "iscsi9999",
		},
	}
	for _, zone := range zones {
		for _, archive := range archives {
			s.Put(ResourceArchive, zone, archive.ID, archive)
		}
	}
}

func initBills(s Store, p *valuePool) {
	bills := []*sacloud.Bill{
		{
			ID:             p.generateID(),
			Amount:         1080,
			Date:           time.Now(),
			MemberID:       "dummy00000",
			Paid:           false,
			PayLimit:       time.Now().AddDate(0, 1, 0),
			PaymentClassID: 999,
		},
	}
	for _, bill := range bills {
		s.Put(ResourceBill, sacloud.APIDefaultZone, bill.ID, bill)
		initBillDetails(s, p, bill.ID)
	}
}

func initBillDetails(s Store, p *valuePool, billID types.ID) {
	details := []*sacloud.BillDetail{
		{
			ID:             p.generateID(),
			Amount:         108,
			Description:    "description",
			ServiceClassID: 999,
			Usage:          100,
			Zone:           "tk1a",
			ContractEndAt:  time.Now(),
		},
	}
	s.Put(ResourceBill+"Details", sacloud.APIDefaultZone, billID, &details)
}

func initCDROMs(s Store, p *valuePool) {
	cdroms := []*sacloud.CDROM{
		{
			ID:           p.generateID(),
			Name:         "dummy",
			Description:  "dummy",
			DisplayOrder: 999,
			Tags:         types.Tags{"current-stable", "os-linux"},
			Availability: types.Availabilities.Available,
			Scope:        types.Scopes.Shared,
			CreatedAt:    time.Now(),
			ModifiedAt:   time.Now(),
		},
	}
	for _, zone := range zones {
		for _, c := range cdroms {
			s.Put(ResourceCDROM, zone, c.ID, c)
		}
	}
}

func initCoupons(s Store, p *valuePool) {
	coupons := []*sacloud.Coupon{
		{
			ID:             p.generateID(),
			MemberID:       "dummy00000",
			ContractID:     p.generateID(),
			ServiceClassID: 999,
			Discount:       20000,
			AppliedAt:      time.Now().AddDate(0, -1, 0),
			UntilAt:        time.Now().AddDate(0, 1, 0),
		},
	}
	for _, c := range coupons {
		s.Put(ResourceCoupon, sacloud.APIDefaultZone, c.ID, c)
	}
}

func initNotes(s Store, p *valuePool) {
	notes := []*sacloud.Note{
		{
			ID:      1,
			Name:    "sys-nfs",
			Class:   "json",
			Content: `{"plans":{"HDD":[{"size": 100,"availability":"available","planId":1}]}}`,
		},
	}
	for _, note := range notes {
		s.Put(ResourceNote, sacloud.APIDefaultZone, note.ID, note)
	}
}

func initSwitch(s Store, p *valuePool) {
	sharedSegmentSwitch = &sacloud.Switch{
		ID:             p.generateID(),
		Name:           "スイッチ",
		Scope:          types.Scopes.Shared,
		Description:    "共有セグメント用スイッチ",
		NetworkMaskLen: p.SharedNetMaskLen,
		DefaultRoute:   p.SharedDefaultGateway.String(),
	}
	for _, zone := range zones {
		s.Put(ResourceSwitch, zone, sharedSegmentSwitch.ID, sharedSegmentSwitch)
	}
}

func initZones(s Store, p *valuePool) {
	// zones
	zones := []*sacloud.Zone{
		{
			ID:           21001,
			Name:         "tk1a",
			Description:  "東京第1ゾーン",
			DisplayOrder: 1,
		},
		{
			ID:           31001,
			Name:         "is1a",
			Description:  "石狩第1ゾーン",
			DisplayOrder: 2,
		},
		{
			ID:           31002,
			Name:         "is1b",
			Description:  "石狩第2ゾーン",
			DisplayOrder: 3,
		},
		{
			ID:           29001,
			Name:         "tk1v",
			Description:  "Sandbox",
			DisplayOrder: 4,
			IsDummy:      true,
		},
	}

	for _, zone := range zones {
		s.Put(ResourceZone, sacloud.APIDefaultZone, zone.ID, zone)
	}
}

func initRegions(s Store, p *valuePool) {
	regions := []*sacloud.Region{
		{
			ID:          210,
			Name:        "東京",
			Description: "東京",
			NameServers: []string{
				"210.188.224.10",
				"210.188.224.11",
			},
		},
		{
			ID:          290,
			Name:        "Sandbox",
			Description: "Sandbox",
			NameServers: []string{
				"133.242.0.3",
				"133.242.0.4",
			},
		},
		{
			ID:          310,
			Name:        "石狩",
			Description: "石狩",
			NameServers: []string{
				"133.242.0.3",
				"133.242.0.4",
			},
		},
	}

	for _, region := range regions {
		s.Put(ResourceRegion, sacloud.APIDefaultZone, region.ID, region)
	}
}

func initPrivateHostPlan(s Store, p *valuePool) {
	plans := []*sacloud.PrivateHostPlan{
		{
			ID:           112900526366,
			Name:         "200Core 224GB 標準",
			Class:        "dynamic",
			CPU:          200,
			MemoryMB:     229376,
			Availability: types.Availabilities.Available,
		},
		{
			ID:           112900526366,
			Name:         "200Core 224GB 標準",
			Class:        "dynamic",
			CPU:          200,
			MemoryMB:     229376,
			Availability: types.Availabilities.Available,
		},
	}
	for _, zone := range zones {
		for _, plan := range plans {
			s.Put(ResourcePrivateHostPlan, zone, plan.ID, plan)
		}
	}
}

func initServerPlan(s Store, p *valuePool) {
	plans := []*sacloud.ServerPlan{
		{
			ID:           p.generateID(),
			Name:         "プラン/1Core-1GB",
			CPU:          1,
			MemoryMB:     1024,
			Commitment:   types.Commitments.Standard,
			Generation:   100,
			Availability: types.Availabilities.Available,
		},
		{
			ID:           p.generateID(),
			Name:         "プラン/2Core-4GB",
			CPU:          2,
			MemoryMB:     4 * 1024,
			Commitment:   types.Commitments.Standard,
			Generation:   100,
			Availability: types.Availabilities.Available,
		},
		// TODO add more plans
	}

	for _, zone := range zones {
		for _, plan := range plans {
			s.Put(ResourceServerPlan, zone, plan.ID, plan)
		}
	}
}

func initInternetPlan(s Store, p *valuePool) {
	bandWidthList := []int{100, 250, 500, 1000, 1500, 2000, 2500, 3000, 5000}

	var plans []*sacloud.InternetPlan

	for _, bw := range bandWidthList {
		plans = append(plans, &sacloud.InternetPlan{
			ID:            types.ID(bw),
			BandWidthMbps: bw,
			Name:          fmt.Sprintf("%dMbps共有", bw),
			Availability:  types.Availabilities.Available,
		})
	}

	for _, zone := range zones {
		for _, plan := range plans {
			s.Put(ResourceInternetPlan, zone, plan.ID, plan)
		}
	}
}

func initDiskPlan(s Store, p *valuePool) {
	plans := []*sacloud.DiskPlan{
		{
			ID:           2,
			Name:         "HDDプラン",
			Availability: types.Availabilities.Available,
			StorageClass: "iscsi1204",
			Size: []*sacloud.DiskPlanSizeInfo{
				{
					Availability:  types.Availabilities.Available,
					DisplaySize:   20,
					DisplaySuffix: "GB",
					SizeMB:        20 * 1024,
				},
				{
					Availability:  types.Availabilities.Available,
					DisplaySize:   40,
					DisplaySuffix: "GB",
					SizeMB:        40 * 1024,
				},
			},
		},
		{
			ID:           4,
			Name:         "SSDプラン",
			Availability: types.Availabilities.Available,
			StorageClass: "iscsi1204",
			Size: []*sacloud.DiskPlanSizeInfo{
				{
					Availability:  types.Availabilities.Available,
					DisplaySize:   20,
					DisplaySuffix: "GB",
					SizeMB:        20 * 1024,
				},
				{
					Availability:  types.Availabilities.Available,
					DisplaySize:   40,
					DisplaySuffix: "GB",
					SizeMB:        40 * 1024,
				},
			},
		},
		// TODO add more size-info
	}

	for _, zone := range zones {
		for _, plan := range plans {
			s.Put(ResourceDiskPlan, zone, plan.ID, plan)
		}
	}
}

func initLicensePlan(s Store, p *valuePool) {
	plans := []*sacloud.LicenseInfo{
		{
			ID:         types.ID(10001),
			Name:       "Windows RDS SAL",
			TermsOfUse: "1ライセンスにつき、1人のユーザが利用できます。",
		},
	}

	for _, zone := range zones {
		for _, plan := range plans {
			s.Put(ResourceLicenseInfo, zone, plan.ID, plan)
		}
	}
}

func initServiceClass(s Store, p *valuePool) {
	classes := []*sacloud.ServiceClass{
		{
			ID:               types.ID(50050),
			ServiceClassName: "plan/1",
			ServiceClassPath: "cloud/plan/1",
			DisplayName:      "プラン1(ディスクなし)",
			IsPublic:         true,
			Price: &sacloud.Price{
				Base:    0,
				Daily:   108,
				Hourly:  10,
				Monthly: 2139,
			},
		},
		{
			ID:               types.ID(50051),
			ServiceClassName: "plan/2",
			ServiceClassPath: "cloud/plan/2",
			DisplayName:      "プラン2(ディスクなし)",
			IsPublic:         true,
			Price: &sacloud.Price{
				Base:    0,
				Daily:   172,
				Hourly:  17,
				Monthly: 3425,
			},
		},
	}

	for _, zone := range zones {
		for _, class := range classes {
			class.Price.Zone = zone
			s.Put(ResourceServiceClass, zone, class.ID, class)
		}
	}
}
