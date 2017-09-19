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

// Define the exceptions types for BCE request and response
package common

type BceError interface { error }

// Define the error struct from the client
type BceClientError struct { Message string }

func (bcr *BceClientError) Error() string { return bcr.Message }

func NewBceClientError(msg string) *BceClientError { return &BceClientError{msg} }

// Define the error struct from the BCE service
type BceServiceError struct {
    Code       string
    Message    string
    RequestId  string
    StatusCode int
}

func (bsr *BceServiceError) Error() string {
    ret := "[Code: " + bsr.Code
    ret += "; Message: " + bsr.Message
    ret += "; RequestId: " + bsr.RequestId + "]"
    return ret
}

func NewBceServiceError(code, msg, reqId string, status int) *BceServiceError {
    return &BceServiceError{code, msg, reqId, status}
}

// Error string constants
var (
    EACCESS_DENIED = "AccessDenied"
    EINAPPROPRIATE_JSON = "InappropriateJSON"
    EINTERNAL_ERROR = "InternalError"
    EINVALID_ACCESS_KEY_ID = "InvalidAccessKeyId"
    EINVALID_HTTP_AUTH_HEADER = "InvalidHTTPAuthHeader"
    EINVALID_HTTP_REQUEST = "InvalidHTTPRequest"
    EINVALID_URI = "InvalidURI"
    EMALFORMED_JSON = "MalformedJSON"
    EINVALID_VERSION = "InvalidVersion"
    EOPT_IN_REQUIRED = "OptInRequired"
    EPRECONDITION_FAILED = "PreconditionFailed"
    EREQUEST_EXPIRED = "RequestExpired"
    ESIGNATURE_DOES_NOT_MATCH = "SignatureDoesNotMatch"
)

