package vpc

import (
	"github.com/baidubce/bce-sdk-go/bce"
)

const (
	URI_PREFIX = bce.URI_PREFIX + "v1"

	REQUEST_VPC_URL      = "/vpc"
	REQUEST_SUBNET_URL   = "/subnet"
	REQUEST_ROUTE_URL    = "/route"
	REQUEST_RULE_URL     = "/rule"
	REQUEST_ACL_URL      = "/acl"
	REQUEST_NAT_URL      = "/nat"
	REQUEST_PEERCONN_URL = "/peerconn"
)

// Client of VPC service is a kind of BceClient, so derived from BceClient
type Client struct {
	*bce.BceClient
}

func NewClient(ak, sk, endPoint string) (*Client, error) {
	client, err := bce.NewBceClientWithAkSk(ak, sk, endPoint)
	if err != nil {
		return nil, err
	}
	return &Client{client}, nil
}

func getURLForVPC() string {
	return URI_PREFIX + REQUEST_VPC_URL
}

func getURLForVPCId(vpcId string) string {
	return getURLForVPC() + "/" + vpcId
}

func getURLForSubnet() string {
	return URI_PREFIX + REQUEST_SUBNET_URL
}

func getURLForSubnetId(subnetId string) string {
	return getURLForSubnet() + "/" + subnetId
}

func getURLForRouteTable() string {
	return URI_PREFIX + REQUEST_ROUTE_URL
}

func getURLForRouteRule() string {
	return getURLForRouteTable() + REQUEST_RULE_URL
}

func getURLForRouteRuleId(routeRuleId string) string {
	return getURLForRouteRule() + "/" + routeRuleId
}

func getURLForAclEntry() string {
	return URI_PREFIX + REQUEST_ACL_URL
}

func getURLForAclRule() string {
	return URI_PREFIX + REQUEST_ACL_URL + REQUEST_RULE_URL
}

func getURLForAclRuleId(aclRuleId string) string {
	return URI_PREFIX + REQUEST_ACL_URL + REQUEST_RULE_URL + "/" + aclRuleId
}

func getURLForNat() string {
	return URI_PREFIX + REQUEST_NAT_URL
}

func getURLForNatId(natId string) string {
	return getURLForNat() + "/" + natId
}

func getURLForPeerConn() string {
	return URI_PREFIX + REQUEST_PEERCONN_URL
}

func getURLForPeerConnId(peerConnId string) string {
	return getURLForPeerConn() + "/" + peerConnId
}
