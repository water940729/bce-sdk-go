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

// list_parts_response.go - the list uploaded parts response definition

package model

import (
    "glog"

    "baidubce/model"
)

type ListPartsResponse struct {
    model.BceResponse
    bucket      string
    key         string
    uploadId    string
    owner       BucketOwnerType
    maxParts    int
    isTruncated bool
    parts       []ListPartType
    partNumberMarker     string
    nextPartNumberMarker int
}

func (resp *ListPartsResponse) Bucket() string { return resp.bucket }
func (resp *ListPartsResponse) Key() string { return resp.key }
func (resp *ListPartsResponse) UploadId() string { return resp.uploadId }
func (resp *ListPartsResponse) Owner() BucketOwnerType { return resp.owner }
func (resp *ListPartsResponse) MaxParts() int { return resp.maxParts }
func (resp *ListPartsResponse) IsTruncated() bool { return resp.isTruncated }
func (resp *ListPartsResponse) Parts() []ListPartType { return resp.parts }
func (resp *ListPartsResponse) PartNumberMarker() string { return resp.partNumberMarker }
func (resp *ListPartsResponse) NextPartNumberMarker() int { return resp.nextPartNumberMarker }

func (resp *ListPartsResponse) ParseResponse() {
    resp.BceResponse.ParseResponse()
    if resp.BceResponse.IsFail() { return }

    jsonBody := &ListPartsOutput{}
    if err := resp.BceResponse.ParseJsonBody(jsonBody); err != nil {
        glog.Error("parse list parts json response failed: %v\n", err)
        return
    }
    resp.bucket = jsonBody.Bucket
    resp.key = jsonBody.Key
    resp.uploadId = jsonBody.UploadId
    resp.owner = jsonBody.Owner
    resp.maxParts = jsonBody.MaxParts
    resp.isTruncated = jsonBody.IsTruncated
    resp.parts = jsonBody.Parts
    resp.partNumberMarker = jsonBody.PartNumberMarker
    resp.nextPartNumberMarker = jsonBody.NextPartNumberMarker
}

func NewListPartsResponse() *ListPartsResponse {
    return &ListPartsResponse{}
}

