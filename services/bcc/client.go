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

// client.go - define the client for BOS service

// Package cfc defines the CFC services of BCE. The supported APIs are all defined in sub-package

package bcc

import (
	"encoding/json"

	"github.com/baidubce/bce-sdk-go/auth"
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/services/bcc/api"
)

// Client of BCC service is a kind of BceClient, so derived from BceClient
type Client struct {
	*bce.BceClient
}

func NewClient(ak, sk, endPoint string) (*Client, error) {
	credentials, err := auth.NewBceCredentials(ak, sk)
	if err != nil {
		return nil, err
	}
	defaultSignOptions := &auth.SignOptions{
		HeadersToSign: auth.DEFAULT_HEADERS_TO_SIGN,
		ExpireSeconds: auth.DEFAULT_EXPIRE_SECONDS}
	defaultConf := &bce.BceClientConfiguration{
		Endpoint:                  endPoint,
		Region:                    bce.DEFAULT_REGION,
		UserAgent:                 bce.DEFAULT_USER_AGENT,
		Credentials:               credentials,
		SignOption:                defaultSignOptions,
		Retry:                     bce.DEFAULT_RETRY_POLICY,
		ConnectionTimeoutInMillis: bce.DEFAULT_CONNECTION_TIMEOUT_IN_MILLIS}
	v1Signer := &auth.BceV1Signer{}

	client := &Client{bce.NewBceClient(defaultConf, v1Signer)}
	return client, nil
}

func (c *Client) CreateInstance(args *api.CreateInstanceArgs) (*api.CreateInstanceResult, error) {
	if len(args.AdminPass) > 0 {
		cryptedPass, err := api.Aes128EncryptUseSecreteKey(c.Config.Credentials.SecretAccessKey, args.AdminPass)
		if err != nil {
			return nil, err
		}

		args.AdminPass = cryptedPass
	}

	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return nil, jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return nil, err
	}

	return api.CreateInstance(c, body)
}

func (c *Client) ListInstances(args *api.ListInstanceArgs) (*api.ListInstanceResult, error) {
	return api.ListInstances(c, args)
}

func (c *Client) GetInstanceDetail(instanceId string) (*api.GetInstanceDetailResult, error) {
	return api.GetInstanceDetail(c, instanceId)
}

func (c *Client) DeleteInstance(instanceId string) error {
	return api.DeleteInstance(c, instanceId)
}

func (c *Client) ResizeInstance(instanceId string, args *api.ResizeInstanceArgs) error {
	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return err
	}

	return api.ResizeInstance(c, instanceId, body)
}

func (c *Client) RebuildInstance(instanceId string, args *api.RebuildInstanceArgs) error {
	cryptedPass, err := api.Aes128EncryptUseSecreteKey(c.Config.Credentials.SecretAccessKey, args.AdminPass)
	if err != nil {
		return err
	}
	args.AdminPass = cryptedPass

	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return err
	}

	return api.RebuildInstance(c, instanceId, body)
}

func (c *Client) StartInstance(instanceId string) error {
	return api.StartInstance(c, instanceId)
}

func (c *Client) StopInstance(instanceId string, forceStop bool) error {
	args := &api.StopInstanceArgs{
		ForceStop: forceStop,
	}

	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return err
	}

	return api.StopInstance(c, instanceId, body)
}

func (c *Client) RebootInstance(instanceId string, forceStop bool) error {
	args := &api.StopInstanceArgs{
		ForceStop: forceStop,
	}

	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return err
	}

	return api.RebootInstance(c, instanceId, body)
}

func (c *Client) ChangeInstancePass(instanceId string, args *api.ChangeInstancePassArgs) error {
	cryptedPass, err := api.Aes128EncryptUseSecreteKey(c.Config.Credentials.SecretAccessKey, args.AdminPass)
	if err != nil {
		return err
	}
	args.AdminPass = cryptedPass

	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return err
	}

	return api.ChangeInstancePass(c, instanceId, body)
}

func (c *Client) ModifyInstanceAttribute(instanceId string, args *api.ModifyInstanceAttributeArgs) error {
	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return err
	}

	return api.ModifyInstanceAttribute(c, instanceId, body)
}

func (c *Client) BindSecurityGroup(instanceId string, securityGroupId string) error {
	args := &api.BindSecurityGroupArgs{
		SecurityGroupId: securityGroupId,
	}

	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return err
	}

	return api.BindSecurityGroup(c, instanceId, body)
}

func (c *Client) UnBindSecurityGroup(instanceId string, securityGroupId string) error {
	args := &api.BindSecurityGroupArgs{
		SecurityGroupId: securityGroupId,
	}

	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return err
	}

	return api.UnBindSecurityGroup(c, instanceId, body)
}

func (c *Client) GetInstanceVNC(instanceId string) (*api.GetInstanceVNCResult, error) {
	return api.GetInstanceVNC(c, instanceId)
}

func (c *Client) InstancePurchaseReserved(instanceId string, args *api.PurchaseReservedArgs) error {
	// 防止用户误设置该参数，将参数设置为空
	args.Billing.PaymentTiming = ""

	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return err
	}

	return api.InstancePurchaseReserved(c, instanceId, body)
}

//cds sdk
func (c *Client) CreateCDSVolume(args *api.CreateCDSVolumeArgs) (*api.CreateCDSVolumeResult, error) {
	return api.CreateCDSVolume(c, args)
}

func (c *Client) ListCDSVolume(queryArgs *api.ListCDSVolumeArgs) (*api.ListCDSVolumeResult, error) {
	return api.ListCDSVolume(c, queryArgs)
}

func (c *Client) GetCDSVolumeDetail(volumeId string) (*api.GetVolumeDetailResult, error) {
	return api.GetCDSVolumeDetail(c, volumeId)
}

func (c *Client) AttachCDSVolume(volumeId string, args *api.AttachVolumeArgs) (*api.AttachVolumeResult, error) {
	return api.AttachCDSVolume(c, volumeId, args)
}

func (c *Client) DetachCDSVolume(volumeId string, args *api.DetachVolumeArgs) error {
	return api.DetachCDSVolume(c, volumeId, args)
}

func (c *Client) DeleteCDSVolume(volumeId string) error {
	return api.DeleteCDSVolume(c, volumeId)
}

func (c *Client) DeleteCDSVolumeNew(volumeId string, args *api.DeleteCSDVolumeArgs) error {
	return api.DeleteCDSVolumeNew(c, volumeId, args)
}

func (c *Client) ResizeCDSVolume(volumeId string, args *api.ResizeCSDVolumeArgs) error {
	return api.ResizeCDSVolume(c, volumeId, args)
}

func (c *Client) RollbackCDSVolume(volumeId string, args *api.RollbackCSDVolumeArgs) error {
	return api.RollbackCDSVolume(c, volumeId, args)
}

func (c *Client) PurchaseReservedCDSVolume(volumeId string, args *api.PurchaseReservedCSDVolumeArgs) error {
	return api.PurchaseReservedCDSVolume(c, volumeId, args)
}

func (c *Client) RenameCDSVolume(volumeId string, args *api.RenameCSDVolumeArgs) error {
	return api.RenameCDSVolume(c, volumeId, args)
}

func (c *Client) ModifyCDSVolume(volumeId string, args *api.ModifyCSDVolumeArgs) error {
	return api.ModifyCDSVolume(c, volumeId, args)
}

func (c *Client) ModifyChargeTypeCDSVolume(volumeId string, args *api.ModifyChargeTypeCSDVolumeArgs) error {
	return api.ModifyChargeTypeCDSVolume(c, volumeId, args)
}

//securityGroup sdk
func (c *Client) CreateSecurityGroup(args *api.CreateSecurityGroupArgs) (*api.CreateSecurityGroupResult, error) {
	return api.CreateSecurityGroup(c, args)
}

func (c *Client) ListSecurityGroup(queryArgs *api.ListSecurityGroupArgs) (*api.ListSecurityGroupResult, error) {
	return api.ListSecurityGroup(c, queryArgs)
}

func (c *Client) AuthorizeSecurityGroupRule(securityGroupId string, args *api.AuthorizeSecurityGroupArgs) error {
	return api.AuthorizeSecurityGroupRule(c, securityGroupId, args)
}

func (c *Client) RevokeSecurityGroupRule(securityGroupId string, args *api.RevokeSecurityGroupArgs) error {
	return api.RevokeSecurityGroupRule(c, securityGroupId, args)
}

func (c *Client) DeleteSecurityGroupRule(securityGroupId string) error {
	return api.DeleteSecurityGroupRule(c, securityGroupId)
}

//image sdk
func (c *Client) CreateImage(args *api.CreateImageArgs) (*api.CreateImageResult, error) {
	return api.CreateImage(c, args)
}

func (c *Client) ListImage(queryArgs *api.ListImageArgs) (*api.ListImageResult, error) {
	return api.ListImage(c, queryArgs)
}

func (c *Client) GetImageDetail(imageId string) (*api.GetImageDetailResult, error) {
	return api.GetImageDetail(c, imageId)
}

func (c *Client) DeleteImage(imageId string) error {
	return api.DeleteImage(c, imageId)
}

func (c *Client) RemoteCopyImage(imageId string, args *api.RemoteCopyImageArgs) error {
	return api.RemoteCopyImage(c, imageId, args)
}

func (c *Client) CancelRemoteCopyImage(imageId string) error {
	return api.CancelRemoteCopyImage(c, imageId)
}

func (c *Client) ShareImage(imageId string, args *api.SharedUser) error {
	return api.ShareImage(c, imageId, args)
}

func (c *Client) UnShareImage(imageId string, args *api.SharedUser) error {
	return api.UnShareImage(c, imageId, args)
}

func (c *Client) GetImageSharedUser(imageId string) (*api.GetImageSharedUserResult, error) {
	return api.GetImageSharedUser(c, imageId)
}

func (c *Client) GetImageOS(args *api.GetImageOsArgs) (*api.GetImageOsResult, error) {
	return api.GetImageOS(c, args)
}

func (c *Client) CreateSnapshot(args *api.CreateSnapshotArgs) (*api.CreateSnapshotResult, error) {
	return api.CreateSnapshot(c, args)
}

func (c *Client) ListSnapshot(args *api.ListSnapshotArgs) (*api.ListSnapshotResult, error) {
	return api.ListSnapshot(c, args)
}

func (c *Client) GetSnapshotDetail(snapshotId string) (*api.GetSnapshotDetailResult, error) {
	return api.GetSnapshotDetail(c, snapshotId)
}

func (c *Client) DeleteSnapshot(snapshotId string) error {
	return api.DeleteSnapshot(c, snapshotId)
}

func (c *Client) CreateAutoSnapshotPolicy(args *api.CreateASPArgs) (*api.CreateASPResult, error) {
	return api.CreateAutoSnapshotPolicy(c, args)
}

func (c *Client) AttachAutoSnapshotPolicy(aspId string, args *api.AttachASPArgs) error {
	return api.AttachAutoSnapshotPolicy(c, aspId, args)
}

func (c *Client) DetachAutoSnapshotPolicy(aspId string, args *api.DetachASPArgs) error {
	return api.DetachAutoSnapshotPolicy(c, aspId, args)
}

func (c *Client) DeleteAutoSnapshotPolicy(aspId string) error {
	return api.DeleteAutoSnapshotPolicy(c, aspId)
}

func (c *Client) ListAutoSnapshotPolicy(args *api.ListASPArgs) (*api.ListASPResult, error) {
	return api.ListAutoSnapshotPolicy(c, args)
}

func (c *Client) GetAutoSnapshotPolicy(aspId string) (*api.GetASPDetailResult, error) {
	return api.GetAutoSnapshotPolicyDetail(c, aspId)
}

func (c *Client) ListInstanceType() (*api.ListInstanceTypeResult, error) {
	return api.ListInstanceType(c)
}

func (c *Client) ListZone() (*api.ListZoneResult, error) {
	return api.ListZone(c)
}