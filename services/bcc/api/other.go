/*
 * Copyright 2017 Baidu, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
 * except in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the
 * License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions
 * and limitations under the License.
 */

// other.go - the other APIs definition supported by the BCC service

// Package api defines all APIs supported by the BCC service of BCE.
package api

import (
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

// ListSpec - get specification list information of the instance
//
// PARAMS:
//     - cli: the client agent which can perform sending request
// RETURNS:
//     - *ListSpecResult: result of the specifications
//     - error: nil if success otherwise the specific error
func ListSpec(cli bce.Client) (*ListSpecResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getSpecUri())
	req.SetMethod(http.GET)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &ListSpecResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}

	return jsonBody, nil
}

// ListZone - get the available zone list in the current region
//
// PARAMS:
//     - cli: the client agent which can perform sending request
// RETURNS:
//     - *ListZoneResult: result of the available zones
//     - error: nil if success otherwise the specific error
func ListZone(cli bce.Client) (*ListZoneResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getZoneUri())
	req.SetMethod(http.GET)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &ListZoneResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}

	return jsonBody, nil
}

// ListFlavorSpec - get flavor specification list information of the instance
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//	   - args: the arguments to list flavor spec
// RETURNS:
//     - *ListBccFlavorSpecResult: result of the flavor specification
//     - error: nil if success otherwise the flavor specific error
func ListFlavorSpec(cli bce.Client, args *ListFlavorSpecRequest) (*ListBccFlavorSpecResponse, error) {
	req := &bce.BceRequest{}
	req.SetUri(getFlavorSpecUri())
	req.SetMethod(http.GET)

	if args != nil {
		if len(args.ZoneName) != 0 {
			req.SetParam("zoneName", args.ZoneName)
		}
	}

	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &ListBccFlavorSpecResponse{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}

	return jsonBody, nil
}

// GetPriceBySpec - get price by spec
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//	   - reqBody: the request body to get price by spec
// RETURNS:
//     - *BccPriceResponse: result of the BccPriceResponse
//     - error: nil if success otherwise the get price by spec error
func GetPriceBySpec(cli bce.Client, clientToken string, reqBody *bce.Body) (*BccPriceResponse, error) {
	req := &bce.BceRequest{}
	req.SetUri(getPriceBySpecUri())
	req.SetMethod(http.POST)
	req.SetBody(reqBody)

	if clientToken != "" {
		req.SetParam("clientToken", clientToken)
	}

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &BccPriceResponse{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}

	return jsonBody, nil
}
