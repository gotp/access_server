/*
 * Copyright 2019 gotp
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
	"testing"
)

const (
	inputBufferCase1 = `{
		"retcode": 0,
		"retmsg": "ok",
		"requestId": "requestid"
	}`

	outputBufferCase1 = `{"retcode":0,"retmsg":"ok","requestId":"requestid"}`
)

func TestResponseHeaderBufferToAccessHeaderBuffer(gtest *testing.T) {
	var buffer string
	err := ResponseHeaderBufferToAccessHeaderBuffer(inputBufferCase1, &buffer)
	if err != nil {
		gtest.Fatal("Test failed!")
	}
	if buffer != outputBufferCase1 {
		gtest.Fatal("Test failed!")
	}
}
