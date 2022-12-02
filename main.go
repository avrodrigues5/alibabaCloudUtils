package main

import (
	"fmt"
	"os"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/bssopenapi"
)

func main() {
	argsWithProg := os.Args[1:]
	accessKey := argsWithProg[0]
	accessSecret := argsWithProg[1]
	regionID := argsWithProg[2]
	ossBucket := argsWithProg[3]
	subscribeType := argsWithProg[4]
	client, err := bssopenapi.NewClientWithAccessKey(regionID, accessKey, accessSecret)

	request := bssopenapi.CreateSubscribeBillToOSSRequest()
	request.Scheme = "https"
	request.SubscribeBucket = ossBucket
	request.SubscribeType = subscribeType

	response, err := client.SubscribeBillToOSS(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("response is %#v\n", response)
}
