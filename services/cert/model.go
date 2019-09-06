package cert

type CreateCertArgs struct {
	CertName        string `json:"certName"`
	CertServerData  string `json:"certServerData"`
	CertPrivateData string `json:"certPrivateData"`
	CertLinkData    string `json:"certLinkData,omitempty"`
	CertType        int    `json:"certType,omitempty"`
}

type CreateCertResult struct {
	CertName string `json:"certName"`
	CertId   string `json:"certId"`
}

type UpdateCertNameArgs struct {
	CertName string `json:"certName"`
}

type CertificateMeta struct {
	CertId         string `json:"certId"`
	CertName       string `json:"certName"`
	CertCommonName string `json:"certCommonName"`
	CertStartTime  string `json:"certStartTime"`
	CertStopTime   string `json:"certStopTime"`
	CertCreateTime string `json:"certCreateTime"`
	CertUpdateTime string `json:"certUpdateTime"`
	CertType       int    `json:"certType"`
}

type ListCertResult struct {
	Certs []CertificateMeta `json:"certs"`
}

type UpdateCertDataArgs struct {
	CertName        string `json:"certName"`
	CertServerData  string `json:"certServerData"`
	CertPrivateData string `json:"certPrivateData"`
	CertLinkData    string `json:"certLinkData,omitempty"`
	CertType        int    `json:"certType,omitempty"`
}