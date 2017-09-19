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

// Put bucket acl request definition
package model

import (
    "encoding/json"

    "baidubce/http"
    "baidubce/common"
)

type PutBucketAclRequest struct {
    BucketRequest
    acl interface{} // `acl string` | body stream | *BucketAclInput
}

func (req *PutBucketAclRequest) SetAcl(val interface{}) {
    req.acl = val
}

func (req *PutBucketAclRequest) BuildHttpRequest() *http.Request {
    httpReq := req.BucketRequest.BuildHttpRequest()
    httpReq.SetMethod(http.PUT)
    httpReq.AddParam("acl", "")

    switch val := req.acl.(type) {
    case string:
        httpReq.AddHeader(http.BCE_ACL, val)
    case *http.BodyStream:
        httpReq.AddHeader(http.CONTENT_TYPE, "application/json; charset=utf-8")
        httpReq.SetBody(val)
    case *BucketAclInput:
        jsonBytes, err := json.Marshal(val)
        if err != nil {
            req.SetClientError(common.NewBceClientError("acl json encode fail: " + err.Error()))
            return nil
        }
        httpReq.AddHeader(http.CONTENT_TYPE, "application/json; charset=utf-8")
        httpReq.SetBody(http.NewBodyStreamFromBytes(jsonBytes))
    default:
        req.SetClientError(common.NewBceClientError("invalid acl parameter type"))
        return nil
    }
    return httpReq
}

func NewPutBucketAclRequest(name string, acl interface{}) *PutBucketAclRequest {
    return &PutBucketAclRequest{BucketRequest{bucketName: name}, acl}
}

