package alicloud

// InstanceInfo represents the simplified structure for ECS instance information
type InstanceInfo struct {
	InstanceName       string   `json:"instanceName"`
	InstanceId         string   `json:"instanceId"`
	InstanceType       string   `json:"InstanceType"`
	InstanceChargeType string   `json:"InstanceChargeType"`
	RegionId           string   `json:"RegionId"`
	ZoneId             string   `json:"ZoneId"`
	HostName           string   `json:"HostName"`
	Status             string   `json:"Status"`
	Cpu                int      `json:"Cpu"`
	Memory             int      `json:"Memory"`
	OSNameEn           string   `json:"OSNameEn"`
	OSType             string   `json:"OSType"`
	PrivateIPAddresses []string `json:"privateIpAddresses"`
	Tags               []string `json:"tags"`
}

type Tag struct {
	Key   string
	Value string
}
