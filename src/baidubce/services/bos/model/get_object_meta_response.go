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

// get_object_meta_response.go - the tet object meta response definition

package model

import (
    "strconv"

    "baidubce/model"
    "baidubce/http"
)

type GetObjectMetaResponse struct {
    model.BceResponse
}

func (resp *GetObjectMetaResponse) CacheControl() string {
    return resp.BceResponse.GetHeader(http.CACHE_CONTROL)
}

func (resp *GetObjectMetaResponse) ContentDisposition() string {
    return resp.BceResponse.GetHeader(http.CONTENT_DISPOSITION)
}

func (resp *GetObjectMetaResponse) ContentLength() int64 {
    length := resp.BceResponse.GetHeader(http.CONTENT_LENGTH)
    val, _ := strconv.ParseInt(length, 10, 64)
    return val
}

func (resp *GetObjectMetaResponse) ContentRange() string {
    return resp.BceResponse.GetHeader(http.CONTENT_RANGE)
}

func (resp *GetObjectMetaResponse) ContentType() string {
    return resp.BceResponse.GetHeader(http.CONTENT_TYPE)
}

func (resp *GetObjectMetaResponse) ContentMD5() string {
    return resp.BceResponse.GetHeader(http.CONTENT_MD5)
}

func (resp *GetObjectMetaResponse) Expires() string {
    return resp.BceResponse.GetHeader(http.EXPIRES)
}

func (resp *GetObjectMetaResponse) ETag() string {
    return resp.BceResponse.GetHeader(http.ETAG)
}

func (resp *GetObjectMetaResponse) LastModified() string {
    return resp.BceResponse.GetHeader(http.LAST_MODIFIED)
}

func (resp *GetObjectMetaResponse) StorageClass() string {
    return resp.BceResponse.GetHeader(http.BCE_STORAGE_CLASS)
}

func (resp *GetObjectMetaResponse) UserMata(key string) string {
    headers := resp.BceResponse.GetHeaders()
    if val, ok := headers[key]; ok {
        return val
    }
    return ""
}

// Only used in appendable object
func (resp *GetObjectMetaResponse) NextAppendOffset() string {
    return resp.BceResponse.GetHeader(http.BCE_PREFIX + "next-append-object")
}
func (resp *GetObjectMetaResponse) ObjectType() string {
    return resp.BceResponse.GetHeader(http.BCE_PREFIX + "object-type")
}

func NewGetObjectMetaResponse() *GetObjectMetaResponse {
    return &GetObjectMetaResponse{}
}

