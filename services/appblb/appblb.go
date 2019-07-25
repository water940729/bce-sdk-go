package appblb

import (
	"fmt"
	"strconv"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

func (c *Client) CreateLoadBalancer(args *CreateLoadBalancerArgs) (*CreateLoadBalanceResult, error) {
	if args == nil || len(args.SubnetId) == 0 {
		return nil, fmt.Errorf("unset subnet id")
	}

	if len(args.VpcId) == 0 {
		return nil, fmt.Errorf("unset vpc id")
	}

	result := &CreateLoadBalanceResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getAppBlbUri()).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

func (c *Client) UpdateLoadBalancer(blbId string, args *UpdateLoadBalancerArgs) error {
	if args == nil {
		args = &UpdateLoadBalancerArgs{}
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getAppBlbUriWithId(blbId)).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		Do()
}

func (c *Client) DescribeLoadBalancers(args *DescribeLoadBalancersArgs) (*DescribeLoadBalancersResult, error) {
	if args == nil {
		args = &DescribeLoadBalancersArgs{}
	}

	if args.MaxKeys > 1000 || args.MaxKeys <= 0 {
		args.MaxKeys = 1000
	}

	result := &DescribeLoadBalancersResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getAppBlbUri()).
		WithQueryParamFilter("address", args.Address).
		WithQueryParamFilter("name", args.Name).
		WithQueryParamFilter("blbId", args.BlbId).
		WithQueryParamFilter("bccId", args.BccId).
		WithQueryParamFilter("exactlyMatch", args.ExactlyMatch).
		WithQueryParamFilter("marker", args.Marker).
		WithQueryParamFilter("maxKeys", strconv.Itoa(args.MaxKeys)).
		WithResult(result).
		Do()

	return result, err
}

func (c *Client) DescribeLoadBalancerDetail(blbId string) (*DescribeLoadBalancerDetailResult, error) {
	result := &DescribeLoadBalancerDetailResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getAppBlbUriWithId(blbId)).
		WithResult(result).
		Do()

	return result, err
}

func (c *Client) DeleteLoadBalancer(blbId string) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getAppBlbUriWithId(blbId)).
		Do()
}
