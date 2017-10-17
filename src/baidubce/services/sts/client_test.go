package sts

import (
    "testing"
)

var (
    AK = "ab3e8280a5ff436eb5c5b9d7fa14fde9"
    SK = "2544e6b152a94241b8af9dbb25245c5e"
    ENDPOINT = "sts.bj.baidubce.com:80"
    CLIENT *Client
)

func init() {
    CLIENT, _ = NewClient(AK, SK, ENDPOINT)
}

func TestGetSessionToken(t *testing.T) {
    res, err := CLIENT.GetSessionToken(10, "")
    if err != nil {
        t.Error(err)
    }
    t.Logf("ak: %v", res.AccessKeyId)
    t.Logf("sk: %v", res.SecretAccessKey)
    t.Logf("sessionToken: %v", res.SessionToken)
    t.Logf("createTime: %v", res.CreateTime)
    t.Logf("expiration: %v", res.Expiration)
    t.Logf("userId: %v", res.UserId)
}
