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

package api

type InstanceStatus string

const (
	InstanceStatusRunning            InstanceStatus = "Running"
	InstanceStatusStarting           InstanceStatus = "Starting"
	InstanceStatusStopping           InstanceStatus = "Stopping"
	InstanceStatusStopped            InstanceStatus = "Stopped"
	InstanceStatusDeleted            InstanceStatus = "Deleted"
	InstanceStatusScaling            InstanceStatus = "Scaling"
	InstanceStatusExpired            InstanceStatus = "Expired"
	InstanceStatusError              InstanceStatus = "Error"
	InstanceStatusSnapshotProcessing InstanceStatus = "SnapshotProcessing"
	InstanceStatusImageProcessing    InstanceStatus = "ImageProcessing"
)

type InstanceType string

const (
	InstanceTypeN1 InstanceType = "N1"
	InstanceTypeN2 InstanceType = "N2"
	InstanceTypeN3 InstanceType = "N3"
	InstanceTypeC1 InstanceType = "C1"
	InstanceTypeC2 InstanceType = "C2"
	InstanceTypeS1 InstanceType = "S1"
	InstanceTypeG1 InstanceType = "G1"
	InstanceTypeF1 InstanceType = "F1"
)

type StorageType string

const (
	StorageTypeStd1     StorageType = "std1"
	StorageTypeHP1      StorageType = "hp1"
	StorageTypeCloudHP1 StorageType = "cloud_hp1"
	StorageTypeLocal    StorageType = "local"
	StorageTypeSATA     StorageType = "sata"
	StorageTypeSSD      StorageType = "ssd"
)

type PaymentTimingType string

const (
	PaymentTimingPrePaid  PaymentTimingType = "Prepaid"
	PaymentTimingPostPaid PaymentTimingType = "Postpaid"
)

// Instance define instance model
type InstanceModel struct {
	InstanceId            string         `json:"id"`
	InstanceName          string         `json:"name"`
	InstanceType          InstanceType   `json:"instanceType"`
	Description           string         `json:"desc"`
	Status                InstanceStatus `json:"status"`
	PaymentTiming         string         `json:"paymentTiming"`
	CreationTime          string         `json:"createTime"`
	ExpireTime            string         `json:"expireTime"`
	PublicIP              string         `json:"publicIp"`
	InternalIP            string         `json:"internalIp"`
	CpuCount              int            `json:"cpuCount"`
	GpuCount              int            `json:"gpuCount"`
	MemoryCapacityInGB    int            `json:"memoryCapacityInGB"`
	LocalDiskSizeInGB     int            `json:"localDiskSizeInGB"`
	ImageId               string         `json:"imageId"`
	NetworkCapacityInMbps int            `json:"networkCapacityInMbps"`
	PlacementPolicy       string         `json:"placementPolicy"`
	ZoneName              string         `json:"zoneName"`
	SubnetId              string         `json:"subnetId"`
	VpcId                 string         `json:"vpcId"`
}

type Reservation struct {
	ReservationLength   int    `json:"reservationLength"`
	ReservationTimeUnit string `json:"reservationTimeUnit"`
}

type Billing struct {
	PaymentTiming PaymentTimingType `json:"paymentTiming,omitempty"`
	Reservation   *Reservation      `json:"reservation,omitempty"`
}

type EphemeralDisk struct {
	StorageType  StorageType `json:"storageType"`
	SizeInGB     int         `json:"sizeInGB"`
	FreeSizeInGB int         `json:"freeSizeInGB"`
}

type CreateCdsModel struct {
	CdsSizeInGB int         `json:"cdsSizeInGB"`
	StorageType StorageType `json:"storageType"`
	SnapShotId  string      `json:"snapshotId,omitempty"`
}

type CreateInstanceArgs struct {
	ImageId               string           `json:"imageId"`
	Billing               Billing          `json:"billing"`
	InstanceType          InstanceType     `json:"instanceType,omitempty"`
	CpuCount              int              `json:"cpuCount"`
	MemoryCapacityInGB    int              `json:"memoryCapacityInGB"`
	RootDiskSizeInGb      int              `json:"rootDiskSizeInGb,omitempty"`
	RootDiskStorageType   StorageType      `json:"rootDiskStorageType,omitempty"`
	LocalDiskSizeInGB     int              `json:"localDiskSizeInGB,omitempty"`
	EphemeralDisks        []EphemeralDisk  `json:"ephemeralDisks,omitempty"`
	CreateCdsList         []CreateCdsModel `json:"creatCdsList,omitempty"`
	NetWorkCapacityInMbps int              `json:"networkCapacityInMbps,omitempty"`
	DedicateHostId        int              `json:"dedicatedHostId,omitempty"`
	PurchaseCount         int              `json:"purchaseCount,omitempty"`
	Name                  string           `json:"name,omitempty"`
	AdminPass             string           `json:"adminPass,omitempty"`
	ZoneName              string           `json:"zoneName,omitempty"`
	SubnetId              string           `json:"subnetId,omitempty"`
	SecurityGroupId       string           `json:"securityGroupId,omitempty"`
	GpuCard               string           `json:"gpuCard,omitempty"`
	FpgaCard              string           `json:"fpgaCard,omitempty"`
	CardCount             string           `json:"cardCount,omitempty"`
}

type CreateInstanceResult struct {
	InstanceIds []string `json:"instanceIds"`
}

type ListInstanceArgs struct {
	Marker          string
	MaxKeys         int
	InternalIp      string
	DedicatedHostId string
	ZoneName        string
}

type ListInstanceResult struct {
	Marker      string          `json:"marker"`
	IsTruncated bool            `json:"isTruncated"`
	NextMarker  string          `json:"nextMarker"`
	MaxKeys     int             `json:"maxKeys"`
	Instances   []InstanceModel `json:"instances"`
}

type GetInstanceDetailResult struct {
	Instance InstanceModel `json:"instance"`
}

type ResizeInstanceArgs struct {
	CpuCount           int             `json:"cpuCount"`
	MemoryCapacityInGB int             `json:"memoryCapacityInGB"`
	EphemeralDisks     []EphemeralDisk `json:"ephemeralDisks,omitempty"`
}

type RebuildInstanceArgs struct {
	ImageId   string `json:"imageId"`
	AdminPass string `json:"adminPass"`
}

type StopInstanceArgs struct {
	ForceStop bool `json:"forceStop"`
}

type ChangeInstancePassArgs struct {
	AdminPass string `json:"adminPass"`
}

type ModifyInstanceAttributeArgs struct {
	Name string `json:"name"`
}

type BindSecurityGroupArgs struct {
	SecurityGroupId string `json:"securityGroupId"`
}

type GetInstanceVNCResult struct {
	VNCUrl string `json:"vncUrl"`
}

type PurchaseReservedArgs struct {
	Billing Billing `json:"billing"`
}

type VolumeStatus string

const (
	VolumeStatusAVAILABLE          VolumeStatus = "Available"
	VolumeStatusINUSE              VolumeStatus = "InUse"
	VolumeStatusSNAPSHOTPROCESSING VolumeStatus = "SnapshotProcessing"
	VolumeStatusRECHARGING         VolumeStatus = "Recharging"
	VolumeStatusDETACHING          VolumeStatus = "Detaching"
	VolumeStatusDELETING           VolumeStatus = "Deleting"
	VolumeStatusEXPIRED            VolumeStatus = "Expired"
	VolumeStatusNOTAVAILABLE       VolumeStatus = "NotAvailable"
	VolumeStatusDELETED            VolumeStatus = "Deleted"
	VolumeStatusSCALING            VolumeStatus = "Scaling"
	VolumeStatusIMAGEPROCESSING    VolumeStatus = "ImageProcessing"
	VolumeStatusCREATING           VolumeStatus = "Creating"
	VolumeStatusATTACHING          VolumeStatus = "Attaching"
	VolumeStatusERROR              VolumeStatus = "Error"
)

type VolumeType string

const (
	VolumeTypeSYSTEM    VolumeType = "System"
	VolumeTypeEPHEMERAL VolumeType = "Ephemeral"
	VolumeTypeCDS       VolumeType = "Cds"
)

type RenameCSDVolumeArgs struct {
	Name string `json:"name"`
}

type ModifyCSDVolumeArgs struct {
	CdsName string `json:"cdsName,omitempty"`
	Desc    string `json:"desc,omitempty"`
}

type DetachVolumeArgs struct {
	InstanceId string `json:"instanceId"`
}

type PurchaseReservedCSDVolumeArgs struct {
	Billing *Billing `json:"billing"`
}

type DeleteCSDVolumeArgs struct {
	ManualSnapshot string `json:"manualSnapshot,omitempty"`
	AutoSnapshot   string `json:"autoSnapshot,omitempty"`
}

type ModifyChargeTypeCSDVolumeArgs struct {
	Billing string `json:"billing"`
}

type ListCDSVolumeResult struct {
	Marker      string         `json:"marker"`
	IsTruncated bool           `json:"isTruncated"`
	NextMarker  string         `json:"nextMarker"`
	MaxKeys     int            `json:"maxKeys"`
	Volumes     *[]VolumeModel `json:"volumes"`
}

type VolumeModel struct {
	Type               VolumeType               `json:"type"`
	StorageType        StorageType              `json:"storageType"`
	Id                 string                   `json:"id"`
	Name               string                   `json:"name"`
	DiskSizeInGB       int                      `json:"diskSizeInGB"`
	PaymentTiming      string                   `json:"paymentTiming"`
	ExpireTime         string                   `json:"expireTime"`
	Status             VolumeStatus             `json:"status"`
	Desc               string                   `json:"desc"`
	Attachments        *[]VolumeAttachmentModel `json:"attachments"`
	ZoneName           string                   `json:"zoneName"`
	AutoSnapshotPolicy *AutoSnapshotPolicyModel `json:"autoSnapshotPolicy"`
	CreateTime         string                   `json:"createTime"`
}

type VolumeAttachmentModel struct {
	VolumeId   string `json:"volumeId"`
	InstanceId string `json:"instanceId"`
	Device     string `json:"device"`
	Serial     string `json:"serial"`
}

type AttachVolumeResult struct {
	VolumeAttachment *VolumeAttachmentModel `json:"volumeAttachment"`
}

type CreateCDSVolumeArgs struct {
	SnapshotId    string      `json:"snapshotId,omitempty"`
	ZoneName      string      `json:"zoneName,omitempty"`
	PurchaseCount int         `json:"purchaseCount,omitempty"`
	CdsSizeInGB   int         `json:"cdsSizeInGB,omitempty"`
	StorageType   StorageType `json:"storageType,omitempty"`
	Billing       *Billing    `json:"billing"`
}

type CreateCDSVolumeResult struct {
	VolumeIds *[]string `json:"volumeIds"`
}

type GetVolumeDetailResult struct {
	Volume *VolumeModel `json:"volume"`
}

type AttachVolumeArgs struct {
	InstanceId string `json:"instanceId"`
}

type ResizeCSDVolumeArgs struct {
	NewCdsSizeInGB int `json:"newCdsSizeInGB"`
}

type RollbackCSDVolumeArgs struct {
	NewCdsSizeInGB string `json:"newCdsSizeInGB"`
}

type ListCDSVolumeArgs struct {
	MaxKeys    int
	InstanceId string
	ZoneName   string
	Marker     string
}

type AutoSnapshotPolicyModel struct {
	CreatedTime     string `json:"createdTime"`
	Id              string `json:"id"`
	Status          string `json:"status"`
	RetentionDays   int    `json:"retentionDays"`
	UpdatedTime     string `json:"updatedTime"`
	DeletedTime     string `json:"deletedTime"`
	LastExecuteTime string `json:"lastExecuteTime"`
	VolumeCount     int    `json:"volumeCount"`
	Name            string `json:"name"`
	TimePoints      *[]int `json:"timePoints"`
	RepeatWeekdays  *[]int `json:"repeatWeekdays"`
}
