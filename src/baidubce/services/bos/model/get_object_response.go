/*
 * Copyright 2017 Baidu, Inc.
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

// get_object_response.go - the get object response definition

package model

import (
    "strconv"

    "baidubce/http"
    "baidubce/model"
)

type GetObjectResponse struct {
    model.BceResponse
    content *http.BodyStream
}

func (resp *GetObjectResponse) CacheControl() string {
    return resp.BceResponse.GetHeader(http.CACHE_CONTROL)
}

func (resp *GetObjectResponse) ContentDisposition() string {
    return resp.BceResponse.GetHeader(http.CONTENT_DISPOSITION)
}

func (resp *GetObjectResponse) ContentLength() int64 {
    length := resp.BceResponse.GetHeader(http.CONTENT_LENGTH)
    val, _ := strconv.ParseInt(length, 10, 64)
    return val
}

func (resp *GetObjectResponse) ContentRange() string {
    return resp.BceResponse.GetHeader(http.CONTENT_RANGE)
}

func (resp *GetObjectResponse) ContentType() string {
    return resp.BceResponse.GetHeader(http.CONTENT_TYPE)
}

func (resp *GetObjectResponse) ContentMD5() string {
    return resp.BceResponse.GetHeader(http.CONTENT_MD5)
}

func (resp *GetObjectResponse) Expires() string {
    return resp.BceResponse.GetHeader(http.EXPIRES)
}

func (resp *GetObjectResponse) ETag() string {
    return resp.BceResponse.GetHeader(http.ETAG)
}

func (resp *GetObjectResponse) LastModified() string {
    return resp.BceResponse.GetHeader(http.LAST_MODIFIED)
}

func (resp *GetObjectResponse) StorageClass() string {
    return resp.BceResponse.GetHeader(http.BCE_STORAGE_CLASS)
}

func (resp *GetObjectResponse) UserMata(key string) string {
    headers := resp.BceResponse.GetHeaders()
    if val, ok := headers[key]; ok {
        return val
    }
    return ""
}

func (resp *GetObjectResponse) Content() *http.BodyStream {
    return resp.content
}

func (resp *GetObjectResponse) ParseResponse() {
    resp.BceResponse.ParseResponse()
    if resp.BceResponse.IsFail() { return }

    size, _ := strconv.ParseInt(resp.BceResponse.GetHeader(http.CONTENT_LENGTH), 10, 64)
    resp.content = &http.BodyStream{resp.BceResponse.Body(), size}
}

func NewGetObjectResponse() *GetObjectResponse {
    return &GetObjectResponse{}
}

