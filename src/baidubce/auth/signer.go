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

// Implement the Signer of BCE V1 protocol
package auth

import (
    "fmt"
    "strings"
    "sort"

    "baidubce/http"
    "baidubce/util"
)

var (
    BCE_AUTH_VERSION        = "bce-auth-v1"
    SIGN_JOINER             = "\n"
    SIGN_HEADER_JOINER      = ";"
    DEFAULT_EXPIRE_SECONDS  = 1800
    DEFAULT_HEADERS_TO_SIGN []string
)

func init() {
    DEFAULT_HEADERS_TO_SIGN = make([]string, 4)
    DEFAULT_HEADERS_TO_SIGN[0] = strings.ToLower(http.HOST)
    DEFAULT_HEADERS_TO_SIGN[1] = strings.ToLower(http.CONTENT_LENGTH)
    DEFAULT_HEADERS_TO_SIGN[2] = strings.ToLower(http.CONTENT_TYPE)
    DEFAULT_HEADERS_TO_SIGN[3] = strings.ToLower(http.CONTENT_MD5)
}

type Signer interface {
    // Sign the given Request with the Credentials and SignOptions
    Sign(*http.Request, BceCredentials, *SignOptions)
}

// Sign optinos data structure used for `Signer` implementation
type SignOptions struct {
    HeadersToSign []string
    Timestamp     int64
    ExpireSeconds int
}

func (opt* SignOptions) String() string {
    return fmt.Sprintf(`SignOptions [
        HeadersToSign=%s;
        Timestamp=%d;
        ExpireSeconds=%d
    ]`, opt.HeadersToSign, opt.Timestamp, opt.ExpireSeconds)
}

// BCE version 1 signer which implements Signer interface
type BceV1Signer struct {}

func (bvs *BceV1Signer) Sign(req *http.Request, cred BceCredentials, opt *SignOptions) {
    if req == nil {
        util.LOGGER.Panic().Println("request should not be null")
    }
    if cred == nil {
        return
    }

    // Prepare parameters
    accessKeyId     := cred.GetAccessKeyId()
    secretAccessKey := cred.GetSecretAccessKey()
    signDate        := util.FormatISO8601Date(opt.Timestamp)

    // Set security token if using session credentials
    if v, ok := cred.(*BceSessionCredentials); ok {
        req.AddHeader(http.BCE_SECURITY_TOKEN, v.GetSessionToken())
    }

    // Prepare the canonical request components
    signKeyInfo := fmt.Sprintf("%s/%s/%s/%d",
                        BCE_AUTH_VERSION,
                        accessKeyId,
                        signDate,
                        opt.ExpireSeconds)
    signKey := util.HmacSha256Hex(secretAccessKey, signKeyInfo)
    canonicalUri := getCanonicalURIPath(req.Uri())
    canonicalQueryString := getCanonicalQueryString(req.Params())
    canonicalHeaders, signedHeadersArr := getCanonicalHeaders(req.Headers(), opt.HeadersToSign)

    // Generate signed headers string
    signedHeaders := ""
    if len(signedHeadersArr) > 0 {
        sort.Strings(signedHeadersArr)
        signedHeaders = strings.Join(signedHeadersArr, SIGN_HEADER_JOINER)
    }

    // Generate signature
    canonicalParts := []string{req.Method(), canonicalUri, canonicalQueryString, canonicalHeaders}
    canonicalReq := strings.Join(canonicalParts, SIGN_JOINER)
    util.LOGGER.Debug().Println("CanonicalRequest data:\n" + canonicalReq)
    signature := util.HmacSha256Hex(signKey, canonicalReq)

    // Generate auth string and add to the reqeust header
    authStr := signKeyInfo + "/" + signedHeaders + "/" + signature
    util.LOGGER.Debug().Println("Authorization=" + authStr)

    req.AddHeader(http.AUTHORIZATION, authStr)
}

func getCanonicalURIPath(path string) string {
    if len(path) == 0 {
        return "/"
    }
    canonical_path := path
    if strings.HasPrefix(path, "/") {
        canonical_path = path[1:]
    }
    canonical_path = util.UriEncode(canonical_path, false)
    return "/" + canonical_path
}

func getCanonicalQueryString(params map[string]string) string {
    if len(params) == 0 {
        return ""
    }

    result := make([]string, 0, len(params))
    for k, v := range params {
        if strings.ToLower(k) == strings.ToLower(http.AUTHORIZATION) {
            continue
        }
        item := ""
        if len(v) == 0 {
            item = fmt.Sprintf("%s=", util.UriEncode(k, true))
        } else {
            item = fmt.Sprintf("%s=%s", util.UriEncode(k, true), util.UriEncode(v, true))
        }
        result = append(result, item)
    }
    sort.Strings(result)
    return strings.Join(result, "&")
}

func getCanonicalHeaders(headers map[string]string, headersToSign []string) (string, []string) {
    canonicalHeaders := make([]string, 0, len(headers))
    signHeaders := make([]string, 0, len(headersToSign))
    finder := func(strArr []string, search string) int {
        for i, h := range strArr {
            if h == search {
                return i
            }
        }
        return -1
    }
    for k, _ := range headers {
        if strings.ToLower(k) == strings.ToLower(http.AUTHORIZATION) {
            continue
        }
        lowerHead := strings.ToLower(k)
        if (finder(headersToSign, lowerHead) != -1) ||
           (strings.HasPrefix(lowerHead, http.BCE_PREFIX) &&
           (lowerHead != http.BCE_REQUEST_ID)) {

            headVal := strings.TrimSpace(headers[k])
            item := util.UriEncode(lowerHead, true) + ":" + util.UriEncode(headVal, true)
            canonicalHeaders = append(canonicalHeaders, item)
            signHeaders = append(signHeaders, lowerHead)
        }
    }
    sort.Strings(canonicalHeaders)
    sort.Strings(signHeaders)
    return strings.Join(canonicalHeaders, SIGN_JOINER), signHeaders
}

