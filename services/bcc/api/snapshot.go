package api

import (
	"encoding/json"
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
	"strconv"
)

func CreateSnapshot(cli bce.Client, args *CreateSnapshotArgs) (*CreateSnapshotResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getSnapshotUri())
	req.SetMethod(http.POST)

	if len(args.ClientToken) != 0 {
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

	jsonBody := &CreateSnapshotResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	return jsonBody, nil
}

func ListSnapshot(cli bce.Client, queryArgs *ListSnapshotArgs) (*ListSnapshotResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getSnapshotUri())
	req.SetMethod(http.GET)

	if queryArgs != nil {
		if len(queryArgs.Marker) != 0 {
			req.SetParam("marker", queryArgs.Marker)
		}
		if queryArgs.MaxKeys != 0 {
			req.SetParam("maxKeys", strconv.Itoa(queryArgs.MaxKeys))
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

	jsonBody := &ListSnapshotResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	return jsonBody, nil
}

func GetSnapshotDetail(cli bce.Client, snapshotId string) (*GetSnapshotDetailResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getSnapshotUriWithId(snapshotId))
	req.SetMethod(http.GET)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &GetSnapshotDetailResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}

	return jsonBody, nil
}

func DeleteSnapshot(cli bce.Client, snapshotId string) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getSnapshotUriWithId(snapshotId))
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
