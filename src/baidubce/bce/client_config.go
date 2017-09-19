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

// Define the common client configuration for BCE
package bce

import (
    "fmt"
    "runtime"
    "reflect"

    . "baidubce/common"
    "baidubce/auth"
)

// Constants for BceClientConfiguration
var (
    DEFAULT_PROTOCOL      = HTTP
    DEFAULT_REGION        = BEIJING
    DEFAULT_USER_AGENT    string
    DEFAULT_RETRY_POLICY  = DefaultRetryPolicy
    DEFAULT_CONNECTION_TIMEOUT_IN_MILLIS = 50 * 1000
)

func init() {
    DEFAULT_USER_AGENT = "bce-sdk-go"
    DEFAULT_USER_AGENT += "/" + SDK_VERSION
    DEFAULT_USER_AGENT += "/" + runtime.Version()
    DEFAULT_USER_AGENT += "/" + runtime.GOOS
    DEFAULT_USER_AGENT += "/" + runtime.GOARCH
}

type BceClientConfiguration struct {
    Protocol     *ProtocolType
    Endpoint     string
    Region       *RegionType
    Credentials  auth.BceCredentials
    SignOption   *auth.SignOptions
    Retry        RetryPolicy
    UserAgent    string
    ConnectionTimeoutInMillis int
}

func (conf *BceClientConfiguration) String() string {
    return fmt.Sprintf(`BceClientConfiguration [
        Protocol=%s;
        Endpoint=%s;
        Region=%s;
        Credentials=%v;
        SignOption=%v;
        RetryPolicy=%v;
        UserAgent=%s;
        ConnectionTimeoutInMillis=%v
    ]`, conf.Protocol, conf.Endpoint, conf.Region, conf.Credentials, conf.SignOption,
    reflect.TypeOf(conf.Retry).Name(), conf.UserAgent, conf.ConnectionTimeoutInMillis)
}

