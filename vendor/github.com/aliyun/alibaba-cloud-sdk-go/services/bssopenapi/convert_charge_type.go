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

// ConvertChargeType invokes the bssopenapi.ConvertChargeType API synchronously
// api document: https://help.aliyun.com/api/bssopenapi/convertchargetype.html
func (client *Client) ConvertChargeType(request *ConvertChargeTypeRequest) (response *ConvertChargeTypeResponse, err error) {
	response = CreateConvertChargeTypeResponse()
	err = client.DoAction(request, response)
	return
}

// ConvertChargeTypeWithChan invokes the bssopenapi.ConvertChargeType API asynchronously
// api document: https://help.aliyun.com/api/bssopenapi/convertchargetype.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ConvertChargeTypeWithChan(request *ConvertChargeTypeRequest) (<-chan *ConvertChargeTypeResponse, <-chan error) {
	responseChan := make(chan *ConvertChargeTypeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ConvertChargeType(request)
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

// ConvertChargeTypeWithCallback invokes the bssopenapi.ConvertChargeType API asynchronously
// api document: https://help.aliyun.com/api/bssopenapi/convertchargetype.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ConvertChargeTypeWithCallback(request *ConvertChargeTypeRequest, callback func(response *ConvertChargeTypeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ConvertChargeTypeResponse
		var err error
		defer close(result)
		response, err = client.ConvertChargeType(request)
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

// ConvertChargeTypeRequest is the request struct for api ConvertChargeType
type ConvertChargeTypeRequest struct {
	*requests.RpcRequest
	Period           requests.Integer `position:"Query" name:"Period"`
	ProductCode      string           `position:"Query" name:"ProductCode"`
	SubscriptionType string           `position:"Query" name:"SubscriptionType"`
	OwnerId          requests.Integer `position:"Query" name:"OwnerId"`
	ProductType      string           `position:"Query" name:"ProductType"`
	InstanceId       string           `position:"Query" name:"InstanceId"`
}

// ConvertChargeTypeResponse is the response struct for api ConvertChargeType
type ConvertChargeTypeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateConvertChargeTypeRequest creates a request to invoke ConvertChargeType API
func CreateConvertChargeTypeRequest() (request *ConvertChargeTypeRequest) {
	request = &ConvertChargeTypeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("BssOpenApi", "2017-12-14", "ConvertChargeType", "bssopenapi", "openAPI")
	return
}

// CreateConvertChargeTypeResponse creates a response to parse from ConvertChargeType response
func CreateConvertChargeTypeResponse() (response *ConvertChargeTypeResponse) {
	response = &ConvertChargeTypeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
