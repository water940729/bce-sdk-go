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

// json_schema.go - Define json schema used by the json request and response for BOS service
// If the struct is used in request, its name suffix is "Input", and if it is used in response its 
// name suffix is "Output". For example "BucketAclInput" is used in the PutBucketAcl request, and
// "BucketAclOutput" is used in the GetBucketAcl response. If the request and response shares the
// same json structure, the suffix will be "InOut" such as "BucketLifecycleInOut".

package model

// Acl schema
type OwnerType struct {
    Id string `json:"id"`
}

type GranteeType struct {
    Id string `json:"id"`
}

type GrantType struct {
    Grantee    []GranteeType `json:"grantee"`
    Permission []string      `json:"permission"`
}

type BucketAclInput struct {
    AccessControlList []GrantType `json:"accessControlList"`
}

type BucketAclOutput struct {
    AccessControlList []GrantType `json:"accessControlList"`
    Owner             OwnerType   `json:"owner"`
}

// Location schema
type LocationInOut struct {
    LocationConstraint string `json:"locationConstraint"`
}

// List buckets schema
type BucketOwnerType struct {
    Id          string `json:"id"`
    DisplayName string `json:"displayName"`
}

type BucketSummary struct {
    Name         string `json:"name"`
    Location     string `json:"location"`
    CreationDate string `json:"creationDate"`
}

type ListBucketsOutput struct {
    Owner   BucketOwnerType `json:"owner"`
    Buckets []BucketSummary `json:"buckets"`
}

// List objects schema
type ObjectOwnerType struct {
    Id          string `json:"id"`
    DisplayName string `json:"displayName"`
}

type ObjectSummary struct {
    Key          string `json:"key"`
    LastModified string `json:"lastModified"`
    ETag         string `json:"eTag"`
    Size         int    `json:"size"`
    StorageClass string `json:"storageClass"`
    Owner        ObjectOwnerType `json:"owner"`
}

type PrefixType struct {
    Prefix string `json:"prefix"`
}

type ListObjectsOutput struct {
    Name           string `json:"name"`
    Prefix         string `json:"prefix"`
    Delimiter      string `json:"delimiter"`
    Marker         string `json:"marker"`
    NextMarker     string `json:"nextMarker,omitempty"`
    MaxKeys        int    `json:"maxKeys"`
    IsTruncated    bool   `json:"isTruncated"`
    Contents       []ObjectSummary `json:"contents"`
    CommonPrefixes []PrefixType    `json:"commonPrefixes"`
}

// Bucket logging schema
type BucketLoggingInput struct {
    TargetBucket string `json:"targetBucket"`
    TargetPrefix string `json:"targetPrefix"`
}

type BucketLoggingOutput struct {
    Status       string `json:"status"`
    TargetBucket string `json:"targetBucket,omitempty"`
    TargetPrefix string `json:targetPrefix, omitempty"`
}

// Bucket storage class schema
type BucketStorageClassInOut struct {
    StorageClass string `json:"storageClass"`
}

// Bucket lifecycle schema
type LifecycleConditionTimeType struct {
    DateGreaterThan string `json:"dateGreaterThan"`
}

type LifecycleConditionType struct {
    Time LifecycleConditionTimeType `json:"time"`
}

type LifecycleActionType struct {
    Name         string `json:"name"`
    StorageClass string `json:"storageClass,omitempty"`
}

type LifecycleRuleType struct {
    Id        string   `json:"id"`
    Status    string   `json:"status"`
    Resource  []string `json:"resource"`
    Condition LifecycleConditionType `json:"condition"`
    Action    LifecycleActionType    `json:"action"`
}

type BucketLifecycleInOut struct {
    Rule []LifecycleRuleType `json:"rule"`
}

// Object copy json schema
type CopyObjectOutput struct {
    LastModified string `json:"lastModified"`
    ETag         string `json:"eTag"`
}

// Fetch object json schema
type FetchObjectOutput struct {
    Code      string `json:"code"`
    Message   string `json:"message"`
    RequestId string `json:"requestId"`
    JobId     string `json:"jobId"`
}

// Delete multiple objects json schema
type DeleteObjectInput struct {
    Key string `json:"key"`
}

type DeleteMultipleObjectsInput struct {
    Objects []DeleteObjectInput `json:"objects"`
}

type DeleteFailedOutput struct {
    Key     string `json:"key"`
    Code    string `json:"code"`
    Message string `json:"message"`
}

type DeleteMultipleObjectsOutput struct {
    Errors []DeleteFailedOutput `json:"errors"`
}

// Multipart upload json schema
type InitiateMultipartUploadOutput struct {
    Bucket   string `json:"bucket"`
    Key      string `json:"key"`
    UploadId string `json:"uploadId"`
}

type CompleteUploadInput struct {
    PartNumber int    `json:"partNumber"`
    ETag       string `json:"eTag"`
}

type CompleteMultipartUploadInput struct {
    Parts []CompleteUploadInput `json:"parts"`
}

type CompleteMultipartUploadOutput struct {
    Location string `json:"location"`
    Bucket   string `json:"bucket"`
    Key      string `json:"key"`
    ETag     string `json:"eTag"`
}

// ListParts json schema
type ListPartType struct {
    PartNumber   string `json:"partNumber"`
    LastModified string `json:"lastModified"`
    ETag         string `json:"ETag"`
    Size         int    `json:"size"`
}

type ListPartsOutput struct {
    Bucket   string `json:"bucket"`
    Key      string `json:"key"`
    UploadId string `json:"uploadId"`
    Owner    BucketOwnerType `json:"owner"`
    PartNumberMarker     string `json:"partNumberMarker"`
    NextPartNumberMarker int `json:"nextPartNumberMarker"`
    MaxParts    int `json:"maxParts"`
    IsTruncated bool `json:"isTruncated"`
    Parts       []ListPartType `json:"parts"`
}

// List multipart uploads json schema
type ListMultipartUploadsType struct {
    Key          string `json:"key"`
    UploadId     string `json:"uploadId"`
    Owner        ObjectOwnerType `json:"owner"`
    Initiated    string `json:"initiated"`
    StorageClass string `json:"storageClass"`
}

type ListMultipartUploadsOutput struct {
    Bucket         string `json:"bucket"`
    CommonPrefixes string `json:"commonPrefixes"`
    Delimiter      string `json:"delimiter"`
    Prefix         string `json:"prefix"`
    IsTruncated    bool `json:"isTruncated"`
    KeyMarker      string `json:"keyMarker"`
    MaxUploads     int `json:"maxUploads"`
    NextKeyMarker  string `json:"nextKeyMarker"`
    Uploads        []ListMultipartUploadsType `json:"uploads"`
}

