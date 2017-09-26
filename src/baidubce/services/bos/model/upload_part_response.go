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

// upload_part_response.go - the upload part response definition

package model

import (
    "baidubce/http"
    "baidubce/model"
)

type UploadPartResponse struct {
    model.BceResponse
}

func (resp *UploadPartResponse) ETag() string {
    return resp.BceResponse.GetHeader(http.ETAG)
}

func NewUploadPartResponse() *UploadPartResponse {
    return &UploadPartResponse{}
}

