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

// util.go - define the utilities for api package of BOS service

package api

import (
    "bytes"

    "baidubce/bce"
)

const (
    METADATA_DIRECTIVE_COPY    = "copy"
    METADATA_DIRECTIVE_REPLACE = "replace"

    STORAGE_CLASS_STANDARD    = "STANDARD"
    STORAGE_CLASS_STANDARD_IA = "STANDARD_IA"
    STORAGE_CLASS_COLD        = "COLD"

    FETCH_MODE_SYNC  = "sync"
    FETCH_MODE_ASYNC = "async"

    CANNED_ACL_PRIVATE = "private"
    CANNED_ACL_PUBLIC_READ = "public-read"
    CANNED_ACL_PUBLIC_READ_WRITE = "public-read-write"

    DEFAULT_CONTENT_TYPE = "application/octet-stream"
)

var (
    GET_OBJECT_ALLOWED_RESPONSE_HEADERS = map[string]bool{
        "ContentDisposition": true,
        "ContentType": true,
        "ContentLanguage": true,
        "Expires": true,
        "CacheControl": true,
        "ContentEncoding": true,
    }
)

func getBucketUri(bucketName string) string {
    return bce.URI_PREFIX  + bucketName
}

func getObjectUri(bucketName, objectName string) string {
    return bce.URI_PREFIX  + bucketName + "/" + objectName
}

func validMetadataDirective(val string) bool {
    if val == METADATA_DIRECTIVE_COPY || val == METADATA_DIRECTIVE_REPLACE {
        return true
    }
    return false
}

func validStorageClass(val string) bool {
    if val == STORAGE_CLASS_STANDARD_IA ||
        val == STORAGE_CLASS_STANDARD ||
        val == STORAGE_CLASS_COLD {
        return true
    }
    return false
}

func validFetchMode(val string) bool {
    if val == FETCH_MODE_SYNC || val == FETCH_MODE_ASYNC {
        return true
    }
    return false
}

func validCannedAcl(val string) bool {
    if val == CANNED_ACL_PRIVATE ||
        val == CANNED_ACL_PUBLIC_READ ||
        val == CANNED_ACL_PUBLIC_READ_WRITE {
        return true
    }
    return false
}

func toHttpHeaderKey(key string) string {
    var result bytes.Buffer
    needToUpper := true
    for _, c := range []byte(key) {
        if needToUpper && (c >= 'a' && c <= 'z') {
            result.WriteByte(c - 32)
            needToUpper = false
        } else if c == '-' {
            result.WriteByte(c)
            needToUpper = true
        } else {
            result.WriteByte(c)
        }
    }
    return result.String()
}

