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
	"gitlab.com/tymonx/go-error/rterror"
	"gitlab.com/tymonx/go-patterns/registry"
)

// Names defines a list of names.
type Names []string

// Factory defines a factory instance that can create registered object types.
type Factory struct {
	registry registry.Registry
}

// New creates a new factory instance.
func New() *Factory {
	return &Factory{
		registry: *registry.New(),
	}
}

// Create creates a new object based on given name.
func (f *Factory) Create(name string, arguments ...interface{}) (object interface{}, err error) {
	var constructor Constructor

	if constructor, err = f.Get(name); err != nil {
		return nil, err
	}

	if object, err = constructor(arguments...); err != nil {
		return nil, rterror.New("cannot create object", name, err)
	}

	if object == nil {
		return nil, rterror.New("object was not created", name)
	}

	return object, nil
}

// Creates creates a list of new objects based on given names.
func (f *Factory) Creates(names []string, arguments ...interface{}) (objects []interface{}, err error) {
	objects = make([]interface{}, 0, len(names))
	errs := make([]interface{}, 0, len(names))

	for _, name := range names {
		var object interface{}

		if object, err = f.Create(name, arguments...); err != nil {
			errs = append(errs, err)

			continue
		}

		objects = append(objects, object)
	}

	if len(errs) != 0 {
		return objects, rterror.New("cannot create objects", errs...)
	}

	return objects, nil
}

// Add adds an object constructor with a given unique id to registry.
func (f *Factory) Add(name string, constructor Constructor) error {
	if constructor == nil {
		return rterror.New("object constructor cannot be nil", name)
	}

	return f.registry.Add(name, constructor)
}

// Adds adds new object constructors with given unique ids to registry.
func (f *Factory) Adds(constructors Constructors) error {
	errs := make([]interface{}, 0, len(constructors))

	for name, constructor := range constructors {
		if err := f.Add(name, constructor); err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) != 0 {
		return rterror.New("cannot add object constructors", errs...)
	}

	return nil
}

// Set sets an object constructor with a given unique id to registry.
func (f *Factory) Set(name string, constructor Constructor) error {
	if constructor == nil {
		return rterror.New("object constructor cannot be nil", name)
	}

	f.registry.Set(name, constructor)

	return nil
}

// Sets sets object constructors with given unique ids to registry.
func (f *Factory) Sets(constructors Constructors) error {
	errs := make([]interface{}, 0, len(constructors))

	for name, constructor := range constructors {
		if err := f.Set(name, constructor); err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) != 0 {
		return rterror.New("cannot set object constructors", errs...)
	}

	return nil
}

// Get returns registered object constructor by given name.
func (f *Factory) Get(name string) (constructor Constructor, err error) {
	var object interface{}

	if object, err = f.registry.Get(name); err != nil {
		return nil, err
	}

	return object.(Constructor), nil
}

// Gets returns registered object constructors by given names.
func (f *Factory) Gets(names []string) (Constructors, error) {
	objects, err := f.registry.Gets(names)

	return toConstructors(objects), err
}

// GetAll returns all registered object constructors.
func (f *Factory) GetAll() Constructors {
	return toConstructors(f.registry.GetAll())
}

// Remove removes registered object constructor.
func (f *Factory) Remove(name string) *Factory {
	f.registry.Remove(name)
	return f
}

// Removes removes registered object constructors.
func (f *Factory) Removes(names []string) *Factory {
	f.registry.Removes(names)
	return f
}

// RemoveAll removes all registered object constructors.
func (f *Factory) RemoveAll() *Factory {
	f.registry.RemoveAll()
	return f
}

// IsExist returns true if object constructor with given name was registered, otherwise it returns false.
func (f *Factory) IsExist(name string) bool {
	return f.registry.IsExist(name)
}

// IsExists returns true if all object constructors with given names were registered, otherwise it returns false.
func (f *Factory) IsExists(names []string) bool {
	return f.registry.IsExists(names)
}

// IsEmpty returns true if there are no registered object constructors, otherwise it returns false.
func (f *Factory) IsEmpty() bool {
	return f.registry.IsEmpty()
}

// Size returns number of registered object constructors.
func (f *Factory) Size() int {
	return f.registry.Size()
}

func toConstructors(objects registry.Objects) Constructors {
	constructors := Constructors{}

	for name, object := range objects {
		var ok bool

		if constructors[name], ok = object.(Constructor); !ok {
			panic("object is not a Constructor type")
		}
	}

	return constructors
}
