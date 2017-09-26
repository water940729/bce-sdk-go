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

// config.go - define the client configuration for BCE

package bce

import (
    "fmt"
    "runtime"
    "reflect"

    "baidubce/auth"
)

// Constants for BceClientConfiguration
var (
    DEFAULT_REGION        = "bj"
    DEFAULT_USER_AGENT    string
    DEFAULT_RETRY_POLICY  = NewBackOffRetryPolicy(3, 20000, 300)
    DEFAULT_CONNECTION_TIMEOUT_IN_MILLIS = 50 * 1000
    PROTOCOL_HTTP         = "http"
    PROTOCOL_HTTPS        = "https"
    PROTOCOL_DEFAULT_PORT = map[string]int{PROTOCOL_HTTP: 80, PROTOCOL_HTTPS: 443}
)

func init() {
    DEFAULT_USER_AGENT = "bce-sdk-go"
    DEFAULT_USER_AGENT += "/" + SDK_VERSION
    DEFAULT_USER_AGENT += "/" + runtime.Version()
    DEFAULT_USER_AGENT += "/" + runtime.GOOS
    DEFAULT_USER_AGENT += "/" + runtime.GOARCH
}

// BceClientConfiguration defines the config components structure.
type BceClientConfiguration struct {
    Endpoint     string
    ProxyUrl     string
    Region       string
    UserAgent    string
    Credentials  *auth.BceCredentials
    SignOption   *auth.SignOptions
    Retry        RetryPolicy
    ConnectionTimeoutInMillis int
}

func (conf *BceClientConfiguration) String() string {
    return fmt.Sprintf(`BceClientConfiguration [
        Endpoint=%s;
        ProxyUrl=%s;
        Region=%s;
        UserAgent=%s;
        Credentials=%v;
        SignOption=%v;
        RetryPolicy=%v;
        ConnectionTimeoutInMillis=%v
    ]`, conf.Endpoint, conf.ProxyUrl, conf.Region, conf.UserAgent, conf.Credentials,
    conf.SignOption, reflect.TypeOf(conf.Retry).Name(), conf.ConnectionTimeoutInMillis)
}

