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

	"gitlab.com/tymonx/go-patterns/guard"
)

var gInstance *Registry // nolint: gochecknoglobals
var gGuard guard.Guard  // nolint: gochecknoglobals
var gOnce sync.Once     // nolint: gochecknoglobals

// Add adds a new object with a given unique id to registry.
func Add(name string, object interface{}) (err error) {
	gGuard.Write(func() {
		err = getInstance().Add(name, object)
	})

	return err
}

// Adds adds new objects with given unique ids to registry.
func Adds(objects Objects) (err error) {
	gGuard.Write(func() {
		err = getInstance().Adds(objects)
	})

	return err
}

// Set sets an object with a given unique id to registry.
func Set(name string, object interface{}) {
	gGuard.Write(func() {
		getInstance().Set(name, object)
	})
}

// Sets sets objects with given unique ids to registry.
func Sets(objects Objects) {
	gGuard.Write(func() {
		getInstance().Sets(objects)
	})
}

// Get returns registered object by given name.
func Get(name string) (object interface{}, err error) {
	gGuard.Read(func() {
		object, err = getInstance().Get(name)
	})

	return object, err
}

// Gets returns registered objects by given names.
func Gets(names []string) (objects Objects, err error) {
	gGuard.Read(func() {
		objects, err = getInstance().Gets(names)
	})

	return objects, err
}

// GetAll returns all registered objects.
func GetAll() (objects Objects) {
	gGuard.Read(func() {
		objects = getInstance().GetAll()
	})

	return objects
}

// Remove removes registered object.
func Remove(name string) {
	gGuard.Write(func() {
		getInstance().Remove(name)
	})
}

// Removes removes registered object.
func Removes(names []string) {
	gGuard.Write(func() {
		getInstance().Removes(names)
	})
}

// RemoveAll removes all registered object.
func RemoveAll() {
	gGuard.Write(func() {
		getInstance().RemoveAll()
	})
}

// IsExist returns true if object with given name was registered, otherwise it returns false.
func IsExist(name string) (value bool) {
	gGuard.Read(func() {
		value = getInstance().IsExist(name)
	})

	return value
}

// IsExists returns true if all objects with given names were registered, otherwise it returns false.
func IsExists(names []string) (value bool) {
	gGuard.Read(func() {
		value = getInstance().IsExists(names)
	})

	return value
}

// IsEmpty returns true if there are no registered objects, otherwise it returns false.
func IsEmpty() (value bool) {
	gGuard.Read(func() {
		value = getInstance().IsEmpty()
	})

	return value
}

// Size returns number of registered objects.
func Size() (value int) {
	gGuard.Read(func() {
		value = getInstance().Size()
	})

	return value
}

// getInstance returns global registry instance.
func getInstance() *Registry {
	gOnce.Do(func() {
		gInstance = New()
	})

	return gInstance
}
