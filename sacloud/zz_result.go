// generated by 'github.com/sacloud/libsacloud/internal/tools/gen-api-result'; DO NOT EDIT

package sacloud

// ArchiveFindResult represents the Result of API
type ArchiveFindResult struct {
	Total int `json:",omitempty"` // Total count of target resources
	From  int `json:",omitempty"` // Current page number
	Count int `json:",omitempty"` // Count of current page

	Archives []*Archive `json:",omitempty" mapconv:"[]Archives,omitempty,recursive"`
}

// archiveCreateResult represents the Result of API
type archiveCreateResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	Archive *Archive `json:",omitempty" mapconv:"Archive,omitempty,recursive"`
}

// archiveCreateBlankResult represents the Result of API
type archiveCreateBlankResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	Archive   *Archive   `json:",omitempty" mapconv:"Archive,omitempty,recursive"`
	FTPServer *FTPServer `json:",omitempty" mapconv:"FTPServer,omitempty,recursive"`
}

// archiveReadResult represents the Result of API
type archiveReadResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	Archive *Archive `json:",omitempty" mapconv:"Archive,omitempty,recursive"`
}

// archiveUpdateResult represents the Result of API
type archiveUpdateResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	Archive *Archive `json:",omitempty" mapconv:"Archive,omitempty,recursive"`
}

// archiveOpenFTPResult represents the Result of API
type archiveOpenFTPResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	FTPServer *FTPServer `json:",omitempty" mapconv:"FTPServer,omitempty,recursive"`
}

// authStatusReadResult represents the Result of API
type authStatusReadResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	AuthStatus *AuthStatus `json:",omitempty" mapconv:"AuthStatus,omitempty,recursive"`
}

// AutoBackupFindResult represents the Result of API
type AutoBackupFindResult struct {
	Total int `json:",omitempty"` // Total count of target resources
	From  int `json:",omitempty"` // Current page number
	Count int `json:",omitempty"` // Count of current page

	AutoBackups []*AutoBackup `json:",omitempty" mapconv:"[]CommonServiceItems,omitempty,recursive"`
}

// autoBackupCreateResult represents the Result of API
type autoBackupCreateResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	AutoBackup *AutoBackup `json:",omitempty" mapconv:"CommonServiceItem,omitempty,recursive"`
}

// autoBackupReadResult represents the Result of API
type autoBackupReadResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	AutoBackup *AutoBackup `json:",omitempty" mapconv:"CommonServiceItem,omitempty,recursive"`
}

// autoBackupUpdateResult represents the Result of API
type autoBackupUpdateResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	AutoBackup *AutoBackup `json:",omitempty" mapconv:"CommonServiceItem,omitempty,recursive"`
}

// BridgeFindResult represents the Result of API
type BridgeFindResult struct {
	Total int `json:",omitempty"` // Total count of target resources
	From  int `json:",omitempty"` // Current page number
	Count int `json:",omitempty"` // Count of current page

	Bridges []*Bridge `json:",omitempty" mapconv:"[]Bridges,omitempty,recursive"`
}

// bridgeCreateResult represents the Result of API
type bridgeCreateResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	Bridge *Bridge `json:",omitempty" mapconv:"Bridge,omitempty,recursive"`
}

// bridgeReadResult represents the Result of API
type bridgeReadResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	Bridge *Bridge `json:",omitempty" mapconv:"Bridge,omitempty,recursive"`
}

// bridgeUpdateResult represents the Result of API
type bridgeUpdateResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	Bridge *Bridge `json:",omitempty" mapconv:"Bridge,omitempty,recursive"`
}

// CDROMFindResult represents the Result of API
type CDROMFindResult struct {
	Total int `json:",omitempty"` // Total count of target resources
	From  int `json:",omitempty"` // Current page number
	Count int `json:",omitempty"` // Count of current page

	CDROMs []*CDROM `json:",omitempty" mapconv:"[]CDROMs,omitempty,recursive"`
}

// cDROMCreateResult represents the Result of API
type cDROMCreateResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	CDROM     *CDROM     `json:",omitempty" mapconv:"CDROM,omitempty,recursive"`
	FTPServer *FTPServer `json:",omitempty" mapconv:"FTPServer,omitempty,recursive"`
}

// cDROMReadResult represents the Result of API
type cDROMReadResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	CDROM *CDROM `json:",omitempty" mapconv:"CDROM,omitempty,recursive"`
}

// cDROMUpdateResult represents the Result of API
type cDROMUpdateResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	CDROM *CDROM `json:",omitempty" mapconv:"CDROM,omitempty,recursive"`
}

// cDROMOpenFTPResult represents the Result of API
type cDROMOpenFTPResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	FTPServer *FTPServer `json:",omitempty" mapconv:"FTPServer,omitempty,recursive"`
}

// DiskFindResult represents the Result of API
type DiskFindResult struct {
	Total int `json:",omitempty"` // Total count of target resources
	From  int `json:",omitempty"` // Current page number
	Count int `json:",omitempty"` // Count of current page

	Disks []*Disk `json:",omitempty" mapconv:"[]Disks,omitempty,recursive"`
}

// diskCreateResult represents the Result of API
type diskCreateResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	Disk *Disk `json:",omitempty" mapconv:"Disk,omitempty,recursive"`
}

// diskCreateDistantlyResult represents the Result of API
type diskCreateDistantlyResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	Disk *Disk `json:",omitempty" mapconv:"Disk,omitempty,recursive"`
}

// diskCreateWithConfigResult represents the Result of API
type diskCreateWithConfigResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	Disk *Disk `json:",omitempty" mapconv:"Disk,omitempty,recursive"`
}

// diskCreateWithConfigDistantlyResult represents the Result of API
type diskCreateWithConfigDistantlyResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	Disk *Disk `json:",omitempty" mapconv:"Disk,omitempty,recursive"`
}

// diskInstallDistantFromResult represents the Result of API
type diskInstallDistantFromResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	Disk *Disk `json:",omitempty" mapconv:"Disk,omitempty,recursive"`
}

// diskInstallResult represents the Result of API
type diskInstallResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	Disk *Disk `json:",omitempty" mapconv:"Disk,omitempty,recursive"`
}

// diskReadResult represents the Result of API
type diskReadResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	Disk *Disk `json:",omitempty" mapconv:"Disk,omitempty,recursive"`
}

// diskUpdateResult represents the Result of API
type diskUpdateResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	Disk *Disk `json:",omitempty" mapconv:"Disk,omitempty,recursive"`
}

// diskMonitorResult represents the Result of API
type diskMonitorResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	DiskActivity *DiskActivity `json:",omitempty" mapconv:"Data,omitempty,recursive"`
}

// DNSFindResult represents the Result of API
type DNSFindResult struct {
	Total int `json:",omitempty"` // Total count of target resources
	From  int `json:",omitempty"` // Current page number
	Count int `json:",omitempty"` // Count of current page

	DNS []*DNS `json:",omitempty" mapconv:"[]CommonServiceItems,omitempty,recursive"`
}

// dNSCreateResult represents the Result of API
type dNSCreateResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	DNS *DNS `json:",omitempty" mapconv:"CommonServiceItem,omitempty,recursive"`
}

// dNSReadResult represents the Result of API
type dNSReadResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	DNS *DNS `json:",omitempty" mapconv:"CommonServiceItem,omitempty,recursive"`
}

// dNSUpdateResult represents the Result of API
type dNSUpdateResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	DNS *DNS `json:",omitempty" mapconv:"CommonServiceItem,omitempty,recursive"`
}

// GSLBFindResult represents the Result of API
type GSLBFindResult struct {
	Total int `json:",omitempty"` // Total count of target resources
	From  int `json:",omitempty"` // Current page number
	Count int `json:",omitempty"` // Count of current page

	GSLBs []*GSLB `json:",omitempty" mapconv:"[]CommonServiceItems,omitempty,recursive"`
}

// gSLBCreateResult represents the Result of API
type gSLBCreateResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	GSLB *GSLB `json:",omitempty" mapconv:"CommonServiceItem,omitempty,recursive"`
}

// gSLBReadResult represents the Result of API
type gSLBReadResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	GSLB *GSLB `json:",omitempty" mapconv:"CommonServiceItem,omitempty,recursive"`
}

// gSLBUpdateResult represents the Result of API
type gSLBUpdateResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	GSLB *GSLB `json:",omitempty" mapconv:"CommonServiceItem,omitempty,recursive"`
}

// IconFindResult represents the Result of API
type IconFindResult struct {
	Total int `json:",omitempty"` // Total count of target resources
	From  int `json:",omitempty"` // Current page number
	Count int `json:",omitempty"` // Count of current page

	Icons []*Icon `json:",omitempty" mapconv:"[]Icons,omitempty,recursive"`
}

// iconCreateResult represents the Result of API
type iconCreateResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	Icon *Icon `json:",omitempty" mapconv:"Icon,omitempty,recursive"`
}

// iconReadResult represents the Result of API
type iconReadResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	Icon *Icon `json:",omitempty" mapconv:"Icon,omitempty,recursive"`
}

// iconUpdateResult represents the Result of API
type iconUpdateResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	Icon *Icon `json:",omitempty" mapconv:"Icon,omitempty,recursive"`
}

// InterfaceFindResult represents the Result of API
type InterfaceFindResult struct {
	Total int `json:",omitempty"` // Total count of target resources
	From  int `json:",omitempty"` // Current page number
	Count int `json:",omitempty"` // Count of current page

	Interfaces []*Interface `json:",omitempty" mapconv:"[]Interfaces,omitempty,recursive"`
}

// interfaceCreateResult represents the Result of API
type interfaceCreateResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	Interface *Interface `json:",omitempty" mapconv:"Interface,omitempty,recursive"`
}

// interfaceReadResult represents the Result of API
type interfaceReadResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	Interface *Interface `json:",omitempty" mapconv:"Interface,omitempty,recursive"`
}

// interfaceUpdateResult represents the Result of API
type interfaceUpdateResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	Interface *Interface `json:",omitempty" mapconv:"Interface,omitempty,recursive"`
}

// interfaceMonitorResult represents the Result of API
type interfaceMonitorResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	InterfaceActivity *InterfaceActivity `json:",omitempty" mapconv:"Data,omitempty,recursive"`
}

// InternetFindResult represents the Result of API
type InternetFindResult struct {
	Total int `json:",omitempty"` // Total count of target resources
	From  int `json:",omitempty"` // Current page number
	Count int `json:",omitempty"` // Count of current page

	Internets []*Internet `json:",omitempty" mapconv:"[]Internets,omitempty,recursive"`
}

// internetCreateResult represents the Result of API
type internetCreateResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	Internet *Internet `json:",omitempty" mapconv:"Internet,omitempty,recursive"`
}

// internetReadResult represents the Result of API
type internetReadResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	Internet *Internet `json:",omitempty" mapconv:"Internet,omitempty,recursive"`
}

// internetUpdateResult represents the Result of API
type internetUpdateResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	Internet *Internet `json:",omitempty" mapconv:"Internet,omitempty,recursive"`
}

// internetUpdateBandWidthResult represents the Result of API
type internetUpdateBandWidthResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	Internet *Internet `json:",omitempty" mapconv:"Internet,omitempty,recursive"`
}

// internetAddSubnetResult represents the Result of API
type internetAddSubnetResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	Subnet *InternetSubnetOperationResult `json:",omitempty" mapconv:"Subnet,omitempty,recursive"`
}

// internetUpdateSubnetResult represents the Result of API
type internetUpdateSubnetResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	Subnet *InternetSubnetOperationResult `json:",omitempty" mapconv:"Subnet,omitempty,recursive"`
}

// internetMonitorResult represents the Result of API
type internetMonitorResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	RouterActivity *RouterActivity `json:",omitempty" mapconv:"Data,omitempty,recursive"`
}

// LoadBalancerFindResult represents the Result of API
type LoadBalancerFindResult struct {
	Total int `json:",omitempty"` // Total count of target resources
	From  int `json:",omitempty"` // Current page number
	Count int `json:",omitempty"` // Count of current page

	LoadBalancers []*LoadBalancer `json:",omitempty" mapconv:"[]Appliances,omitempty,recursive"`
}

// loadBalancerCreateResult represents the Result of API
type loadBalancerCreateResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	LoadBalancer *LoadBalancer `json:",omitempty" mapconv:"Appliance,omitempty,recursive"`
}

// loadBalancerReadResult represents the Result of API
type loadBalancerReadResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	LoadBalancer *LoadBalancer `json:",omitempty" mapconv:"Appliance,omitempty,recursive"`
}

// loadBalancerUpdateResult represents the Result of API
type loadBalancerUpdateResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	LoadBalancer *LoadBalancer `json:",omitempty" mapconv:"Appliance,omitempty,recursive"`
}

// loadBalancerMonitorInterfaceResult represents the Result of API
type loadBalancerMonitorInterfaceResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	InterfaceActivity *InterfaceActivity `json:",omitempty" mapconv:"Data,omitempty,recursive"`
}

// LoadBalancerStatusResult represents the Result of API
type LoadBalancerStatusResult struct {
	Total int `json:",omitempty"` // Total count of target resources
	From  int `json:",omitempty"` // Current page number
	Count int `json:",omitempty"` // Count of current page

	Status []*LoadBalancerStatus `json:",omitempty" mapconv:"[]LoadBalancer,omitempty,recursive"`
}

// NFSFindResult represents the Result of API
type NFSFindResult struct {
	Total int `json:",omitempty"` // Total count of target resources
	From  int `json:",omitempty"` // Current page number
	Count int `json:",omitempty"` // Count of current page

	NFS []*NFS `json:",omitempty" mapconv:"[]Appliances,omitempty,recursive"`
}

// nFSCreateResult represents the Result of API
type nFSCreateResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	NFS *NFS `json:",omitempty" mapconv:"Appliance,omitempty,recursive"`
}

// nFSReadResult represents the Result of API
type nFSReadResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	NFS *NFS `json:",omitempty" mapconv:"Appliance,omitempty,recursive"`
}

// nFSUpdateResult represents the Result of API
type nFSUpdateResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	NFS *NFS `json:",omitempty" mapconv:"Appliance,omitempty,recursive"`
}

// nFSMonitorFreeDiskSizeResult represents the Result of API
type nFSMonitorFreeDiskSizeResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	FreeDiskSizeActivity *FreeDiskSizeActivity `json:",omitempty" mapconv:"Data,omitempty,recursive"`
}

// nFSMonitorInterfaceResult represents the Result of API
type nFSMonitorInterfaceResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	InterfaceActivity *InterfaceActivity `json:",omitempty" mapconv:"Data,omitempty,recursive"`
}

// NoteFindResult represents the Result of API
type NoteFindResult struct {
	Total int `json:",omitempty"` // Total count of target resources
	From  int `json:",omitempty"` // Current page number
	Count int `json:",omitempty"` // Count of current page

	Notes []*Note `json:",omitempty" mapconv:"[]Notes,omitempty,recursive"`
}

// noteCreateResult represents the Result of API
type noteCreateResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	Note *Note `json:",omitempty" mapconv:"Note,omitempty,recursive"`
}

// noteReadResult represents the Result of API
type noteReadResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	Note *Note `json:",omitempty" mapconv:"Note,omitempty,recursive"`
}

// noteUpdateResult represents the Result of API
type noteUpdateResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	Note *Note `json:",omitempty" mapconv:"Note,omitempty,recursive"`
}

// PacketFilterFindResult represents the Result of API
type PacketFilterFindResult struct {
	Total int `json:",omitempty"` // Total count of target resources
	From  int `json:",omitempty"` // Current page number
	Count int `json:",omitempty"` // Count of current page

	PacketFilters []*PacketFilter `json:",omitempty" mapconv:"[]PacketFilters,omitempty,recursive"`
}

// packetFilterCreateResult represents the Result of API
type packetFilterCreateResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	PacketFilter *PacketFilter `json:",omitempty" mapconv:"PacketFilter,omitempty,recursive"`
}

// packetFilterReadResult represents the Result of API
type packetFilterReadResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	PacketFilter *PacketFilter `json:",omitempty" mapconv:"PacketFilter,omitempty,recursive"`
}

// packetFilterUpdateResult represents the Result of API
type packetFilterUpdateResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	PacketFilter *PacketFilter `json:",omitempty" mapconv:"PacketFilter,omitempty,recursive"`
}

// ProxyLBFindResult represents the Result of API
type ProxyLBFindResult struct {
	Total int `json:",omitempty"` // Total count of target resources
	From  int `json:",omitempty"` // Current page number
	Count int `json:",omitempty"` // Count of current page

	ProxyLBs []*ProxyLB `json:",omitempty" mapconv:"[]CommonServiceItems,omitempty,recursive"`
}

// proxyLBCreateResult represents the Result of API
type proxyLBCreateResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	ProxyLB *ProxyLB `json:",omitempty" mapconv:"CommonServiceItem,omitempty,recursive"`
}

// proxyLBReadResult represents the Result of API
type proxyLBReadResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	ProxyLB *ProxyLB `json:",omitempty" mapconv:"CommonServiceItem,omitempty,recursive"`
}

// proxyLBUpdateResult represents the Result of API
type proxyLBUpdateResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	ProxyLB *ProxyLB `json:",omitempty" mapconv:"CommonServiceItem,omitempty,recursive"`
}

// proxyLBChangePlanResult represents the Result of API
type proxyLBChangePlanResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	ProxyLB *ProxyLB `json:",omitempty" mapconv:"CommonServiceItem,omitempty,recursive"`
}

// proxyLBGetCertificatesResult represents the Result of API
type proxyLBGetCertificatesResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	ProxyLBCertificates *ProxyLBCertificates `json:",omitempty" mapconv:"ProxyLB,omitempty,recursive"`
}

// proxyLBSetCertificatesResult represents the Result of API
type proxyLBSetCertificatesResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	ProxyLBCertificates *ProxyLBCertificates `json:",omitempty" mapconv:"ProxyLB,omitempty,recursive"`
}

// proxyLBHealthStatusResult represents the Result of API
type proxyLBHealthStatusResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	ProxyLBHealth *ProxyLBHealth `json:",omitempty" mapconv:"ProxyLB,omitempty,recursive"`
}

// ServerFindResult represents the Result of API
type ServerFindResult struct {
	Total int `json:",omitempty"` // Total count of target resources
	From  int `json:",omitempty"` // Current page number
	Count int `json:",omitempty"` // Count of current page

	Servers []*Server `json:",omitempty" mapconv:"[]Servers,omitempty,recursive"`
}

// serverCreateResult represents the Result of API
type serverCreateResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	Server *Server `json:",omitempty" mapconv:"Server,omitempty,recursive"`
}

// serverReadResult represents the Result of API
type serverReadResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	Server *Server `json:",omitempty" mapconv:"Server,omitempty,recursive"`
}

// serverUpdateResult represents the Result of API
type serverUpdateResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	Server *Server `json:",omitempty" mapconv:"Server,omitempty,recursive"`
}

// serverChangePlanResult represents the Result of API
type serverChangePlanResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	Server *Server `json:",omitempty" mapconv:"Server,omitempty,recursive"`
}

// serverMonitorResult represents the Result of API
type serverMonitorResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	CPUTimeActivity *CPUTimeActivity `json:",omitempty" mapconv:"Data,omitempty,recursive"`
}

// SIMFindResult represents the Result of API
type SIMFindResult struct {
	Total int `json:",omitempty"` // Total count of target resources
	From  int `json:",omitempty"` // Current page number
	Count int `json:",omitempty"` // Count of current page

	SIMs []*SIM `json:",omitempty" mapconv:"[]CommonServiceItems,omitempty,recursive"`
}

// sIMCreateResult represents the Result of API
type sIMCreateResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	SIM *SIM `json:",omitempty" mapconv:"CommonServiceItem,omitempty,recursive"`
}

// sIMReadResult represents the Result of API
type sIMReadResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	SIM *SIM `json:",omitempty" mapconv:"CommonServiceItem,omitempty,recursive"`
}

// sIMUpdateResult represents the Result of API
type sIMUpdateResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	SIM *SIM `json:",omitempty" mapconv:"CommonServiceItem,omitempty,recursive"`
}

// SIMLogsResult represents the Result of API
type SIMLogsResult struct {
	Total int `json:",omitempty"` // Total count of target resources
	From  int `json:",omitempty"` // Current page number
	Count int `json:",omitempty"` // Count of current page

	Logs []*SIMLog `json:",omitempty" mapconv:"[]Logs,omitempty,recursive"`
}

// SIMGetNetworkOperatorResult represents the Result of API
type SIMGetNetworkOperatorResult struct {
	Total int `json:",omitempty"` // Total count of target resources
	From  int `json:",omitempty"` // Current page number
	Count int `json:",omitempty"` // Count of current page

	Configs []*SIMNetworkOperatorConfig `json:",omitempty" mapconv:"[]NetworkOperationConfigs,omitempty,recursive"`
}

// sIMMonitorSIMResult represents the Result of API
type sIMMonitorSIMResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	LinkActivity *LinkActivity `json:",omitempty" mapconv:"Data,omitempty,recursive"`
}

// SimpleMonitorFindResult represents the Result of API
type SimpleMonitorFindResult struct {
	Total int `json:",omitempty"` // Total count of target resources
	From  int `json:",omitempty"` // Current page number
	Count int `json:",omitempty"` // Count of current page

	SimpleMonitors []*SimpleMonitor `json:",omitempty" mapconv:"[]CommonServiceItems,omitempty,recursive"`
}

// simpleMonitorCreateResult represents the Result of API
type simpleMonitorCreateResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	SimpleMonitor *SimpleMonitor `json:",omitempty" mapconv:"CommonServiceItem,omitempty,recursive"`
}

// simpleMonitorReadResult represents the Result of API
type simpleMonitorReadResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	SimpleMonitor *SimpleMonitor `json:",omitempty" mapconv:"CommonServiceItem,omitempty,recursive"`
}

// simpleMonitorUpdateResult represents the Result of API
type simpleMonitorUpdateResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	SimpleMonitor *SimpleMonitor `json:",omitempty" mapconv:"CommonServiceItem,omitempty,recursive"`
}

// simpleMonitorMonitorResponseTimeResult represents the Result of API
type simpleMonitorMonitorResponseTimeResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	ResponseTimeSecActivity *ResponseTimeSecActivity `json:",omitempty" mapconv:"Data,omitempty,recursive"`
}

// simpleMonitorHealthStatusResult represents the Result of API
type simpleMonitorHealthStatusResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	SimpleMonitorHealthStatus *SimpleMonitorHealthStatus `json:",omitempty" mapconv:"SimpleMonitor,omitempty,recursive"`
}

// SSHKeyFindResult represents the Result of API
type SSHKeyFindResult struct {
	Total int `json:",omitempty"` // Total count of target resources
	From  int `json:",omitempty"` // Current page number
	Count int `json:",omitempty"` // Count of current page

	SSHKeys []*SSHKey `json:",omitempty" mapconv:"[]SSHKeys,omitempty,recursive"`
}

// sSHKeyCreateResult represents the Result of API
type sSHKeyCreateResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	SSHKey *SSHKey `json:",omitempty" mapconv:"SSHKey,omitempty,recursive"`
}

// sSHKeyGenerateResult represents the Result of API
type sSHKeyGenerateResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	SSHKeyGenerated *SSHKeyGenerated `json:",omitempty" mapconv:"SSHKey,omitempty,recursive"`
}

// sSHKeyReadResult represents the Result of API
type sSHKeyReadResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	SSHKey *SSHKey `json:",omitempty" mapconv:"SSHKey,omitempty,recursive"`
}

// sSHKeyUpdateResult represents the Result of API
type sSHKeyUpdateResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	SSHKey *SSHKey `json:",omitempty" mapconv:"SSHKey,omitempty,recursive"`
}

// SwitchFindResult represents the Result of API
type SwitchFindResult struct {
	Total int `json:",omitempty"` // Total count of target resources
	From  int `json:",omitempty"` // Current page number
	Count int `json:",omitempty"` // Count of current page

	Switches []*Switch `json:",omitempty" mapconv:"[]Switches,omitempty,recursive"`
}

// switchCreateResult represents the Result of API
type switchCreateResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	Switch *Switch `json:",omitempty" mapconv:"Switch,omitempty,recursive"`
}

// switchReadResult represents the Result of API
type switchReadResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	Switch *Switch `json:",omitempty" mapconv:"Switch,omitempty,recursive"`
}

// switchUpdateResult represents the Result of API
type switchUpdateResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	Switch *Switch `json:",omitempty" mapconv:"Switch,omitempty,recursive"`
}

// VPCRouterFindResult represents the Result of API
type VPCRouterFindResult struct {
	Total int `json:",omitempty"` // Total count of target resources
	From  int `json:",omitempty"` // Current page number
	Count int `json:",omitempty"` // Count of current page

	VPCRouters []*VPCRouter `json:",omitempty" mapconv:"[]Appliances,omitempty,recursive"`
}

// vPCRouterCreateResult represents the Result of API
type vPCRouterCreateResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	VPCRouter *VPCRouter `json:",omitempty" mapconv:"Appliance,omitempty,recursive"`
}

// vPCRouterReadResult represents the Result of API
type vPCRouterReadResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	VPCRouter *VPCRouter `json:",omitempty" mapconv:"Appliance,omitempty,recursive"`
}

// vPCRouterUpdateResult represents the Result of API
type vPCRouterUpdateResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	VPCRouter *VPCRouter `json:",omitempty" mapconv:"Appliance,omitempty,recursive"`
}

// vPCRouterMonitorInterfaceResult represents the Result of API
type vPCRouterMonitorInterfaceResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	InterfaceActivity *InterfaceActivity `json:",omitempty" mapconv:"Data,omitempty,recursive"`
}

// ZoneFindResult represents the Result of API
type ZoneFindResult struct {
	Total int `json:",omitempty"` // Total count of target resources
	From  int `json:",omitempty"` // Current page number
	Count int `json:",omitempty"` // Count of current page

	Zones []*Zone `json:",omitempty" mapconv:"[]Zones,omitempty,recursive"`
}

// zoneReadResult represents the Result of API
type zoneReadResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	Zone *Zone `json:",omitempty" mapconv:"Zone,omitempty,recursive"`
}
