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

// Abort multipart upload request definition
package model

import (
    "baidubce/http"
)

type AbortMultipartUploadRequest struct {
    ObjectRequest
    uploadId string
}

func (req *AbortMultipartUploadRequest) BuildHttpRequest() *http.Request {
    httpReq := req.ObjectRequest.BuildHttpRequest()
    httpReq.SetMethod(http.DELETE)
    httpReq.AddParam("uploadId", req.uploadId)
    return httpReq
}

func NewAbortMultipartUploadRequest(bucket, object, uploadId string) *AbortMultipartUploadRequest {
    return &AbortMultipartUploadRequest{
        ObjectRequest{bucketName: bucket, objectName: object}, uploadId}
}

