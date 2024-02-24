package model

import (
	entityAlibaba "cloud-scrapper/internal/entity/alicloud"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
)

// DescribeInstances calls the ECS DescribeInstances API and processes the response
func DescribeInstances(client *ecs.Client) ([]entityAlibaba.InstanceInfo, error) {
	request := ecs.CreateDescribeInstancesRequest()
	request.Scheme = "https"

	response, err := client.DescribeInstances(request)
	if err != nil {
		return nil, err
	}

	var instances []entityAlibaba.InstanceInfo
	for _, instance := range response.Instances.Instance {
		// Assuming the API response structure has InnerIpAddress containing an IpAddress list
		privateIPs := append([]string{}, instance.VpcAttributes.PrivateIpAddress.IpAddress...)
		// Assuming the API response structure has InnerIpAddress containing an IpAddress list
		var tags []string
		for _, tag := range instance.Tags.Tag {
			tagString := tag.TagKey + ":" + tag.TagValue // Example format, adjust as needed
			tags = append(tags, tagString)
		}
		info := entityAlibaba.InstanceInfo{
			InstanceName:       instance.InstanceName,
			InstanceId:         instance.InstanceId,
			InstanceType:       instance.InstanceType,
			InstanceChargeType: instance.InstanceChargeType,
			RegionId:           instance.RegionId,
			ZoneId:             instance.ZoneId,
			HostName:           instance.HostName,
			Status:             instance.Status,
			Cpu:                instance.Cpu,
			Memory:             instance.Memory,
			OSNameEn:           instance.OSNameEn,
			OSType:             instance.OSType,
			PrivateIPAddresses: privateIPs,
			Tags:               tags,
		}
		instances = append(instances, info)
	}

	return instances, nil
}
