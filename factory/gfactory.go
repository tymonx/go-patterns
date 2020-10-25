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

package factory

import (
	"sync"
)

var gInstance *Factory  // nolint: gochecknoglobals
var gMutex sync.RWMutex // nolint: gochecknoglobals
var gOnce sync.Once     // nolint: gochecknoglobals

// Create creates a new object based on given name.
func Create(name string, arguments ...interface{}) (object interface{}, err error) {
	gMutex.RLock()
	defer gMutex.RUnlock()

	return getInstance().Create(name, arguments...)
}

// Creates creates a list of new objects based on given names.
func Creates(names []string, arguments ...interface{}) (objects []interface{}, err error) {
	gMutex.RLock()
	defer gMutex.RUnlock()

	return getInstance().Creates(names, arguments...)
}

// Add adds a new constructor with a given unique id to factory.
func Add(name string, constructor Constructor) error {
	gMutex.Lock()
	defer gMutex.Unlock()

	return getInstance().Add(name, constructor)
}

// Adds adds new constructors with given unique ids to factory.
func Adds(constructors Constructors) error {
	gMutex.Lock()
	defer gMutex.Unlock()

	return getInstance().Adds(constructors)
}

// Set sets an constructor with a given unique id to factory.
func Set(name string, constructor Constructor) error {
	gMutex.Lock()
	defer gMutex.Unlock()

	return getInstance().Set(name, constructor)
}

// Sets sets constructors with given unique ids to factory.
func Sets(constructors Constructors) error {
	gMutex.Lock()
	defer gMutex.Unlock()

	return getInstance().Sets(constructors)
}

// Get returns registered constructor by given name.
func Get(name string) (constructor Constructor, err error) {
	gMutex.RLock()
	defer gMutex.RUnlock()

	return getInstance().Get(name)
}

// Gets returns registered constructors by given names.
func Gets(names []string) (constructors Constructors, err error) {
	gMutex.RLock()
	defer gMutex.RUnlock()

	return getInstance().Gets(names)
}

// GetAll returns all registered constructors.
func GetAll() Constructors {
	gMutex.RLock()
	defer gMutex.RUnlock()

	return getInstance().GetAll()
}

// Remove removes registered constructor.
func Remove(name string) {
	gMutex.Lock()
	defer gMutex.Unlock()

	getInstance().Remove(name)
}

// Removes removes registered constructor.
func Removes(names []string) {
	gMutex.Lock()
	defer gMutex.Unlock()

	getInstance().Removes(names)
}

// RemoveAll removes all registered constructor.
func RemoveAll() {
	gMutex.Lock()
	defer gMutex.Unlock()

	getInstance().RemoveAll()
}

// IsExist returns true if constructor with given name was registered, otherwise it returns false.
func IsExist(name string) bool {
	gMutex.RLock()
	defer gMutex.RUnlock()

	return getInstance().IsExist(name)
}

// IsExists returns true if all object constructors with given names were registered, otherwise it returns false.
func IsExists(names []string) bool {
	gMutex.RLock()
	defer gMutex.RUnlock()

	return getInstance().IsExists(names)
}

// IsEmpty returns true if there are no registered object constructors, otherwise it returns false.
func IsEmpty() bool {
	gMutex.RLock()
	defer gMutex.RUnlock()

	return getInstance().IsEmpty()
}

// Size returns number of registered object constructors.
func Size() int {
	gMutex.RLock()
	defer gMutex.RUnlock()

	return getInstance().Size()
}

// getInstance returns global factory instance.
func getInstance() *Factory {
	gOnce.Do(func() {
		gInstance = New()
	})

	return gInstance
}
