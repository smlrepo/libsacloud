// generated by 'github.com/sacloud/libsacloud/internal/tools/gen-api-interfaces'; DO NOT EDIT

package sacloud

import (
	"context"

	"github.com/sacloud/libsacloud/sacloud/types"
)

/*************************************************
* ArchiveAPI
*************************************************/

// ArchiveAPI is interface for operate Archive resource
type ArchiveAPI interface {
	Find(ctx context.Context, zone string, conditions *FindCondition) ([]*Archive, error)
	Create(ctx context.Context, zone string, param *ArchiveCreateRequest) (*Archive, error)
	CreateBlank(ctx context.Context, zone string, param *ArchiveCreateBlankRequest) (*Archive, *FTPServer, error)
	Read(ctx context.Context, zone string, id types.ID) (*Archive, error)
	Update(ctx context.Context, zone string, id types.ID, param *ArchiveUpdateRequest) (*Archive, error)
	Delete(ctx context.Context, zone string, id types.ID) error
	OpenFTP(ctx context.Context, zone string, id types.ID, openOption *OpenFTPRequest) (*FTPServer, error)
	CloseFTP(ctx context.Context, zone string, id types.ID) error
}

/*************************************************
* BridgeAPI
*************************************************/

// BridgeAPI is interface for operate Bridge resource
type BridgeAPI interface {
	Find(ctx context.Context, zone string, conditions *FindCondition) ([]*Bridge, error)
	Create(ctx context.Context, zone string, param *BridgeCreateRequest) (*Bridge, error)
	Read(ctx context.Context, zone string, id types.ID) (*Bridge, error)
	Update(ctx context.Context, zone string, id types.ID, param *BridgeUpdateRequest) (*Bridge, error)
	Delete(ctx context.Context, zone string, id types.ID) error
}

/*************************************************
* CDROMAPI
*************************************************/

// CDROMAPI is interface for operate CDROM resource
type CDROMAPI interface {
	Find(ctx context.Context, zone string, conditions *FindCondition) ([]*CDROM, error)
	Create(ctx context.Context, zone string, param *CDROMCreateRequest) (*CDROM, *FTPServer, error)
	Read(ctx context.Context, zone string, id types.ID) (*CDROM, error)
	Update(ctx context.Context, zone string, id types.ID, param *CDROMUpdateRequest) (*CDROM, error)
	Delete(ctx context.Context, zone string, id types.ID) error
	OpenFTP(ctx context.Context, zone string, id types.ID, openOption *OpenFTPRequest) (*FTPServer, error)
	CloseFTP(ctx context.Context, zone string, id types.ID) error
}

/*************************************************
* DiskAPI
*************************************************/

// DiskAPI is interface for operate Disk resource
type DiskAPI interface {
	Find(ctx context.Context, zone string, conditions *FindCondition) ([]*Disk, error)
	Create(ctx context.Context, zone string, param *DiskCreateRequest) (*Disk, error)
	CreateDistantly(ctx context.Context, zone string, createParam *DiskCreateRequest, distantFrom []types.ID) (*Disk, error)
	Config(ctx context.Context, zone string, id types.ID, edit *DiskEditRequest) error
	CreateWithConfig(ctx context.Context, zone string, createParam *DiskCreateRequest, editParam *DiskEditRequest, bootAtAvailable bool) (*Disk, error)
	CreateWithConfigDistantly(ctx context.Context, zone string, createParam *DiskCreateRequest, editParam *DiskEditRequest, bootAtAvailable bool, distantFrom []types.ID) (*Disk, error)
	ToBlank(ctx context.Context, zone string, id types.ID) error
	ResizePartition(ctx context.Context, zone string, id types.ID) error
	ConnectToServer(ctx context.Context, zone string, id types.ID, serverID types.ID) error
	DisconnectFromServer(ctx context.Context, zone string, id types.ID) error
	InstallDistantFrom(ctx context.Context, zone string, id types.ID, installParam *DiskInstallRequest, distantFrom []types.ID) (*Disk, error)
	Install(ctx context.Context, zone string, id types.ID, installParam *DiskInstallRequest) (*Disk, error)
	Read(ctx context.Context, zone string, id types.ID) (*Disk, error)
	Update(ctx context.Context, zone string, id types.ID, param *DiskUpdateRequest) (*Disk, error)
	Delete(ctx context.Context, zone string, id types.ID) error
	Monitor(ctx context.Context, zone string, id types.ID, condition *MonitorCondition) (*DiskActivity, error)
}

/*************************************************
* GSLBAPI
*************************************************/

// GSLBAPI is interface for operate GSLB resource
type GSLBAPI interface {
	Find(ctx context.Context, zone string, conditions *FindCondition) ([]*GSLB, error)
	Create(ctx context.Context, zone string, param *GSLBCreateRequest) (*GSLB, error)
	Read(ctx context.Context, zone string, id types.ID) (*GSLB, error)
	Update(ctx context.Context, zone string, id types.ID, param *GSLBUpdateRequest) (*GSLB, error)
	Delete(ctx context.Context, zone string, id types.ID) error
}

/*************************************************
* InterfaceAPI
*************************************************/

// InterfaceAPI is interface for operate Interface resource
type InterfaceAPI interface {
	Find(ctx context.Context, zone string, conditions *FindCondition) ([]*Interface, error)
	Create(ctx context.Context, zone string, param *InterfaceCreateRequest) (*Interface, error)
	Read(ctx context.Context, zone string, id types.ID) (*Interface, error)
	Update(ctx context.Context, zone string, id types.ID, param *InterfaceUpdateRequest) (*Interface, error)
	Delete(ctx context.Context, zone string, id types.ID) error
	Monitor(ctx context.Context, zone string, id types.ID, condition *MonitorCondition) (*InterfaceActivity, error)
	ConnectToSharedSegment(ctx context.Context, zone string, id types.ID) error
	ConnectToSwitch(ctx context.Context, zone string, id types.ID, switchID types.ID) error
	DisconnectFromSwitch(ctx context.Context, zone string, id types.ID) error
	ConnectToPacketFilter(ctx context.Context, zone string, id types.ID, packetFilterID types.ID) error
	DisconnectFromPacketFilter(ctx context.Context, zone string, id types.ID) error
}

/*************************************************
* InternetAPI
*************************************************/

// InternetAPI is interface for operate Internet resource
type InternetAPI interface {
	Find(ctx context.Context, zone string, conditions *FindCondition) ([]*Internet, error)
	Create(ctx context.Context, zone string, param *InternetCreateRequest) (*Internet, error)
	Read(ctx context.Context, zone string, id types.ID) (*Internet, error)
	Update(ctx context.Context, zone string, id types.ID, param *InternetUpdateRequest) (*Internet, error)
	Delete(ctx context.Context, zone string, id types.ID) error
	UpdateBandWidth(ctx context.Context, zone string, id types.ID, param *InternetUpdateBandWidthRequest) (*Internet, error)
	AddSubnet(ctx context.Context, zone string, id types.ID, param *InternetAddSubnetRequest) (*InternetSubnetOperationResult, error)
	UpdateSubnet(ctx context.Context, zone string, id types.ID, subnetID types.ID, param *InternetUpdateSubnetRequest) (*InternetSubnetOperationResult, error)
	DeleteSubnet(ctx context.Context, zone string, id types.ID, subnetID types.ID) error
	Monitor(ctx context.Context, zone string, id types.ID, condition *MonitorCondition) (*RouterActivity, error)
}

/*************************************************
* LoadBalancerAPI
*************************************************/

// LoadBalancerAPI is interface for operate LoadBalancer resource
type LoadBalancerAPI interface {
	Find(ctx context.Context, zone string, conditions *FindCondition) ([]*LoadBalancer, error)
	Create(ctx context.Context, zone string, param *LoadBalancerCreateRequest) (*LoadBalancer, error)
	Read(ctx context.Context, zone string, id types.ID) (*LoadBalancer, error)
	Update(ctx context.Context, zone string, id types.ID, param *LoadBalancerUpdateRequest) (*LoadBalancer, error)
	Delete(ctx context.Context, zone string, id types.ID) error
	Config(ctx context.Context, zone string, id types.ID) error
	Boot(ctx context.Context, zone string, id types.ID) error
	Shutdown(ctx context.Context, zone string, id types.ID, shutdownOption *ShutdownOption) error
	Reset(ctx context.Context, zone string, id types.ID) error
	MonitorInterface(ctx context.Context, zone string, id types.ID, condition *MonitorCondition) (*InterfaceActivity, error)
	Status(ctx context.Context, zone string, id types.ID) ([]*LoadBalancerStatus, error)
}

/*************************************************
* NFSAPI
*************************************************/

// NFSAPI is interface for operate NFS resource
type NFSAPI interface {
	Find(ctx context.Context, zone string, conditions *FindCondition) ([]*NFS, error)
	Create(ctx context.Context, zone string, param *NFSCreateRequest) (*NFS, error)
	Read(ctx context.Context, zone string, id types.ID) (*NFS, error)
	Update(ctx context.Context, zone string, id types.ID, param *NFSUpdateRequest) (*NFS, error)
	Delete(ctx context.Context, zone string, id types.ID) error
	Boot(ctx context.Context, zone string, id types.ID) error
	Shutdown(ctx context.Context, zone string, id types.ID, shutdownOption *ShutdownOption) error
	Reset(ctx context.Context, zone string, id types.ID) error
	MonitorFreeDiskSize(ctx context.Context, zone string, id types.ID, condition *MonitorCondition) (*FreeDiskSizeActivity, error)
	MonitorInterface(ctx context.Context, zone string, id types.ID, condition *MonitorCondition) (*InterfaceActivity, error)
}

/*************************************************
* NoteAPI
*************************************************/

// NoteAPI is interface for operate Note resource
type NoteAPI interface {
	Find(ctx context.Context, zone string, conditions *FindCondition) ([]*Note, error)
	Create(ctx context.Context, zone string, param *NoteCreateRequest) (*Note, error)
	Read(ctx context.Context, zone string, id types.ID) (*Note, error)
	Update(ctx context.Context, zone string, id types.ID, param *NoteUpdateRequest) (*Note, error)
	Delete(ctx context.Context, zone string, id types.ID) error
}

/*************************************************
* PacketFilterAPI
*************************************************/

// PacketFilterAPI is interface for operate PacketFilter resource
type PacketFilterAPI interface {
	Find(ctx context.Context, zone string, conditions *FindCondition) ([]*PacketFilter, error)
	Create(ctx context.Context, zone string, param *PacketFilterCreateRequest) (*PacketFilter, error)
	Read(ctx context.Context, zone string, id types.ID) (*PacketFilter, error)
	Update(ctx context.Context, zone string, id types.ID, param *PacketFilterUpdateRequest) (*PacketFilter, error)
	Delete(ctx context.Context, zone string, id types.ID) error
}

/*************************************************
* ServerAPI
*************************************************/

// ServerAPI is interface for operate Server resource
type ServerAPI interface {
	Find(ctx context.Context, zone string, conditions *FindCondition) ([]*Server, error)
	Create(ctx context.Context, zone string, param *ServerCreateRequest) (*Server, error)
	Read(ctx context.Context, zone string, id types.ID) (*Server, error)
	Update(ctx context.Context, zone string, id types.ID, param *ServerUpdateRequest) (*Server, error)
	Delete(ctx context.Context, zone string, id types.ID) error
	ChangePlan(ctx context.Context, zone string, id types.ID, plan *ServerChangePlanRequest) (*Server, error)
	InsertCDROM(ctx context.Context, zone string, id types.ID, insertParam *InsertCDROMRequest) error
	EjectCDROM(ctx context.Context, zone string, id types.ID, insertParam *EjectCDROMRequest) error
	Boot(ctx context.Context, zone string, id types.ID) error
	Shutdown(ctx context.Context, zone string, id types.ID, shutdownOption *ShutdownOption) error
	Reset(ctx context.Context, zone string, id types.ID) error
	Monitor(ctx context.Context, zone string, id types.ID, condition *MonitorCondition) (*CPUTimeActivity, error)
}

/*************************************************
* SIMAPI
*************************************************/

// SIMAPI is interface for operate SIM resource
type SIMAPI interface {
	Find(ctx context.Context, zone string, conditions *FindCondition) ([]*SIM, error)
	Create(ctx context.Context, zone string, param *SIMCreateRequest) (*SIM, error)
	Read(ctx context.Context, zone string, id types.ID) (*SIM, error)
	Update(ctx context.Context, zone string, id types.ID, param *SIMUpdateRequest) (*SIM, error)
	Delete(ctx context.Context, zone string, id types.ID) error
	Activate(ctx context.Context, zone string, id types.ID) error
	Deactivate(ctx context.Context, zone string, id types.ID) error
	AssignIP(ctx context.Context, zone string, id types.ID, param *SIMAssignIPRequest) error
	ClearIP(ctx context.Context, zone string, id types.ID) error
	IMEILock(ctx context.Context, zone string, id types.ID, param *SIMIMEILockRequest) error
	IMEIUnlock(ctx context.Context, zone string, id types.ID) error
	Logs(ctx context.Context, zone string, id types.ID) ([]*SIMLog, error)
	GetNetworkOperator(ctx context.Context, zone string, id types.ID) ([]*SIMNetworkOperatorConfig, error)
	SetNetworkOperator(ctx context.Context, zone string, id types.ID, configs *SIMNetworkOperatorConfigs) error
	MonitorSIM(ctx context.Context, zone string, id types.ID, condition *MonitorCondition) (*LinkActivity, error)
}

/*************************************************
* SwitchAPI
*************************************************/

// SwitchAPI is interface for operate Switch resource
type SwitchAPI interface {
	Find(ctx context.Context, zone string, conditions *FindCondition) ([]*Switch, error)
	Create(ctx context.Context, zone string, param *SwitchCreateRequest) (*Switch, error)
	Read(ctx context.Context, zone string, id types.ID) (*Switch, error)
	Update(ctx context.Context, zone string, id types.ID, param *SwitchUpdateRequest) (*Switch, error)
	Delete(ctx context.Context, zone string, id types.ID) error
	ConnectToBridge(ctx context.Context, zone string, id types.ID, bridgeID types.ID) error
	DisconnectFromBridge(ctx context.Context, zone string, id types.ID) error
}

/*************************************************
* VPCRouterAPI
*************************************************/

// VPCRouterAPI is interface for operate VPCRouter resource
type VPCRouterAPI interface {
	Find(ctx context.Context, zone string, conditions *FindCondition) ([]*VPCRouter, error)
	Create(ctx context.Context, zone string, param *VPCRouterCreateRequest) (*VPCRouter, error)
	Read(ctx context.Context, zone string, id types.ID) (*VPCRouter, error)
	Update(ctx context.Context, zone string, id types.ID, param *VPCRouterUpdateRequest) (*VPCRouter, error)
	Delete(ctx context.Context, zone string, id types.ID) error
	Config(ctx context.Context, zone string, id types.ID) error
	Boot(ctx context.Context, zone string, id types.ID) error
	Shutdown(ctx context.Context, zone string, id types.ID, shutdownOption *ShutdownOption) error
	Reset(ctx context.Context, zone string, id types.ID) error
	ConnectToSwitch(ctx context.Context, zone string, id types.ID, nicIndex int, switchID types.ID) error
	DisconnectFromSwitch(ctx context.Context, zone string, id types.ID, nicIndex int) error
	MonitorInterface(ctx context.Context, zone string, id types.ID, index int, condition *MonitorCondition) (*InterfaceActivity, error)
}

/*************************************************
* ZoneAPI
*************************************************/

// ZoneAPI is interface for operate Zone resource
type ZoneAPI interface {
	Find(ctx context.Context, zone string, conditions *FindCondition) ([]*Zone, error)
	Read(ctx context.Context, zone string, id types.ID) (*Zone, error)
}
