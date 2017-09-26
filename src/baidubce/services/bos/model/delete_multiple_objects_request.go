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

// delete_multiple_objects_request.go - the delete multiple objects request definition

package model

import (
    "encoding/json"

    "baidubce/bce"
    "baidubce/http"
)

type DeleteMultipleObjectsRequest struct {
    BucketRequest
    body interface{}  // json string or *BodyStream
}

func (req *DeleteMultipleObjectsRequest) BuildHttpRequest() *http.Request {
    httpReq := req.BucketRequest.BuildHttpRequest()
    httpReq.SetMethod(http.POST)
    httpReq.AddParam("delete", "")
    httpReq.AddHeader(http.CONTENT_TYPE, "application/json; charset=utf-8")

    switch val := req.body.(type) {
    case string:
        httpReq.SetBody(http.NewBodyStreamFromString(val))
    case *http.BodyStream:
        httpReq.SetBody(val)
    case *DeleteMultipleObjectsInput:
        jsonBytes, err := json.Marshal(val)
        if err != nil {
            req.SetClientError(
                bce.NewBceClientError("multi delete json encode fail: " + err.Error()))
            return nil
        }
        httpReq.SetBody(http.NewBodyStreamFromBytes(jsonBytes))
    default:
        req.SetClientError(bce.NewBceClientError("invalid delete multi objects body type"))
        return nil
    }
    return httpReq
}

func NewDeleteMultipleObjectsRequest(
        bucket string, body interface{}) *DeleteMultipleObjectsRequest {
    return &DeleteMultipleObjectsRequest{BucketRequest{bucketName: bucket}, body}
}

