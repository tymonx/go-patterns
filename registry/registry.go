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
	"gitlab.com/tymonx/go-error/rterror"
)

// Names defines a list of names.
type Names []string

// Objects defines a list of objects.
type Objects map[string]interface{}

// Registry defines a registry object that can register objects.
type Registry struct {
	objects Objects
}

// New creates a new registry object.
func New() *Registry {
	return &Registry{
		objects: Objects{},
	}
}

// Add adds an object with a given unique id to registry.
func (r *Registry) Add(name string, object interface{}) error {
	if r.IsExist(name) {
		return rterror.New("object was already registered", name)
	}

	r.objects[name] = object

	return nil
}

// Adds adds new objects with given unique ids to registry.
func (r *Registry) Adds(objects Objects) error {
	errs := make([]interface{}, 0, len(objects))

	for name, object := range objects {
		if err := r.Add(name, object); err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) != 0 {
		return rterror.New("cannot add objects", errs...)
	}

	return nil
}

// Set sets an object with a given unique id to registry.
func (r *Registry) Set(name string, object interface{}) *Registry {
	r.objects[name] = object
	return r
}

// Sets sets objects with given unique ids to registry.
func (r *Registry) Sets(objects Objects) *Registry {
	for name, object := range objects {
		r.Set(name, object)
	}

	return r
}

// Get returns registered object by given name.
func (r *Registry) Get(name string) (object interface{}, err error) {
	var ok bool

	if object, ok = r.objects[name]; !ok {
		return nil, rterror.New("object was not registered", name)
	}

	return object, nil
}

// Gets returns registered objects by given names.
func (r *Registry) Gets(names []string) (objects Objects, err error) {
	errs := make([]interface{}, 0, len(objects))

	objects = Objects{}

	for _, name := range names {
		var object interface{}

		if object, err = r.Get(name); err != nil {
			errs = append(errs, err)

			continue
		}

		objects[name] = object
	}

	if len(errs) != 0 {
		return objects, rterror.New("cannot get objects", errs...)
	}

	return objects, nil
}

// GetAll returns all registered objects.
func (r *Registry) GetAll() Objects {
	objects := Objects{}

	for name, object := range r.objects {
		objects[name] = object
	}

	return objects
}

// Remove removes registered object.
func (r *Registry) Remove(name string) *Registry {
	if _, ok := r.objects[name]; !ok {
		return r
	}

	delete(r.objects, name)

	return r
}

// Removes removes registered objects.
func (r *Registry) Removes(names []string) *Registry {
	for _, name := range names {
		r.Remove(name)
	}

	return r
}

// RemoveAll removes all registered objects.
func (r *Registry) RemoveAll() *Registry {
	r.objects = Objects{}
	return r
}

// IsExist returns true if object with given name was registered, otherwise it returns false.
func (r *Registry) IsExist(name string) bool {
	_, ok := r.objects[name]
	return ok
}

// IsExists returns true if all objects with given names were registered, otherwise it returns false.
func (r *Registry) IsExists(names []string) bool {
	for _, name := range names {
		if !r.IsExist(name) {
			return false
		}
	}

	return true
}

// IsEmpty returns true if there are no registered objects, otherwise it returns false.
func (r *Registry) IsEmpty() bool {
	return len(r.objects) == 0
}

// Size returns number of registered objects.
func (r *Registry) Size() int {
	return len(r.objects)
}
