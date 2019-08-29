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

// Task is a nested struct in emr response
type Task struct {
	DataSourceId     string `json:"DataSourceId" xml:"DataSourceId"`
	TaskResultDetail string `json:"TaskResultDetail" xml:"TaskResultDetail"`
	TaskStatus       string `json:"TaskStatus" xml:"TaskStatus"`
	EndTime          int64  `json:"EndTime" xml:"EndTime"`
	BizId            string `json:"BizId" xml:"BizId"`
	StartTime        int64  `json:"StartTime" xml:"StartTime"`
	GmtCreate        int64  `json:"GmtCreate" xml:"GmtCreate"`
	ExecuteTime      int64  `json:"ExecuteTime" xml:"ExecuteTime"`
	TaskProcess      int    `json:"TaskProcess" xml:"TaskProcess"`
	TaskObject       string `json:"TaskObject" xml:"TaskObject"`
	GmtModified      int64  `json:"GmtModified" xml:"GmtModified"`
	TaskType         string `json:"TaskType" xml:"TaskType"`
	TriggerUser      string `json:"TriggerUser" xml:"TriggerUser"`
	TriggerType      string `json:"TriggerType" xml:"TriggerType"`
	TaskDetail       string `json:"TaskDetail" xml:"TaskDetail"`
}
