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

// request.go - defines the BCE servies request

package bce

import (
    "baidubce/http"
    "baidubce/thirdlib/uuid"
)

// BceRequest defines the request structure for accessing BCE services
type BceRequest struct {
    http.Request
    requestId   string
    clientError *BceClientError
}

func (bceReq *BceRequest) RequestId() string { return bceReq.requestId }

func (bceReq *BceRequest) SetRequestId(val string) { bceReq.requestId = val }

func (bceReq *BceRequest) ClientError() *BceClientError { return bceReq.clientError }

func (bceReq *BceRequest) SetClientError(err *BceClientError) { bceReq.clientError = err }

func (bceReq *BceRequest) BuildHttpRequest() {
    // Only need to build the specific `requestId` field for BCE, other fields are same as the
    // `http.Request` as well as its methods.
    if len(bceReq.requestId) == 0 {
        // Construct the request ID with UUID V1
        bceReq.requestId = uuid.NewV1().String()
    }
    bceReq.SetHeader(http.BCE_REQUEST_ID, bceReq.requestId)
}

func (bceReq *BceRequest) String() string {
    requestIdStr := "requestId=" + bceReq.requestId
    if bceReq.clientError != nil {
        return requestIdStr + ", client error: "+ bceReq.ClientError().Error()
    }
    return requestIdStr + "\n" + bceReq.Request.String()
}

