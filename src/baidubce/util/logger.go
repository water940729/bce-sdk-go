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

// Construct the logger infrastructure
package util

import (
    "log"
    "io"
    "os"
)

var (
    DEFAULT_LOGGER_OUT = os.Stderr
    DEFAULT_LOGGER_PREFIX = "[DEBUG]"
)

type Logger interface {
    Debug() *BceLogger
    Info() *BceLogger
    Error() *BceLogger
    Fatal() *BceLogger
    Panic() *BceLogger
}

type BceLogger struct { log.Logger }

func (log *BceLogger) Debug() *BceLogger {
    log.SetPrefix("[DEBUG]")
    return log
}

func (log *BceLogger) Info() *BceLogger {
    log.SetPrefix("[INFO]")
    return log
}

func (log *BceLogger) Error() *BceLogger {
    log.SetPrefix("[ERROR]")
    return log
}

func (log *BceLogger) Fatal() *BceLogger {
    log.SetPrefix("[FATAL]")
    return log
}

func (log *BceLogger) Panic() *BceLogger {
    log.SetPrefix("[PANIC]")
    return log
}

func GetLogger(out io.Writer, prefix string) *BceLogger {
    logger := log.New(out, prefix, log.Lshortfile)
    return &BceLogger{*logger}
}

var LOGGER = GetLogger(DEFAULT_LOGGER_OUT, DEFAULT_LOGGER_PREFIX)

