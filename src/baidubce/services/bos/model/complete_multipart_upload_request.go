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

// complete_multipart_upload_request.go - the complete multipart upload request definition

package model

import (
    "encoding/json"

    "baidubce/bce"
    "baidubce/http"
)

type CompleteMultipartUploadRequest struct {
    ObjectRequest
    uploadId string
    userMeta map[string]string
    body     interface{}
}

func (req *CompleteMultipartUploadRequest) SetUploadId(val string) { req.uploadId = val }

func (req *CompleteMultipartUploadRequest) SetUserMeta(key, val string) {
    if req.userMeta == nil {
        req.userMeta = make(map[string]string)
    }
    req.userMeta[key] = val
}

func (req *CompleteMultipartUploadRequest) BuildHttpRequest() *http.Request {
    httpReq := req.ObjectRequest.BuildHttpRequest()
    httpReq.SetMethod(http.POST)
    httpReq.AddParam("uploadId", req.uploadId)

    if req.userMeta != nil {
        for k, v := range req.userMeta {
            httpReq.AddHeader(k, v)
        }
    }

    httpReq.AddHeader(http.CONTENT_TYPE, "application/json; charset=utf-8")
    switch val := req.body.(type) {
    case string:
        httpReq.SetBody(http.NewBodyStreamFromString(val))
    case *http.BodyStream:
        httpReq.SetBody(val)
    case *CompleteMultipartUploadInput:
        jsonBytes, err := json.Marshal(val)
        if err != nil {
            req.SetClientError(bce.NewBceClientError("upload part json encode fail: " + err.Error()))
            return nil
        }
        httpReq.SetBody(http.NewBodyStreamFromBytes(jsonBytes))
    default:
        req.SetClientError(bce.NewBceClientError("invalid complete multiupload parameter type"))
        return nil
    }

    return httpReq
}

func NewCompleteMultipartUploadRequest(
        bucket, object string, body interface{}) *CompleteMultipartUploadRequest {
    return &CompleteMultipartUploadRequest{
                ObjectRequest: ObjectRequest{bucketName: bucket, objectName: object}, body: body}
}

