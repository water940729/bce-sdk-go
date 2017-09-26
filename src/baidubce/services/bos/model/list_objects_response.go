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

// list_objects_response.go - the list objects response definition

package model

import (
    "glog"

    "baidubce/model"
)

type ListObjectsResponse struct {
    model.BceResponse
    name        string
    prefix      string
    delimiter   string
    marker      string
    nextMarker  string
    maxKeys     int
    isTruncated bool
    contents    []ObjectSummary
}

func (resp *ListObjectsResponse) Name() string { return resp.name }
func (resp *ListObjectsResponse) Prefix() string { return resp.prefix }
func (resp *ListObjectsResponse) Delimiter() string { return resp.delimiter }
func (resp *ListObjectsResponse) Marker() string { return resp.marker }
func (resp *ListObjectsResponse) NextMarker() string { return resp.nextMarker }
func (resp *ListObjectsResponse) MaxKeys() int { return resp.maxKeys }
func (resp *ListObjectsResponse) IsTruncated() bool { return resp.isTruncated }
func (resp *ListObjectsResponse) Contents() []ObjectSummary { return resp.contents }

func (resp *ListObjectsResponse) ParseResponse() {
    resp.BceResponse.ParseResponse()
    if resp.BceResponse.IsFail() { return }

    jsonBody := &ListObjectsOutput{}
    if err := resp.BceResponse.ParseJsonBody(jsonBody); err != nil {
        glog.Error("parse list objects json response failed: %v\n", err)
        return
    }
    resp.name = jsonBody.Name
    resp.prefix = jsonBody.Prefix
    resp.delimiter = jsonBody.Delimiter
    resp.marker = jsonBody.Marker
    resp.nextMarker = jsonBody.NextMarker
    resp.maxKeys = jsonBody.MaxKeys
    resp.isTruncated = jsonBody.IsTruncated
    resp.contents = jsonBody.Contents
}

func NewListObjectsResponse() *ListObjectsResponse {
    return &ListObjectsResponse{}
}

