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

	"gitlab.com/tymonx/go-patterns/guard"
)

var gInstance *Factory // nolint: gochecknoglobals
var gGuard guard.Guard // nolint: gochecknoglobals
var gOnce sync.Once    // nolint: gochecknoglobals

// Create creates a new object based on given name.
func Create(name string, arguments ...interface{}) (object interface{}, err error) {
	gGuard.Read(func() {
		object, err = getInstance().Create(name, arguments...)
	})

	return object, err
}

// Creates creates a list of new objects based on given names.
func Creates(names []string, arguments ...interface{}) (objects []interface{}, err error) {
	gGuard.Read(func() {
		objects, err = getInstance().Creates(names, arguments...)
	})

	return objects, err
}

// Add adds a new constructor with a given unique id to factory.
func Add(name string, constructor Constructor) (err error) {
	gGuard.Write(func() {
		err = getInstance().Add(name, constructor)
	})

	return err
}

// Adds adds new constructors with given unique ids to factory.
func Adds(constructors Constructors) (err error) {
	gGuard.Write(func() {
		err = getInstance().Adds(constructors)
	})

	return err
}

// Set sets an constructor with a given unique id to factory.
func Set(name string, constructor Constructor) (err error) {
	gGuard.Write(func() {
		err = getInstance().Set(name, constructor)
	})

	return err
}

// Sets sets constructors with given unique ids to factory.
func Sets(constructors Constructors) (err error) {
	gGuard.Write(func() {
		err = getInstance().Sets(constructors)
	})

	return err
}

// Get returns registered constructor by given name.
func Get(name string) (constructor Constructor, err error) {
	gGuard.Read(func() {
		constructor, err = getInstance().Get(name)
	})

	return constructor, err
}

// Gets returns registered constructors by given names.
func Gets(names []string) (constructors Constructors, err error) {
	gGuard.Read(func() {
		constructors, err = getInstance().Gets(names)
	})

	return constructors, err
}

// GetAll returns all registered constructors.
func GetAll() (constructors Constructors) {
	gGuard.Read(func() {
		constructors = getInstance().GetAll()
	})

	return constructors
}

// Remove removes registered constructor.
func Remove(name string) {
	gGuard.Write(func() {
		getInstance().Remove(name)
	})
}

// Removes removes registered constructor.
func Removes(names []string) {
	gGuard.Write(func() {
		getInstance().Removes(names)
	})
}

// RemoveAll removes all registered constructor.
func RemoveAll() {
	gGuard.Write(func() {
		getInstance().RemoveAll()
	})
}

// IsExist returns true if constructor with given name was registered, otherwise it returns false.
func IsExist(name string) (value bool) {
	gGuard.Read(func() {
		value = getInstance().IsExist(name)
	})

	return value
}

// IsExists returns true if all object constructors with given names were registered, otherwise it returns false.
func IsExists(names []string) (value bool) {
	gGuard.Read(func() {
		value = getInstance().IsExists(names)
	})

	return value
}

// IsEmpty returns true if there are no registered object constructors, otherwise it returns false.
func IsEmpty() (value bool) {
	gGuard.Read(func() {
		value = getInstance().IsEmpty()
	})

	return value
}

// Size returns number of registered object constructors.
func Size() (value int) {
	gGuard.Read(func() {
		value = getInstance().Size()
	})

	return value
}

// getInstance returns global factory instance.
func getInstance() *Factory {
	gOnce.Do(func() {
		gInstance = New()
	})

	return gInstance
}
