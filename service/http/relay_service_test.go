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

package http_service

import (
	"testing"
)

const (
	inputBufferCase1 = `{
		"clientId": "cid",
		"clientType": 1,
		"version": "v1",
		"token": "token",
		"data": {
			"str": "str",
			"int": 1,
			"intList": [1,2,3],
			"strList": ["1","2","3"],
			"map": {
				"dummy": 0
			}
		}
	}`
	outputHeaderCase1 = `{"clientId":"cid","clientType":1,"token":"token","version":"v1"}`
	outputDataCase1   = `{"int":1,"intList":[1,2,3],"map":{"dummy":0},"str":"str","strList":["1","2","3"]}`

	inputBufferCase2 = `{
		"clientId": "cid",
		"clientType": 1,
		"version": "v1",
		"token": "token",
		"data": {
		}
	}`

	inputBufferCase3 = `{
		"clientId": "cid",
		"clientType": 1,
		"version": "v1",
		"token": "token"
	}`

	inputBufferCase4 = `{
		"clientId": "cid",
		"clientType": 1,
		"version": "v1",
		"token": "token",
	}`

	inputBufferCase5 = `{
		"clientId": "cid",
		"clientType": 1,
		"version": "v1",
		"token": "token",
		"data": "data"
	}`

	inputBufferCase6 = `{
		"clientId": "cid",
		"clientType": 1,
		"version": "v1",
		"token": "token",
		"data": 1
	}`
)

func TestParseRequest_Normal_Case01(gtest *testing.T) {
	var header, data string

	err := relayService.parseRequest(inputBufferCase1, &header, &data)
	if err != nil {
		gtest.Fatal("Test failed!")
	}
	if header != outputHeaderCase1 {
		gtest.Fatal("Parse header failed! ", header, " != ", outputHeaderCase1)
	}
	if data != outputDataCase1 {
		gtest.Fatal("Parse data failed! ", data, " != ", outputDataCase1)
	}
}

func TestParseRequest_EmptyData_Case02(gtest *testing.T) {
	var header, data string

	err := relayService.parseRequest(inputBufferCase2, &header, &data)
	if err == nil {
		gtest.Fatal("Test failed!")
	} else {
		gtest.Log(err)
	}

	err = relayService.parseRequest(inputBufferCase3, &header, &data)
	if err == nil {
		gtest.Fatal("Test failed!")
	} else {
		gtest.Log(err)
	}

	err = relayService.parseRequest(inputBufferCase4, &header, &data)
	if err == nil {
		gtest.Fatal("Test failed!")
	} else {
		gtest.Log(err)
	}
}

func TestParseRequest_InvaildFormat_Case03(gtest *testing.T) {
	var header, data string

	err := relayService.parseRequest(inputBufferCase5, &header, &data)
	if err == nil {
		gtest.Fatal("Test failed!")
	} else {
		gtest.Log(err)
	}

	err = relayService.parseRequest(inputBufferCase6, &header, &data)
	if err == nil {
		gtest.Fatal("Test failed!")
	} else {
		gtest.Log(err)
	}
}
