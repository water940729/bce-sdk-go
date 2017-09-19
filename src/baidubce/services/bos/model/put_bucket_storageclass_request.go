/*
 * Copyright 2014 Baidu, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
 * the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
 * an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations under the License.
 */

// Put bucket storageClass request definition
package model

import (
    "encoding/json"

    "baidubce/http"
    "baidubce/common"
)

type PutBucketStorageClassRequest struct {
    BucketRequest
    storageClass interface{} // storage string | stream
}

func (req *PutBucketStorageClassRequest) SetStorageClass(val interface{}) {
    req.storageClass = val
}

func (req *PutBucketStorageClassRequest) BuildHttpRequest() *http.Request {
    httpReq := req.BucketRequest.BuildHttpRequest()
    httpReq.SetMethod(http.PUT)
    httpReq.AddParam("storageClass", "")
    httpReq.AddHeader(http.CONTENT_TYPE, "application/json; charset=utf-8")

    httpReq.AddHeader(http.CONTENT_TYPE, "application/json; charset=utf-8")
    var body *http.BodyStream
    switch val := req.storageClass.(type) {
    case string:
        body = http.NewBodyStreamFromString(val)
    case *http.BodyStream:
        body = val
    case *BucketStorageClassInOut:
        jsonBytes, err := json.Marshal(val)
        if err != nil {
            req.SetClientError(
                common.NewBceClientError("storage class json encode fail: " + err.Error()))
            return nil
        }
        body = http.NewBodyStreamFromBytes(jsonBytes)
    default:
        req.SetClientError(common.NewBceClientError("invalid storage class param type"))
        return nil
    }
    httpReq.SetBody(body)
    return httpReq
}

func NewPutBucketStorageClassRequest(name string, storage interface{}) *PutBucketStorageClassRequest {
    return &PutBucketStorageClassRequest{BucketRequest{bucketName: name}, storage}
}

