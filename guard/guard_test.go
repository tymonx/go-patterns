// Copyright 2020 Tymoteusz Blazejczyk
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package guard_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/tymonx/go-patterns/guard"
)

func TestGuardRead(test *testing.T) {
	var g guard.Guard

	count := 0

	g.Read(func() {
		count = 1
	})

	assert.Equal(test, 1, count)
}

func TestGuardWrite(test *testing.T) {
	var g guard.Guard

	count := 0

	g.Write(func() {
		count++
	})

	assert.Equal(test, 1, count)
}
