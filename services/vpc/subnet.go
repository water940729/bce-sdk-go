package vpc

import (
	"fmt"
	"strconv"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

func (c *Client) CreateSubnet(args *CreateSubnetArgs) (*CreateSubnetResult, error) {
	if args == nil {
		return nil, fmt.Errorf("CreateSubnetArgs cannot be nil.")
	}

	result := &CreateSubnetResult{}
	err := bce.NewRequestBuilder(c).
		WithURL(getURLForSubnet()).
		WithMethod(http.POST).
		WithBody(args).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithResult(result).
		Do()

	return result, err
}

func (c *Client) ListSubnets(args *ListSubnetArgs) (*ListSubnetResult, error) {
	if args == nil {
		args = &ListSubnetArgs{}
	}
	if args.MaxKeys < 0 || args.MaxKeys > 1000 {
		return nil, fmt.Errorf("The field maxKeys is out of range [0, 1000]")
	} else if args.MaxKeys == 0 {
		args.MaxKeys = 1000
	}

	result := &ListSubnetResult{}
	err := bce.NewRequestBuilder(c).
		WithURL(getURLForSubnet()).
		WithMethod(http.GET).
		WithQueryParamFilter("marker", args.Marker).
		WithQueryParamFilter("vpcId", args.VpcId).
		WithQueryParamFilter("zoneName", args.ZoneName).
		WithQueryParamFilter("subnetType", args.SubnetType).
		WithQueryParamFilter("maxKeys", strconv.Itoa(args.MaxKeys)).
		WithResult(result).
		Do()

	return result, err
}

func (c *Client) GetSubnetDetail(subnetId string) (*GetSubnetDetailResult, error) {
	if subnetId == "" {
		return nil, fmt.Errorf("The subnetId cannot be blank.")
	}

	result := &GetSubnetDetailResult{}
	err := bce.NewRequestBuilder(c).
		WithURL(getURLForSubnetId(subnetId)).
		WithMethod(http.GET).
		WithResult(result).
		Do()

	return result, err
}

func (c *Client) UpdateSubnet(subnetId string, args *UpdateSubnetArgs) error {
	if subnetId == "" {
		return fmt.Errorf("The subnetId cannot be blank.")
	}
	if args == nil {
		return fmt.Errorf("The UpdateSubnetArgs cannot be nil.")
	}

	return bce.NewRequestBuilder(c).
		WithURL(getURLForSubnetId(subnetId)).
		WithMethod(http.PUT).
		WithBody(args).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithQueryParam("modifyAttribute", "").
		Do()
}

func (c *Client) DeleteSubnet(subnetId string, clientToken string) error {
	if subnetId == "" {
		return fmt.Errorf("The subnetId cannot be blank.")
	}

	return bce.NewRequestBuilder(c).
		WithURL(getURLForSubnetId(subnetId)).
		WithMethod(http.DELETE).
		WithQueryParamFilter("clientToken", clientToken).
		Do()
}
