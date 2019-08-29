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

// QueryFinanceUnitResource invokes the bssopenapi.QueryFinanceUnitResource API synchronously
// api document: https://help.aliyun.com/api/bssopenapi/queryfinanceunitresource.html
func (client *Client) QueryFinanceUnitResource(request *QueryFinanceUnitResourceRequest) (response *QueryFinanceUnitResourceResponse, err error) {
	response = CreateQueryFinanceUnitResourceResponse()
	err = client.DoAction(request, response)
	return
}

// QueryFinanceUnitResourceWithChan invokes the bssopenapi.QueryFinanceUnitResource API asynchronously
// api document: https://help.aliyun.com/api/bssopenapi/queryfinanceunitresource.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryFinanceUnitResourceWithChan(request *QueryFinanceUnitResourceRequest) (<-chan *QueryFinanceUnitResourceResponse, <-chan error) {
	responseChan := make(chan *QueryFinanceUnitResourceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryFinanceUnitResource(request)
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

// QueryFinanceUnitResourceWithCallback invokes the bssopenapi.QueryFinanceUnitResource API asynchronously
// api document: https://help.aliyun.com/api/bssopenapi/queryfinanceunitresource.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryFinanceUnitResourceWithCallback(request *QueryFinanceUnitResourceRequest, callback func(response *QueryFinanceUnitResourceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryFinanceUnitResourceResponse
		var err error
		defer close(result)
		response, err = client.QueryFinanceUnitResource(request)
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

// QueryFinanceUnitResourceRequest is the request struct for api QueryFinanceUnitResource
type QueryFinanceUnitResourceRequest struct {
	*requests.RpcRequest
	PageNum  requests.Integer `position:"Query" name:"PageNum"`
	PageSize requests.Integer `position:"Query" name:"PageSize"`
	UnitId   requests.Integer `position:"Query" name:"UnitId"`
	OwnerUid requests.Integer `position:"Query" name:"OwnerUid"`
}

// QueryFinanceUnitResourceResponse is the response struct for api QueryFinanceUnitResource
type QueryFinanceUnitResourceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateQueryFinanceUnitResourceRequest creates a request to invoke QueryFinanceUnitResource API
func CreateQueryFinanceUnitResourceRequest() (request *QueryFinanceUnitResourceRequest) {
	request = &QueryFinanceUnitResourceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("BssOpenApi", "2017-12-14", "QueryFinanceUnitResource", "bssopenapi", "openAPI")
	return
}

// CreateQueryFinanceUnitResourceResponse creates a response to parse from QueryFinanceUnitResource response
func CreateQueryFinanceUnitResourceResponse() (response *QueryFinanceUnitResourceResponse) {
	response = &QueryFinanceUnitResourceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
