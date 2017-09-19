package sts

import (
    "testing"
    "os"

    "baidubce/util"
    "baidubce/services/sts/model"
)

var (
    AK = "ab3e8280a5ff436eb5c5b9d7fa14fde9"
    SK = "2544e6b152a94241b8af9dbb25245c5e"
    ENDPOINT = "sts.bj.baidubce.com"
    CLIENT *Client
)

func init() {
    CLIENT, _ = NewClient(AK, SK, ENDPOINT)

    // Clear the debug info
    null, _ := os.Open("/dev/null")
    util.LOGGER = util.GetLogger(null, util.DEFAULT_LOGGER_PREFIX)
}

func TestGetSessionToken(t *testing.T) {
    request  := model.NewGetSessionTokenRequest(30, "")
    response := model.NewGetSessionTokenResponse()
    err := CLIENT.GetSessionToken(request, response)

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
    t.Logf("%v", response.AccessKeyId())
    t.Logf("%v", response.SecretAccessKey())
    t.Logf("%v", response.SessionToken())
    t.Logf("%v", response.CreateTime())
    t.Logf("%v", response.Expiration())
    t.Logf("%v", response.UserId())
}

