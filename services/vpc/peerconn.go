package vpc

import (
	"fmt"
	"strconv"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

func (c *Client) CreatePeerConn(args *CreatePeerConnArgs) (*CreatePeerConnResult, error) {
	if args == nil {
		return nil, fmt.Errorf("The createPeerConnArgs cannot be nil.")
	}

	result := &CreatePeerConnResult{}
	err := bce.NewRequestBuilder(c).
		WithURL(getURLForPeerConn()).
		WithMethod(http.POST).
		WithBody(args).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithResult(result).
		Do()

	return result, err
}

func (c *Client) ListPeerConn(args *ListPeerConnsArgs) (*ListPeerConnsResult, error) {
	if args == nil {
		args = &ListPeerConnsArgs{}
	}
	if args.MaxKeys == 0 {
		args.MaxKeys = 1000
	}

	result := &ListPeerConnsResult{}
	err := bce.NewRequestBuilder(c).
		WithURL(getURLForPeerConn()).
		WithMethod(http.GET).
		WithQueryParamFilter("vpcId", args.VpcId).
		WithQueryParamFilter("marker", args.Marker).
		WithQueryParamFilter("maxKeys", strconv.Itoa(args.MaxKeys)).
		WithResult(result).
		Do()

	return result, err
}

func (c *Client) GetPeerConnDetail(peerConnId string, role PeerConnRoleType) (*PeerConn, error) {
	result := &PeerConn{}
	err := bce.NewRequestBuilder(c).
		WithURL(getURLForPeerConnId(peerConnId)).
		WithMethod(http.GET).
		WithQueryParamFilter("role", string(role)).
		WithResult(result).
		Do()

	return result, err
}

func (c *Client) UpdatePeerConn(peerConnId string, args *UpdatePeerConnArgs) error {
	if args == nil {
		return fmt.Errorf("The updatePeerConnArgs cannot be nil.")
	}

	return bce.NewRequestBuilder(c).
		WithURL(getURLForPeerConnId(peerConnId)).
		WithMethod(http.PUT).
		WithBody(args).
		Do()
}

func (c *Client) AcceptPeerConnApply(peerConnId, clientToken string) error {
	return bce.NewRequestBuilder(c).
		WithURL(getURLForPeerConnId(peerConnId)).
		WithMethod(http.PUT).
		WithQueryParam("accept", "").
		WithQueryParamFilter("clientToken", clientToken).
		Do()
}

func (c *Client) RejectPeerConnApply(peerConnId, clientToken string) error {
	return bce.NewRequestBuilder(c).
		WithURL(getURLForPeerConnId(peerConnId)).
		WithMethod(http.PUT).
		WithQueryParam("reject", "").
		WithQueryParamFilter("clientToken", clientToken).
		Do()
}

func (c *Client) DeletePeerConn(peerConnId, clientToken string) error {
	return bce.NewRequestBuilder(c).
		WithURL(getURLForPeerConnId(peerConnId)).
		WithMethod(http.DELETE).
		WithQueryParamFilter("clientToken", clientToken).
		Do()
}

func (c *Client) ResizePeerConn(peerConnId string, args *ResizePeerConnArgs) error {
	if args == nil {
		return fmt.Errorf("The resizePeerConnArgs cannot be nil.")
	}

	return bce.NewRequestBuilder(c).
		WithURL(getURLForPeerConnId(peerConnId)).
		WithMethod(http.PUT).
		WithBody(args).
		WithQueryParam("resize", "").
		WithQueryParamFilter("clientToken", args.ClientToken).
		Do()
}

func (c *Client) RenewPeerConn(peerConnId string, args *RenewPeerConnArgs) error {
	if args == nil {
		return fmt.Errorf("The renewPeerConnArgs cannot be nil.")
	}

	return bce.NewRequestBuilder(c).
		WithURL(getURLForPeerConnId(peerConnId)).
		WithMethod(http.PUT).
		WithBody(args).
		WithQueryParam("purchaseReserved", "").
		WithQueryParamFilter("clientToken", args.ClientToken).
		Do()
}

func (c *Client) OpenPeerConnSyncDNS(peerConnId string, args *PeerConnSyncDNSArgs) error {
	if args == nil {
		return fmt.Errorf("The peerConnSyncDNS cannot be nil.")
	}

	return bce.NewRequestBuilder(c).
		WithURL(getURLForPeerConnId(peerConnId)).
		WithMethod(http.PUT).
		WithQueryParam("open", "").
		WithQueryParamFilter("role", string(args.Role)).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		Do()
}

func (c *Client) ClosePeerConnSyncDNS(peerConnId string, args *PeerConnSyncDNSArgs) error {
	if args == nil {
		return fmt.Errorf("The peerConnSyncDNS cannot be nil.")
	}

	return bce.NewRequestBuilder(c).
		WithURL(getURLForPeerConnId(peerConnId)).
		WithMethod(http.PUT).
		WithQueryParam("close", "").
		WithQueryParamFilter("role", string(args.Role)).
		WithQueryParamFilter("clientToken", args.ClientToken).
		Do()
}
