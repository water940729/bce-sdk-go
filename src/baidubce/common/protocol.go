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

// Communication protocol used to sending requests for BCE
package common

type ProtocolType struct {
    protocol string
    defaultPort int
}

var (
    HTTP  = &ProtocolType{"http", 80}
    HTTPS = &ProtocolType{"https", 443}
)

func (p ProtocolType) DefaultPort() int {
    return p.defaultPort
}

func (p ProtocolType) String() string {
    return p.protocol
}

