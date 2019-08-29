package bssopenapi

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// QuerySettlementBill invokes the bssopenapi.QuerySettlementBill API synchronously
// api document: https://help.aliyun.com/api/bssopenapi/querysettlementbill.html
func (client *Client) QuerySettlementBill(request *QuerySettlementBillRequest) (response *QuerySettlementBillResponse, err error) {
	response = CreateQuerySettlementBillResponse()
	err = client.DoAction(request, response)
	return
}

// QuerySettlementBillWithChan invokes the bssopenapi.QuerySettlementBill API asynchronously
// api document: https://help.aliyun.com/api/bssopenapi/querysettlementbill.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QuerySettlementBillWithChan(request *QuerySettlementBillRequest) (<-chan *QuerySettlementBillResponse, <-chan error) {
	responseChan := make(chan *QuerySettlementBillResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QuerySettlementBill(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// QuerySettlementBillWithCallback invokes the bssopenapi.QuerySettlementBill API asynchronously
// api document: https://help.aliyun.com/api/bssopenapi/querysettlementbill.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QuerySettlementBillWithCallback(request *QuerySettlementBillRequest, callback func(response *QuerySettlementBillResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QuerySettlementBillResponse
		var err error
		defer close(result)
		response, err = client.QuerySettlementBill(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// QuerySettlementBillRequest is the request struct for api QuerySettlementBill
type QuerySettlementBillRequest struct {
	*requests.RpcRequest
	ProductCode      string           `position:"Query" name:"ProductCode"`
	IsHideZeroCharge requests.Boolean `position:"Query" name:"IsHideZeroCharge"`
	SubscriptionType string           `position:"Query" name:"SubscriptionType"`
	EndTime          string           `position:"Query" name:"EndTime"`
	BillingCycle     string           `position:"Query" name:"BillingCycle"`
	StartTime        string           `position:"Query" name:"StartTime"`
	OwnerId          requests.Integer `position:"Query" name:"OwnerId"`
	PageNum          requests.Integer `position:"Query" name:"PageNum"`
	Type             string           `position:"Query" name:"Type"`
	ProductType      string           `position:"Query" name:"ProductType"`
	PageSize         requests.Integer `position:"Query" name:"PageSize"`
}

// QuerySettlementBillResponse is the response struct for api QuerySettlementBill
type QuerySettlementBillResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateQuerySettlementBillRequest creates a request to invoke QuerySettlementBill API
func CreateQuerySettlementBillRequest() (request *QuerySettlementBillRequest) {
	request = &QuerySettlementBillRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("BssOpenApi", "2017-12-14", "QuerySettlementBill", "bssopenapi", "openAPI")
	return
}

// CreateQuerySettlementBillResponse creates a response to parse from QuerySettlementBill response
func CreateQuerySettlementBillResponse() (response *QuerySettlementBillResponse) {
	response = &QuerySettlementBillResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
