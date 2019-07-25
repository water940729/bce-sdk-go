package appblb

import "github.com/baidubce/bce-sdk-go/bce"

const (
	URI_PREFIX         = bce.URI_PREFIX + "v1"
	REQUEST_APPBLB_URL = "/appblb"

	APP_SERVER_GROUP_URL      = "/appservergroup"
	APP_SERVER_GROUP_PORT_URL = "/appservergroupport"
	BLB_RS_URL                = "/blbrs"
	BLB_RS_MOUNT_URL          = "/blbrsmount"
	BLB_RS_UNMOUNT_URL        = "/blbrsunmount"

	APP_LISTENER_URL      = "/listener"
	APP_TCPLISTENER_URL   = "/TCPlistener"
	APP_UDPLISTENER_URL   = "/UDPlistener"
	APP_HTTPLISTENER_URL  = "/HTTPlistener"
	APP_HTTPSLISTENER_URL = "/HTTPSlistener"
	APP_SSLLISTENER_URL   = "/SSLlistener"

	POLICYS_URL = "/policys"
)

// Client of EIP service is a kind of BceClient, so derived from BceClient
type Client struct {
	*bce.BceClient
}

func NewClient(ak, sk, endPoint string) (*Client, error) {
	client, err := bce.NewBceClientWithAkSk(ak, sk, endPoint)
	if err != nil {
		return nil, err
	}
	return &Client{client}, nil
}

func getAppBlbUri() string {
	return URI_PREFIX + REQUEST_APPBLB_URL
}

func getAppBlbUriWithId(id string) string {
	return URI_PREFIX + REQUEST_APPBLB_URL + "/" + id
}

func getAppServerGroupUri(id string) string {
	return URI_PREFIX + REQUEST_APPBLB_URL + "/" + id + APP_SERVER_GROUP_URL
}

func getAppServerGroupPortUri(id string) string {
	return URI_PREFIX + REQUEST_APPBLB_URL + "/" + id + APP_SERVER_GROUP_PORT_URL
}

func getBlbRsUri(id string) string {
	return URI_PREFIX + REQUEST_APPBLB_URL + "/" + id + BLB_RS_URL
}

func getBlbRsMountUri(id string) string {
	return URI_PREFIX + REQUEST_APPBLB_URL + "/" + id + BLB_RS_MOUNT_URL
}

func getBlbRsUnMountUri(id string) string {
	return URI_PREFIX + REQUEST_APPBLB_URL + "/" + id + BLB_RS_UNMOUNT_URL
}

func getAppListenerUri(id string) string {
	return URI_PREFIX + REQUEST_APPBLB_URL + "/" + id + APP_LISTENER_URL
}

func getAppTCPListenerUri(id string) string {
	return URI_PREFIX + REQUEST_APPBLB_URL + "/" + id + APP_TCPLISTENER_URL
}

func getAppUDPListenerUri(id string) string {
	return URI_PREFIX + REQUEST_APPBLB_URL + "/" + id + APP_UDPLISTENER_URL
}

func getAppHTTPListenerUri(id string) string {
	return URI_PREFIX + REQUEST_APPBLB_URL + "/" + id + APP_HTTPLISTENER_URL
}

func getAppHTTPSListenerUri(id string) string {
	return URI_PREFIX + REQUEST_APPBLB_URL + "/" + id + APP_HTTPSLISTENER_URL
}

func getAppSSLListenerUri(id string) string {
	return URI_PREFIX + REQUEST_APPBLB_URL + "/" + id + APP_SSLLISTENER_URL
}

func getPolicysUrl(id string) string {
	return URI_PREFIX + REQUEST_APPBLB_URL + "/" + id + POLICYS_URL
}
