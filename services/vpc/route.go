package vpc

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

func (c *Client) GetRouteTableDetail(routeTableId, vpcId string) (*GetRouteTableResult, error) {
	if routeTableId == "" && vpcId == "" {
		return nil, fmt.Errorf("The routeTableId and vpcId cannot be blank at the same time.")
	}

	result := &GetRouteTableResult{}
	err := bce.NewRequestBuilder(c).
		WithURL(getURLForRouteTable()).
		WithMethod(http.GET).
		WithQueryParamFilter("routeTableId", routeTableId).
		WithQueryParamFilter("vpcId", vpcId).
		WithResult(result).
		Do()

	return result, err
}

func (c *Client) CreateRouteRule(args *CreateRouteRuleArgs) (*CreateRouteRuleResult, error) {
	if args == nil {
		return nil, fmt.Errorf("CreateRouteRuleArgs cannot be nil.")
	}

	result := &CreateRouteRuleResult{}
	err := bce.NewRequestBuilder(c).
		WithURL(getURLForRouteRule()).
		WithMethod(http.POST).
		WithBody(args).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithResult(result).
		Do()

	return result, err
}

func (c *Client) DeleteRouteRule(routeRuleId, clientToken string) error {
	return bce.NewRequestBuilder(c).
		WithURL(getURLForRouteRuleId(routeRuleId)).
		WithMethod(http.DELETE).
		WithQueryParamFilter("clientToken", clientToken).
		Do()
}
