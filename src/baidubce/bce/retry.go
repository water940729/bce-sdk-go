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

// retry.go - define the retry policy when making requests to BCE services

package bce

import (
    "time"
    "net"
    "net/http"

    "baidubce/thirdlib/glog"
)

// RetryPolicy defines the two methods to retry for sending request.
type RetryPolicy interface {
    ShouldRetry(BceError, int) bool
    GetDelayBeforeNextRetryInMillis(BceError, int) time.Duration
}

// NoRetryPolicy just does not retry.
type NoRetryPolicy struct {}

func (rp *NoRetryPolicy) ShouldRetry(err BceError, attempts int) bool {
    return false
}

func (rp *NoRetryPolicy) GetDelayBeforeNextRetryInMillis(
        err BceError, attempts int) time.Duration {
    return 0 * time.Millisecond
}

func NewNoRetryPolicy() *NoRetryPolicy {
    return &NoRetryPolicy{}
}

// BackOffRetryPolicy implements a policy that retries with exponential back-off strategy.
// This policy will keep retrying until the maximum number of retries is reached. The delay time
// will be a fixed interval for the first time then 2 * interval for the second, 4 * internal for
// the third, and so on.
// In general, the delay time will be 2^number_of_retries_attempted*interval. When a maximum of
// delay time is specified, the delay time will never exceed this limit.
type BackOffRetryPolicy struct {
    maxErrorRetry        int
    maxDelayInMillis     int64
    baseIntervalInMillis int64
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
            glog.Warning("retry for internal server error(500)")
            return true
        }
        if realErr.StatusCode == http.StatusBadGateway {
            glog.Warning("retry for bad gateway(502)")
            return true
        }
        if realErr.StatusCode == http.StatusServiceUnavailable {
            glog.Warning("retry for service unavailable(503)")
            return true
        }

        if realErr.Code == EREQUEST_EXPIRED {
            glog.Warning("retry for request expired")
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
        return time.Duration(rp.maxDelayInMillis) * time.Millisecond
    }
    return time.Duration(delayInMillis) * time.Millisecond
}

func NewBackOffRetryPolicy(maxRetry int, maxDelay, base int64) *BackOffRetryPolicy {
    return &BackOffRetryPolicy{maxRetry, maxDelay, base}
}

