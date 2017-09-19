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

// Copy object response definition
package model

import (
    "baidubce/model"
    "baidubce/util"
)

type CopyObjectResponse struct {
    model.BceResponse
    lastModified string
    eTag         string
}

func (resp *CopyObjectResponse) LastModified() string { return resp.lastModified }

func (resp *CopyObjectResponse) ETag() string { return resp.eTag }

func (resp *CopyObjectResponse) ParseResponse() {
    resp.BceResponse.ParseResponse()
    if resp.BceResponse.IsFail() { return }

    jsonBody := &CopyObjectOutput{}
    if err := resp.BceResponse.ParseJsonBody(jsonBody); err != nil {
        util.LOGGER.Error().Printf("parse copy object json response failed: %v\n", err)
        return
    }
    resp.lastModified = jsonBody.LastModified
    resp.eTag = jsonBody.ETag
}

func NewCopyObjectResponse() *CopyObjectResponse {
    return &CopyObjectResponse{}
}

