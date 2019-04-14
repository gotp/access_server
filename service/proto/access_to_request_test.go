/*
 * Copyright 2019 juzhongguoji
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package service_proto

import (
	"fmt"
	"testing"

	proto "github.com/gotp/proto"
)

const (
	accessHeaderBufferCase1 = `{
		"clientId": "cid",
		"clientType": 1,
		"version": "v1",
		"token": "token"
	}`

	requestHeaderBufferCase1 = `{"requestId":"requestid","clientId":"cid","clientType":1,"version":"v1","testFlag":false}`
)

func TestAccessHeaderBufferToRequestHeaderBuffer(gtest *testing.T) {
	var buffer string
	err := AccessHeaderBufferToRequestHeaderBuffer(accessHeaderBufferCase1, &buffer)
	if err != nil {
		gtest.Fatal("Test failed!")
	}
	fmt.Println(requestHeaderBufferCase1)
	if buffer != requestHeaderBufferCase1 {
		gtest.Fatal("Test failed!")
	}
}

func TestAccessHeaderToRequestHeader(gtest *testing.T) {
	var accessHeader proto.AccessRequestHeader
	var requestHeader proto.RequestHeader

	accessHeaderToRequestHeader(accessHeader, &requestHeader)
}
