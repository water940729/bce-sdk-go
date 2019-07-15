package vpc

type CreateVPCArgs struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Cidr        string `json:"cidr"`
}

type CreateVPCResult struct {
	VPCID string `json:"vpcId"`
}

type ListVPCArgs struct {
	Marker    string
	MaxKeys   int

	// IsDefault is a string type,
	// so we can identify if it has been setted when the value is false.
	// NOTE: it can be only true or false.
	IsDefault string
}

type ListVPCResult struct {
	Marker      string `json:"marker"`
	IsTruncated bool   `json:"isTruncated"`
	NextMarker  string `json:"nextMarker"`
	MaxKeys     int    `json:"maxKeys"`
	VPCs        []VPC  `json:"vpcs"`
}

type VPC struct {
	VPCID       string `json:"vpcId"`
	Name        string `json:"name"`
	Cidr        string `json:"cidr"`
	Description string `json:"description"`
	IsDefault   bool   `json:"isDefault"`
}

type GetVPCDetailResult struct {
	VPC ShowVPCModel `json:"vpc"`
}

type ShowVPCModel struct {
	VPCId       string   `json:"vpcId"`
	Name        string   `json:"name"`
	Cidr        string   `json:"cidr"`
	Description string   `json:"description"`
	IsDefault   bool     `json:"isDefault"`
	Subnets     []Subnet `json:"subnets"`
}

type Subnet struct {
	SubnetId    string `json:"subnetId"`
	Name        string `json:"name"`
	ZoneName    string `json:"zoneName"`
	Cidr        string `json:"cidr"`
	VPCId       string `json:"vpcId"`
	SubnetType  string `json:"subnetType"`
	Description string `json:"description"`
	AvailableIp int    `json:"availableIp"`
}

type UpdateVPCArgs struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
