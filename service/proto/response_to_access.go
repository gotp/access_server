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

func ResponseHeaderBufferToAccessHeaderBuffer(responseHeaderBuffer string, accessHeaderBuffer *string) error {
	var responseHeader proto.ResponseHeader
	var accessHeader proto.AccessResponseHeader
	var jsonMarshaler = jsonpb.Marshaler{EmitDefaults: true}

	err := jsonpb.UnmarshalString(responseHeaderBuffer, &responseHeader)
	if err != nil {
		return err
	}
	glog.V(2).Infoln(responseHeader)

	responseHeaderToAccessHeader(responseHeader, &accessHeader)
	glog.V(2).Infoln(accessHeader)

	*accessHeaderBuffer, err = jsonMarshaler.MarshalToString(&accessHeader)
	if err != nil {
		return err
	}
	glog.V(2).Infoln(*accessHeaderBuffer)

	return nil
}

func responseHeaderToAccessHeader(responseHeader proto.ResponseHeader, accessHeader *proto.AccessResponseHeader) {
	accessHeader.RequestId = responseHeader.RequestId
	accessHeader.Retcode = int32(responseHeader.Retcode)
	accessHeader.Retmsg = responseHeader.Retmsg
}
