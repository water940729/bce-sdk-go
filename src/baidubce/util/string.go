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

// The string utilities
package util

import (
    "fmt"
    "bytes"
    "encoding/hex"
    "crypto/hmac"
    "crypto/sha256"
)

func HmacSha256Hex(key, str_to_sign string) string {
    hasher := hmac.New(sha256.New, []byte(key))
    hasher.Write([]byte(str_to_sign))
    return hex.EncodeToString(hasher.Sum(nil))
}

func UriEncode(uri string, encodeSlash bool) string {
    ascii := make([]byte, 0, 256)
    for i := 0; i < 256; i++ {
        ascii = append(ascii, byte(i))
    }

    unreserved_chars := make([]byte, 0, 10 + 26 + 26 + 4)
    for i := '0'; i <= '9'; i++ {
        unreserved_chars = append(unreserved_chars, byte(i))
    }
    for i := 'a'; i <= 'z'; i++ {
        unreserved_chars = append(unreserved_chars, byte(i))
    }
    for i := 'A'; i <= 'Z'; i++ {
        unreserved_chars = append(unreserved_chars, byte(i))
    }
    unreserved_chars = append(unreserved_chars, '-')
    unreserved_chars = append(unreserved_chars, '_')
    unreserved_chars = append(unreserved_chars, '.')
    unreserved_chars = append(unreserved_chars, '~')

    var byte_buf bytes.Buffer
    for _, b := range []byte(uri) {
        if b == '/' && encodeSlash == false {
            byte_buf.WriteByte(b)
        } else if bytes.IndexByte(unreserved_chars, b) != -1 {
            byte_buf.WriteByte(b)
        } else {
            byte_buf.WriteString(fmt.Sprintf("%%%02X", b))
        }
    }
    return byte_buf.String()
}

