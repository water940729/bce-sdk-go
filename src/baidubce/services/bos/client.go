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

// Define client for BOS service
package bos

import (
    "fmt"
    "os"

    "baidubce/auth"
    "baidubce/bce"
    "baidubce/common"
    "baidubce/http"
    "baidubce/util"
    . "baidubce/services/bos/model"
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

type Client struct {
    // Client of BOS service is a kind of BceClient, so derived from BceClient
    *bce.BceClient

    // Fileds that used in parallel operation for BOS service
    MaxParallel   int
    MultipartSize int
}

// Make the BOS service client with default configuration
// Use `cli.Config.xxx` to access the config or change it to non-default value
func NewClient(ak, sk string) (*Client, error) {
    credentials, err := auth.NewBceDefaultCredentials(ak, sk)
    if err != nil {
        return nil, err
    }
    defaultSignOptions := &auth.SignOptions{
        auth.DEFAULT_HEADERS_TO_SIGN,
        util.NowUTCSeconds(),
        auth.DEFAULT_EXPIRE_SECONDS}
    defaultConf := &bce.BceClientConfiguration{
        bce.DEFAULT_PROTOCOL,
        fmt.Sprintf("%s.%s", bce.DEFAULT_REGION, common.DEFAULT_SERVICE_DOMAIN),
        bce.DEFAULT_REGION,
        credentials,
        defaultSignOptions,
        bce.DEFAULT_RETRY_POLICY,
        bce.DEFAULT_USER_AGENT,
        bce.DEFAULT_CONNECTION_TIMEOUT_IN_MILLIS}
    v1Signer := &auth.BceV1Signer{}

    client := &Client{bce.NewBceClient(defaultConf, v1Signer),
        DEFAULT_MAX_PARALLEL, DEFAULT_MULTIPART_SIZE}
    return client, nil
}

// The raw api of the BOS client
func (cli *Client) ApiHeadBucket(req *HeadBucketRequest, resp *HeadBucketResponse) error {
    return cli.SendRequest(req, resp)
}

func (cli *Client) ApiPutBucket(req *PutBucketRequest, resp *PutBucketResponse) error {
    return cli.SendRequest(req, resp)
}

func (cli *Client) ApiDeleteBucket(req *DeleteBucketRequest, resp *DeleteBucketResponse) error {
    return cli.SendRequest(req, resp)
}

func (cli *Client) ApiGetBucketAcl(req *GetBucketAclRequest, resp *GetBucketAclResponse) error {
    return cli.SendRequest(req, resp)
}

func (cli *Client) ApiPutBucketAcl(req *PutBucketAclRequest, resp *PutBucketAclResponse) error {
    return cli.SendRequest(req, resp)
}

func (cli *Client) ApiGetBucketLocation(req *GetBucketLocationRequest,
        resp *GetBucketLocationResponse) error {
    return cli.SendRequest(req, resp)
}

func (cli *Client) ApiListBuckets(req *ListBucketsRequest, resp *ListBucketsResponse) error {
    return cli.SendRequest(req, resp)
}

func (cli *Client) ApiListObjects(req *ListObjectsRequest, resp *ListObjectsResponse) error {
    return cli.SendRequest(req, resp)
}

func (cli *Client) ApiDeleteBucketLifecycle(req *DeleteBucketLifecycleRequest,
        resp *DeleteBucketLifecycleResponse) error {
    return cli.SendRequest(req, resp)
}

func (cli *Client) ApiDeleteBucketLogging(req *DeleteBucketLoggingRequest,
        resp *DeleteBucketLoggingResponse) error {
    return cli.SendRequest(req, resp)
}

func (cli *Client) ApiGetBucketLifecycle(req *GetBucketLifecycleRequest,
        resp *GetBucketLifecycleResponse) error {
    return cli.SendRequest(req, resp)
}

func (cli *Client) ApiGetBucketLogging(req *GetBucketLoggingRequest,
        resp *GetBucketLoggingResponse) error {
    return cli.SendRequest(req, resp)
}

func (cli *Client) ApiGetBucketStorageClass(req *GetBucketStorageClassRequest,
        resp *GetBucketStorageClassResponse) error {
    return cli.SendRequest(req, resp)
}

func (cli *Client) ApiPutBucketLifecycle(req *PutBucketLifecycleRequest,
        resp *PutBucketLifecycleResponse) error {
    return cli.SendRequest(req, resp)
}

func (cli *Client) ApiPutBucketLogging(req *PutBucketLoggingRequest,
        resp *PutBucketLoggingResponse) error {
    return cli.SendRequest(req, resp)
}

func (cli *Client) ApiPutBucketStorageClass(req *PutBucketStorageClassRequest,
        resp *PutBucketStorageClassResponse) error {
    return cli.SendRequest(req, resp)
}

func (cli *Client) ApiGetObject(req *GetObjectRequest, resp *GetObjectResponse) error {
    return cli.SendRequest(req, resp)
}

func (cli *Client) ApiPutObject(req *PutObjectRequest, resp *PutObjectResponse) error {
    return cli.SendRequest(req, resp)
}

func (cli *Client) ApiDeleteObject(req *DeleteObjectRequest, resp *DeleteObjectResponse) error {
    return cli.SendRequest(req, resp)
}

func (cli *Client) ApiCopyObject(req *CopyObjectRequest, resp *CopyObjectResponse) error {
    return cli.SendRequest(req, resp)
}

func (cli *Client) ApiGetObjectMeta(req *GetObjectMetaRequest, resp *GetObjectMetaResponse) error {
    return cli.SendRequest(req, resp)
}

func (cli *Client) ApiFetchObject(req *FetchObjectRequest, resp *FetchObjectResponse) error {
    return cli.SendRequest(req, resp)
}

func (cli *Client) ApiAppendObject(req *AppendObjectRequest, resp *AppendObjectResponse) error {
    return cli.SendRequest(req, resp)
}

func (cli *Client) ApiDeleteMultipleObjects(req *DeleteMultipleObjectsRequest,
        resp *DeleteMultipleObjectsResponse) error {
    return cli.SendRequest(req, resp)
}

func (cli *Client) ApiInitiateMultipartUpload(req *InitiateMultipartUploadRequest,
        resp *InitiateMultipartUploadResponse) error {
    return cli.SendRequest(req, resp)
}

func (cli *Client) ApiCompleteMultipartUpload(req *CompleteMultipartUploadRequest,
        resp *CompleteMultipartUploadResponse) error {
    return cli.SendRequest(req, resp)
}

func (cli *Client) ApiUploadPart(req *UploadPartRequest, resp *UploadPartResponse) error {
    return cli.SendRequest(req, resp)
}

func (cli *Client) ApiAbortMultipartUpload(req *AbortMultipartUploadRequest,
        resp *AbortMultipartUploadResponse) error {
    return cli.SendRequest(req, resp)
}

func (cli *Client) ApiListParts(req *ListPartsRequest, resp *ListPartsResponse) error {
    return cli.SendRequest(req, resp)
}

func (cli *Client) ApiListMultipartUploads(req *ListMultipartUploadsRequest,
        resp *ListMultipartUploadsResponse) error {
    return cli.SendRequest(req, resp)
}

// Wrapper methods for easily use
func (cli *Client) HeadBucket(bucket string) error {
    req := NewHeadBucketRequest(bucket)
    resp := NewHeadBucketResponse()
    return cli.ApiHeadBucket(req, resp)
}

func (cli *Client) PutBucket(bucket string) error {
    req := NewPutBucketRequest(bucket)
    resp := NewPutBucketResponse()
    return cli.ApiPutBucket(req, resp)
}

func (cli *Client) DeleteBucket(bucket string) error {
    req := NewDeleteBucketRequest(bucket)
    resp := NewDeleteBucketResponse()
    return cli.ApiDeleteBucket(req, resp)
}

func (cli *Client) GetObject(bucket, object, fileName string, meta *ObjectMetadata) error {
    req := NewGetObjectRequest(bucket, object)
    resp := NewGetObjectResponse()
    if err := cli.ApiGetObject(req, resp); err != nil {
        return err
    }
    body := resp.Content()
    defer body.Close()

    file, fileErr := os.OpenFile(fileName,
        os.O_TRUNC | os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0755)
    if fileErr != nil {
        return fileErr
    }

    buf := make([]byte, 4096)
    for {
        n, e := body.Read(buf)
        file.Write(buf[:n])
        if e != nil {
            break
        }
    }

    if meta != nil {
        meta.CacheControl = resp.CacheControl()
        meta.ContentDisposition = resp.ContentDisposition()
        meta.ContentLength = resp.ContentLength()
        meta.ContentRange = resp.ContentRange()
        meta.ContentType = resp.ContentType()
        meta.ContentMD5 = resp.ContentMD5()
        meta.Expires = resp.Expires()
        meta.ETag = resp.ETag()
        meta.LastModified = resp.LastModified()
        meta.StorageClass = resp.StorageClass()
    }
    return nil
}

func (cli *Client) PutObjectFromFile(bucket, object, fileName string,
        etag, version *string) error {
    stream, err := http.NewBodyStreamFromFile(fileName)
    if err != nil {
        return err
    }
    req := NewPutObjectRequest(bucket, object, stream)
    resp := NewPutObjectResponse()
    if err := cli.ApiPutObject(req, resp); err != nil {
        return err
    }
    if etag != nil {
        *etag = resp.ETag()
    }
    if version != nil {
        *version = resp.Version()
    }
    return nil
}

func (cli *Client) PutObjectFromString(bucket, object, content string,
        etag, version *string) error {
    req := NewPutObjectRequest(bucket, object, content)
    resp := NewPutObjectResponse()
    if err := cli.ApiPutObject(req, resp); err != nil {
        return err
    }
    if etag != nil {
        *etag = resp.ETag()
    }
    if version != nil {
        *version = resp.Version()
    }
    return nil
}

func (cli *Client) DeleteObject(bucket, object string) error {
    req := NewDeleteObjectRequest(bucket, object)
    resp := NewDeleteObjectResponse()
    return cli.ApiDeleteObject(req, resp)
}

func (cli *Client) CopyObject(bucket, object, srcBucket, srcObject, storageClass string,
        meta map[string]string, lastModified, etag *string) error {
    req := NewCopyObjectRequest(bucket, object)
    req.SetSourceBucket(srcBucket)
    req.SetSourceObject(srcObject)
    req.SetStorageClass(storageClass)
    if val, ok := meta[http.BCE_COPY_SOURCE_IF_MATCH]; ok {
        req.SetMatchETag(val)
    }
    if val, ok := meta[http.BCE_COPY_SOURCE_IF_NONE_MATCH]; ok {
        req.SetNoneMatchETag(val)
    }
    if val, ok := meta[http.BCE_COPY_SOURCE_IF_MODIFIED_SINCE]; ok {
        req.SetModifiedSince(val)
    }
    if val, ok := meta[http.BCE_COPY_SOURCE_IF_UNMODIFIED_SINCE]; ok {
        req.SetUnmodifiedSince(val)
    }
    if val, ok := meta[http.BCE_COPY_METADATA_DIRECTIVE]; ok {
        req.SetMetadataDirective(val)
    }
    resp := NewCopyObjectResponse()
    if err := cli.ApiCopyObject(req, resp); err != nil {
        return err
    }
    if lastModified != nil {
        *lastModified = resp.LastModified()
    }
    if etag != nil {
        *etag = resp.ETag()
    }
    return nil
}

