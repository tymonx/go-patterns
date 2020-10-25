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

package registry

import (
	"sync"
)

var gInstance *Registry // nolint: gochecknoglobals
var gMutex sync.RWMutex // nolint: gochecknoglobals
var gOnce sync.Once     // nolint: gochecknoglobals

// Add adds a new object with a given unique id to registry.
func Add(name string, object interface{}) error {
	gMutex.Lock()
	defer gMutex.Unlock()

	return getInstance().Add(name, object)
}

// Adds adds new objects with given unique ids to registry.
func Adds(objects Objects) error {
	gMutex.Lock()
	defer gMutex.Unlock()

	return getInstance().Adds(objects)
}

// Set sets an object with a given unique id to registry.
func Set(name string, object interface{}) {
	gMutex.Lock()
	defer gMutex.Unlock()

	getInstance().Set(name, object)
}

// Sets sets objects with given unique ids to registry.
func Sets(objects Objects) {
	gMutex.Lock()
	defer gMutex.Unlock()

	getInstance().Sets(objects)
}

// Get returns registered object by given name.
func Get(name string) (object interface{}, err error) {
	gMutex.RLock()
	defer gMutex.RUnlock()

	return getInstance().Get(name)
}

// Gets returns registered objects by given names.
func Gets(names []string) (objects Objects, err error) {
	gMutex.RLock()
	defer gMutex.RUnlock()

	return getInstance().Gets(names)
}

// GetAll returns all registered objects.
func GetAll() Objects {
	gMutex.RLock()
	defer gMutex.RUnlock()

	return getInstance().GetAll()
}

// Remove removes registered object.
func Remove(name string) {
	gMutex.Lock()
	defer gMutex.Unlock()

	getInstance().Remove(name)
}

// Removes removes registered object.
func Removes(names []string) {
	gMutex.Lock()
	defer gMutex.Unlock()

	getInstance().Removes(names)
}

// RemoveAll removes all registered object.
func RemoveAll() {
	gMutex.Lock()
	defer gMutex.Unlock()

	getInstance().RemoveAll()
}

// IsExist returns true if object with given name was registered, otherwise it returns false.
func IsExist(name string) bool {
	gMutex.RLock()
	defer gMutex.RUnlock()

	return getInstance().IsExist(name)
}

// IsExists returns true if all objects with given names were registered, otherwise it returns false.
func IsExists(names []string) bool {
	gMutex.RLock()
	defer gMutex.RUnlock()

	return getInstance().IsExists(names)
}

// IsEmpty returns true if there are no registered objects, otherwise it returns false.
func IsEmpty() bool {
	gMutex.RLock()
	defer gMutex.RUnlock()

	return getInstance().IsEmpty()
}

// Size returns number of registered objects.
func Size() int {
	gMutex.RLock()
	defer gMutex.RUnlock()

	return getInstance().Size()
}

// getInstance returns global registry instance.
func getInstance() *Registry {
	gOnce.Do(func() {
		gInstance = New()
	})

	return gInstance
}
