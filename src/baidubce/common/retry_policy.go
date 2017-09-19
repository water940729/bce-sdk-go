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

// Define the retry policy when making requests to BCE services
package common

import (
    "time"
    "net"
    "net/http"

    "baidubce/util"
)

type RetryPolicy interface {
    ShouldRetry(BceError, int) bool
    GetDelayBeforeNextRetryInMillis(BceError, int) time.Duration
}

type NoRetryPolicy struct {}

func (rp *NoRetryPolicy) ShouldRetry(err BceError, attempts int) bool {
    return false
}

func (rp *NoRetryPolicy) GetDelayBeforeNextRetryInMillis(
    err BceError, attempts int) time.Duration {
    return 0 * time.Millisecond
}

type BackOffRetryPolicy struct {
    maxErrorRetry        int
    maxDelayInMillis     time.Duration
    baseIntervalInMillis time.Duration
}

func (rp *BackOffRetryPolicy) ShouldRetry(err BceError, attempts int) bool {
    // Do not retry any more when retry the max times
    if attempts >= rp.maxErrorRetry {
        return false
    }

    // Always retry on IO error
    if _, ok := err.(net.Error); ok {
        return true
    }

    // Only retry on a service error
    if realErr, ok := err.(*BceServiceError); ok {
        if realErr.StatusCode == http.StatusInternalServerError {
            util.LOGGER.Debug().Println("retry for internal server error(500)")
            return true
        }
        if realErr.StatusCode == http.StatusBadGateway {
            util.LOGGER.Debug().Println("retry for bad gateway(502)")
            return true
        }
        if realErr.StatusCode == http.StatusServiceUnavailable {
            util.LOGGER.Debug().Println("retry for service unavailable(503)")
            return true
        }

        if realErr.Code == EREQUEST_EXPIRED {
            util.LOGGER.Debug().Println("retry for request expired")
            return true;
        }
    }
    return false
}

func (rp *BackOffRetryPolicy) GetDelayBeforeNextRetryInMillis(
    err BceError, attempts int) time.Duration {
    if attempts < 0 {
        return 0 * time.Millisecond
    }
    delayInMillis := (1 << uint64(attempts)) * rp.baseIntervalInMillis
    if delayInMillis > rp.maxDelayInMillis {
        return rp.maxDelayInMillis
    }
    return delayInMillis
}

var DefaultRetryPolicy = &BackOffRetryPolicy{
    3,
    20000 * time.Millisecond,
    300 * time.Millisecond}

