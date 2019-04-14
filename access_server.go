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

package main

import (
	"flag"

	glog "github.com/golang/glog"

	config "github.com/gotp/access_server/config"
	httpservice "github.com/gotp/access_server/service/http"
)

const (
	defaultConfigFilePath string = "../config/access_server.conf"
)

func init() {
	configManager := config.GetConfigManager()
	flag.StringVar(&configManager.ConfigFilePath, "config", defaultConfigFilePath, "Config file path")
	flag.Parse()

	// Load config
	if !configManager.Init() {
		glog.Fatal("Load server config failed!")
	}
	glog.Info("Load server config success")
	// Load router table
	routerTable := config.GetRouterTable()
	if !routerTable.Init(configManager.RouterTableFilePath) {
		glog.Fatal("Load router table failed!")
	}
	glog.Info("Load router table success")
}

func main() {
	// Create servers
	httpServer := httpservice.NewServer()
	// Start server
	glog.Info("Start server...")
	httpServer.Start()
}
