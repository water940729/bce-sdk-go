package bos

import (
    "testing"

    "baidubce/http"
    . "baidubce/services/bos/model"
)

var (
    AK = "ab3e8280a5ff436eb5c5b9d7fa14fde9"
    SK = "2544e6b152a94241b8af9dbb25245c5e"
    BOS_CLIENT *Client

    EXISTS_BUCKET = "bos-rd-ssy"
    NEW_BUCKET    = "bos-rd-ssy1"
)

func init() {
    BOS_CLIENT, _ = NewClient(AK, SK)
}

func TestHeadBucket(t *testing.T) {
    request  := NewHeadBucketRequest(EXISTS_BUCKET)
    response := NewHeadBucketResponse()
    err := BOS_CLIENT.ApiHeadBucket(request, response)
    if err != nil {
        t.Error(err)
    }
    t.Logf("time cost: %v", response.ElapsedTime())
    if response.IsFail() {
        t.Error(response.StatusText())
    }
    t.Logf("Status: %v", response.StatusText())
    for k, v := range response.GetHeaders() {
        t.Logf("%s: %s", k, v)
    }
}

func TestPutBucket(t *testing.T) {
    request  := NewPutBucketRequest(NEW_BUCKET)
    response := NewPutBucketResponse()
    err := BOS_CLIENT.ApiPutBucket(request, response)
    if err != nil {
        t.Error(err)
    }
    t.Logf("time cost: %v", response.ElapsedTime())
    if response.IsFail() {
        t.Error(response.StatusText())
    }
    t.Logf("Status: %v", response.StatusText())
    for k, v := range response.GetHeaders() {
        t.Logf("%s: %s", k, v)
    }
    t.Logf("Location=%v", response.GetLocation())
}

func TestDeleteBucket(t *testing.T) {
    request  := NewDeleteBucketRequest(NEW_BUCKET)
    response := NewDeleteBucketResponse()
    err := BOS_CLIENT.ApiDeleteBucket(request, response)
    if err != nil {
        t.Error(err)
    }
    t.Logf("time cost: %v", response.ElapsedTime())
    if response.IsFail() {
        t.Error(response.StatusText())
    }
    t.Logf("Status: %v", response.StatusText())
    for k, v := range response.GetHeaders() {
        t.Logf("%s: %s", k, v)
    }
}

func TestGetBucketAcl(t *testing.T) {
    request  := NewGetBucketAclRequest(EXISTS_BUCKET)
    response := NewGetBucketAclResponse()
    err := BOS_CLIENT.ApiGetBucketAcl(request, response)
    if err != nil {
        t.Error(err)
    }
    t.Logf("time cost: %v", response.ElapsedTime())
    if response.IsFail() {
        t.Error(response.StatusText())
    }
    t.Logf("Status: %v", response.StatusText())
    for k, v := range response.GetHeaders() {
        t.Logf("%s: %s", k, v)
    }
    t.Logf("Acl=%v", response.Acl())
}

func TestPutBucketAcl(t *testing.T) {
    acl := `{
    "accessControlList":[
        {
            "grantee":[{
                "id":"e13b12d0131b4c8bae959df4969387b8"
            }],
            "permission":["FULL_CONTROL"]
        }
    ]
}`
    stream := http.NewBodyStreamFromString(acl)
    request  := NewPutBucketAclRequest(EXISTS_BUCKET, stream)//"private")
    response := NewPutBucketAclResponse()
    err := BOS_CLIENT.ApiPutBucketAcl(request, response)
    if err != nil {
        t.Error(err)
    }
    t.Logf("time cost: %v", response.ElapsedTime())
    if response.IsFail() {
        t.Error(response.StatusText())
    }
    t.Logf("Status: %v", response.StatusText())
    for k, v := range response.GetHeaders() {
        t.Logf("%s: %s", k, v)
    }
}

func TestGetBucketLocation(t *testing.T) {
    request  := NewGetBucketLocationRequest(EXISTS_BUCKET)
    response := NewGetBucketLocationResponse()
    err := BOS_CLIENT.ApiGetBucketLocation(request, response)
    if err != nil {
        t.Error(err)
    }
    t.Logf("time cost: %v", response.ElapsedTime())
    if response.IsFail() {
        t.Error(response.StatusText())
    }
    t.Logf("Status: %v", response.StatusText())
    for k, v := range response.GetHeaders() {
        t.Logf("%s: %s", k, v)
    }
    t.Logf("Location=%v", response.Location())
}

func TestListBuckets(t *testing.T) {
    request  := NewListBucketsRequest()
    response := NewListBucketsResponse()
    err := BOS_CLIENT.ApiListBuckets(request, response)
    if err != nil {
        t.Error(err)
    }
    t.Logf("time cost: %v", response.ElapsedTime())
    if response.IsFail() {
        t.Error(response.StatusText())
    }
    t.Logf("Status: %v", response.StatusText())
    for k, v := range response.GetHeaders() {
        t.Logf("%s: %s", k, v)
    }
    t.Logf("%v", response.Buckets())
}

func TestListObjects(t *testing.T) {
    request  := NewListObjectsRequest(EXISTS_BUCKET)
    request.SetPrefix("d")
    response := NewListObjectsResponse()
    err := BOS_CLIENT.ApiListObjects(request, response)
    if err != nil {
        t.Error(err)
    }
    t.Logf("time cost: %v", response.ElapsedTime())
    if response.IsFail() {
        t.Error(response.StatusText())
    }
    t.Logf("Status: %v", response.StatusText())
    for k, v := range response.GetHeaders() {
        t.Logf("%s: %s", k, v)
    }
    t.Logf("%v", response.Contents())
}

func TestGetBucketLifecycle(t *testing.T) {
    request  := NewGetBucketLifecycleRequest(EXISTS_BUCKET)
    response := NewGetBucketLifecycleResponse()
    err := BOS_CLIENT.ApiGetBucketLifecycle(request, response)
    if err != nil {
        t.Error(err)
    }
    t.Logf("time cost: %v", response.ElapsedTime())
    if response.IsFail() {
        if response.StatusCode() == 404 {
            t.Logf("%+v", response.ServiceError())
        } else {
            t.Error(response.StatusText())
        }
    }
    for k, v := range response.GetHeaders() {
        t.Logf("%s: %s", k, v)
    }
    if !response.IsFail() {
        t.Logf("%v", response.Rule())
    }
}

func TestGetBucketLogging(t *testing.T) {
    request  := NewGetBucketLoggingRequest(EXISTS_BUCKET)
    response := NewGetBucketLoggingResponse()
    err := BOS_CLIENT.ApiGetBucketLogging(request, response)
    if err != nil {
        t.Error(err)
    }
    t.Logf("time cost: %v", response.ElapsedTime())
    if response.IsFail() {
        t.Error(response.StatusText())
    }
    t.Logf("Status: %v", response.StatusText())
    for k, v := range response.GetHeaders() {
        t.Logf("%s: %s", k, v)
    }
    t.Logf("%+v", response)
}

func TestPutBucketStorageClass(t *testing.T) {
    storage := `{
    "storageClass":"STANDARD_IA"
}`
    request  := NewPutBucketStorageClassRequest(EXISTS_BUCKET, storage)
    response := NewPutBucketStorageClassResponse()
    err := BOS_CLIENT.ApiPutBucketStorageClass(request, response)
    if err != nil {
        t.Error(err)
    }
    t.Logf("time cost: %v", response.ElapsedTime())
    if response.IsFail() {
        t.Error(response.StatusText())
    }
    t.Logf("Status: %v", response.StatusText())
    for k, v := range response.GetHeaders() {
        t.Logf("%s: %s", k, v)
    }
}

func TestGetBucketStorageClass(t *testing.T) {
    request  := NewGetBucketStorageClassRequest(EXISTS_BUCKET)
    response := NewGetBucketStorageClassResponse()
    err := BOS_CLIENT.ApiGetBucketStorageClass(request, response)
    if err != nil {
        t.Error(err)
    }
    t.Logf("time cost: %v", response.ElapsedTime())
    if response.IsFail() {
        t.Error(response.StatusText())
    }
    t.Logf("Status: %v", response.StatusText())
    for k, v := range response.GetHeaders() {
        t.Logf("%s: %s", k, v)
    }
    t.Logf("%v", response.StorageClass())
}

// Object API testing
func TestGetObject(t *testing.T) {
    request  := NewGetObjectRequest(EXISTS_BUCKET, "testsdk")
    request.AddResponseHeader("ContentType", "text/json")
    request.SetRange(10, 11)
    response := NewGetObjectResponse()
    //err := BOS_CLIENT.SendRequest(request, response)
    err := BOS_CLIENT.ApiGetObject(request, response)
    if err != nil {
        t.Error(err)
    }
    t.Logf("time cost: %v", response.ElapsedTime())
    if response.IsFail() {
        t.Error(response.StatusText())
    }
    t.Logf("Status: %v", response.StatusText())
    for k, v := range response.GetHeaders() {
        t.Logf("%s: %s", k, v)
    }
    t.Logf("%#v", response.ContentType())
    t.Logf("%#v", response.ContentLength())
    t.Logf("%#v", response.ContentMD5())

    body := response.Content()
    defer body.Close()
    t.Logf("%d", body.Len())
    for {
        buf := make([]byte, 1024)
        n, e := body.Read(buf)
        t.Logf("%s", buf[0:n])
        if e != nil {
            break
        }
    }
}

func TestPutObject(t *testing.T) {
    request  := NewPutObjectRequest(EXISTS_BUCKET, "testsdk-putobject", "1234567890")
    response := NewPutObjectResponse()
    err := BOS_CLIENT.ApiPutObject(request, response)
    if err != nil {
        t.Error(err)
    }
    t.Logf("time cost: %v", response.ElapsedTime())
    if response.IsFail() {
        t.Error(response.StatusText())
        t.Logf("%v", response.ServiceError())
        return
    }
    t.Logf("Status: %v", response.StatusText())
    for k, v := range response.GetHeaders() {
        t.Logf("%s: %s", k, v)
    }
    t.Logf("%v", response.ETag())
    t.Logf("%v", response.Version())
}

func TestCopyObject(t *testing.T) {
    request  := NewCopyObjectRequest(EXISTS_BUCKET, "testsdk-putobject")
    request.SetSourceBucket(EXISTS_BUCKET)
    request.SetSourceObject("testsdk-putobject")
    request.SetMetadataDirective(METADATA_DIRECTIVE_REPLACE)
    request.SetStorageClass(STORAGE_CLASS_COLD)

    response := NewCopyObjectResponse()
    err := BOS_CLIENT.ApiCopyObject(request, response)
    if err != nil {
        t.Logf("%v", response.ServiceError())
        t.Error(err)
        return
    }
    t.Logf("time cost: %v", response.ElapsedTime())
    if response.IsFail() {
        t.Error(response.StatusText())
    }
    t.Logf("Status: %v", response.StatusText())
    for k, v := range response.GetHeaders() {
        t.Logf("%s: %s", k, v)
    }
    t.Logf("%v", response.ETag())
    t.Logf("%v", response.LastModified())
}

func TestGetObjectMeta(t *testing.T) {
    request  := NewGetObjectMetaRequest(EXISTS_BUCKET, "testsdk-putobject")
    response := NewGetObjectMetaResponse()
    err := BOS_CLIENT.ApiGetObjectMeta(request, response)
    if err != nil {
        t.Logf("%v", response.ServiceError())
        t.Error(err)
        return
    }
    t.Logf("time cost: %v", response.ElapsedTime())
    if response.IsFail() {
        t.Error(response.StatusText())
    }
    t.Logf("Status: %v", response.StatusText())
    for k, v := range response.GetHeaders() {
        t.Logf("%s: %s", k, v)
    }
    t.Logf("%v", response.CacheControl())
    t.Logf("%v", response.ContentLength())
    t.Logf("%v", response.ContentMD5())
    t.Logf("%v", response.ETag())
    t.Logf("%v", response.StorageClass())
    t.Logf("%v", response.Expires())
}

func TestDeleteObject(t *testing.T) {
    request  := NewDeleteObjectRequest(EXISTS_BUCKET, "testsdk-putobject")
    response := NewDeleteObjectResponse()
    err := BOS_CLIENT.ApiDeleteObject(request, response)
    if err != nil {
        t.Error(err)
        t.Logf("%v", response.ServiceError())
        return
    }
    t.Logf("time cost: %v", response.ElapsedTime())
    if response.IsFail() {
        t.Error(response.StatusText())
    }
    t.Logf("Status: %v", response.StatusText())
    for k, v := range response.GetHeaders() {
        t.Logf("%s: %s", k, v)
    }
}

func TestFetchObject(t *testing.T) {
    request  := NewFetchObjectRequest(EXISTS_BUCKET, "testsdk-fetchobject")
    request.SetFetchMode(FETCH_MODE_ASYNC)
    request.SetFetchSource("https://docs.python.org/2/library/argparse.html")

    response := NewFetchObjectResponse()
    err := BOS_CLIENT.ApiFetchObject(request, response)
    if err != nil {
        t.Error(err)
    }
    t.Logf("time cost: %v", response.ElapsedTime())
    if response.IsFail() {
        t.Error(response.StatusText())
        t.Logf("%v", response.ServiceError())
        return
    }
    t.Logf("Status: %v", response.StatusText())
    for k, v := range response.GetHeaders() {
        t.Logf("%s: %s", k, v)
    }
    t.Logf("%v", response.Code())
    t.Logf("%v", response.Message())
    t.Logf("%v", response.RequestId())
    t.Logf("%v", response.JobId())
}

func TestAppendObject(t *testing.T) {
    request  := NewAppendObjectRequest(EXISTS_BUCKET, "testsdk-appendobject", "append")
    //request.SetOffset(6)
    response := NewAppendObjectResponse()
    err := BOS_CLIENT.ApiAppendObject(request, response)
    if err != nil {
        t.Error(err)
    }
    t.Logf("time cost: %v", response.ElapsedTime())
    if response.IsFail() {
        t.Error(response.StatusText())
        t.Logf("%v", response.ServiceError())
        return
    }
    t.Logf("Status: %v", response.StatusText())
    for k, v := range response.GetHeaders() {
        t.Logf("%s: %s", k, v)
    }
    t.Logf("%v", response.ContentMD5())
    t.Logf("%v", response.ETag())
    t.Logf("%v", response.NextAppendOffset())
}

func TestDeleteMultipleObjects(t *testing.T) {
//    multiDeleteStr := `{
//    "objects":[
//        {"key": "aaaa"},
//        {"key": "bbbb"},
//        {"key": "testsdk-appendobject"},
//        {"key": "cccc"}
//    ]
//}`
//    request  := NewDeleteMultipleObjectsRequest(EXISTS_BUCKET, multiDeleteStr)

    multiDeleteObj := &DeleteMultipleObjectsInput{
            []DeleteObjectInput{DeleteObjectInput{"1"}, DeleteObjectInput{"2"}}}
    request  := NewDeleteMultipleObjectsRequest(EXISTS_BUCKET, multiDeleteObj)
    response := NewDeleteMultipleObjectsResponse()
    err := BOS_CLIENT.SendRequest(request, response)
    if err != nil { // client error occurs
        t.Error(err)
        return
    }
    t.Logf("time cost: %v", response.ElapsedTime())
    if response.IsFail() { // service error occurs
        t.Error(response.StatusText())
        t.Logf("%+v", response.ServiceError())
        return
    }
    t.Logf("Status: %+v", response.StatusText())
    for k, v := range response.GetHeaders() {
        t.Logf("%s: %s", k, v)
    }
    t.Logf("%+v", response.FailedObjects())
}

