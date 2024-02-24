package model

import (
	entityAlibaba "cloud-scrapper/internal/entity/alicloud"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/bssopenapi"
	"github.com/labstack/gommon/log"
)

func DescribePrices(client *bssopenapi.Client) ([]entityAlibaba.PriceInfo, error) {
	request := bssopenapi.CreateQueryInstanceBillRequest()
	request.Scheme = "https"
	request.BillingCycle = "2024-02"
	request.BillingDate = "2024-02-01"
	request.Granularity = "DAILY"
	request.ProductCode = "ECS"

	var prices []entityAlibaba.PriceInfo

	for count := request.PageSize; count == request.PageSize; {
		resp, err := client.QueryInstanceBill(request)
		if err != nil {
			log.Errorf("Error QueryInstanceBill: ", err)
			return nil, err
		}

		for i := range resp.Data.Items.Item {
			item := resp.Data.Items.Item[i]

			// Assuming prices is a slice of entityAlibaba.PriceInfo and item is of type Item
			prices = append(prices, entityAlibaba.PriceInfo{
				// Map fields from item to the PriceInfo struct
				// Example:
				InstanceID:        item.InstanceID,
				InvoiceDiscount:   item.InvoiceDiscount,
				PretaxGrossAmount: item.PretaxGrossAmount,
				Item:              item.Item,
				OutstandingAmount: item.OutstandingAmount,
				ProductCode:       item.ProductCode,
				InstanceSpec:      item.InstanceSpec,
				BillingType:       item.BillingType,
				Tag:               item.Tag,
				Region:            item.Region,
				NickName:          item.NickName,
				Currency:          item.Currency,
				DeductedByCoupons: item.DeductedByCoupons,
				Zone:              item.Zone,
				PretaxAmount:      item.PretaxAmount,
				IntranetIP:        item.IntranetIP,
				// Add any other fields here
			})

		}
		request.PageNum = requests.NewInteger(resp.Data.PageNum + 1)
		count = requests.NewInteger(len(resp.Data.Items.Item))
	}

	return prices, nil
}
