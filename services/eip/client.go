package eip

import "github.com/baidubce/bce-sdk-go/bce"

const (
	URI_PREFIX = bce.URI_PREFIX + "v1"

	REQUEST_EIP_URL = "/eip"
)

// Client of EIP service is a kind of BceClient, so derived from BceClient
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

func getEipUri() string {
	return URI_PREFIX + REQUEST_EIP_URL
}

func getEipUriWithEip(eip string) string {
	return URI_PREFIX + REQUEST_EIP_URL + "/" + eip
}
