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

package http_service

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	config "github.com/gotp/access_server/config"
	proto "github.com/gotp/access_server/service/proto"

	glog "github.com/golang/glog"
	//"github.com/golang/protobuf/jsonpb"
)

//var jsonMarshaler = jsonpb.Marshaler{EmitDefaults: true}

var relayService RelayService

const (
	urlPathSegmentNum = 5 // /Project/Server/Service/Function -> ['', 'Project', 'Server', 'Service', 'Function']
)

type RelayService struct {
	balancer map[string]RoundRobinPicker
}

func InitRelayService() {
	relayService.initBalancer()

	RegisterServiceHandler("/",
		func(response http.ResponseWriter, request *http.Request) {
			relayService.RelayRequest(response, request)
		},
	)
}

func (this *RelayService) initBalancer() {
	this.balancer = make(map[string]RoundRobinPicker)
	routerTable := config.GetRouterTable()
	for name, addrs := range routerTable.Addrs {
		this.balancer[name] = RoundRobinPicker{
			list: addrs,
			next: 0,
		}
	}
}

func (this *RelayService) RelayRequest(response http.ResponseWriter, request *http.Request) {
	var accessHeader, requestData string
	var requestHeader string

	glog.V(2).Info("Get request: ", request.URL)
	requestBody, err := ioutil.ReadAll(request.Body)
	if err != nil {
		http.Error(response, err.Error(), http.StatusServiceUnavailable)
		return
	}
	glog.V(2).Info("Parse request")
	err = this.parseRequest(string(requestBody), &accessHeader, &requestData)
	if err != nil {
		http.Error(response, err.Error(), http.StatusServiceUnavailable)
		return
	}
	glog.V(2).Info("Relay address")
	relayAddress := this.getRelayAddress(request.URL)
	glog.Infoln("Get Relay address: ", relayAddress)
	glog.V(2).Info("Prepare relay request")
	err = proto.AccessHeaderBufferToRequestHeaderBuffer(accessHeader, &requestHeader)
	if err != nil {
		http.Error(response, err.Error(), http.StatusServiceUnavailable)
		return
	}
	glog.V(2).Info("Relay request")
	client := &http.Client{}
	relayRequest, err := http.NewRequest("POST",
		relayAddress,
		request.Body)
	if err != nil {
		http.Error(response, err.Error(), http.StatusServiceUnavailable)
		return
	}
	relayResponse, err := client.Do(relayRequest)
	if err != nil {
		http.Error(response, err.Error(), http.StatusServiceUnavailable)
		return
	}
	defer relayResponse.Body.Close()
	//this.copyHeader(response.Header(), relayResponse.Header)
	response.WriteHeader(relayResponse.StatusCode)
	// TODO: 转换为外部请求头
	io.Copy(response, relayResponse.Body)
	glog.V(2).Info("Send response")
}

func (this *RelayService) parseRequest(requestBody string, header *string, data *string) error {
	var jsonObj map[string]interface{}

	err := json.Unmarshal([]byte(requestBody), &jsonObj)
	if err != nil {
		return err
	}
	dataObj, found := jsonObj["data"]
	if !found {
		return errors.New("No data field find in request!")
	}
	dataMap, ok := dataObj.(map[string]interface{})
	if !ok {
		return errors.New("Invaild data field type in request!")
	}
	if len(dataMap) == 0 {
		return errors.New("No empty data field in request!")
	}
	dataBytes, err := json.Marshal(dataObj)
	if err != nil {
		return err
	}
	*data = string(dataBytes)
	delete(jsonObj, "data")
	headerBytes, err := json.Marshal(jsonObj)
	if err != nil {
		return err
	}
	*header = string(headerBytes)
	return nil
}

func (this *RelayService) getRelayAddress(reqUrl *url.URL) string {
	routerName := this.getRouterName(reqUrl.Path)
	glog.V(2).Info(routerName)
	picker, found := this.balancer[routerName]
	if !found {
		return ""
	}
	reqUrl.Host = picker.Pick()

	if reqUrl.Scheme == "" {
		reqUrl.Scheme = "http"
	}

	return reqUrl.String()
}

func (this *RelayService) copyHeader(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}

func (this *RelayService) getRouterName(urlPath string) string {
	segments := strings.Split(urlPath, "/")
	if len(segments) != urlPathSegmentNum {
		glog.V(2).Infoln("Invalid url, path segment num ", len(segments),
			"(expect ", urlPathSegmentNum, ")")
		return ""
	}

	return segments[1] + "." + segments[2] + "." + segments[3]
}
