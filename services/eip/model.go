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

// model.go - definitions of the request arguments and results data structure model
package eip

import "github.com/baidubce/bce-sdk-go/util"

type Reservation struct {
	ReservationLength   int    `json:"reservationLength,omitempty"`
	ReservationTimeUnit string `json:"reservationTimeUnit,omitempty"`
}

type Billing struct {
	PaymentTiming string       `json:"paymentTiming,omitempty"`
	BillingMethod string       `json:"billingMethod,omitempty"`
	Reservation   *Reservation `json:"reservation,omitempty"`
}

type CreateEipArgs struct {
	Name            string          `json:"name,omitempty"`
	BandWidthInMbps int             `json:"bandwidthInMbps"`
	Billing         *Billing        `json:"billing"`
	Tags            []util.TagModel `json:"tags"`
	ClientToken     string          `json:"-"`
}

type CreateEipResult struct {
	Eip string `json:"eip"`
}

type ResizeEipArgs struct {
	NewBandWidthInMbps int    `json:"newBandwidthInMbps"`
	ClientToken        string `json:"-"`
}

type BindEipArgs struct {
	InstanceType string `json:"instanceType"`
	InstanceId   string `json:"instanceId"`
	ClientToken  string `json:"-"`
}

type ListEipArgs struct {
	Eip          string
	InstanceType string
	InstanceId   string
	Marker       string
	MaxKeys      int
	Status       string
}

type ListEipResult struct {
	Marker      string     `json:"marker"`
	MaxKeys     int        `json:"maxKeys"`
	NextMarker  string     `json:"nextMarker"`
	IsTruncated bool       `json:"isTruncated"`
	EipList     []EipModel `json:"eipList"`
}

type EipModel struct {
	Name            string          `json:"name"`
	Eip             string          `json:"eip"`
	Status          string          `json:"status"`
	EipInstanceType string          `json:"eipInstanceType"`
	InstanceType    string          `json:"instanceType"`
	InstanceId      string          `json:"instanceId"`
	ShareGroupId    string          `json:"shareGroupId"`
	BandWidthInMbps int             `json:"bandwidthInMbps"`
	PaymentTiming   string          `json:"paymentTiming"`
	BillingMethod   string          `json:"billingMethod"`
	CreateTime      string          `json:"createTime"`
	ExpireTime      string          `json:"expireTime"`
	Tags            []util.TagModel `json:"tags"`
}

type PurchaseReservedEipArgs struct {
	Billing     *Billing `json:"billing"`
	ClientToken string   `json:"clientToken"`
}
