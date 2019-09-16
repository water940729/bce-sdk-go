package vpc

import "github.com/baidubce/bce-sdk-go/util"

type (
	SubnetType           string
	NexthopType          string
	AclRuleProtocolType  string
	AclRuleDirectionType string
	AclRuleActionType    string
	AclRulePortType      string
	NatGatewaySpecType   string
	PaymentTimingType    string
	PeerConnRoleType     string
	NatStatusType        string
	PeerConnStatusType   string
	DnsStatusType        string
)

const (
	SUBNET_TYPE_BCC    SubnetType = "BCC"
	SUBNET_TYPE_BCCNAT SubnetType = "BCC_NAT"
	SUBNET_TYPE_BBC    SubnetType = "BBC"

	NEXTHOP_TYPE_CUSTOM NexthopType = "custom"
	NEXTHOP_TYPE_VPN    NexthopType = "vpn"
	NEXTHOP_TYPE_NAT    NexthopType = "nat"

	ACL_RULE_PROTOCOL_TCP  AclRuleProtocolType = "tcp"
	ACL_RULE_PROTOCOL_UDP  AclRuleProtocolType = "udp"
	ACL_RULE_PROTOCOL_ICMP AclRuleProtocolType = "icmp"

	ACL_RULE_DIRECTION_INGRESS AclRuleDirectionType = "ingress"
	ACL_RULE_DIRECTION_EGRESS  AclRuleDirectionType = "egress"

	ACL_RULE_ACTION_ALLOW AclRuleActionType = "allow"
	ACL_RULE_ACTION_DENY  AclRuleActionType = "deny"

	ACL_RULE_PORT_ALL AclRulePortType = "all"

	NAT_GATEWAY_SPEC_SMALL  NatGatewaySpecType = "small"
	NAT_GATEWAY_SPEC_MEDIUM NatGatewaySpecType = "medium"
	NAT_GATEWAY_SPEC_LARGE  NatGatewaySpecType = "large"

	PAYMENT_TIMING_PREPAID  PaymentTimingType = "Prepaid"
	PAYMENT_TIMING_POSTPAID PaymentTimingType = "Postpaid"

	PEERCONN_ROLE_INITIATOR PeerConnRoleType = "initiator"
	PEERCONN_ROLE_ACCEPTOR  PeerConnRoleType = "acceptor"

	NAT_STATUS_BUILDING     NatStatusType = "building"
	NAT_STATUS_UNCONFIGURED NatStatusType = "unconfigured"
	NAT_STATUS_CONFIGURING  NatStatusType = "configuring"
	NAT_STATUS_ACTIVE       NatStatusType = "active"
	NAT_STATUS_STOPPING     NatStatusType = "stopping"
	NAT_STATUS_DOWN         NatStatusType = "down"
	NAT_STATUS_STARTING     NatStatusType = "starting"
	NAT_STATUS_DELETING     NatStatusType = "deleting"
	NAT_STATUS_DELETED      NatStatusType = "deleted"

	PEERCONN_STATUS_CREATING       PeerConnStatusType = "creating"
	PEERCONN_STATUS_CONSULTING     PeerConnStatusType = "consulting"
	PEERCONN_STATUS_CONSULT_FAILED PeerConnStatusType = "consult_failed"
	PEERCONN_STATUS_ACTIVE         PeerConnStatusType = "active"
	PEERCONN_STATUS_DOWN           PeerConnStatusType = "down"
	PEERCONN_STATUS_STARTING       PeerConnStatusType = "starting"
	PEERCONN_STATUS_STOPPING       PeerConnStatusType = "stopping"
	PEERCONN_STATUS_DELETING       PeerConnStatusType = "deleting"
	PEERCONN_STATUS_DELETED        PeerConnStatusType = "deleted"
	PEERCONN_STATUS_EXPIRED        PeerConnStatusType = "expired"
	PEERCONN_STATUS_ERROR          PeerConnStatusType = "error"
	PEERCONN_STATUS_UPDATING       PeerConnStatusType = "updating"

	DNS_STATUS_CLOSE   DnsStatusType = "close"
	DNS_STATUS_WAIT    DnsStatusType = "wait"
	DNS_STATUS_SYNCING DnsStatusType = "syncing"
	DNS_STATUS_OPEN    DnsStatusType = "open"
	DNS_STATUS_CLOSING DnsStatusType = "closing"
)

type CreateVPCArgs struct {
	ClientToken string          `json:"-"`
	Name        string          `json:"name"`
	Description string          `json:"description,omitempty"`
	Cidr        string          `json:"cidr"`
	Tags        []util.TagModel `json:"tags,omitempty"`
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
	VPCID       string          `json:"vpcId"`
	Name        string          `json:"name"`
	Cidr        string          `json:"cidr"`
	Description string          `json:"description"`
	IsDefault   bool            `json:"isDefault"`
	Tags        []util.TagModel `json:"tags"`
}

type GetVPCDetailResult struct {
	VPC ShowVPCModel `json:"vpc"`
}

type ShowVPCModel struct {
	VPCId       string          `json:"vpcId"`
	Name        string          `json:"name"`
	Cidr        string          `json:"cidr"`
	Description string          `json:"description"`
	IsDefault   bool            `json:"isDefault"`
	Subnets     []Subnet        `json:"subnets"`
	Tags        []util.TagModel `json:"tags"`
}

type Subnet struct {
	SubnetId    string          `json:"subnetId"`
	Name        string          `json:"name"`
	ZoneName    string          `json:"zoneName"`
	Cidr        string          `json:"cidr"`
	VPCId       string          `json:"vpcId"`
	SubnetType  SubnetType      `json:"subnetType"`
	Description string          `json:"description"`
	AvailableIp int             `json:"availableIp"`
	Tags        []util.TagModel `json:"tags"`
}

type UpdateVPCArgs struct {
	ClientToken string `json:"-"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

type CreateSubnetArgs struct {
	ClientToken string          `json:"-"`
	Name        string          `json:"name"`
	ZoneName    string          `json:"zoneName"`
	Cidr        string          `json:"cidr"`
	VpcId       string          `json:"vpcId"`
	SubnetType  SubnetType      `json:"subnetType,omitempty"`
	Description string          `json:"description,omitempty"`
	Tags        []util.TagModel `json:"tags,omitempty"`
}

type CreateSubnetResult struct {
	SubnetId string `json:"subnetId"`
}

type ListSubnetArgs struct {
	Marker     string
	MaxKeys    int
	VpcId      string
	ZoneName   string
	SubnetType SubnetType
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
	RouteRuleId        string      `json:"routeRuleId"`
	RouteTableId       string      `json:"routeTableId"`
	SourceAddress      string      `json:"sourceAddress"`
	DestinationAddress string      `json:"destinationAddress"`
	NexthopId          string      `json:"nexthopId"`
	NexthopType        NexthopType `json:"nexthopType"`
	Description        string      `json:"description"`
}

type GetRouteTableResult struct {
	RouteTableId string      `json:"routeTableId"`
	VpcId        string      `json:"vpcId"`
	RouteRules   []RouteRule `json:"routeRules"`
}

type CreateRouteRuleArgs struct {
	ClientToken        string
	RouteTableId       string      `json:"routeTableId"`
	SourceAddress      string      `json:"sourceAddress"`
	DestinationAddress string      `json:"destinationAddress"`
	NexthopId          string      `json:"nexthopId"`
	NexthopType        NexthopType `json:"nexthopType"`
	Description        string      `json:"description,omitempty"`
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
	Id                   string               `json:"id"`
	SubnetId             string               `json:"subnetId"`
	Description          string               `json:"description"`
	Protocol             AclRuleProtocolType  `json:"protocol"`
	SourceIpAddress      string               `json:"sourceIpAddress"`
	DestinationIpAddress string               `json:"destinationIpAddress"`
	SourcePort           string               `json:"sourcePort"`
	DestinationPort      string               `json:"destinationPort"`
	Position             int                  `json:"position"`
	Direction            AclRuleDirectionType `json:"direction"`
	Action               AclRuleActionType    `json:"action"`
}

type CreateAclRuleArgs struct {
	AclRules []AclRuleRequest `json:"aclRules"`
}

type AclRuleRequest struct {
	SubnetId             string               `json:"subnetId"`
	Description          string               `json:"description,omitempty"`
	Protocol             AclRuleProtocolType  `json:"protocol"`
	SourceIpAddress      string               `json:"sourceIpAddress"`
	DestinationIpAddress string               `json:"destinationIpAddress"`
	SourcePort           string               `json:"sourcePort"`
	DestinationPort      string               `json:"destinationPort"`
	Position             int                  `json:"position"`
	Direction            AclRuleDirectionType `json:"direction"`
	Action               AclRuleActionType    `json:"action"`
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
	ClientToken          string              `json:"-"`
	Description          string              `json:"description,omitempty"`
	Protocol             AclRuleProtocolType `json:"protocol,omitempty"`
	SourceIpAddress      string              `json:"sourceIpAddress,omitempty"`
	DestinationIpAddress string              `json:"destinationIpAddress,omitempty"`
	SourcePort           string              `json:"sourcePort,omitempty"`
	DestinationPort      string              `json:"destinationPort,omitempty"`
	Position             int                 `json:"position,omitempty"`
	Action               AclRuleActionType   `json:"action,omitempty"`
}

type CreateNatGatewayArgs struct {
	ClientToken string             `json:"-"`
	Name        string             `json:"name"`
	VpcId       string             `json:"vpcId"`
	Spec        NatGatewaySpecType `json:"spec"`
	Eips        []string           `json:"eips,omitempty"`
	Billing     *Billing           `json:"billing"`
}

type Reservation struct {
	ReservationLength   int    `json:"reservationLength"`
	ReservationTimeUnit string `json:"reservationTimeUnit"`
}

type Billing struct {
	PaymentTiming PaymentTimingType `json:"paymentTiming,omitempty"`
	Reservation   *Reservation      `json:"reservation,omitempty"`
}

type CreateNatGatewayResult struct {
	NatId string `json:"natId"`
}

type ListNatGatewayArgs struct {
	VpcId   string
	NatId   string
	Name    string
	Ip      string
	Marker  string
	MaxKeys int
}

type ListNatGatewayResult struct {
	Nats        []NAT  `json:"nats"`
	Marker      string `json:"marker"`
	IsTruncated bool   `json:"isTruncated"`
	NextMarker  string `json:"nextMarker"`
	MaxKeys     int    `json:"maxKeys"`
}

// NAT is the result for getNatGatewayDetail api.
type NAT struct {
	Id            string        `json:"id"`
	Name          string        `json:"name"`
	VpcId         string        `json:"vpcId"`
	Spec          string        `json:"spec"`
	Status        NatStatusType `json:"status"`
	Eips          []string      `json:"eips"`
	PaymentTiming string        `json:"paymentTiming"`
	ExpiredTime   string        `json:"expiredTime"`
}

type UpdateNatGatewayArgs struct {
	ClientToken string `json:"-"`
	Name        string `json:"name"`
}

type BindEipsArgs struct {
	ClientToken string   `json:"-"`
	Eips        []string `json:"eips"`
}

type UnBindEipsArgs struct {
	ClientToken string   `json:"-"`
	Eips        []string `json:"eips"`
}

type RenewNatGatewayArgs struct {
	ClientToken string   `json:"-"`
	Billing     *Billing `json:"billing"`
}

type CreatePeerConnArgs struct {
	ClientToken     string   `json:"-"`
	BandwidthInMbps int      `json:"bandwidthInMbps"`
	Description     string   `json:"description,omitempty"`
	LocalIfName     string   `json:"localIfName,omitempty"`
	LocalVpcId      string   `json:"localVpcId"`
	PeerAccountId   string   `json:"peerAccountId,omitempty"`
	PeerVpcId       string   `json:"peerVpcId"`
	PeerRegion      string   `json:"peerRegion"`
	PeerIfName      string   `json:"peerIfName,omitempty"`
	Billing         *Billing `json:"billing"`
}

type CreatePeerConnResult struct {
	PeerConnId string `json:"peerConnId"`
}

type ListPeerConnsArgs struct {
	VpcId   string
	Marker  string
	MaxKeys int
}

type ListPeerConnsResult struct {
	PeerConns   []PeerConn `json:"peerConns"`
	Marker      string     `json:"marker"`
	IsTruncated bool       `json:"isTruncated"`
	NextMarker  string     `json:"nextMarker"`
	MaxKeys     int        `json:"maxKeys"`
}

type PeerConn struct {
	PeerConnId      string             `json:"peerConnId"`
	Role            PeerConnRoleType   `json:"role"`
	Status          PeerConnStatusType `json:"status"`
	BandwidthInMbps int                `json:"bandwidthInMbps"`
	Description     string             `json:"description"`
	LocalIfId       string             `json:"localIfId"`
	LocalIfName     string             `json:"localIfName"`
	LocalVpcId      string             `json:"localVpcId"`
	LocalRegion     string             `json:"localRegion"`
	PeerVpcId       string             `json:"peerVpcId"`
	PeerRegion      string             `json:"peerRegion"`
	PeerAccountId   string             `json:"peerAccountId"`
	PaymentTiming   string             `json:"paymentTiming"`
	DnsStatus       DnsStatusType      `json:"dnsStatus"`
	CreatedTime     string             `json:"createdTime"`
	ExpiredTime     string             `json:"expiredTime"`
}

type UpdatePeerConnArgs struct {
	LocalIfId   string `json:"localIfId"`
	Description string `json:"description,omitempty"`
	LocalIfName string `json:"localIfName,omitempty"`
}

type ResizePeerConnArgs struct {
	NewBandwidthInMbps int    `json:"newBandwidthInMbps"`
	ClientToken        string `json:"-"`
}

type RenewPeerConnArgs struct {
	Billing     *Billing `json:"billing"`
	ClientToken string   `json:"-"`
}

type PeerConnSyncDNSArgs struct {
	Role        PeerConnRoleType `json:"role"`
	ClientToken string           `json:"-"`
}
