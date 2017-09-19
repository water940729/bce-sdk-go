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

// Access to BCE services by : BCE accessing key ID and secret access key
package auth

import "errors"

type BceCredentials interface {
    GetAccessKeyId() string
    GetSecretAccessKey() string
}

// Default BCE credentials
type BceDefaultCredentials struct {
    accessKeyId     string
    secretAccessKey string
}

func (bdc *BceDefaultCredentials) GetAccessKeyId() string {
    return bdc.accessKeyId
}

func (bdc *BceDefaultCredentials) GetSecretAccessKey() string {
    return bdc.secretAccessKey
}

func (bdc *BceDefaultCredentials) String() string {
    return "ak:" + bdc.GetAccessKeyId() + ",sk:" + bdc.GetSecretAccessKey()
}

func NewBceDefaultCredentials(ak, sk string) (*BceDefaultCredentials, error) {
    if len(ak) == 0 {
        return nil, errors.New("accessKeyId should not be empty")
    }
    if len(sk) == 0 {
        return nil, errors.New("secretKey should not be empty")
    }

    return &BceDefaultCredentials{ak, sk}, nil
}

// BCE credentials which use additional session token
type BceSessionCredentials struct {
    BceDefaultCredentials
    sessionToken string
}

func (bsc *BceSessionCredentials) GetAccessKeyId() string {
    return bsc.accessKeyId
}

func (bsc *BceSessionCredentials) GetSecretAccessKey() string {
    return bsc.secretAccessKey
}

func (bsc *BceSessionCredentials) GetSessionToken() string {
    return bsc.sessionToken
}

func (bsc *BceSessionCredentials) String() string {
    return bsc.BceDefaultCredentials.String() + ",token:" + bsc.GetSessionToken()
}

func NewBceSessionCredentials(ak, sk, token string) (*BceSessionCredentials, error) {
    var cred *BceDefaultCredentials
    var err error
    cred, err = NewBceDefaultCredentials(ak, sk)
    if err != nil {
        return nil, err
    }
    if len(token) == 0 {
        return nil, errors.New("SessionToken should not be empty")
    }

    return &BceSessionCredentials{*cred, token}, nil
}

