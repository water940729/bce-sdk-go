package api

import (
	"encoding/hex"
	"fmt"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/util/crypto"
)

const (
	URI_PREFIX = bce.URI_PREFIX + "v2"

	REQUEST_INSTANCE_URI = "/instance"
	REQUEST_VNC_SUFFIX   = "/vnc"

	REQUEST_VOLUME_URI = "/volume"
)

func getInstanceUri() string {
	return URI_PREFIX + REQUEST_INSTANCE_URI
}

func getInstanceUriWithId(id string) string {
	return URI_PREFIX + REQUEST_INSTANCE_URI + "/" + id
}

func getInstanceVNCUri(id string) string {
	return URI_PREFIX + REQUEST_INSTANCE_URI + "/" + id + REQUEST_VNC_SUFFIX
}

func Aes128EncryptUseSecreteKey(sk string, data string) (string, error) {
	if len(sk) < 16 {
		return "", fmt.Errorf("error secrete key")
	}

	crypted, err := crypto.EBCEncrypto([]byte(sk[:16]), []byte(data))
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(crypted), nil
}

func getVolumeUri() string {
	return URI_PREFIX + REQUEST_VOLUME_URI
}

func getVolumeUriWithId(id string) string {
	return URI_PREFIX + REQUEST_VOLUME_URI + "/" + id
}
