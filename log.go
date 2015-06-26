// Copyright 2015 Google Inc. All rights reserved
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package kati

import (
	"bytes"
	"fmt"
	"sync"
)

var logMu sync.Mutex

func logAlways(f string, a ...interface{}) {
	var buf bytes.Buffer
	buf.WriteString("*kati*: ")
	buf.WriteString(f)
	buf.WriteByte('\n')
	logMu.Lock()
	fmt.Printf(buf.String(), a...)
	logMu.Unlock()
}

func logStats(f string, a ...interface{}) {
	if !LogFlag && !StatsFlag {
		return
	}
	logAlways(f, a...)
}

func logf(f string, a ...interface{}) {
	if !LogFlag {
		return
	}
	logAlways(f, a...)
}

func warn(loc srcpos, f string, a ...interface{}) {
	f = fmt.Sprintf("%s: warning: %s\n", loc, f)
	fmt.Printf(f, a...)
}

func warnNoPrefix(loc srcpos, f string, a ...interface{}) {
	f = fmt.Sprintf("%s: %s\n", loc, f)
	fmt.Printf(f, a...)
}
