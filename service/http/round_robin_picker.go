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
	"sync"
)

type RoundRobinPicker struct {
	list  []string
	mutex sync.Mutex
	next  int
}

func (picker *RoundRobinPicker) Pick() string {
	if len(picker.list) <= 0 {
		return ""
	}

	picker.mutex.Lock()
	item := picker.list[picker.next]
	picker.next = (picker.next + 1) % len(picker.list)
	picker.mutex.Unlock()
	return item
}
