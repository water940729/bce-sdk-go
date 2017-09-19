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

// Define the BCE client
package bce

import (
    "fmt"
    "time"

    "baidubce/auth"
    "baidubce/common"
    "baidubce/model"
    "baidubce/http"
    "baidubce/util"
)

type BceClient struct {
    Config *BceClientConfiguration
    Signer auth.Signer
}

func (cli *BceClient) BuildHttpRequest(request model.BceRequester) *http.Request {
    // Construct the http request instance
    httpReq := request.BuildHttpRequest()
    if httpReq == nil {
        return nil
    }

    // Set configuration
    httpReq.SetEndpoint(cli.Config.Endpoint)
    if httpReq.Protocol() == "" {
        httpReq.SetProtocol(cli.Config.Protocol.String())
    }
    if httpReq.Port() == 0 {
        httpReq.SetPort(cli.Config.Protocol.DefaultPort())
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

func (cli *BceClient) SendRequest(req model.BceRequester, resp model.BceResponser) error {
    httpReq := cli.BuildHttpRequest(req)
    if httpReq == nil {
        return req.ClientError()
    }
    util.LOGGER.Info().Printf("Send http request: %v\n", httpReq)

    retries := 0
    for {
        httpResp, err := http.Execute(httpReq)
        if err != nil {
            if cli.Config.Retry.ShouldRetry(err, retries) {
                delay_in_mills := cli.Config.Retry.GetDelayBeforeNextRetryInMillis(err, retries)
                time.Sleep(delay_in_mills)
            } else {
                return &common.BceClientError{
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

