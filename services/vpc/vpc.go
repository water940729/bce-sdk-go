package vpc

import (
	"fmt"
	"strconv"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

func (c *Client) CreateVPC(args *CreateVPCArgs) (*CreateVPCResult, error) {
	result := &CreateVPCResult{}

	err := bce.NewRequestBuilder(c).
		WithURL(getURLForVPC()).
		WithMethod(http.POST).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

func (c *Client) ListVPC(args *ListVPCArgs) (*ListVPCResult, error) {
	if args == nil {
		args = &ListVPCArgs{}
	}
	if args.IsDefault != "" && args.IsDefault != "true" && args.IsDefault != "false" {
		return nil, fmt.Errorf("The field isDefault can only be true or false.")
	}
	if args.MaxKeys < 0 || args.MaxKeys > 1000 {
		return nil, fmt.Errorf("The field maxKeys is out of range [0, 1000]")
	}

	result := &ListVPCResult{}
	builder := bce.NewRequestBuilder(c).
		WithURL(getURLForVPC()).
		WithMethod(http.GET).
		WithResult(result).
		WithQueryParamFilter("marker", args.Marker).
		WithQueryParamFilter("isDefault", args.IsDefault)
	if args.MaxKeys != 0 {
		builder.WithQueryParamFilter("maxKeys", strconv.Itoa(args.MaxKeys))
	}
	err := builder.Do()

	return result, err
}

func (c *Client) GetVPCDetail(vpcId string) (*GetVPCDetailResult, error) {
	result := &GetVPCDetailResult{}

	err := bce.NewRequestBuilder(c).
		WithURL(getURLForVPCId(vpcId)).
		WithMethod(http.GET).
		WithResult(result).
		Do()

	return result, err
}

func (c *Client) UpdateVPC(vpcId string, updateVPCArgs *UpdateVPCArgs) error {
	return bce.NewRequestBuilder(c).
		WithURL(getURLForVPCId(vpcId)).
		WithMethod(http.PUT).
		WithQueryParam("modifyAttribute", "").
		WithBody(updateVPCArgs).
		Do()
}

func (c *Client) DeleteVPC(vpcId string) error {
	return bce.NewRequestBuilder(c).
		WithURL(getURLForVPCId(vpcId)).
		WithMethod(http.DELETE).
		Do()
}
