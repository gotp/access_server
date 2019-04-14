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
	"testing"
)

var picker RoundRobinPicker

func TestPick_Normal_Case01(gtest *testing.T) {
	var data string

	picker.list = []string{"1", "2", "3", "4", "5"}
	data = picker.Pick()
	if data != "1" {
		gtest.Fatal(data, " != 1")
	}
	data = picker.Pick()
	if data != "2" {
		gtest.Fatal(data, " != 2")
	}
}
