package api

import (
	"encoding/json"
	"strconv"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

func CreateCDSVolume(cli bce.Client, args *CreateCDSVolumeArgs) (*CreateCDSVolumeResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getVolumeUri())
	req.SetMethod(http.POST)

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

	jsonBody := &CreateCDSVolumeResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	return jsonBody, nil
}

func ListCDSVolume(cli bce.Client, queryArgs *ListCDSVolumeArgs) (*ListCDSVolumeResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getVolumeUri())
	req.SetMethod(http.GET)

	if queryArgs != nil {
		if len(queryArgs.InstanceId) != 0 {
			req.SetParam("instanceId", queryArgs.InstanceId)
		}
		if len(queryArgs.ZoneName) != 0 {
			req.SetParam("zoneName", queryArgs.ZoneName)
		}
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

	jsonBody := &ListCDSVolumeResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	return jsonBody, nil
}

func GetCDSVolumeDetail(cli bce.Client, volumeId string) (*GetVolumeDetailResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getVolumeUriWithId(volumeId))
	req.SetMethod(http.GET)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &GetVolumeDetailResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	return jsonBody, nil
}

func AttachCDSVolume(cli bce.Client, volumeId string, args *AttachVolumeArgs) (*AttachVolumeResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getVolumeUriWithId(volumeId))
	req.SetMethod(http.PUT)

	req.SetParam("attach", "")

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

	jsonBody := &AttachVolumeResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	return jsonBody, nil
}

func DetachCDSVolume(cli bce.Client, volumeId string, args *DetachVolumeArgs) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getVolumeUriWithId(volumeId))
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

func DeleteCDSVolume(cli bce.Client, volumeId string) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getVolumeUriWithId(volumeId))
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

func DeleteCDSVolumeNew(cli bce.Client, volumeId string, args *DeleteCSDVolumeArgs) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getVolumeUriWithId(volumeId))
	req.SetMethod(http.POST)

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

func ResizeCDSVolume(cli bce.Client, volumeId string, args *ResizeCSDVolumeArgs) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getVolumeUriWithId(volumeId))
	req.SetMethod(http.PUT)

	req.SetParam("resize", "")

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

func RollbackCDSVolume(cli bce.Client, volumeId string, args *RollbackCSDVolumeArgs) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getVolumeUriWithId(volumeId))
	req.SetMethod(http.PUT)

	req.SetParam("rollback", "")

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

func PurchaseReservedCDSVolume(cli bce.Client, volumeId string, args *PurchaseReservedCSDVolumeArgs) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getVolumeUriWithId(volumeId))
	req.SetMethod(http.PUT)

	req.SetParam("purchaseReserved", "")

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

func RenameCDSVolume(cli bce.Client, volumeId string, args *RenameCSDVolumeArgs) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getVolumeUriWithId(volumeId))
	req.SetMethod(http.PUT)

	req.SetParam("rename", "")

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

func ModifyCDSVolume(cli bce.Client, volumeId string, args *ModifyCSDVolumeArgs) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getVolumeUriWithId(volumeId))
	req.SetMethod(http.PUT)

	req.SetParam("modify", "")

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

func ModifyChargeTypeCDSVolume(cli bce.Client, volumeId string, args *ModifyChargeTypeCSDVolumeArgs) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getVolumeUriWithId(volumeId))
	req.SetMethod(http.PUT)

	req.SetParam("modify", "")

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
