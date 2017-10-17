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

// Package bos defines the BOS services of BCE. The supported APIs are all defined in sub-package
// model with three types: 16 bucket APIs, 9 object APIs and 7 multipart APIs.
package bos

import (
    "encoding/json"
    "fmt"
    "io"
    "os"

    "baidubce/auth"
    "baidubce/bce"
    "baidubce/http"
    "baidubce/util"
    "baidubce/services/bos/api"
)

const (
    DEFAULT_MAX_PARALLEL   = 10
    MULTIPART_ALIGN        = 1 << 20        // 1MB
    MIN_MULTIPART_SIZE     = 5 * (1 << 20)  // 5MB
    DEFAULT_MULTIPART_SIZE = 10 * (1 << 20) // 10MB

    MAX_PART_NUMBER        = 10000
    MAX_SINGLE_PART_SIZE   = 5 * (1 << 30) // 5GB
    MAX_SINGLE_OBJECT_SIZE = 5 * (1 << 40) // 5TB
)

// Client of BOS service is a kind of BceClient, so derived from BceClient
type Client struct {
    *bce.BceClient

    // Fileds that used in parallel operation for BOS service
    MaxParallel   int64
    MultipartSize int64
}

// NewClient make the BOS service client with default configuration.
// Use `cli.Config.xxx` to access the config or change it to non-default value.
func NewClient(ak, sk string) (*Client, error) {
    credentials, err := auth.NewBceCredentials(ak, sk)
    if err != nil {
        return nil, err
    }
    defaultSignOptions := &auth.SignOptions{
        auth.DEFAULT_HEADERS_TO_SIGN,
        util.NowUTCSeconds(),
        auth.DEFAULT_EXPIRE_SECONDS}
    defaultConf := &bce.BceClientConfiguration{
        Endpoint: fmt.Sprintf("%s.%s", bce.DEFAULT_REGION, bce.DEFAULT_SERVICE_DOMAIN),
        Region: bce.DEFAULT_REGION,
        UserAgent: bce.DEFAULT_USER_AGENT,
        Credentials: credentials,
        SignOption: defaultSignOptions,
        Retry: bce.DEFAULT_RETRY_POLICY,
        ConnectionTimeoutInMillis: bce.DEFAULT_CONNECTION_TIMEOUT_IN_MILLIS}
    v1Signer := &auth.BceV1Signer{}

    client := &Client{bce.NewBceClient(defaultConf, v1Signer),
        DEFAULT_MAX_PARALLEL, DEFAULT_MULTIPART_SIZE}
    return client, nil
}

/*
 * ListBuckets - list all buckets
 *
 * RETURNS:
 *     - *api.ListBucketsResult: the all buckets
 *     - error: the uploaded error if any occurs
 */
func (cli *Client) ListBuckets() (*api.ListBucketsResult, error) {
    return api.ListBuckets(cli)
}

/*
 * ListObjects - list all objects of the given bucket
 *
 * PARAMS:
 *     - bucket: the bucket name
 *     - args: the optional arguments to list objects
 * RETURNS:
 *     - *api.ListObjectsResult: the all objects of the bucket
 *     - error: the uploaded error if any occurs
 */
func (cli *Client) ListObjects(bucket string,
        args *api.ListObjectsArgs) (*api.ListObjectsResult, error) {
    return api.ListObjects(cli, bucket, args)
}

/*
 * HeadBucket - test the given bucket existed and access authority
 *
 * PARAMS:
 *     - bucket: the bucket name
 * RETURNS:
 *     - error: nil if exists and have authority otherwise the specific error
 */
func (cli *Client) HeadBucket(bucket string) error {
    return api.HeadBucket(cli, bucket)
}

/*
 * PutBucket - create a new bucket
 *
 * PARAMS:
 *     - bucket: the new bucket name
 * RETURNS:
 *     - string: the location of the new bucket if create success
 *     - error: nil if create success otherwise the specific error
 */
func (cli *Client) PutBucket(bucket string) (string, error) {
    return api.PutBucket(cli, bucket)
}

/*
 * DeleteBucket - delete a empty bucket
 *
 * PARAMS:
 *     - bucket: the bucket name to be deleted
 * RETURNS:
 *     - error: nil if delete success otherwise the specific error
 */
func (cli *Client) DeleteBucket(bucket string) error {
    return api.DeleteBucket(cli, bucket)
}

/*
 * GetBucketLocation - get the location fo the given bucket
 *
 * PARAMS:
 *     - bucket: the bucket name
 * RETURNS:
 *     - string: the location of the bucket
 *     - error: nil if success otherwise the specific error
 */
func (cli *Client) GetBucketLocation(bucket string) (string, error) {
    return api.GetBucketLocation(cli, bucket)
}

/*
 * PutBucketAcl - set the acl of the given bucket with acl json file stream
 *
 * PARAMS:
 *     - bucket: the bucket name
 *     - aclStream: the acl json file stream
 * RETURNS:
 *     - error: nil if success otherwise the specific error
 */
func (cli *Client) PutBucketAcl(bucket string, aclStream *http.BodyStream) error {
    return api.PutBucketAcl(cli, bucket, "", aclStream)
}

/*
 * PutBucketAclFromCanned - set the canned acl of the given bucket
 *
 * PARAMS:
 *     - bucket: the bucket name
 *     - cannedAcl: the cannedAcl string
 * RETURNS:
 *     - error: nil if success otherwise the specific error
 */
func (cli *Client) PutBucketAclFromCanned(bucket, cannedAcl string) error {
    return api.PutBucketAcl(cli, bucket, cannedAcl, nil)
}

/*
 * PutBucketAclFromFile - set the acl of the given bucket with acl json file name
 *
 * PARAMS:
 *     - bucket: the bucket name
 *     - aclFile: the acl file name
 * RETURNS:
 *     - error: nil if success otherwise the specific error
 */
func (cli *Client) PutBucketAclFromFile(bucket, aclFile string) error {
    stream, err := http.NewBodyStreamFromFile(aclFile)
    if err != nil {
        return err
    }
    return api.PutBucketAcl(cli, bucket, "", stream)
}

/*
 * PutBucketAclFromStruct - set the acl of the given bucket with acl data structure
 *
 * PARAMS:
 *     - bucket: the bucket name
 *     - aclObj: the acl struct object
 * RETURNS:
 *     - error: nil if success otherwise the specific error
 */
func (cli *Client) PutBucketAclFromStruct(bucket string, aclObj *api.PutBucketAclArgs) error {
    jsonBytes, jsonErr := json.Marshal(aclObj)
    if jsonErr != nil {
        return jsonErr
    }
    stream := http.NewBodyStreamFromBytes(jsonBytes)
    return api.PutBucketAcl(cli, bucket, "", stream)
}

/*
 * GetBucketAcl - get the acl of the given bucket
 *
 * PARAMS:
 *     - bucket: the bucket name
 * RETURNS:
 *     - *api.GetBucketAclResult: the result of the bucket acl
 *     - error: nil if success otherwise the specific error
 */
func (cli *Client) GetBucketAcl(bucket string) (*api.GetBucketAclResult, error) {
    return api.GetBucketAcl(cli, bucket)
}

/*
 * PutBucketLogging - set the loging setting of the given bucket with json stream
 *
 * PARAMS:
 *     - bucket: the bucket name
 *     - stream: the json stream
 * RETURNS:
 *     - error: nil if success otherwise the specific error
 */
func (cli *Client) PutBucketLogging(bucket string, stream *http.BodyStream) error {
    return api.PutBucketLogging(cli, bucket, stream)
}

/*
 * PutBucketLoggingFromString - set the loging setting of the given bucket with json string
 *
 * PARAMS:
 *     - bucket: the bucket name
 *     - logging: the json format string
 * RETURNS:
 *     - error: nil if success otherwise the specific error
 */
func (cli *Client) PutBucketLoggingFromString(bucket, logging string) error {
    stream := http.NewBodyStreamFromString(logging)
    return api.PutBucketLogging(cli, bucket, stream)
}

/*
 * PutBucketLoggingFromStruct - set the loging setting of the given bucket with args object
 *
 * PARAMS:
 *     - bucket: the bucket name
 *     - obj: the logging setting object
 * RETURNS:
 *     - error: nil if success otherwise the specific error
 */
func (cli *Client) PutBucketLoggingFromStruct(bucket string, obj *api.PutBucketLoggingArgs) error {
    jsonBytes, jsonErr := json.Marshal(obj)
    if jsonErr != nil {
        return jsonErr
    }
    stream := http.NewBodyStreamFromBytes(jsonBytes)
    return api.PutBucketLogging(cli, bucket, stream)
}

/*
 * GetBucketLogging - get the logging setting of the given bucket
 *
 * PARAMS:
 *     - bucket: the bucket name
 * RETURNS:
 *     - *api.GetBucketLoggingResult: the logging setting of the bucket
 *     - error: nil if success otherwise the specific error
 */
func (cli *Client) GetBucketLogging(bucket string) (*api.GetBucketLoggingResult, error) {
    return api.GetBucketLogging(cli, bucket)
}

/*
 * DeleteBucketLogging - delete the logging setting of the given bucket
 *
 * PARAMS:
 *     - bucket: the bucket name
 * RETURNS:
 *     - error: nil if success otherwise the specific error
 */
func (cli *Client) DeleteBucketLogging(bucket string) error {
    return api.DeleteBucketLogging(cli, bucket)
}

/*
 * PutBucketLifecycle - set the lifecycle rule of the given bucket with raw stream
 *
 * PARAMS:
 *     - bucket: the bucket name
 *     - lifecycle: the lifecycle rule json stream
 * RETURNS:
 *     - error: nil if success otherwise the specific error
 */
func (cli *Client) PutBucketLifecycle(bucket string, lifecycle *http.BodyStream) error {
    return api.PutBucketLifecycle(cli, bucket, lifecycle)
}

/*
 * PutBucketLifecycleFromString - set the lifecycle rule of the given bucket with string
 *
 * PARAMS:
 *     - bucket: the bucket name
 *     - lifecycle: the lifecycle rule json format string
 * RETURNS:
 *     - error: nil if success otherwise the specific error
 */
func (cli *Client) PutBucketLifecycleFromString(bucket, lifecycle string) error {
    stream := http.NewBodyStreamFromString(lifecycle)
    return api.PutBucketLifecycle(cli, bucket, stream)
}

/*
 * GetBucketLifecycle - get the lifecycle rule of the given bucket
 *
 * PARAMS:
 *     - bucket: the bucket name
 * RETURNS:
 *     - *api.GetBucketLifecycleResult: the lifecycle rule of the bucket
 *     - error: nil if success otherwise the specific error
 */
func (cli *Client) GetBucketLifecycle(bucket string) (*api.GetBucketLifecycleResult, error) {
    return api.GetBucketLifecycle(cli, bucket)
}

/*
 * DeleteBucketLifecycle - delete the lifecycle rule of the given bucket
 *
 * PARAMS:
 *     - bucket: the bucket name
 * RETURNS:
 *     - error: nil if success otherwise the specific error
 */
func (cli *Client) DeleteBucketLifecycle(bucket string) error {
    return api.DeleteBucketLifecycle(cli, bucket)
}

/*
 * PutBucketStorageclass - set the storage class of the given bucket
 *
 * PARAMS:
 *     - bucket: the bucket name
 *     - storageClass: the storage class string value
 * RETURNS:
 *     - error: nil if success otherwise the specific error
 */
func (cli *Client) PutBucketStorageclass(bucket, storageClass string) error {
    return api.PutBucketStorageclass(cli, bucket, storageClass)
}

/*
 * GetBucketStorageclass - get the storage class of the given bucket
 *
 * PARAMS:
 *     - bucket: the bucket name
 * RETURNS:
 *     - string: the storage class string value
 *     - error: nil if success otherwise the specific error
 */
func (cli *Client) GetBucketStorageclass(bucket string) (string, error) {
    return api.GetBucketStorageclass(cli, bucket)
}

/*
 * PutObject - upload a new object or rewrite the existed object with raw stream
 *
 * PARAMS:
 *     - bucket: the name of the bucket to store the object
 *     - object: the name of the object
 *     - stream: the content stream
 * RETURNS:
 *     - string: etag of the uploaded object
 *     - error: the uploaded error if any occurs
 */
func (cli *Client) PutObject(bucket, object string, stream *http.BodyStream) (string, error) {
    return api.PutObject(cli, bucket, object, stream)
}

/*
 * PutObjectFromString - upload a new object or rewrite the existed object from a string
 *
 * PARAMS:
 *     - bucket: the name of the bucket to store the object
 *     - object: the name of the object
 *     - content: the content string
 * RETURNS:
 *     - string: etag of the uploaded object
 *     - error: the uploaded error if any occurs
 */
func (cli *Client) PutObjectFromString(bucket, object, content string) (string, error) {
    body := http.NewBodyStreamFromString(content)
    return api.PutObject(cli, bucket, object, body)
}

/*
 * PutObjectFromFile - upload a new object or rewrite the existed object from a local file
 *
 * PARAMS:
 *     - bucket: the name of the bucket to store the object
 *     - object: the name of the object
 *     - fileName: the local file full path name
 * RETURNS:
 *     - string: etag of the uploaded object
 *     - error: the uploaded error if any occurs
 */
func (cli *Client) PutObjectFromFile(bucket, object, fileName string) (string, error) {
    body, err := http.NewBodyStreamFromFile(fileName)
    if err != nil {
        return "", err
    }
    return api.PutObject(cli, bucket, object, body)
}

/*
 * CopyObject - copy a remote object to another one
 *
 * PARAMS:
 *     - bucket: the name of the destination bucket
 *     - object: the name of the destination object
 *     - srcBucket: the name of the source bucket
 *     - srcObject: the name of the source object
 *     - args: the optional arguments for copying object
 *         MetadataDirective, StorageClass, IfMatch,
 *         IfNoneMatch, ifModifiedSince, IfUnmodifiedSince
 * RETURNS:
 *     - *api.CopyObjectResult: result struct which contains "ETag" and "LastModified" fields
 *     - error: any error if it occurs
 */
func (cli *Client) CopyObject(bucket, object, srcBucket, srcObject string,
        args *api.CopyObjectArgs) (*api.CopyObjectResult, error) {
    source := fmt.Sprintf("/%s/%s", srcBucket, srcObject)
    return api.CopyObject(cli, bucket, object, source, args)
}

/*
 * BasicCopyObject - the basic interface of copying a object to another one
 *
 * PARAMS:
 *     - bucket: the name of the destination bucket
 *     - object: the name of the destination object
 *     - srcBucket: the name of the source bucket
 *     - srcObject: the name of the source object
 * RETURNS:
 *     - *api.CopyObjectResult: result struct which contains "ETag" and "LastModified" fields
 *     - error: any error if it occurs
 */
func (cli *Client) BasicCopyObject(bucket, object, srcBucket,
        srcObject string) (*api.CopyObjectResult, error) {
    source := fmt.Sprintf("/%s/%s", srcBucket, srcObject)
    return api.CopyObject(cli, bucket, object, source, nil)
}

/*
 * GetObject - get the given object with raw stream return
 *
 * PARAMS:
 *     - bucket: the name of the bucket
 *     - object: the name of the object
 *     - args: the optional arguments of getting the object
 * RETURNS:
 *     - *api.CopyObjectResult: result struct which contains "Body" and header fields
 *       for details reference https://cloud.baidu.com/doc/BOS/API.html#GetObject.E6.8E.A5.E5.8F.A3
 *     - error: any error if it occurs
 */
func (cli *Client) GetObject(bucket, object string,
        args *api.GetObjectArgs) (*api.GetObjectResult, error) {
    return api.GetObject(cli, bucket, object, args)
}

/*
 * BasicGetObject - the basic interface of geting the given object
 *
 * PARAMS:
 *     - bucket: the name of the bucket
 *     - object: the name of the object
 * RETURNS:
 *     - *api.CopyObjectResult: result struct which contains "Body" and header fields
 *       for details reference https://cloud.baidu.com/doc/BOS/API.html#GetObject.E6.8E.A5.E5.8F.A3
 *     - error: any error if it occurs
 */
func (cli *Client) BasicGetObject(bucket, object string) (*api.GetObjectResult, error) {
    return api.GetObject(cli, bucket, object, nil)
}

/*
 * SimpleGetObject - get the given object with simple arguments interface
 *
 * PARAMS:
 *     - bucket: the name of the bucket
 *     - object: the name of the object
 *     - rangeStart: start offset to get the object
 *     - rangeEnd: end offset to get the object
 *     - responseHeaders: the user specified headers when return, only support 5 headers:
 *       ContentDisposition, ContentType, ContentLanguage, Expires, CacheControl, ContentEncoding
 * RETURNS:
 *     - *api.CopyObjectResult: result struct which contains "Body" and header fields
 *       for details reference https://cloud.baidu.com/doc/BOS/API.html#GetObject.E6.8E.A5.E5.8F.A3
 *     - error: any error if it occurs
 */
func (cli *Client) SimpleGetObject(bucket, object string, rangeStart, rangeEnd int64,
        responseHeaders map[string]string) (*api.GetObjectResult, error) {
    args := &api.GetObjectArgs{
        rangeStart,
        rangeEnd,
        responseHeaders}
    return api.GetObject(cli, bucket, object, args)
}

/*
 * GetObjectToFile - get the given object to the given file path
 *
 * PARAMS:
 *     - bucket: the name of the bucket
 *     - object: the name of the object
 *     - filePath: the file path to store the object content
 * RETURNS:
 *     - error: any error if it occurs
 */
func (cli *Client) GetObjectToFile(bucket, object, filePath string) error {
    res, err := api.GetObject(cli, bucket, object, nil)
    if err != nil {
        return err
    }
    defer res.Body.Close()

    file, fileErr := os.OpenFile(filePath, os.O_TRUNC | os.O_CREATE | os.O_WRONLY, 0644)
    if fileErr != nil {
        return fileErr
    }
    defer file.Close()

    written, writeErr := io.CopyN(file, res.Body, res.Body.Len())
    if writeErr != nil {
        return writeErr
    }
    if written != res.Body.Len() {
        return fmt.Errorf("written content size does not match the response content")
    }
    return nil
}

/*
 * GetObjectMeta - get the given object metadata
 *
 * PARAMS:
 *     - bucket: the name of the bucket
 *     - object: the name of the object
 * RETURNS:
 *     - *api.GetObjectMetaResult: metadata result, for details reference
 *       https://cloud.baidu.com/doc/BOS/API.html#GetObjectMeta.E6.8E.A5.E5.8F.A3
 *     - error: any error if it occurs
 */
func (cli *Client) GetObjectMeta(bucket, object string) (*api.GetObjectMetaResult, error) {
    return api.GetObjectMeta(cli, bucket, object)
}

/*
 * FetchObject - fetch the object content from the given source and store
 *
 * PARAMS:
 *     - bucket: the name of the bucket to store
 *     - object: the name of the object to store
 *     - source: fetch source url
 *     - args: the optional arguments to fetch the object
 * RETURNS:
 *     - *api.FetchObjectResult: result struct with Code, Message, RequestId and JobId fields
 *     - error: any error if it occurs
 */
func (cli *Client) FetchObject(bucket, object, source string,
        args *api.FetchObjectArgs) (*api.FetchObjectResult, error) {
    return api.FetchObject(cli, bucket, object, source, args)
}

/*
 * BasicFetchObject - the basic interface of the fetch object api
 *
 * PARAMS:
 *     - bucket: the name of the bucket to store
 *     - object: the name of the object to store
 *     - source: fetch source url
 * RETURNS:
 *     - *api.FetchObjectResult: result struct with Code, Message, RequestId and JobId fields
 *     - error: any error if it occurs
 */
func (cli *Client) BasicFetchObject(bucket, object, source string) (*api.FetchObjectResult, error) {
    return api.FetchObject(cli, bucket, object, source, nil)
}

/*
 * AppendObject - append the gievn content to a new or existed object which is appendable
 *
 * PARAMS:
 *     - bucket: the name of the bucket
 *     - object: the name of the object
 *     - content: the append object stream
 *     - args: the optional arguments to append object
 * RETURNS:
 *     - *api.FetchObjectResult: result struct with Code, Message, RequestId and JobId fields
 *     - error: any error if it occurs
 */
func (cli *Client) AppendObject(bucket, object string, content *http.BodyStream,
        args *api.AppendObjectArgs) (*api.AppendObjectResult, error) {
    return api.AppendObject(cli, bucket, object, content, args)
}

/*
 * BasicAppendObject - basic interface to append object
 *
 * PARAMS:
 *     - bucket: the name of the bucket
 *     - object: the name of the object
 *     - content: the append object stream
 * RETURNS:
 *     - *api.FetchObjectResult: result struct with Code, Message, RequestId and JobId fields
 *     - error: any error if it occurs
 */
func (cli *Client) BasicAppendObject(bucket, object string,
        content *http.BodyStream) (*api.AppendObjectResult, error) {
    return api.AppendObject(cli, bucket, object, content, nil)
}

/*
 * BasicAppendObjectFromString - basic interface of appending an object from a string
 *
 * PARAMS:
 *     - bucket: the name of the bucket
 *     - object: the name of the object
 *     - content: the object string to append
 * RETURNS:
 *     - *api.FetchObjectResult: result struct with Code, Message, RequestId and JobId fields
 *     - error: any error if it occurs
 */
func (cli *Client) BasicAppendObjectFromString(bucket, object,
        content string) (*api.AppendObjectResult, error) {
    stream := http.NewBodyStreamFromString(content)
    return api.AppendObject(cli, bucket, object, stream, nil)
}

/*
 * BasicAppendObjectFromFile - basic interface of appending an object from a file
 *
 * PARAMS:
 *     - bucket: the name of the bucket
 *     - object: the name of the object
 *     - filePath: the full file path
 * RETURNS:
 *     - *api.FetchObjectResult: result struct with Code, Message, RequestId and JobId fields
 *     - error: any error if it occurs
 */
func (cli *Client) BasicAppendObjectFromFile(bucket, object,
        filePath string) (*api.AppendObjectResult, error) {
    stream, err := http.NewBodyStreamFromFile(filePath)
    if err != nil {
        return nil, err
    }
    return api.AppendObject(cli, bucket, object, stream, nil)
}

/*
 * DeleteObject - delete the given object
 *
 * PARAMS:
 *     - bucket: the name of the bucket to delete
 *     - object: the name of the object to delete
 * RETURNS:
 *     - error: any error if it occurs
 */
func (cli *Client) DeleteObject(bucket, object string) error {
    return api.DeleteObject(cli, bucket, object)
}

/*
 * DeleteMultipleObjects - delete a list of objects
 *
 * PARAMS:
 *     - bucket: the name of the bucket to delete
 *     - objectListStream: the object list stream to be deleted
 * RETURNS:
 *     - *api.DeleteMultipleObjectsResult: the delete information
 *     - error: any error if it occurs
 */
func (cli *Client) DeleteMultipleObjects(bucket string,
        objectListStream *http.BodyStream) (*api.DeleteMultipleObjectsResult, error) {
    return api.DeleteMultipleObjects(cli, bucket, objectListStream)
}

/*
 * DeleteMultipleObjectsFromString - delete a list of objects with json format string
 *
 * PARAMS:
 *     - bucket: the name of the bucket to delete
 *     - objectListString: the object list string to be deleted
 * RETURNS:
 *     - error: any error if it occurs
 */
func (cli *Client) DeleteMultipleObjectsFromString(bucket,
        objectListString string) (*api.DeleteMultipleObjectsResult, error) {
    stream := http.NewBodyStreamFromString(objectListString)
    return api.DeleteMultipleObjects(cli, bucket, stream)
}

/*
 * DeleteMultipleObjectsFromStruct - delete a list of objects with object list struct
 *
 * PARAMS:
 *     - bucket: the name of the bucket to delete
 *     - objectListStruct: the object list struct to be deleted
 * RETURNS:
 *     - error: any error if it occurs
 */
func (cli *Client) DeleteMultipleObjectsFromStruct(bucket string,
        objectListStruct *api.DeleteMultipleObjectsArgs) (*api.DeleteMultipleObjectsResult, error) {
    jsonBytes, jsonErr := json.Marshal(objectListStruct)
    if jsonErr != nil {
        return nil, jsonErr
    }
    stream := http.NewBodyStreamFromBytes(jsonBytes)
    return api.DeleteMultipleObjects(cli, bucket, stream)
}

/*
 * InitiateMultipartUpload - initiate a multipart upload to get a upload ID
 *
 * PARAMS:
 *     - bucket: the bucket name
 *     - object: the object name
 *     - contentType: the content type of the object to be uploaded which should be specified,
 *       otherwise use the default(application/octet-stream)
 *     - args: the optional arguments
 * RETURNS:
 *     - *InitiateMultipartUploadResult: the result data structure
 *     - error: nil if ok otherwise the specific error
 */
func (cli *Client) InitiateMultipartUpload(bucket, object, contentType string,
        args *api.InitiateMultipartUploadArgs) (*api.InitiateMultipartUploadResult, error) {
    return api.InitiateMultipartUpload(cli, bucket, object, contentType, args)
}

/*
 * BasicInitiateMultipartUpload - basic interface to initiate a multipart upload
 *
 * PARAMS:
 *     - bucket: the bucket name
 *     - object: the object name
 * RETURNS:
 *     - *InitiateMultipartUploadResult: the result data structure
 *     - error: nil if ok otherwise the specific error
 */
func (cli *Client) BasicInitiateMultipartUpload(bucket,
        object string) (*api.InitiateMultipartUploadResult, error) {
    return api.InitiateMultipartUpload(cli, bucket, object, "", nil)
}

/*
 * UploadPart - upload the single part in the multipart upload process
 *
 * PARAMS:
 *     - bucket: the bucket name
 *     - object: the object name
 *     - uploadId: the multipart upload id
 *     - partNumber: the current part number
 *     - content: the uploaded part content
 *     - args: the optional arguments
 * RETURNS:
 *     - string: the etag of the uploaded part
 *     - error: nil if ok otherwise the specific error
 */
func (cli *Client) UploadPart(bucket, object, uploadId string, partNumber int,
        content *http.BodyStream, args *api.UploadPartArgs) (string, error) {
    return api.UploadPart(cli, bucket, object, uploadId, partNumber, content, args)
}

/*
 * BasicUploadPart - basic interface to upload the single part in the multipart upload process
 *
 * PARAMS:
 *     - bucket: the bucket name
 *     - object: the object name
 *     - uploadId: the multipart upload id
 *     - partNumber: the current part number
 *     - content: the uploaded part content
 * RETURNS:
 *     - string: the etag of the uploaded part
 *     - error: nil if ok otherwise the specific error
 */
func (cli *Client) BasicUploadPart(bucket, object, uploadId string, partNumber int,
        content *http.BodyStream) (string, error) {
    return api.UploadPart(cli, bucket, object, uploadId, partNumber, content, nil)
}

/*
 * UploadPartCopy - copy the multipart object
 *
 * PARAMS:
 *     - bucket: the destination bucket name
 *     - object: the destination object name
 *     - srcBucket: the source bucket
 *     - srcObject: the source object
 *     - uploadId: the multipart upload id
 *     - partNumber: the current part number
 *     - args: the optional arguments
 * RETURNS:
 *     - *CopyObjectResult: the lastModified and eTag of the part
 *     - error: nil if ok otherwise the specific error
 */
func (cli *Client) UploadPartCopy(bucket, object, srcBucket, srcObject, uploadId string,
        partNumber int, args *api.UploadPartCopyArgs) (*api.CopyObjectResult, error) {
    source := fmt.Sprintf("/%s/%s", srcBucket, srcObject)
    return api.UploadPartCopy(cli, bucket, object, source, uploadId, partNumber, args)
}

/*
 * BasicUploadPartCopy - basic interface to copy the multipart object
 *
 * PARAMS:
 *     - bucket: the destination bucket name
 *     - object: the destination object name
 *     - srcBucket: the source bucket
 *     - srcObject: the source object
 *     - uploadId: the multipart upload id
 *     - partNumber: the current part number
 * RETURNS:
 *     - *CopyObjectResult: the lastModified and eTag of the part
 *     - error: nil if ok otherwise the specific error
 */
func (cli *Client) BasicUploadPartCopy(bucket, object, srcBucket, srcObject, uploadId string,
        partNumber int) (*api.CopyObjectResult, error) {
    source := fmt.Sprintf("/%s/%s", srcBucket, srcObject)
    return api.UploadPartCopy(cli, bucket, object, source, uploadId, partNumber, nil)
}

/*
 * CompleteMultipartUpload - finish a multipart upload operation with parts stream
 *
 * PARAMS:
 *     - bucket: the destination bucket name
 *     - object: the destination object name
 *     - uploadId: the multipart upload id
 *     - parts: all parts info stream
 *     - meta: user defined meta data
 * RETURNS:
 *     - *CompleteMultipartUploadResult: the result data
 *     - error: nil if ok otherwise the specific error
 */
func (cli *Client) CompleteMultipartUpload(bucket, object, uploadId string,
        parts *http.BodyStream,
        meta map[string]string) (*api.CompleteMultipartUploadResult, error) {
    return api.CompleteMultipartUpload(cli, bucket, object, uploadId, parts, meta)
}

/*
 * CompleteMultipartUploadFromStruct - finish a multipart upload operation with parts struct
 *
 * PARAMS:
 *     - bucket: the destination bucket name
 *     - object: the destination object name
 *     - uploadId: the multipart upload id
 *     - parts: all parts info struct object
 *     - meta: user defined meta data
 * RETURNS:
 *     - *CompleteMultipartUploadResult: the result data
 *     - error: nil if ok otherwise the specific error
 */
func (cli *Client) CompleteMultipartUploadFromStruct(bucket, object, uploadId string,
        parts *api.CompleteMultipartUploadArgs,
        meta map[string]string) (*api.CompleteMultipartUploadResult, error) {
    jsonBytes, jsonErr := json.Marshal(parts)
    if jsonErr != nil {
        return nil, jsonErr
    }
    stream := http.NewBodyStreamFromBytes(jsonBytes)
    return api.CompleteMultipartUpload(cli, bucket, object, uploadId, stream, meta)
}

/*
 * AbortMultipartUpload - abort a multipart upload operation
 *
 * PARAMS:
 *     - bucket: the destination bucket name
 *     - object: the destination object name
 *     - uploadId: the multipart upload id
 * RETURNS:
 *     - error: nil if ok otherwise the specific error
 */
func (cli *Client) AbortMultipartUpload(bucket, object, uploadId string) error {
    return api.AbortMultipartUpload(cli, bucket, object, uploadId)
}

/*
 * ListParts - list the successfully uploaded parts info by upload id
 *
 * PARAMS:
 *     - bucket: the destination bucket name
 *     - object: the destination object name
 *     - uploadId: the multipart upload id
 *     - args: the optional arguments
 * RETURNS:
 *     - *ListPartsResult: the uploaded parts info result
 *     - error: nil if ok otherwise the specific error
 */
func (cli *Client) ListParts(bucket, object, uploadId string,
        args *api.ListPartsArgs) (*api.ListPartsResult, error) {
    return api.ListParts(cli, bucket, object, uploadId, args)
}

/*
 * BasicListParts - basic interface to list the successfully uploaded parts info by upload id
 *
 * PARAMS:
 *     - bucket: the destination bucket name
 *     - object: the destination object name
 *     - uploadId: the multipart upload id
 * RETURNS:
 *     - *ListPartsResult: the uploaded parts info result
 *     - error: nil if ok otherwise the specific error
 */
func (cli *Client) BasicListParts(bucket, object, uploadId string) (*api.ListPartsResult, error) {
    return api.ListParts(cli, bucket, object, uploadId, nil)
}

/*
 * ListMultipartUploads - list the unfinished uploaded parts of the given bucket
 *
 * PARAMS:
 *     - bucket: the destination bucket name
 *     - args: the optional arguments
 * RETURNS:
 *     - *ListMultipartUploadsResult: the unfinished uploaded parts info result
 *     - error: nil if ok otherwise the specific error
 */
func (cli *Client) ListMultipartUploads(bucket string,
        args *api.ListMultipartUploadsArgs) (*api.ListMultipartUploadsResult, error) {
    return api.ListMultipartUploads(cli, bucket, args)
}

/*
 * BasicListMultipartUploads - basic interface to list the unfinished uploaded parts
 *
 * PARAMS:
 *     - bucket: the destination bucket name
 * RETURNS:
 *     - *ListMultipartUploadsResult: the unfinished uploaded parts info result
 *     - error: nil if ok otherwise the specific error
 */
func (cli *Client) BasicListMultipartUploads(bucket string) (
        *api.ListMultipartUploadsResult, error) {
    return api.ListMultipartUploads(cli, bucket, nil)
}

