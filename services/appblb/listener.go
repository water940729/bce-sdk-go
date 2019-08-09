package appblb

import (
	"fmt"
	"strconv"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

func (c *Client) CreateAppTCPListener(blbId string, args *CreateAppTCPListenerArgs) error {
	if args == nil {
		return fmt.Errorf("unset args")
	}

	if args.ListenerPort == 0 {
		return fmt.Errorf("unsupport listener port")
	}

	if len(args.Scheduler) == 0 {
		return fmt.Errorf("unset scheduler")
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getAppTCPListenerUri(blbId)).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		Do()
}

func (c *Client) CreateAppUDPListener(blbId string, args *CreateAppUDPListenerArgs) error {
	if args == nil {
		return fmt.Errorf("unset args")
	}

	if args.ListenerPort == 0 {
		return fmt.Errorf("unsupport listener port")
	}

	if len(args.Scheduler) == 0 {
		return fmt.Errorf("unset scheduler")
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getAppUDPListenerUri(blbId)).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		Do()
}

func (c *Client) CreateAppHTTPListener(blbId string, args *CreateAppHTTPListenerArgs) error {
	if args == nil {
		return fmt.Errorf("unset args")
	}

	if args.ListenerPort == 0 {
		return fmt.Errorf("unsupport listener port")
	}

	if len(args.Scheduler) == 0 {
		return fmt.Errorf("unset scheduler")
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getAppHTTPListenerUri(blbId)).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		Do()
}

func (c *Client) CreateAppHTTPSListener(blbId string, args *CreateAppHTTPSListenerArgs) error {
	if args == nil {
		return fmt.Errorf("unset args")
	}

	if args.ListenerPort == 0 {
		return fmt.Errorf("unsupport listener port")
	}

	if len(args.Scheduler) == 0 {
		return fmt.Errorf("unset scheduler")
	}

	if len(args.CertIds) == 0 {
		return fmt.Errorf("unset certIds")
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getAppHTTPSListenerUri(blbId)).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		Do()
}

func (c *Client) CreateAppSSLListener(blbId string, args *CreateAppSSLListenerArgs) error {
	if args == nil {
		return fmt.Errorf("unset args")
	}

	if args.ListenerPort == 0 {
		return fmt.Errorf("unsupport listener port")
	}

	if len(args.Scheduler) == 0 {
		return fmt.Errorf("unset scheduler")
	}

	if len(args.CertIds) == 0 {
		return fmt.Errorf("unset certIds")
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getAppSSLListenerUri(blbId)).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		Do()
}

func (c *Client) UpdateAppTCPListener(blbId string, args *UpdateAppTCPListenerArgs) error {
	if args == nil || args.ListenerPort == 0 {
		return fmt.Errorf("unset listener port")
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getAppTCPListenerUri(blbId)).
		WithQueryParam("listenerPort", strconv.Itoa(int(args.ListenerPort))).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		Do()
}

func (c *Client) UpdateAppUDPListener(blbId string, args *UpdateAppUDPListenerArgs) error {
	if args == nil || args.ListenerPort == 0 {
		return fmt.Errorf("unset listener port")
	}

	if len(args.Scheduler) == 0 {
		return fmt.Errorf("unset scheduler")
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getAppUDPListenerUri(blbId)).
		WithQueryParam("listenerPort", strconv.Itoa(int(args.ListenerPort))).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		Do()
}

func (c *Client) UpdateAppHTTPListener(blbId string, args *UpdateAppHTTPListenerArgs) error {
	if args == nil || args.ListenerPort == 0 {
		return fmt.Errorf("unset listener port")
	}

	if len(args.Scheduler) == 0 {
		return fmt.Errorf("unset scheduler")
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getAppHTTPListenerUri(blbId)).
		WithQueryParam("listenerPort", strconv.Itoa(int(args.ListenerPort))).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		Do()
}

func (c *Client) UpdateAppHTTPSListener(blbId string, args *UpdateAppHTTPSListenerArgs) error {
	if args == nil || args.ListenerPort == 0 {
		return fmt.Errorf("unset listener port")
	}

	if len(args.Scheduler) == 0 {
		return fmt.Errorf("unset scheduler")
	}

	if len(args.CertIds) == 0 {
		return fmt.Errorf("unset certIds")
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getAppHTTPSListenerUri(blbId)).
		WithQueryParam("listenerPort", strconv.Itoa(int(args.ListenerPort))).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		Do()
}

func (c *Client) UpdateAppSSLListener(blbId string, args *UpdateAppSSLListenerArgs) error {
	if args == nil || args.ListenerPort == 0 {
		return fmt.Errorf("unset listener port")
	}

	if len(args.Scheduler) == 0 {
		return fmt.Errorf("unset scheduler")
	}

	if len(args.CertIds) == 0 {
		return fmt.Errorf("unset certIds")
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getAppSSLListenerUri(blbId)).
		WithQueryParam("listenerPort", strconv.Itoa(int(args.ListenerPort))).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		Do()
}

func (c *Client) DescribeAppTCPListeners(blbId string, args *DescribeAppListenerArgs) (*DescribeAppTCPListenersResult, error) {
	if args == nil {
		args = &DescribeAppListenerArgs{}
	}

	if args.MaxKeys <= 0 || args.MaxKeys > 1000 {
		args.MaxKeys = 1000
	}

	result := &DescribeAppTCPListenersResult{}
	request := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getAppTCPListenerUri(blbId)).
		WithQueryParamFilter("marker", args.Marker).
		WithQueryParam("maxKeys", strconv.Itoa(args.MaxKeys)).
		WithResult(result)

	if args.ListenerPort != 0 {
		request.WithQueryParam("listenerPort", strconv.Itoa(int(args.ListenerPort)))
	}

	err := request.Do()
	return result, err
}

func (c *Client) DescribeAppUDPListeners(blbId string, args *DescribeAppListenerArgs) (*DescribeAppUDPListenersResult, error) {
	if args == nil {
		args = &DescribeAppListenerArgs{}
	}

	if args.MaxKeys <= 0 || args.MaxKeys > 1000 {
		args.MaxKeys = 1000
	}

	result := &DescribeAppUDPListenersResult{}
	request := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getAppUDPListenerUri(blbId)).
		WithQueryParamFilter("marker", args.Marker).
		WithQueryParam("maxKeys", strconv.Itoa(args.MaxKeys)).
		WithResult(result)

	if args.ListenerPort != 0 {
		request.WithQueryParam("listenerPort", strconv.Itoa(int(args.ListenerPort)))
	}

	err := request.Do()
	return result, err
}

func (c *Client) DescribeAppHTTPListeners(blbId string, args *DescribeAppListenerArgs) (*DescribeAppHTTPListenersResult, error) {
	if args == nil {
		args = &DescribeAppListenerArgs{}
	}

	if args.MaxKeys <= 0 || args.MaxKeys > 1000 {
		args.MaxKeys = 1000
	}

	result := &DescribeAppHTTPListenersResult{}
	request := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getAppHTTPListenerUri(blbId)).
		WithQueryParamFilter("marker", args.Marker).
		WithQueryParam("maxKeys", strconv.Itoa(args.MaxKeys)).
		WithResult(result)

	if args.ListenerPort != 0 {
		request.WithQueryParam("listenerPort", strconv.Itoa(int(args.ListenerPort)))
	}

	err := request.Do()
	return result, err
}

func (c *Client) DescribeAppHTTPSListeners(blbId string, args *DescribeAppListenerArgs) (*DescribeAppHTTPSListenersResult, error) {
	if args == nil {
		args = &DescribeAppListenerArgs{}
	}

	if args.MaxKeys <= 0 || args.MaxKeys > 1000 {
		args.MaxKeys = 1000
	}

	result := &DescribeAppHTTPSListenersResult{}
	request := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getAppHTTPSListenerUri(blbId)).
		WithQueryParamFilter("marker", args.Marker).
		WithQueryParam("maxKeys", strconv.Itoa(args.MaxKeys)).
		WithResult(result)

	if args.ListenerPort != 0 {
		request.WithQueryParam("listenerPort", strconv.Itoa(int(args.ListenerPort)))
	}

	err := request.Do()
	return result, err
}

func (c *Client) DescribeAppSSLListeners(blbId string, args *DescribeAppListenerArgs) (*DescribeAppSSLListenersResult, error) {
	if args == nil {
		args = &DescribeAppListenerArgs{}
	}

	if args.MaxKeys <= 0 || args.MaxKeys > 1000 {
		args.MaxKeys = 1000
	}

	result := &DescribeAppSSLListenersResult{}
	request := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getAppSSLListenerUri(blbId)).
		WithQueryParamFilter("marker", args.Marker).
		WithQueryParam("maxKeys", strconv.Itoa(args.MaxKeys)).
		WithResult(result)

	if args.ListenerPort != 0 {
		request.WithQueryParam("listenerPort", strconv.Itoa(int(args.ListenerPort)))
	}

	err := request.Do()
	return result, err
}

func (c *Client) DeleteAppListeners(blbId string, args *DeleteAppListenersArgs) error {
	if args == nil {
		return fmt.Errorf("unset args")
	}

	if len(args.PortList) == 0 {
		return fmt.Errorf("unset port list")
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getAppListenerUri(blbId)).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithQueryParam("batchdelete", "").
		WithBody(args).
		Do()
}

func (c *Client) CreatePolicys(blbId string, args *CreatePolicysArgs) error {
	if args == nil {
		return fmt.Errorf("unset args")
	}

	if args.ListenerPort == 0 {
		return fmt.Errorf("unset listen port")
	}

	if len(args.AppPolicyVos) == 0 {
		return fmt.Errorf("unset App Policy")
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getPolicysUrl(blbId)).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		Do()
}

func (c *Client) DescribePolicys(blbId string, args *DescribePolicysArgs) (*DescribePolicysResult, error) {
	if args == nil {
		return nil, fmt.Errorf("unset args")
	}

	if args.Port == 0 {
		return nil, fmt.Errorf("unset port")
	}

	if args.MaxKeys > 1000 || args.MaxKeys <= 0 {
		args.MaxKeys = 1000
	}

	result := &DescribePolicysResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getPolicysUrl(blbId)).
		WithQueryParam("port", strconv.Itoa(int(args.Port))).
		WithQueryParamFilter("marker", args.Marker).
		WithQueryParamFilter("maxKeys", strconv.Itoa(args.MaxKeys)).
		WithResult(result).
		Do()

	return result, err
}

func (c *Client) DeletePolicys(blbId string, args *DeletePolicysArgs) error {
	if args == nil {
		return fmt.Errorf("unset args")
	}

	if args.Port == 0 {
		return fmt.Errorf("unset port")
	}

	if len(args.PolicyIdList) == 0 {
		return fmt.Errorf("unset policy id list")
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getPolicysUrl(blbId)).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithQueryParam("batchdelete", "").
		WithBody(args).
		Do()
}
