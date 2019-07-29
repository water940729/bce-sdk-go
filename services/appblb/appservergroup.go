package appblb

import (
	"fmt"
	"strconv"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

func (c *Client) CreateAppServerGroup(blbId string, args *CreateAppServerGroupArgs) (*CreateAppServerGroupResult, error) {
	if args == nil {
		args = &CreateAppServerGroupArgs{}
	}

	result := &CreateAppServerGroupResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getAppServerGroupUri(blbId)).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

func (c *Client) UpdateAppServerGroup(blbId string, args *UpdateAppServerGroupArgs) error {
	if args == nil || len(args.SgId) == 0 {
		return fmt.Errorf("unset server group id")
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getAppServerGroupUri(blbId)).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		Do()
}

func (c *Client) DescribeAppServerGroup(blbId string, args *DescribeAppServerGroupArgs) (*DescribeAppServerGroupResult, error) {
	if args == nil {
		args = &DescribeAppServerGroupArgs{}
	}

	if args.MaxKeys > 1000 || args.MaxKeys <= 0 {
		args.MaxKeys = 1000
	}

	result := &DescribeAppServerGroupResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getAppServerGroupUri(blbId)).
		WithQueryParamFilter("name", args.Name).
		WithQueryParamFilter("exactlyMatch", args.ExactlyMatch).
		WithQueryParamFilter("marker", args.Marker).
		WithQueryParamFilter("maxKeys", strconv.Itoa(args.MaxKeys)).
		WithResult(result).
		Do()

	return result, err
}

func (c *Client) DeleteAppServerGroup(blbId string, args *DeleteAppServerGroupArgs) error {
	if args == nil || len(args.SgId) == 0 {
		return fmt.Errorf("unset server group id")
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getAppServerGroupUri(blbId)).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithQueryParam("delete", "").
		WithBody(args).
		Do()
}

func (c *Client) CreateAppServerGroupPort(blbId string, args *CreateAppServerGroupPortArgs) (*CreateAppServerGroupPortResult, error) {
	if args == nil || len(args.SgId) == 0 {
		return nil, fmt.Errorf("unset server group id")
	}

	if len(args.Type) == 0 {
		return nil, fmt.Errorf("unset type")
	}

	if args.Type == "UDP" && len(args.UdpHealthCheckString) == 0 {
		return nil, fmt.Errorf("unset udpHealthCheckString")
	}

	result := &CreateAppServerGroupPortResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getAppServerGroupPortUri(blbId)).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

func (c *Client) UpdateAppServerGroupPort(blbId string, args *UpdateAppServerGroupPortArgs) error {
	if args == nil || len(args.SgId) == 0 {
		return fmt.Errorf("unset server group id")
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getAppServerGroupPortUri(blbId)).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		Do()
}

func (c *Client) DeleteAppServerGroupPort(blbId string, args *DeleteAppServerGroupPortArgs) error {
	if args == nil || len(args.SgId) == 0 {
		return fmt.Errorf("unset server group id")
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getAppServerGroupPortUri(blbId)).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithQueryParam("batchdelete", "").
		WithBody(args).
		Do()
}

func (c *Client) CreateBlbRs(blbId string, args *CreateBlbRsArgs) error {
	if args == nil || len(args.SgId) == 0 {
		return fmt.Errorf("unset server group id")
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getBlbRsUri(blbId)).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		Do()
}

func (c *Client) UpdateBlbRs(blbId string, args *UpdateBlbRsArgs) error {
	if args == nil || len(args.SgId) == 0 {
		return fmt.Errorf("unset server group id")
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getBlbRsUri(blbId)).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		Do()
}

func (c *Client) DescribeBlbRs(blbId string, args *DescribeBlbRsArgs) (*DescribeBlbRsResult, error) {
	if args == nil || len(args.SgId) == 0 {
		return nil, fmt.Errorf("unset server group id")
	}

	if args.MaxKeys > 1000 || args.MaxKeys <= 0 {
		args.MaxKeys = 1000
	}

	result := &DescribeBlbRsResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getBlbRsUri(blbId)).
		WithQueryParamFilter("marker", args.Marker).
		WithQueryParamFilter("maxKeys", strconv.Itoa(args.MaxKeys)).
		WithQueryParam("sgId", args.SgId).
		WithResult(result).
		Do()

	return result, err
}

func (c *Client) DeleteBlbRs(blbId string, args *DeleteBlbRsArgs) error {
	if args == nil || len(args.SgId) == 0 {
		return fmt.Errorf("unset server group id")
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getBlbRsUri(blbId)).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithQueryParam("batchdelete", "").
		WithBody(args).
		Do()
}

func (c *Client) DescribeRsMount(blbId, sgId string) (*DescribeRsMountResult, error) {
	if len(sgId) == 0 {
		return nil, fmt.Errorf("unset server group id")
	}

	result := &DescribeRsMountResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getBlbRsMountUri(blbId)).
		WithQueryParam("sgId", sgId).
		WithResult(result).
		Do()

	return result, err
}

func (c *Client) DescribeRsUnMount(blbId, sgId string) (*DescribeRsMountResult, error) {
	if len(sgId) == 0 {
		return nil, fmt.Errorf("unset server group id")
	}

	result := &DescribeRsMountResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getBlbRsUnMountUri(blbId)).
		WithQueryParam("sgId", sgId).
		WithResult(result).
		Do()

	return result, err
}
