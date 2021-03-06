package emr

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

// DeleteBatchResourceUsers invokes the emr.DeleteBatchResourceUsers API synchronously
// api document: https://help.aliyun.com/api/emr/deletebatchresourceusers.html
func (client *Client) DeleteBatchResourceUsers(request *DeleteBatchResourceUsersRequest) (response *DeleteBatchResourceUsersResponse, err error) {
	response = CreateDeleteBatchResourceUsersResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteBatchResourceUsersWithChan invokes the emr.DeleteBatchResourceUsers API asynchronously
// api document: https://help.aliyun.com/api/emr/deletebatchresourceusers.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteBatchResourceUsersWithChan(request *DeleteBatchResourceUsersRequest) (<-chan *DeleteBatchResourceUsersResponse, <-chan error) {
	responseChan := make(chan *DeleteBatchResourceUsersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteBatchResourceUsers(request)
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

// DeleteBatchResourceUsersWithCallback invokes the emr.DeleteBatchResourceUsers API asynchronously
// api document: https://help.aliyun.com/api/emr/deletebatchresourceusers.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteBatchResourceUsersWithCallback(request *DeleteBatchResourceUsersRequest, callback func(response *DeleteBatchResourceUsersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteBatchResourceUsersResponse
		var err error
		defer close(result)
		response, err = client.DeleteBatchResourceUsers(request)
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

// DeleteBatchResourceUsersRequest is the request struct for api DeleteBatchResourceUsers
type DeleteBatchResourceUsersRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceId      string           `position:"Query" name:"ResourceId"`
	UserIdList      *[]string        `position:"Query" name:"UserIdList"  type:"Repeated"`
	ResourceType    string           `position:"Query" name:"ResourceType"`
}

// DeleteBatchResourceUsersResponse is the response struct for api DeleteBatchResourceUsers
type DeleteBatchResourceUsersResponse struct {
	*responses.BaseResponse
	Paging    bool   `json:"Paging" xml:"Paging"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      bool   `json:"Data" xml:"Data"`
}

// CreateDeleteBatchResourceUsersRequest creates a request to invoke DeleteBatchResourceUsers API
func CreateDeleteBatchResourceUsersRequest() (request *DeleteBatchResourceUsersRequest) {
	request = &DeleteBatchResourceUsersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "DeleteBatchResourceUsers", "emr", "openAPI")
	return
}

// CreateDeleteBatchResourceUsersResponse creates a response to parse from DeleteBatchResourceUsers response
func CreateDeleteBatchResourceUsersResponse() (response *DeleteBatchResourceUsersResponse) {
	response = &DeleteBatchResourceUsersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
