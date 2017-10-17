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

// Package bce implements the infrastructure to access BCE services.
// - BceClient:
//     It is the general client of BCE to access all services. It builds http request to access the
//     services based on the given client configuration.
// - BceClientConfiguration:
//     The client configuration data structure which contains endpoint, region, credentials, retry
//     policy, sign options and so on. It supports most of the default value and user can also
//     access or change the default with its public fields' name.
// - Error types:
//     The error types when making request or receiving response to the BCE services contains two
//     types: the BceClientError when making request to BCE services and the BceServiceError when
//     recieving response from them.
// - BceRequest:
//     The request instance stands for an request to access the BCE services.
// - BceResponse:
//     The response instance stands for an response from the BCE services.
// - Log facility functions:
//     Define some global functions to set the log handler, log level and log directory.
package bce

import (
    "fmt"
    "time"

    "baidubce/auth"
    "baidubce/http"
    "baidubce/thirdlib/glog"
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

// Client is the general interface which can perform sending request. Different service
// will define its own client in case of specific extension.
type Client interface {
    SendRequest(*BceRequest, *BceResponse) error
}

// BceClient defines the general client to access the BCE services.
type BceClient struct {
    Config    *BceClientConfiguration
    Signer    auth.Signer // the sign algorithm
}

/*
 * BuildHttpRequest - the helper method for the client to build http request
 *
 * PARAMS:
 *     - request: the input request object to be built
 */
func (cli *BceClient) buildHttpRequest(request *BceRequest) {
    // Construct the http request instance for the special fields
    request.BuildHttpRequest()

    // Set the client specific configurations
    request.SetEndpoint(cli.Config.Endpoint)
    if request.Protocol() == "" {
        request.SetProtocol(DEFAULT_PROTOCOL)
    }
    if len(cli.Config.ProxyUrl) != 0 {
        request.SetProxyUrl(cli.Config.ProxyUrl)
    }
    request.SetTimeout(cli.Config.ConnectionTimeoutInMillis / 1000)

    // Set the BCE request headers
    request.SetHeader(http.HOST, request.Host())
    request.SetHeader(http.USER_AGENT, cli.Config.UserAgent)
    request.SetHeader(http.BCE_DATE, util.FormatISO8601Date(cli.Config.SignOption.Timestamp))
    if request.Body() != nil {
        request.SetHeader(http.CONTENT_LENGTH, fmt.Sprintf("%d", request.Body().Len()))
    }

    // Generate the auth string
    cli.Signer.Sign(&request.Request, cli.Config.Credentials, cli.Config.SignOption)
}

/*
 * SendRequest - the client performs sending the http request with retry policy and receive the
 *               response from the BCE services.
 * PARAMS:
 *     - req: the request object to be sent to the BCE service
 *     - resp: the response object to receive the content from BCE service
 * RETURNS:
 *     - error: nil if ok otherwise the specific error
 */
func (cli *BceClient) SendRequest(req *BceRequest, resp *BceResponse) error {
    // Return client error if it is not nil
    if req.ClientError() != nil {
        return req.ClientError()
    }

    // Build the http request and prepare to send
    cli.buildHttpRequest(req)
    glog.Infof("send http request: %v", req)

    // Send request with the given retry policy
    retries := 0
    for {
        httpResp, err := http.Execute(&req.Request)
        if err != nil {
            if cli.Config.Retry.ShouldRetry(err, retries) {
                delay_in_mills := cli.Config.Retry.GetDelayBeforeNextRetryInMillis(err, retries)
                time.Sleep(delay_in_mills)
            } else {
                return &BceClientError{
                        fmt.Sprintf("execute http request failed! Retried %d times, error: %v",
                                    retries, err)}
            }
            retries++
            glog.Warningf("send request failed, retry for %d time(s)", retries)
            continue
        }
        resp.SetHttpResponse(httpResp)
        resp.ParseResponse()

        glog.Infof("receive http response: status: %s, debugId: %s, requestId: %s, elapsed: %v",
            resp.StatusText(), resp.DebugId(), resp.RequestId(), resp.ElapsedTime())
        if resp.IsFail() {
            return resp.ServiceError()
        }
        for k, v := range resp.Headers() {
            glog.Debugf("%s=%s", k, v)
        }
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

