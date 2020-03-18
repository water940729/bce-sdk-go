/*
 * Copyright 2020 Baidu, Inc.
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

// util.go - define the utilities for api package of BBC service
package bbc

import (
	"encoding/hex"
	"fmt"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/util/crypto"
)

const (
	URI_PREFIX_V1 = bce.URI_PREFIX + "v1"
	URI_PREFIX_V2 = bce.URI_PREFIX + "v2"

	REQUEST_INSTANCE_URI = "/instance"
	REQUEST_SUBNET_URI   = "/vpcSubnet"

	REQUEST_IMAGE_URI = "/image"

	REQUEST_FLAVOR_URI      = "/flavor"
	REQUEST_FLAVOR_RAID_URI = "/flavorRaid"

	REQUEST_OPERATION_LOG_URI = "/operationLog"

	REQUEST_DEPLOY_SET_URI = "/deployset"
)

func Aes128EncryptUseSecreteKey(sk string, data string) (string, error) {
	if len(sk) < 16 {
		return "", fmt.Errorf("error secrete key")
	}

	crypted, err := crypto.EBCEncrypto([]byte(sk[:16]), []byte(data))
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(crypted), nil
}
