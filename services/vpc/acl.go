package vpc

import (
	"fmt"
	"strconv"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

func (c *Client) ListAclEntrys(vpcId string) (*ListAclEntrysResult, error) {
	result := &ListAclEntrysResult{}
	err := bce.NewRequestBuilder(c).
		WithURL(getURLForAclEntry()).
		WithMethod(http.GET).
		WithQueryParam("vpcId", vpcId).
		WithResult(result).
		Do()

	return result, err
}

func (c *Client) CreateAclRule(args *CreateAclRuleArgs, clientToken string) error {
	return bce.NewRequestBuilder(c).
		WithURL(getURLForAclRule()).
		WithMethod(http.POST).
		WithBody(args).
		WithQueryParamFilter("clientToken", clientToken).
		Do()
}

func (c *Client) ListAclRules(args *ListAclRulesArgs) (*ListAclRulesResult, error) {
	if args == nil {
		return nil, fmt.Errorf("The listAclRulesArgs cannot be nil.")
	}
	if args.MaxKeys == 0 {
		args.MaxKeys = 1000
	}

	result := &ListAclRulesResult{}
	err := bce.NewRequestBuilder(c).
		WithURL(getURLForAclRule()).
		WithMethod(http.GET).
		WithQueryParamFilter("marker", args.Marker).
		WithQueryParam("maxKeys", strconv.Itoa(args.MaxKeys)).
		WithQueryParam("subnetId", args.SubnetId).
		WithResult(result).
		Do()

	return result, err
}

func (c *Client) UpdateAclRule(aclRuleId string, args *UpdateAclRuleArgs) error {
	if args == nil {
		args = &UpdateAclRuleArgs{}
	}

	return bce.NewRequestBuilder(c).
		WithURL(getURLForAclRuleId(aclRuleId)).
		WithMethod(http.PUT).
		WithBody(args).
		WithQueryParamFilter("clientToken", args.ClientToken).
		Do()
}

func (c *Client) DeleteAclRule(aclRuleId, clientToken string) error {
	return bce.NewRequestBuilder(c).
		WithURL(getURLForAclRuleId(aclRuleId)).
		WithMethod(http.DELETE).
		WithQueryParamFilter("clientToken", clientToken).
		Do()
}
