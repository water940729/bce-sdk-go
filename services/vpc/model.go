package vpc

const (
	SUBNET_TYPE_BCC    = "BCC"
	SUBNET_TYPE_BCCNAT = "BCC_NAT"
	SUBNET_TYPE_BBC    = "BBC"

	NEXTHOP_TYPE_CUSTOM = "custom"
	NEXTHOP_TYPE_VPN    = "vpn"
	NEXTHOP_TYPE_NAT    = "nat"

	ACL_RULE_PROTOCOL_TCP  = "tcp"
	ACL_RULE_PROTOCOL_UDP  = "udp"
	ACL_RULE_PROTOCOL_ICMP = "icmp"

	ACL_RULE_DIRECTION_INGRESS = "ingress"
	ACL_RULE_DIRECTION_EGRESS  = "egress"

	ACL_RULE_ACTION_ALLOW = "allow"
	ACL_RULE_ACTION_DENY  = "deny"

	ACL_RULE_PORT_ALL = "all"
)

type CreateVPCArgs struct {
	ClientToken string `json:"-"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Cidr        string `json:"cidr"`
}

type CreateVPCResult struct {
	VPCID string `json:"vpcId"`
}

type ListVPCArgs struct {
	Marker  string
	MaxKeys int

	// IsDefault is a string type,
	// so we can identify if it has been setted when the value is false.
	// NOTE: it can be only true or false.
	IsDefault string
}

type ListVPCResult struct {
	Marker      string `json:"marker"`
	IsTruncated bool   `json:"isTruncated"`
	NextMarker  string `json:"nextMarker"`
	MaxKeys     int    `json:"maxKeys"`
	VPCs        []VPC  `json:"vpcs"`
}

type VPC struct {
	VPCID       string `json:"vpcId"`
	Name        string `json:"name"`
	Cidr        string `json:"cidr"`
	Description string `json:"description"`
	IsDefault   bool   `json:"isDefault"`
}

type GetVPCDetailResult struct {
	VPC ShowVPCModel `json:"vpc"`
}

type ShowVPCModel struct {
	VPCId       string   `json:"vpcId"`
	Name        string   `json:"name"`
	Cidr        string   `json:"cidr"`
	Description string   `json:"description"`
	IsDefault   bool     `json:"isDefault"`
	Subnets     []Subnet `json:"subnets"`
}

type Subnet struct {
	SubnetId    string `json:"subnetId"`
	Name        string `json:"name"`
	ZoneName    string `json:"zoneName"`
	Cidr        string `json:"cidr"`
	VPCId       string `json:"vpcId"`
	SubnetType  string `json:"subnetType"`
	Description string `json:"description"`
	AvailableIp int    `json:"availableIp"`
}

type UpdateVPCArgs struct {
	ClientToken string `json:"-"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateSubnetArgs struct {
	ClientToken string `json:"-"`
	Name        string `json:"name"`
	ZoneName    string `json:"zoneName"`
	Cidr        string `json:"cidr"`
	VpcId       string `json:"vpcId"`
	SubnetType  string `json:"subnetType,omitempty"`
	Description string `json:"description,omitempty"`
}

type CreateSubnetResult struct {
	SubnetId string `json:"subnetId"`
}

type ListSubnetArgs struct {
	Marker     string
	MaxKeys    int
	VpcId      string
	ZoneName   string
	SubnetType string
}

type ListSubnetResult struct {
	Marker      string   `json:"marker"`
	IsTruncated bool     `json:"isTruncated"`
	NextMarker  string   `json:"nextMarker"`
	MaxKeys     int      `json:"maxKeys"`
	Subnets     []Subnet `json:"subnets"`
}

type GetSubnetDetailResult struct {
	Subnet Subnet `json:"subnet"`
}

type UpdateSubnetArgs struct {
	ClientToken string `json:"-"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

type RouteRule struct {
	RouteRuleId        string `json:"routeRuleId"`
	RouteTableId       string `json:"routeTableId"`
	SourceAddress      string `json:"sourceAddress"`
	DestinationAddress string `json:"destinationAddress"`
	NexthopId          string `json:"nexthopId"`
	NexthopType        string `json:"nexthopType"`
	Description        string `json:"description"`
}

type GetRouteTableResult struct {
	RouteTableId string      `json:"routeTableId"`
	VpcId        string      `json:"vpcId"`
	RouteRules   []RouteRule `json:"routeRules"`
}

type CreateRouteRuleArgs struct {
	ClientToken        string
	RouteTableId       string `json:"routeTableId"`
	SourceAddress      string `json:"sourceAddress"`
	DestinationAddress string `json:"destinationAddress"`
	NexthopId          string `json:"nexthopId,omitempty"`
	NexthopType        string `json:"nexthopType"`
	Description        string `json:"description"`
}

type CreateRouteRuleResult struct {
	RouteRuleId string `json:"routeRuleId"`
}

type ListAclEntrysResult struct {
	VpcId     string     `json:"vpcId"`
	VpcName   string     `json:"vpcName"`
	VpcCidr   string     `json:"vpcCidr"`
	AclEntrys []AclEntry `json:"aclEntrys"`
}

type AclEntry struct {
	SubnetId   string    `json:"subnetId"`
	SubnetName string    `json:"subnetName"`
	SubnetCidr string    `json:"subnetCidr"`
	AclRules   []AclRule `json:"aclRules"`
}

type AclRule struct {
	Id                   string `json:"id"`
	SubnetId             string `json:"subnetId"`
	Description          string `json:"description"`
	Protocol             string `json:"protocol"`
	SourceIpAddress      string `json:"sourceIpAddress"`
	DestinationIpAddress string `json:"destinationIpAddress"`
	SourcePort           string `json:"sourcePort"`
	DestinationPort      string `json:"destinationPort"`
	Position             int    `json:"position"`
	Direction            string `json:"direction"`
	Action               string `json:"action"`
}

type CreateAclRuleArgs struct {
	AclRules []AclRuleRequest `json:"aclRules"`
}

type AclRuleRequest struct {
	SubnetId             string `json:"subnetId"`
	Description          string `json:"description"`
	Protocol             string `json:"protocol"`
	SourceIpAddress      string `json:"sourceIpAddress"`
	DestinationIpAddress string `json:"destinationIpAddress"`
	SourcePort           string `json:"sourcePort"`
	DestinationPort      string `json:"destinationPort"`
	Position             int    `json:"position"`
	Direction            string `json:"direction"`
	Action               string `json:"action"`
}

type ListAclRulesArgs struct {
	Marker   string `json:"marker"`
	MaxKeys  int    `json:"maxKeys"`
	SubnetId string `json:"subnetId"`
}

type ListAclRulesResult struct {
	Marker      string    `json:"marker"`
	IsTruncated bool      `json:"isTruncated"`
	NextMarker  string    `json:"nextMarker"`
	MaxKeys     int       `json:"maxKeys"`
	AclRules    []AclRule `json:"aclRules"`
}

type UpdateAclRuleArgs struct {
	ClientToken          string `json:"-"`
	Description          string `json:"description"`
	Protocol             string `json:"protocol"`
	SourceIpAddress      string `json:"sourceIpAddress"`
	DestinationIpAddress string `json:"destinationIpAddress"`
	SourcePort           string `json:"sourcePort"`
	DestinationPort      string `json:"destinationPort"`
	Position             int    `json:"position"`
	Action               string `json:"action"`
}
