package cert

import "github.com/baidubce/bce-sdk-go/bce"

const (
	URI_PREFIX       = bce.URI_PREFIX + "v1"
	REQUEST_CERT_URL = "/certificate"
)

// Client of Cert service is a kind of BceClient, so derived from BceClient
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

func getCertUri() string {
	return URI_PREFIX + REQUEST_CERT_URL
}

func getCertUriWithId(id string) string {
	return URI_PREFIX + REQUEST_CERT_URL + "/" + id
}
