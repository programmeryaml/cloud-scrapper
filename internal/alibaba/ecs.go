package client

import (
	"os"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
)

func ECS() (*ecs.Client, error) {
	return ecs.NewClientWithAccessKey(os.Getenv("ALIBABA_CLOUD_REGION"), os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_ID"), os.Getenv("ALIBABA_CLOUD_ACCESS_SECRET"))
}
