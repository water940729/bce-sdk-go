package api

import (
	"encoding/json"
	"strconv"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

func CreateAutoSnapshotPolicy(cli bce.Client, args *CreateASPArgs) (*CreateASPResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getASPUri())
	req.SetMethod(http.POST)

	if args != nil && len(args.ClientToken) != 0 {
		req.SetParam("clientToken", args.ClientToken)
	}

	jsonBytes, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return nil, err
	}
	req.SetBody(body)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &CreateASPResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	return jsonBody, nil
}

func AttachAutoSnapshotPolicy(cli bce.Client, aspId string, args *AttachASPArgs) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getASPUriWithId(aspId))
	req.SetMethod(http.PUT)

	req.SetParam("attach", "")

	jsonBytes, err := json.Marshal(args)
	if err != nil {
		return err
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return err
	}
	req.SetBody(body)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return err
	}
	if resp.IsFail() {
		return resp.ServiceError()
	}

	defer func() { resp.Body().Close() }()
	return nil
}

func DetachAutoSnapshotPolicy(cli bce.Client, aspId string, args *DetachASPArgs) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getASPUriWithId(aspId))
	req.SetMethod(http.PUT)

	req.SetParam("detach", "")

	jsonBytes, err := json.Marshal(args)
	if err != nil {
		return err
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return err
	}
	req.SetBody(body)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return err
	}
	if resp.IsFail() {
		return resp.ServiceError()
	}

	defer func() { resp.Body().Close() }()
	return nil
}

func DeleteAutoSnapshotPolicy(cli bce.Client, aspId string) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getASPUriWithId(aspId))
	req.SetMethod(http.DELETE)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return err
	}
	if resp.IsFail() {
		return resp.ServiceError()
	}

	defer func() { resp.Body().Close() }()
	return nil
}

func ListAutoSnapshotPolicy(cli bce.Client, queryArgs *ListASPArgs) (*ListASPResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getASPUri())
	req.SetMethod(http.GET)

	if queryArgs != nil {
		if len(queryArgs.Marker) != 0 {
			req.SetParam("marker", queryArgs.Marker)
		}
		if queryArgs.MaxKeys != 0 {
			req.SetParam("maxKeys", strconv.Itoa(queryArgs.MaxKeys))
		}
		if len(queryArgs.AspName) != 0 {
			req.SetParam("aspName", queryArgs.AspName)
		}
		if len(queryArgs.VolumeName) != 0 {
			req.SetParam("volumeName", queryArgs.VolumeName)
		}
	}

	if queryArgs == nil || queryArgs.MaxKeys == 0 {
		req.SetParam("maxKeys", "1000")
	}

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &ListASPResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	return jsonBody, nil
}

func GetAutoSnapshotPolicyDetail(cli bce.Client, aspId string) (*GetASPDetailResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getASPUriWithId(aspId))
	req.SetMethod(http.GET)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &GetASPDetailResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}

	return jsonBody, nil
}
