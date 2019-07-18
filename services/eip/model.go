package eip

type Reservation struct {
	ReservationLength   int    `json:"reservationLength,omitempty"`
	ReservationTimeUnit string `json:"reservationTimeUnit,omitempty"`
}

type Billing struct {
	PaymentTiming string       `json:"paymentTiming,omitempty"`
	BillingMethod string       `json:"billingMethod,omitempty"`
	Reservation   *Reservation `json:"reservation,omitempty"`
}

type CreateEipArgs struct {
	Name            string   `json:"name,omitempty"`
	BandWidthInMbps int      `json:"bandwidthInMbps"`
	Billing         *Billing `json:"billing"`
	ClientToken     string   `json:"-"`
}

type CreateEipResult struct {
	Eip string `json:"eip"`
}

type ResizeEipArgs struct {
	NewBandWidthInMbps int    `json:"newBandwidthInMbps"`
	ClientToken        string `json:"-"`
}

type BindEipArgs struct {
	InstanceType string `json:"instanceType"`
	InstanceId   string `json:"instanceId"`
	ClientToken  string `json:"-"`
}

type ListEipArgs struct {
	Eip          string
	InstanceType string
	InstanceId   string
	Marker       string
	MaxKeys      int
	Status       string
}

type ListEipResult struct {
	Marker      string     `json:"marker"`
	MaxKeys     int        `json:"maxKeys"`
	NextMarker  string     `json:"nextMarker"`
	IsTruncated bool       `json:"isTruncated"`
	EipList     []EipModel `json:"eipList"`
}

type EipModel struct {
	Name            string `json:"name"`
	Eip             string `json:"eip"`
	Status          string `json:"status"`
	EipInstanceType string `json:"eipInstanceType"`
	InstanceType    string `json:"instanceType"`
	InstanceId      string `json:"instanceId"`
	ShareGroupId    string `json:"shareGroupId"`
	BandWidthInMbps int    `json:"bandwidthInMbps"`
	PaymentTiming   string `json:"paymentTiming"`
	BillingMethod   string `json:"billingMethod"`
	CreateTime      string `json:"createTime"`
	ExpireTime      string `json:"expireTime"`
}

type PurchaseReservedEipArgs struct {
	Billing     *Billing `json:"billing"`
	ClientToken string   `json:"clientToken"`
}
