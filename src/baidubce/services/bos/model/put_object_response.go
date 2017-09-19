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

// Put object response definition
package model

import (
    "baidubce/model"
    "baidubce/http"
)

type PutObjectResponse struct {
    model.BceResponse
}

func (resp *PutObjectResponse) ETag() string {
    return resp.BceResponse.GetHeader(http.ETAG)
}

func (resp *PutObjectResponse) Version() string {
    return resp.BceResponse.GetHeader(http.BCE_PREFIX + "version")
}

func NewPutObjectResponse() *PutObjectResponse {
    return &PutObjectResponse{}
}

