package vpc

import (
	"github.com/baidubce/bce-sdk-go/bce"
)

const (
	URI_PREFIX = bce.URI_PREFIX + "v1"

	REQUEST_VPC_URL    = "/vpc"
	REQUEST_SUBNET_URL = "/subnet"
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
