package accessor

import "github.com/sacloud/libsacloud/sacloud/types"

// InstanceStatus InstanceStatusを持つリソース向けのインターフェース
type InstanceStatus interface {
	GetInstanceStatus() types.EServerInstanceStatus
	SetInstanceStatus(types.EServerInstanceStatus)
}
