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

// Upload part request definition
package model

import (
    "fmt"

    "baidubce/http"
    "baidubce/common"
)

type UploadPartRequest struct {
    ObjectRequest
    uploadId      string
    partNumber    int
    contentMD5    string
    contentSha256 string
    content       *http.BodyStream
}

func (req *UploadPartRequest) SetUploadId(val string) { req.uploadId = val }
func (req *UploadPartRequest) SetPartNumber(val int) { req.partNumber = val }
func (req *UploadPartRequest) SetContentMD5(val string) { req.contentMD5 = val }
func (req *UploadPartRequest) SetContentSha256(val string) { req.contentSha256 = val }
func (req *UploadPartRequest) SetContent(val *http.BodyStream) { req.content = val }

func (req *UploadPartRequest) BuildHttpRequest() *http.Request {
    httpReq := req.ObjectRequest.BuildHttpRequest()
    httpReq.SetMethod(http.PUT)

    httpReq.AddParam("uploadId", req.uploadId)
    httpReq.AddParam("partNumber", fmt.Sprintf("%d", req.partNumber))

    if req.content == nil {
        req.SetClientError(common.NewBceClientError("no content for upload part"))
        return nil
    }
    httpReq.SetBody(req.content)

    if len(req.contentMD5) != 0 {
        httpReq.AddParam(http.CONTENT_MD5, req.contentMD5)
    }
    if len(req.contentSha256) != 0 {
        httpReq.AddParam(http.BCE_CONTENT_SHA256, req.contentSha256)
    }

    return httpReq
}

func NewUploadPartRequest(bucket, object string, body *http.BodyStream) *UploadPartRequest {
    return &UploadPartRequest{
            ObjectRequest: ObjectRequest{bucketName: bucket, objectName: object}, content: body}
}

