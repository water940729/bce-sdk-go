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

// delete_multiple_objects_responsel.go - the delete multiple objects response definition

package model

import (
    "glog"

    "baidubce/model"
)

type DeleteMultipleObjectsResponse struct {
    model.BceResponse
    failedObjects []DeleteFailedOutput
}

func (resp *DeleteMultipleObjectsResponse) FailedObjects() []DeleteFailedOutput {
    return resp.failedObjects
}

func (resp *DeleteMultipleObjectsResponse) ParseResponse() {
    resp.BceResponse.ParseResponse()
    if resp.BceResponse.IsFail() { return }

    jsonBody := &DeleteMultipleObjectsOutput{}
    if err := resp.BceResponse.ParseJsonBody(jsonBody); err != nil {
        glog.Error("parse delete multiple objects response failed: %v\n", err)
        return
    }
    resp.failedObjects = jsonBody.Errors
}

func NewDeleteMultipleObjectsResponse() *DeleteMultipleObjectsResponse {
    return &DeleteMultipleObjectsResponse{}
}

