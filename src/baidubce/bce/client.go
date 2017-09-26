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

// client.go - definiton the BceClientConfiguration and BceClient structure

// Package bce implements the BCE client to access BCE services.
// This package defines the general client of BCE - BceClient - to access all services.
// The BceClient makes http request to access the services based on the given client configuration,
// which contains endpoint, region, credentials, retry policy, sign options and so on.
// The interface of BCE Request and response are defined for the concrete implementation.
// This package also defines the error types when making request or receiving response. The error
// for BCE contains two types: the BceClientError when making request to BCE services and the Bce-
// ServiceError when recieving response from them.
package bce

import (
    "fmt"
    "io"
    "time"

    "glog"

    "baidubce/auth"
    "baidubce/http"
    "baidubce/util"
)

const (
    SDK_VERSION            = "0.1.0"
    URI_PREFIX             = "/v1/"
    DEFAULT_SERVICE_DOMAIN = "bcebos.com"

    LOG_DEBUG int32 = iota
    LOG_INFO
    LOG_WARN
    LOG_ERROR
    LOG_FATAL
)

// BceRequester is an interface for setting up a contract to make request.
type BceRequester interface {
    GetUri() string
    SetHeader(key, val string)
    SetBody(*http.BodyStream)
    BuildHttpRequest() *http.Request  // May be override by specific request
    ClientError() *BceClientError
}

// BceResponser is an interface for setting up a contract to reveive response.
type BceResponser interface {
    IsFail() bool
    StatusCode() int
    StatusText() string
    RequestId() string
    DebugId() string
    GetHeader(key string) string
    GetHeaders() map[string]string
    Body() io.ReadCloser
    SetHttpResponse(*http.Response)
    ElapsedTime() time.Duration
    ParseResponse()
    ServiceError() *BceServiceError
}

// BceClient defines a client to access the BCE services.
type BceClient struct {
    Config    *BceClientConfiguration
    Signer    auth.Signer // the sign algorithm
}

func (cli *BceClient) BuildHttpRequest(request BceRequester) *http.Request {
    // Construct the http request instance
    httpReq := request.BuildHttpRequest()
    if httpReq == nil {
        return nil
    }

    // Set configuration
    httpReq.SetEndpoint(cli.Config.Endpoint)
    if httpReq.Protocol() == "" {
        httpReq.SetProtocol(PROTOCOL_HTTP)
    }
    if httpReq.Port() == 0 {
        port, err := PROTOCOL_DEFAULT_PORT[httpReq.Protocol()]
        if !err {
            return nil
        }
        httpReq.SetPort(port)
    }
    if len(cli.Config.ProxyUrl) != 0 {
        httpReq.SetProxyUrl(cli.Config.ProxyUrl)
    }
    httpReq.SetTimeout(cli.Config.ConnectionTimeoutInMillis / 1000)

    // Add common headers headers
    httpReq.AddHeader(http.HOST, httpReq.Host())
    httpReq.AddHeader(http.USER_AGENT, cli.Config.UserAgent)
    httpReq.AddHeader(http.BCE_DATE, util.FormatISO8601Date(cli.Config.SignOption.Timestamp))
    if httpReq.Body() != nil {
        httpReq.AddHeader(http.CONTENT_LENGTH, fmt.Sprintf("%d", httpReq.Body().Len()))
    }

    // Generate the auth string
    cli.Signer.Sign(httpReq, cli.Config.Credentials, cli.Config.SignOption)

    return httpReq
}

func (cli *BceClient) SendRequest(req BceRequester, resp BceResponser) error {
    httpReq := cli.BuildHttpRequest(req)
    if httpReq == nil {
        return req.ClientError()
    }
    glog.Info("Send http request: %v\n", httpReq)

    retries := 0
    for {
        httpResp, err := http.Execute(httpReq)
        if err != nil {
            if cli.Config.Retry.ShouldRetry(err, retries) {
                delay_in_mills := cli.Config.Retry.GetDelayBeforeNextRetryInMillis(err, retries)
                time.Sleep(delay_in_mills)
            } else {
                return &BceClientError{
                        fmt.Sprintf("Execute http request failed! Retried %d times, error: %v",
                                    retries, err)}
            }
            retries++
            continue
        }

        resp.SetHttpResponse(httpResp)
        resp.ParseResponse()
        return nil
    }
}

func NewBceClient(conf *BceClientConfiguration, sign auth.Signer) *BceClient {
    return &BceClient{conf, sign}
}

/*
 * SetLogHandler - set the log handler of where to log, default is to stderr
 *
 * PARAMS:
 *     -toStderr: set weather logging to stderr(default true)
 *     -toFile: set weather logging to file(default false) and if set to file all levels are logged
 */
func SetLogHandler(toStderr, toFile bool) {
    glog.SetLogToStderr(toStderr)

    if toFile {
        glog.SetLogAlsoToStderr(toStderr)
    }

    if !toStderr && !toFile {
        glog.SetLogStderrThreshold(5)
    }
}

/*
 * SetLogLevel - if log to stderr and file, set the level to be logged(default error)
 *
 * PARAMS:
 *     -level: if set logging to stderr and file, only the message whose level is greater or equal
 *             than this value will be logged to stderr. Default is 3(error).
 *             available value: 0=debug, 1=info, 2=warning, 3=error, 4=fatal
 */
func SetLogLevel(level int32) {
    if level >= LOG_DEBUG && level <= LOG_FATAL {
        glog.SetLogStderrThreshold(level)
    }
    glog.SetLogStderrThreshold(LOG_ERROR)
}

/*
 * SetLogDir - set the directory in which puts the log files
 *
 * PARAMS:
 *     -logDir: the log directory value
 */
func SetLogDir(logDir string) {
    glog.SetLogToDir(logDir)
}

