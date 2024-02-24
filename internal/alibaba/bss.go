package client

import (
	"os"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/bssopenapi"
)

func BSS() (*bssopenapi.Client, error) {
	return bssopenapi.NewClientWithAccessKey(os.Getenv("ALIBABA_CLOUD_REGION"), os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_ID"), os.Getenv("ALIBABA_CLOUD_ACCESS_SECRET"))
}
