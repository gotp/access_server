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
	proto "github.com/gotp/proto"

	glog "github.com/golang/glog"
	jsonpb "github.com/golang/protobuf/jsonpb"
)

func AccessHeaderBufferToRequestHeaderBuffer(accessHeaderBuffer string, requestHeaderBuffer *string) error {
	var accessHeader proto.AccessRequestHeader
	var requestHeader proto.RequestHeader
	var jsonMarshaler = jsonpb.Marshaler{EmitDefaults: true}

	err := jsonpb.UnmarshalString(accessHeaderBuffer, &accessHeader)
	if err != nil {
		return err
	}
	glog.V(2).Infoln(accessHeader)

	accessHeaderToRequestHeader(accessHeader, &requestHeader)
	glog.V(2).Infoln(requestHeader)

	*requestHeaderBuffer, err = jsonMarshaler.MarshalToString(&requestHeader)
	if err != nil {
		return err
	}
	glog.V(2).Infoln(*requestHeaderBuffer)

	return nil
}

func accessHeaderToRequestHeader(accessHeader proto.AccessRequestHeader, requestHeader *proto.RequestHeader) {
	requestHeader.RequestId = "requestid"
	requestHeader.ClientId = accessHeader.ClientId
	requestHeader.ClientType = accessHeader.ClientType
	requestHeader.Version = accessHeader.Version
	requestHeader.TestFlag = false
}
