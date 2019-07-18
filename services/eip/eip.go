package eip

import (
	"fmt"
	"strconv"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

func (c *Client) CreateEip(args *CreateEipArgs) (*CreateEipResult, error) {
	if args == nil {
		return nil, fmt.Errorf("please set create eip argments")
	}

	if args.BandWidthInMbps <= 0 || args.BandWidthInMbps > 1000 {
		return nil, fmt.Errorf("unsupport bandwidthInMbps value: %d", args.BandWidthInMbps)
	}

	if args.Billing == nil {
		return nil, fmt.Errorf("please set billing")
	}

	result := &CreateEipResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getEipUri()).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

func (c *Client) ResizeEip(eip string, args *ResizeEipArgs) error {
	if args == nil {
		return fmt.Errorf("please set resize eip argments")
	}

	if args.NewBandWidthInMbps <= 0 || args.NewBandWidthInMbps > 1000 {
		return fmt.Errorf("unsupport bandwidthInMbps value: %d", args.NewBandWidthInMbps)
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getEipUriWithEip(eip)).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithQueryParam("resize", "").
		WithBody(args).
		Do()
}

func (c *Client) BindEip(eip string, args *BindEipArgs) error {
	if args == nil {
		return fmt.Errorf("please set bind eip argments")
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getEipUriWithEip(eip)).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithQueryParam("bind", "").
		WithBody(args).
		Do()
}

func (c *Client) UnBindEip(eip, clientToken string) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getEipUriWithEip(eip)).
		WithQueryParamFilter("clientToken", clientToken).
		WithQueryParam("unbind", "").
		Do()
}

func (c *Client) DeleteEip(eip, clientToken string) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getEipUriWithEip(eip)).
		WithQueryParamFilter("clientToken", clientToken).
		Do()
}

func (c *Client) ListEip(args *ListEipArgs) (*ListEipResult, error) {
	if args == nil {
		args = &ListEipArgs{}
	}

	if args.MaxKeys <= 0 || args.MaxKeys > 1000 {
		args.MaxKeys = 1000
	}

	queryMap := map[string]string{
		"eip":          args.Eip,
		"instanceType": args.InstanceType,
		"instanceId":   args.InstanceId,
		"marker":       args.Marker,
		"maxKeys":      strconv.Itoa(args.MaxKeys),
		"status":       args.Status,
	}

	result := &ListEipResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getEipUri()).
		WithQueryParams(queryMap).
		WithResult(result).
		Do()

	return result, err
}

func (c *Client) PurchaseReservedEip(eip string, args *PurchaseReservedEipArgs) error {
	if args == nil {
		return fmt.Errorf("please set purchase reserved eip argments")
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getEipUriWithEip(eip)).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		Do()
}
