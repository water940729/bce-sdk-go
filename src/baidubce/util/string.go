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

// string.go - define the string util function

// Package util defines the common utilities including string and time.
package util

import (
    "bytes"
    "crypto/hmac"
    "crypto/sha256"
    "encoding/hex"
    "fmt"
)

func HmacSha256Hex(key, str_to_sign string) string {
    hasher := hmac.New(sha256.New, []byte(key))
    hasher.Write([]byte(str_to_sign))
    return hex.EncodeToString(hasher.Sum(nil))
}

func UriEncode(uri string, encodeSlash bool) string {
    var byte_buf bytes.Buffer
    for _, b := range []byte(uri) {
        if (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') || (b >= '0' && b <= '9') ||
            b == '-' || b == '_' || b == '.' || b == '~' || (b == '/' && !encodeSlash) {
            byte_buf.WriteByte(b)
        } else {
            byte_buf.WriteString(fmt.Sprintf("%%%02X", b))
        }
    }
    return byte_buf.String()
}

