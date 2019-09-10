package cert

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

func (c *Client) CreateCert(args *CreateCertArgs) (*CreateCertResult, error) {
	if args == nil {
		return nil, fmt.Errorf("unset args")
	}

	if args.CertName == "" {
		return nil, fmt.Errorf("unset CertName")
	}

	if args.CertServerData == "" {
		return nil, fmt.Errorf("unset CertServerData")
	}

	if args.CertPrivateData == "" {
		return nil, fmt.Errorf("unset CertPrivateData")
	}

	result := &CreateCertResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCertUri()).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

func (c *Client) UpdateCertName(id string, args *UpdateCertNameArgs) error {
	if args == nil || args.CertName == "" {
		return fmt.Errorf("unset CertName")
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getCertUriWithId(id)).
		WithQueryParam("certName", "").
		WithBody(args).
		Do()
}

func (c *Client) ListCerts() (*ListCertResult, error) {
	result := &ListCertResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getCertUri()).
		WithResult(result).
		Do()

	return result, err
}

func (c *Client) GetCertMeta(id string) (*CertificateMeta, error) {
	result := &CertificateMeta{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getCertUriWithId(id)).
		WithResult(result).
		Do()

	return result, err
}

func (c *Client) DeleteCert(id string) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getCertUriWithId(id)).
		Do()
}

func (c *Client) UpdateCertData(id string, args *UpdateCertDataArgs) error {
	if args == nil {
		return fmt.Errorf("unset args")
	}

	if args.CertName == "" {
		return fmt.Errorf("unset CertName")
	}

	if args.CertServerData == "" {
		return fmt.Errorf("unset CertServerData")
	}

	if args.CertPrivateData == "" {
		return fmt.Errorf("unset CertPrivateData")
	}

	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getCertUriWithId(id)).
		WithQueryParam("certData", "").
		WithBody(args).
		Do()

	return err
}
