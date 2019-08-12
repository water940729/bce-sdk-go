package appblb

type BLBStatus string

const (
	BLBStatusCreating    BLBStatus = "creating"
	BLBStatusAvailable   BLBStatus = "available"
	BLBStatusUpdating    BLBStatus = "updating"
	BLBStatusPaused      BLBStatus = "paused"
	BLBStatusUnavailable BLBStatus = "unavailable"
)

type AppRsPortModel struct {
	ListenerPort        int    `json:"listenerPort"`
	BackendPort         int    `json:"backendPort"`
	PortType            string `json:"portType"`
	HealthCheckPortType string `json:"healthCheckPortType"`
	Status              string `json:"status"`
	PortId              string `json:"portId"`
	PolicyId            string `json:"policyId"`
}

type AppBackendServer struct {
	InstanceId string           `json:"instanceId,omitempty"`
	Weight     int              `json:"weight,omitempty"`
	PrivateIp  string           `json:"privateIp,omitempty"`
	PortList   []AppRsPortModel `json:"portList,omitempty"`
}

type DescribeResultMeta struct {
	Marker      string `json:"marker"`
	IsTruncated bool   `json:"isTruncated"`
	NextMarker  string `json:"nextMarker"`
	MaxKeys     int    `json:"maxKeys"`
}

type CreateAppServerGroupArgs struct {
	Name              string             `json:"name,omitempty"`
	Description       string             `json:"desc,omitempty"`
	BackendServerList []AppBackendServer `json:"backendServerList,omitempty"`
	ClientToken       string             `json:"-"`
}

type CreateAppServerGroupResult struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"desc"`
	Status      BLBStatus `json:"status"`
}

type UpdateAppServerGroupArgs struct {
	SgId        string `json:"sgId"`
	Name        string `json:"name,omitempty"`
	Description string `json:"desc,omitempty"`
	ClientToken string `json:"-"`
}

type DescribeAppServerGroupArgs struct {
	Name         string
	ExactlyMatch string
	Marker       string
	MaxKeys      int
}

type AppServerGroupPort struct {
	Id                          string    `json:"id"`
	Port                        int       `json:"port"`
	Type                        string    `json:"type"`
	Status                      BLBStatus `json:"status"`
	HealthCheck                 string    `json:"healthCheck"`
	HealthCheckPort             int       `json:"healthCheckPort"`
	HealthCheckTimeoutInSecond  int       `json:"healthCheckTimeoutInSecond"`
	HealthCheckIntervalInSecond int       `json:"healthCheckIntervalInSecond"`
	HealthCheckDownRetry        int       `json:"healthCheckDownRetry"`
	HealthCheckUpRetry          int       `json:"healthCheckUpRetry"`
	HealthCheckNormalStatus     string    `json:"healthCheckNormalStatus"`
	HealthCheckUrlPath          string    `json:"healthCheckUrlPath"`
	UdpHealthCheckString        string    `json:"udpHealthCheckString"`
}

type AppServerGroup struct {
	Id          string               `json:"id"`
	Name        string               `json:"name"`
	Description string               `json:"desc"`
	Status      BLBStatus            `json:"status"`
	PortList    []AppServerGroupPort `json:"portList"`
}

type DescribeAppServerGroupResult struct {
	DescribeResultMeta
	AppServerGroupList []AppServerGroup `json:"appServerGroupList"`
}

type DeleteAppServerGroupArgs struct {
	SgId        string `json:"sgId"`
	ClientToken string `json:"-"`
}

type CreateAppServerGroupPortArgs struct {
	ClientToken                 string `json:"-"`
	SgId                        string `json:"sgId"`
	Port                        uint16 `json:"port"`
	Type                        string `json:"type"`
	HealthCheck                 string `json:"healthCheck,omitempty"`
	HealthCheckPort             int    `json:"healthCheckPort,omitempty"`
	HealthCheckTimeoutInSecond  int    `json:"healthCheckTimeoutInSecond,omitempty"`
	HealthCheckIntervalInSecond int    `json:"healthCheckIntervalInSecond,omitempty"`
	HealthCheckDownRetry        int    `json:"healthCheckDownRetry,omitempty"`
	HealthCheckUpRetry          int    `json:"healthCheckUpRetry,omitempty"`
	HealthCheckNormalStatus     string `json:"healthCheckNormalStatus,omitempty"`
	HealthCheckUrlPath          string `json:"healthCheckUrlPath,omitempty"`
	UdpHealthCheckString        string `json:"udpHealthCheckString,omitempty"`
}

type CreateAppServerGroupPortResult struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"desc"`
	Status      BLBStatus `json:"status"`
}

type UpdateAppServerGroupPortArgs struct {
	ClientToken                 string `json:"-"`
	SgId                        string `json:"sgId"`
	PortId                      string `json:"portId"`
	HealthCheck                 string `json:"healthCheck,omitempty"`
	HealthCheckPort             int    `json:"healthCheckPort,omitempty"`
	HealthCheckUrlPath          string `json:"healthCheckUrlPath,omitempty"`
	HealthCheckTimeoutInSecond  int    `json:"healthCheckTimeoutInSecond,omitempty"`
	HealthCheckIntervalInSecond int    `json:"healthCheckIntervalInSecond,omitempty"`
	HealthCheckDownRetry        int    `json:"healthCheckDownRetry,omitempty"`
	HealthCheckUpRetry          int    `json:"healthCheckUpRetry,omitempty"`
	HealthCheckNormalStatus     string `json:"healthCheckNormalStatus,omitempty"`
	UdpHealthCheckString        string `json:"udpHealthCheckString,omitempty"`
}

type DeleteAppServerGroupPortArgs struct {
	SgId        string   `json:"sgId"`
	PortIdList  []string `json:"portIdList"`
	ClientToken string   `json:"-"`
}

type BlbRsWriteOpArgs struct {
	SgId              string             `json:"sgId"`
	BackendServerList []AppBackendServer `json:"backendServerList"`
	ClientToken       string             `json:"-"`
}

type CreateBlbRsArgs struct {
	BlbRsWriteOpArgs
}

type UpdateBlbRsArgs struct {
	BlbRsWriteOpArgs
}

type DescribeBlbRsArgs struct {
	Marker  string
	MaxKeys int
	SgId    string
}

type DescribeBlbRsResult struct {
	BackendServerList []AppBackendServer `json:"backendServerList"`
	DescribeResultMeta
}

type DeleteBlbRsArgs struct {
	SgId              string   `json:"sgId"`
	BackendServerList []string `json:"backendServerIdList"`
	ClientToken       string   `json:"-"`
}

type DescribeRsMountResult struct {
	BackendServerList []AppBackendServer `json:"backendServerList"`
}

type TagModel struct {
	TagKey   string `json:"tagKey"`
	TagValue string `json:"tagValue"`
}

type CreateLoadBalancerArgs struct {
	ClientToken string     `json:"-"`
	Name        string     `json:"name,omitempty"`
	Description string     `json:"desc,omitempty"`
	SubnetId    string     `json:"subnetId"`
	VpcId       string     `json:"vpcId"`
	Tags        []TagModel `json:"tags,omitempty"`
}

type CreateLoadBalanceResult struct {
	Address     string `json:"address"`
	Name        string `json:"name"`
	Description string `json:"desc"`
	BlbId       string `json:"blbId"`
}

type UpdateLoadBalancerArgs struct {
	ClientToken string `json:"-"`
	Name        string `json:"name,omitempty"`
	Description string `json:"desc,omitempty"`
}

type DescribeLoadBalancersArgs struct {
	Address      string
	Name         string
	BlbId        string
	BccId        string
	ExactlyMatch string
	Marker       string
	MaxKeys      int
}

type AppBLBModel struct {
	BlbId       string     `json:"blbId"`
	Name        string     `json:"name"`
	Description string     `json:"desc"`
	Address     string     `json:"address"`
	Status      BLBStatus  `json:"status"`
	VpcId       string     `json:"vpcId"`
	PublicIp    string     `json:"publicIp"`
	Tags        []TagModel `json:"tags"`
}

type DescribeLoadBalancersResult struct {
	BlbList []AppBLBModel `json:"blbList"`
	DescribeResultMeta
}

type ListenerModel struct {
	Port int    `json:"port"`
	Type string `json:"type"`
}

type DescribeLoadBalancerDetailResult struct {
	BlbId       string          `json:"blbId"`
	Status      BLBStatus       `json:"status"`
	Description string          `json:"desc"`
	Address     string          `json:"address"`
	PublicIp    string          `json:"publicIp"`
	Cidr        string          `json:"cidr"`
	VpcName     string          `json:"vpcName"`
	SubnetCider string          `json:"subnetCider"`
	SubnetName  string          `json:"subnetName"`
	CreateTime  string          `json:"createTime"`
	ReleaseTime string          `json:"releaseTime"`
	Listener    []ListenerModel `json:"listener"`
	Tags        []TagModel      `json:"tags"`
}

type CreateAppTCPListenerArgs struct {
	ListenerPort uint16 `json:"listenerPort"`
	Scheduler    string `json:"scheduler"`
	ClientToken  string `json:"-"`
}

type CreateAppUDPListenerArgs struct {
	ListenerPort uint16 `json:"listenerPort"`
	Scheduler    string `json:"scheduler"`
	ClientToken  string `json:"-"`
}

type CreateAppHTTPListenerArgs struct {
	ClientToken           string `json:"-"`
	ListenerPort          uint16 `json:"listenerPort"`
	Scheduler             string `json:"scheduler"`
	KeepSession           bool   `json:"keepSession,omitempty"`
	KeepSessionType       string `json:"keepSessionType,omitempty"`
	KeepSessionTimeout    int    `json:"keepSessionTimeout,omitempty"`
	KeepSessionCookieName string `json:"keepSessionCookieName,omitempty"`
	XForwardedFor         bool   `json:"xForwardedFor,omitempty"`
	ServerTimeout         int    `json:"serverTimeout,omitempty"`
	RedirectPort          uint16 `json:"redirectPort,omitempty"`
}

type CreateAppHTTPSListenerArgs struct {
	ClientToken           string   `json:"-"`
	ListenerPort          uint16   `json:"listenerPort"`
	Scheduler             string   `json:"scheduler"`
	KeepSession           bool     `json:"keepSession,omitempty"`
	KeepSessionType       string   `json:"keepSessionType,omitempty"`
	KeepSessionTimeout    int      `json:"keepSessionTimeout,omitempty"`
	KeepSessionCookieName string   `json:"keepSessionCookieName,omitempty"`
	XForwardedFor         bool     `json:"xForwardedFor,omitempty"`
	ServerTimeout         int      `json:"serverTimeout,omitempty"`
	CertIds               []string `json:"certIds"`
	Ie6Compatible         bool     `json:"ie6Compatible,omitempty"`
	EncryptionType        string   `json:"encryptionType,omitempty"`
	EncryptionProtocols   []string `json:"encryptionProtocols,omitempty"`
	DualAuth              bool     `json:"dualAuth,omitempty"`
	ClientCertIds         []string `json:"clientCertIds,omitempty"`
}

type CreateAppSSLListenerArgs struct {
	ClientToken         string   `json:"-"`
	ListenerPort        uint16   `json:"listenerPort"`
	Scheduler           string   `json:"scheduler"`
	CertIds             []string `json:"certIds"`
	Ie6Compatible       bool     `json:"ie6Compatible,omitempty"`
	EncryptionType      string   `json:"encryptionType,omitempty"`
	EncryptionProtocols []string `json:"encryptionProtocols,omitempty"`
	DualAuth            bool     `json:"dualAuth,omitempty"`
	ClientCertIds       []string `json:"clientCertIds,omitempty"`
}

type UpdateAppListenerArgs struct {
	ClientToken  string `json:"-"`
	ListenerPort uint16 `json:"-"`
	Scheduler    string `json:"scheduler,omitempty"`
}

type UpdateAppTCPListenerArgs struct {
	UpdateAppListenerArgs
}

type UpdateAppUDPListenerArgs struct {
	UpdateAppListenerArgs
}

type UpdateAppHTTPListenerArgs struct {
	ClientToken           string `json:"-"`
	ListenerPort          uint16 `json:"-"`
	Scheduler             string `json:"scheduler"`
	KeepSession           bool   `json:"keepSession,omitempty"`
	KeepSessionType       string `json:"keepSessionType,omitempty"`
	KeepSessionTimeout    int    `json:"keepSessionTimeout,omitempty"`
	KeepSessionCookieName string `json:"keepSessionCookieName,omitempty"`
	XForwardedFor         bool   `json:"xForwardedFor,omitempty"`
	ServerTimeout         int    `json:"serverTimeout,omitempty"`
	RedirectPort          uint16 `json:"redirectPort,omitempty"`
}

type UpdateAppHTTPSListenerArgs struct {
	ClientToken           string   `json:"-"`
	ListenerPort          uint16   `json:"listenerPort"`
	Scheduler             string   `json:"scheduler"`
	KeepSession           bool     `json:"keepSession,omitempty"`
	KeepSessionType       string   `json:"keepSessionType,omitempty"`
	KeepSessionTimeout    int      `json:"keepSessionTimeout,omitempty"`
	KeepSessionCookieName string   `json:"keepSessionCookieName,omitempty"`
	XForwardedFor         bool     `json:"xForwardedFor,omitempty"`
	ServerTimeout         int      `json:"serverTimeout,omitempty"`
	CertIds               []string `json:"certIds"`
	Ie6Compatible         bool     `json:"ie6Compatible,omitempty"`
	EncryptionType        string   `json:"encryptionType,omitempty"`
	EncryptionProtocols   []string `json:"encryptionProtocols,omitempty"`
	DualAuth              bool     `json:"dualAuth,omitempty"`
	ClientCertIds         []string `json:"clientCertIds,omitempty"`
}

type UpdateAppSSLListenerArgs struct {
	ClientToken         string   `json:"-"`
	ListenerPort        uint16   `json:"-"`
	Scheduler           string   `json:"scheduler"`
	CertIds             []string `json:"certIds"`
	Ie6Compatible       bool     `json:"ie6Compatible,omitempty"`
	EncryptionType      string   `json:"encryptionType,omitempty"`
	EncryptionProtocols []string `json:"encryptionProtocols,omitempty"`
	DualAuth            bool     `json:"dualAuth,omitempty"`
	ClientCertIds       []string `json:"clientCertIds,omitempty"`
}

type AppListenerModel struct {
	Port      uint16 `json:"listenerPort"`
	Scheduler string `json:"scheduler"`
}

type AppTCPListenerModel struct {
	AppListenerModel
}

type AppUDPListenerModel struct {
	AppListenerModel
}

type AppHTTPListenerModel struct {
	ListenerPort          uint16 `json:"listenerPort"`
	Scheduler             string `json:"scheduler"`
	KeepSession           bool   `json:"keepSession"`
	KeepSessionType       string `json:"keepSessionType"`
	KeepSessionTimeout    int    `json:"keepSessionTimeout"`
	KeepSessionCookieName string `json:"keepSessionCookieName"`
	XForwardedFor         bool   `json:"xForwardedFor"`
	ServerTimeout         int    `json:"serverTimeout"`
	RedirectPort          int    `json:"redirectPort"`
}

type AppHTTPSListenerModel struct {
	ListenerPort          uint16   `json:"listenerPort"`
	Scheduler             string   `json:"scheduler"`
	KeepSession           bool     `json:"keepSession"`
	KeepSessionType       string   `json:"keepSessionType"`
	KeepSessionTimeout    int      `json:"keepSessionTimeout"`
	KeepSessionCookieName string   `json:"keepSessionCookieName"`
	XForwardedFor         bool     `json:"xForwardedFor"`
	ServerTimeout         int      `json:"serverTimeout"`
	CertIds               []string `json:"certIds"`
	Ie6Compatible         bool     `json:"ie6Compatible"`
	EncryptionType        string   `json:"encryptionType"`
	EncryptionProtocols   []string `json:"encryptionProtocols"`
	DualAuth              bool     `json:"dualAuth"`
	ClientCertIds         []string `json:"clientCertIds"`
}

type AppSSLListenerModel struct {
	ListenerPort        uint16   `json:"listenerPort"`
	Scheduler           string   `json:"scheduler"`
	CertIds             []string `json:"certIds"`
	Ie6Compatible       bool     `json:"ie6Compatible"`
	EncryptionType      string   `json:"encryptionType"`
	EncryptionProtocols []string `json:"encryptionProtocols"`
	DualAuth            bool     `json:"dualAuth"`
	ClientCertIds       []string `json:"clientCertIds"`
}

type DescribeAppListenerArgs struct {
	ListenerPort uint16
	Marker       string
	MaxKeys      int
}

type DescribeAppTCPListenersResult struct {
	ListenerList []AppTCPListenerModel `json:"listenerList"`
	DescribeResultMeta
}

type DescribeAppUDPListenersResult struct {
	ListenerList []AppUDPListenerModel `json:"listenerList"`
	DescribeResultMeta
}

type DescribeAppHTTPListenersResult struct {
	ListenerList []AppHTTPListenerModel `json:"listenerList"`
	DescribeResultMeta
}

type DescribeAppHTTPSListenersResult struct {
	ListenerList []AppHTTPSListenerModel `json:"listenerList"`
	DescribeResultMeta
}

type DescribeAppSSLListenersResult struct {
	ListenerList []AppSSLListenerModel `json:"listenerList"`
	DescribeResultMeta
}

type DeleteAppListenersArgs struct {
	ClientToken string   `json:"-"`
	PortList    []uint16 `json:"portList"`
}

type AppRule struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type AppPolicy struct {
	Description      string    `json:"desc"`
	AppServerGroupId string    `json:"appServerGroupId"`
	BackendPort      uint16    `json:"backendPort"`
	Priority         int       `json:"priority"`
	RuleList         []AppRule `json:"ruleList"`

	Id                 string `json:"id"`
	FrontendPort       int    `json:"frontendPort"`
	AppServerGroupName string `json:"appServerGroupName"`
	PortType           string `json:"portType"`
}

type CreatePolicysArgs struct {
	ClientToken  string      `json:"-"`
	ListenerPort uint16      `json:"listenerPort"`
	AppPolicyVos []AppPolicy `json:"appPolicyVos"`
}

type DescribePolicysArgs struct {
	Port    uint16
	Marker  string
	MaxKeys int
}

type DescribePolicysResult struct {
	PolicyList []AppPolicy `json:"policyList"`
	DescribeResultMeta
}

type DeletePolicysArgs struct {
	ClientToken  string   `json:"-"`
	Port         uint16   `json:"port"`
	PolicyIdList []string `json:"policyIdList"`
}
