package vpc

import (
	"fmt"
	"strconv"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

func (c *Client) CreateNatGateway(args *CreateNatGatewayArgs) (*CreateNatGatewayResult, error) {
	if args == nil {
		return nil, fmt.Errorf("The createNatGatewayArgs cannot be nil.")
	}

	result := &CreateNatGatewayResult{}
	err := bce.NewRequestBuilder(c).
		WithURL(getURLForNat()).
		WithMethod(http.POST).
		WithBody(args).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithResult(result).
		Do()

	return result, err
}

func (c *Client) ListNatGateway(args *ListNatGatewayArgs) (*ListNatGatewayResult, error) {
	if args == nil {
		return nil, fmt.Errorf("The listNatGatewayArgs cannot be nil.")
	}
	if args.MaxKeys == 0 {
		args.MaxKeys = 1000
	}

	result := &ListNatGatewayResult{}
	err := bce.NewRequestBuilder(c).
		WithURL(getURLForNat()).
		WithMethod(http.GET).
		WithQueryParam("vpcId", args.VpcId).
		WithQueryParamFilter("natId", args.NatId).
		WithQueryParamFilter("name", args.Name).
		WithQueryParamFilter("ip", args.Ip).
		WithQueryParamFilter("marker", args.Marker).
		WithQueryParamFilter("maxKeys", strconv.Itoa(args.MaxKeys)).
		WithResult(result).
		Do()

	return result, err
}

func (c *Client) GetNatGatewayDetail(natId string) (*NAT, error) {
	result := &NAT{}
	err := bce.NewRequestBuilder(c).
		WithURL(getURLForNatId(natId)).
		WithMethod(http.GET).
		WithResult(result).
		Do()

	return result, err
}

func (c *Client) UpdateNatGateway(natId string, args *UpdateNatGatewayArgs) error {
	if args == nil {
		return fmt.Errorf("The updateNatGatewayArgs cannot be nil.")
	}

	return bce.NewRequestBuilder(c).
		WithURL(getURLForNatId(natId)).
		WithMethod(http.PUT).
		WithBody(args).
		WithQueryParamFilter("clientToken", args.ClientToken).
		Do()
}

func (c *Client) BindEips(natId string, args *BindEipsArgs) error {
	if args == nil {
		return fmt.Errorf("The bindEipArgs cannot be nil.")
	}

	return bce.NewRequestBuilder(c).
		WithURL(getURLForNatId(natId)).
		WithMethod(http.PUT).
		WithBody(args).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithQueryParam("bind", "").
		Do()
}

func (c *Client) UnBindEips(natId string, args *UnBindEipsArgs) error {
	if args == nil {
		return fmt.Errorf("The unBindEipArgs cannot be nil.")
	}

	return bce.NewRequestBuilder(c).
		WithURL(getURLForNatId(natId)).
		WithMethod(http.PUT).
		WithBody(args).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithQueryParam("unbind", "").
		Do()
}

func (c *Client) DeleteNatGateway(natId, clientToken string) error {
	return bce.NewRequestBuilder(c).
		WithURL(getURLForNatId(natId)).
		WithMethod(http.DELETE).
		WithQueryParamFilter("clientToken", clientToken).
		Do()
}

func (c *Client) RenewNatGateway(natId string, args *RenewNatGatewayArgs) error {
	if args == nil {
		return fmt.Errorf("The renewNatGatewayArgs cannot be nil.")
	}

	return bce.NewRequestBuilder(c).
		WithURL(getURLForNatId(natId)).
		WithMethod(http.PUT).
		WithBody(args).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithQueryParam("purchaseReserved", "").
		Do()
}
