package alibaba

import (
    "os"
    "github.com/aliyun/alibaba-cloud-sdk-go/services/bssopenapi"
    "github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
)

// Generalized function for creating Alibaba clients
func NewAlibabaClient(createClientFunc func(string, string, string) (interface{}, error)) (interface{}, error) {
    region := os.Getenv("ALIBABA_CLOUD_REGION")
    accessKey := os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_ID")
    secretKey := os.Getenv("ALIBABA_CLOUD_ACCESS_SECRET")

    return createClientFunc(region, accessKey, secretKey)
}

// BSS Client using generalized function
func BSS() (*bssopenapi.Client, error) {
    client, err := bssopenapi.NewClientWithAccessKey(os.Getenv("ALIBABA_CLOUD_REGION"), os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_ID"), os.Getenv("ALIBABA_CLOUD_ACCESS_SECRET"))
    if err != nil {
        return nil, err
    }
    return client, nil
}

// ECS Client using generalized function
func ECS() (*ecs.Client, error) {
    client, err := ecs.NewClientWithAccessKey(os.Getenv("ALIBABA_CLOUD_REGION"), os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_ID"), os.Getenv("ALIBABA_CLOUD_ACCESS_SECRET"))
    if err != nil {
        return nil, err
    }
    return client, nil
}