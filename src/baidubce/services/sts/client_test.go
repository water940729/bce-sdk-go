package sts

import (
    "encoding/json"
    "os"
    "path/filepath"
    "runtime"
    "testing"

    "baidubce/util/log"
)

var CLIENT *Client

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

    CLIENT, _ = NewClient(confObj.AK, confObj.SK)
    log.SetLogHandler(log.STDERR)
    log.SetLogLevel(log.INFO)
}

func TestGetSessionToken(t *testing.T) {
    acl := `
{
    "id":"10eb6f5ff6ff4605bf044313e8f3ffa5",
    "accessControlList": [
    {
        "eid": "10eb6f5ff6ff4605bf044313e8f3ffa5-1",
        "effect": "Deny",
        "resource": ["bos-rd-ssy/*"],
        "region": "bj",
        "service": "bce:bos",
        "permission": ["WRITE"]
    }
    ]
}`
    res, err := CLIENT.GetSessionToken(10, acl)
    if err != nil {
        t.Error(err)
    } else {
        t.Logf("ak: %v", res.AccessKeyId)
        t.Logf("sk: %v", res.SecretAccessKey)
        t.Logf("sessionToken: %v", res.SessionToken)
        t.Logf("createTime: %v", res.CreateTime)
        t.Logf("expiration: %v", res.Expiration)
        t.Logf("userId: %v", res.UserId)
    }
}
