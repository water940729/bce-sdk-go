package bos

import (
    "encoding/json"
    "os"
    "path/filepath"
    "runtime"
    "testing"

    "baidubce/bce"
    "baidubce/services/bos/api"
    "baidubce/util/log"
)

var (
    BOS_CLIENT *Client
    EXISTS_BUCKET = "bos-rd-ssy"
)

// For security reason, ak/sk should not hard write here.
type Conf struct {
    AK string
    SK string
}

func init() {
    _, f, _, _ := runtime.Caller(0)
    for i := 0; i < 5; i++ {
        f = filepath.Dir(f)
    }
    conf := filepath.Join(f, "config.json")
    fp, err := os.Open(conf)
    if err != nil {
        log.Fatal("config json file of ak/sk not given:", conf)
        os.Exit(1)
    }
    decoder := json.NewDecoder(fp)
    confObj := &Conf{}
    decoder.Decode(confObj)

    BOS_CLIENT, _ = NewClient(confObj.AK, confObj.SK)
    log.SetLogHandler(log.STDERR | log.FILE)
    log.SetRotateType(log.ROTATE_SIZE)
    //log.SetLogLevel(log.WARN)
}

func TestListBuckets(t *testing.T) {
    res, err := BOS_CLIENT.ListBuckets()
    if err != nil {
        t.Error(err)
    }
    t.Logf("%+v", res)
}

func TestListObjects(t *testing.T) {
    args := &api.ListObjectsArgs{Prefix: "test", MaxKeys: 10}
    res, err := BOS_CLIENT.ListObjects(EXISTS_BUCKET, args)
    if err != nil {
        t.Error(err)
    }
    t.Logf("%+v", res)
}

func TestHeadBucket(t *testing.T) {
    err := BOS_CLIENT.HeadBucket(EXISTS_BUCKET)
    if err != nil {
        t.Error(err)
    }
}

func TestPutBucket(t *testing.T) {
    res, err := BOS_CLIENT.PutBucket("test-put-bucket")
    if err != nil {
        t.Error(err)
    }
    t.Logf("%v", res)
}

func TestDeleteBucket(t *testing.T) {
    err := BOS_CLIENT.DeleteBucket("test-put-bucket")
    if err != nil {
        t.Error(err)
    }
}

func TestGetBucketLocation(t *testing.T) {
    res, err := BOS_CLIENT.GetBucketLocation(EXISTS_BUCKET)
    if err != nil {
        t.Error(err)
    }
    t.Logf("%v", res)
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
    body, _ := bce.NewBodyFromString(acl)
    err := BOS_CLIENT.PutBucketAcl(EXISTS_BUCKET, body)
    if err != nil {
        t.Error(err)
    }
    res, _ := BOS_CLIENT.GetBucketAcl(EXISTS_BUCKET)
    t.Logf("%+v", res)
}

func TestPutBucketAclFromCanned(t *testing.T) {
    err := BOS_CLIENT.PutBucketAclFromCanned(EXISTS_BUCKET, api.CANNED_ACL_PUBLIC_READ)
    if err != nil {
        t.Error(err)
    }
    res, _ := BOS_CLIENT.GetBucketAcl(EXISTS_BUCKET)
    t.Logf("%+v", res)
}

func TestPutBucketAclFromFile(t *testing.T) {
    acl := `{
    "accessControlList":[
        {
            "grantee":[
                {"id":"e13b12d0131b4c8bae959df4969387b8"},
                {"id":"a13b12d0131b4c8bae959df4969387b8"}
            ],
            "permission":["FULL_CONTROL"]
        }
    ]
}`
    fname := "/tmp/test-put-bucket-acl-by-file"
    f, _ := os.Create(fname)
    f.WriteString(acl)
    f.Close()
    err := BOS_CLIENT.PutBucketAclFromFile(EXISTS_BUCKET, fname)
    if err != nil {
        t.Error(err)
    }
    res, _ := BOS_CLIENT.GetBucketAcl(EXISTS_BUCKET)
    t.Logf("%+v", res)
    os.Remove(fname)
}

func TestPutBucketAclFromStruct(t *testing.T) {
    args := &api.PutBucketAclArgs{
        []api.GrantType{
            api.GrantType{
                Grantee: []api.GranteeType{
                    api.GranteeType{"e13b12d0131b4c8bae959df4969387b8"},
                },
                Permission: []string{
                    "FULL_CONTROL",
                },
            },
        },
    }
    err := BOS_CLIENT.PutBucketAclFromStruct(EXISTS_BUCKET, args)
    if err != nil {
        t.Error(err)
    }
    res, _ := BOS_CLIENT.GetBucketAcl(EXISTS_BUCKET)
    t.Logf("%+v", res)
}

func TestPutBucketLogging(t *testing.T) {
    body, _ := bce.NewBodyFromString(
        `{"targetBucket": "bos-rd-ssy", "targetPrefix": "my-log/"}`)
    err := BOS_CLIENT.PutBucketLogging(EXISTS_BUCKET, body)
    if err != nil {
        t.Error(err)
    }
    res, _ := BOS_CLIENT.GetBucketLogging(EXISTS_BUCKET)
    t.Logf("%+v", res)
}

func TestPutBucketLoggingFromString(t *testing.T) {
    logging := `{"targetBucket": "bos-rd-ssy", "targetPrefix": "my-log2/"}`
    err := BOS_CLIENT.PutBucketLoggingFromString(EXISTS_BUCKET, logging)
    if err != nil {
        t.Error(err)
    }
    res, _ := BOS_CLIENT.GetBucketLogging(EXISTS_BUCKET)
    t.Logf("%+v", res)
}

func TestPutBucketLoggingFromStruct(t *testing.T) {
    obj := &api.PutBucketLoggingArgs{"bos-rd-ssy", "my-log3/"}
    err := BOS_CLIENT.PutBucketLoggingFromStruct(EXISTS_BUCKET, obj)
    if err != nil {
        t.Error(err)
    }
    res, _ := BOS_CLIENT.GetBucketLogging(EXISTS_BUCKET)
    t.Logf("%+v", res)
}

func TestDeleteBucketLogging(t *testing.T) {
    err := BOS_CLIENT.DeleteBucketLogging(EXISTS_BUCKET)
    if err != nil {
        t.Error(err)
    }
    res, _ := BOS_CLIENT.GetBucketLogging(EXISTS_BUCKET)
    t.Logf("%+v", res)
}

func TestPutBucketLifecycle(t *testing.T) {
    str := `{
    "rule": [
        {
            "id": "transition-to-cold",
            "status": "enabled",
            "resource": ["bos-rd-ssy/test*"],
            "condition": {
                "time": {
                    "dateGreaterThan": "2018-09-07T00:00:00Z"
                }
            },
            "action": {
                "name": "DeleteObject"
            }
        }
    ]
}`
    body, _ := bce.NewBodyFromString(str)
    err := BOS_CLIENT.PutBucketLifecycle(EXISTS_BUCKET, body)
    if err != nil {
        t.Error(err)
    }
    res, _ := BOS_CLIENT.GetBucketLifecycle(EXISTS_BUCKET)
    t.Logf("%+v", res)
}

func TestPutBucketLifecycleFromString(t *testing.T) {
    obj := `{
    "rule": [
        {
            "id": "transition-to-cold",
            "status": "enabled",
            "resource": ["bos-rd-ssy/test*"],
            "condition": {
                "time": {
                    "dateGreaterThan": "2018-09-07T00:00:00Z"
                }
            },
            "action": {
                "name": "DeleteObject"
            }
        }
    ]
}`
    err := BOS_CLIENT.PutBucketLifecycleFromString(EXISTS_BUCKET, obj)
    if err != nil {
        t.Error(err)
    }
    res, _ := BOS_CLIENT.GetBucketLifecycle(EXISTS_BUCKET)
    t.Logf("%+v", res)
}

func TestDeleteBucketLifecycle(t *testing.T) {
    err := BOS_CLIENT.DeleteBucketLifecycle(EXISTS_BUCKET)
    if err != nil {
        t.Error(err)
    }
    res, _ := BOS_CLIENT.GetBucketLifecycle(EXISTS_BUCKET)
    if res != nil {
        t.Error("delete failed")
    }
}

func TestPutBucketStorageClass(t *testing.T) {
    err := BOS_CLIENT.PutBucketStorageclass(EXISTS_BUCKET, api.STORAGE_CLASS_STANDARD_IA)
    if err != nil {
        t.Error(err)
    }
}

func TestGetBucketStorageClass(t *testing.T) {
    res, err := BOS_CLIENT.GetBucketStorageclass(EXISTS_BUCKET)
    if err != nil {
        t.Error(err)
    }
    t.Logf("%+v", res)
}

func TestPutObject(t *testing.T) {
    body, _ := bce.NewBodyFromString("aaaaaaaaaaa")
    res, err := BOS_CLIENT.PutObject(EXISTS_BUCKET, "test-put-object", body)
    if err != nil {
        t.Error(err)
    }
    t.Logf("etag: %v", res)
}

func TestPutObjectFromString(t *testing.T) {
    res, err := BOS_CLIENT.PutObjectFromString(EXISTS_BUCKET, "test-put-object", "123")
    if err != nil {
        t.Error(err)
    }
    t.Logf("etag: %v", res)
}

func TestPutObjectFromFile(t *testing.T) {
    fname := "/tmp/test-put-file"
    f, _ := os.Create(fname)
    f.WriteString("12345")
    f.Close()
    res, err := BOS_CLIENT.PutObjectFromFile(EXISTS_BUCKET, "test-put-object", fname)
    if err != nil {
        t.Error(err)
    }
    t.Logf("etag: %v", res)
    os.Remove(fname)
}

func TestCopyObject(t *testing.T) {
    args := &api.CopyObjectArgs{StorageClass: api.STORAGE_CLASS_COLD}
    res, err := BOS_CLIENT.CopyObject(EXISTS_BUCKET, "test-copy-object",
        EXISTS_BUCKET, "test-put-object", args)
    if err != nil {
        t.Error(err)
    }
    t.Logf("copy result: %+v", res)
}

func TestBasicCopyObject(t *testing.T) {
    res, err := BOS_CLIENT.BasicCopyObject(EXISTS_BUCKET, "test-copy-object",
        EXISTS_BUCKET, "test-put-object")
    if err != nil {
        t.Error(err)
    }
    t.Logf("copy result: %+v", res)
}

func TestGetObject(t *testing.T) {
    respHeaders := map[string]string{"ContentEncoding" : "qqqqqqqqqqqqq"}
    args := &api.GetObjectArgs{2, 4, respHeaders}
    res, err := BOS_CLIENT.GetObject(EXISTS_BUCKET, "test-put-object", args)
    if err != nil {
        t.Error(err)
    }
    t.Logf("%+v", res)

    defer res.Body.Close()
    t.Logf("%v", res.ContentLength)
    for {
        buf := make([]byte, 1024)
        n, e := res.Body.Read(buf)
        t.Logf("%s", buf[0:n])
        if e != nil {
            break
        }
    }
}

func TestBasicGetObject(t *testing.T) {
    res, err := BOS_CLIENT.BasicGetObject(EXISTS_BUCKET, "test-put-object")
    if err != nil {
        t.Error(err)
    }
    t.Logf("%+v", res)

    defer res.Body.Close()
    t.Logf("%v", res.ContentLength)
    for {
        buf := make([]byte, 1024)
        n, e := res.Body.Read(buf)
        t.Logf("%s", buf[0:n])
        if e != nil {
            break
        }
    }
}

func TestSimpleGetObject(t *testing.T) {
    respHeaders := map[string]string{"ContentEncoding" : "qqqqqqqqqqqqq"}
    res, err := BOS_CLIENT.SimpleGetObject(EXISTS_BUCKET, "test-put-object", 0, 5, respHeaders)
    if err != nil {
        t.Error(err)
    }
    t.Logf("%+v", res)

    defer res.Body.Close()
    t.Logf("%v", res.ContentLength)
    for {
        buf := make([]byte, 1024)
        n, e := res.Body.Read(buf)
        t.Logf("%s", buf[0:n])
        if e != nil {
            break
        }
    }
}

func TestGetObjectToFile(t *testing.T) {
    fname := "/tmp/test-get-object"
    err := BOS_CLIENT.GetObjectToFile(EXISTS_BUCKET, "test-put-object", fname)
    if err != nil {
        t.Error(err)
    }
    os.Remove(fname)
}

func TestGetObjectMeta(t *testing.T) {
    res, err := BOS_CLIENT.GetObjectMeta(EXISTS_BUCKET, "test-put-object")
    if err != nil {
        t.Error(err)
    }
    t.Logf("get object meta result: %+v", res)
}

func TestFetchObject(t *testing.T) {
    args := &api.FetchObjectArgs{api.FETCH_MODE_ASYNC, api.STORAGE_CLASS_COLD}
    res, err := BOS_CLIENT.FetchObject(EXISTS_BUCKET, "test-fetch-object",
        "https://cloud.baidu.com/doc/BOS/API.html", args)
    if err != nil {
        t.Error(err)
    }
    t.Logf("result: %+v", res)
}

func TestBasicFetchObject(t *testing.T) {
    res, err := BOS_CLIENT.BasicFetchObject(EXISTS_BUCKET, "test-fetch-object",
        "https://cloud.baidu.com/doc/BOS/API.html")
    if err != nil {
        t.Error(err)
    }
    t.Logf("result: %+v", res)

    res1, err1 := BOS_CLIENT.GetObjectMeta(EXISTS_BUCKET, "test-fetch-object")
    if err1 != nil {
        t.Error(err1)
    }
    t.Logf("meta: %+v", res1)
}

func TestAppendObject(t *testing.T) {
    args := &api.AppendObjectArgs{}
    body, _ := bce.NewBodyFromString("aaaaaaaaaaa")
    res, err := BOS_CLIENT.AppendObject(EXISTS_BUCKET, "test-append-object", body, args)
    if err != nil {
        t.Error(err)
    }
    t.Logf("%+v", res)
}

func TestBasicAppendObject(t *testing.T) {
    body, _ := bce.NewBodyFromString("bbbbbbbbbbb")
    res, err := BOS_CLIENT.BasicAppendObject(EXISTS_BUCKET, "test-append-object", body)
    if err != nil {
        t.Error(err)
    }
    t.Logf("%+v", res)
}

func TestBasicAppendObjectFromString(t *testing.T) {
    res, err := BOS_CLIENT.BasicAppendObjectFromString(EXISTS_BUCKET, "test-append-object", "123")
    if err != nil {
        t.Error(err)
    }
    t.Logf("%+v", res)
}

func TestBasicAppendObjectFromFile(t *testing.T) {
    fname := "/tmp/test-append-file"
    f, _ := os.Create(fname)
    f.WriteString("12345")
    f.Close()
    res, err := BOS_CLIENT.BasicAppendObjectFromFile(EXISTS_BUCKET, "test-append-object", fname)
    if err != nil {
        t.Error(err)
    }
    t.Logf("%+v", res)
    os.Remove(fname)
}

func TestDeleteObject(t *testing.T) {
    err := BOS_CLIENT.DeleteObject(EXISTS_BUCKET, "test-put-object")
    if err != nil {
        t.Error(err)
    }
}

func TestDeleteMultipleObjectsFromString(t *testing.T) {
    multiDeleteStr := `{
    "objects":[
        {"key": "aaaa"},
        {"key": "test-copy-object"},
        {"key": "test-append-object"},
        {"key": "cccc"}
    ]
}`
    res, err := BOS_CLIENT.DeleteMultipleObjectsFromString(EXISTS_BUCKET, multiDeleteStr)
    if err != nil {
        t.Error(err)
    }
    t.Logf("%+v", res)
}

func TestDeleteMultipleObjectsFromStruct(t *testing.T) {
    multiDeleteObj := &api.DeleteMultipleObjectsArgs{[]api.DeleteObjectArgs{
        api.DeleteObjectArgs{"1"}, api.DeleteObjectArgs{"test-fetch-object"}}}
    res, err := BOS_CLIENT.DeleteMultipleObjectsFromStruct(EXISTS_BUCKET, multiDeleteObj)
    if err != nil {
        t.Error(err)
    }
    t.Logf("%+v", res)
}

func TestInitiateMultipartUpload(t *testing.T) {
    args := &api.InitiateMultipartUploadArgs{Expires: "aaaaaaa"}
    res, err := BOS_CLIENT.InitiateMultipartUpload(EXISTS_BUCKET, "test-multipart-upload", "", args)
    if err != nil {
        t.Error(err)
    }
    t.Logf("%+v", res)

    err1 := BOS_CLIENT.AbortMultipartUpload(EXISTS_BUCKET,
        "test-multipart-upload", res.UploadId)
    if err1 != nil {
        t.Error(err1)
    }
}

func TestBasicInitiateMultipartUpload(t *testing.T) {
    res, err := BOS_CLIENT.BasicInitiateMultipartUpload(EXISTS_BUCKET, "test-multipart-upload")
    if err != nil {
        t.Error(err)
    }
    t.Logf("%+v", res)

    err1 := BOS_CLIENT.AbortMultipartUpload(EXISTS_BUCKET,
        "test-multipart-upload", res.UploadId)
    if err1 != nil {
        t.Error(err1)
    }
}

func TestListMultipartUploads(t *testing.T) {
    args := &api.ListMultipartUploadsArgs{MaxUploads: 10}
    res, err := BOS_CLIENT.ListMultipartUploads(EXISTS_BUCKET, args)
    if err != nil {
        t.Error(err)
    }
    t.Logf("%+v", res)
}

func TestBasicListMultipartUploads(t *testing.T) {
    res, err := BOS_CLIENT.BasicListMultipartUploads(EXISTS_BUCKET)
    if err != nil {
        t.Error(err)
    }
    t.Logf("%+v", res)
}

func TestUploadSuperFile(t *testing.T) {
    err := BOS_CLIENT.UploadSuperFile(EXISTS_BUCKET, "super-object", "/tmp/super-file", "")
    if err != nil {
        t.Error(err)
    }
}

