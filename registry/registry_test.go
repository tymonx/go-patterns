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

package registry_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/tymonx/go-patterns/registry"
)

func TestRegistryNew(test *testing.T) {
	r := registry.New()

	assert.NotNil(test, r)
	assert.Empty(test, r.GetAll())
}

func TestRegistryAdd(test *testing.T) {
	var object struct{}

	r := registry.New()

	assert.NoError(test, r.Add("object", object))
	assert.Len(test, r.GetAll(), 1)
}

func TestRegistryAdds(test *testing.T) {
	var objectA, objectB struct{}

	r := registry.New()

	assert.NoError(test, r.Adds(registry.Objects{
		"objectA": objectA,
		"objectB": objectB,
	}))

	assert.Len(test, r.GetAll(), 2)
}

func TestRegistryAddsError(test *testing.T) {
	var object struct{}

	r := registry.New()

	assert.NoError(test, r.Add("object", object))
	assert.Error(test, r.Adds(registry.Objects{
		"object": object,
	}))
}

func TestRegistrySet(test *testing.T) {
	var object struct{}

	r := registry.New()

	assert.Same(test, r, r.Set("object", object))
	assert.Same(test, r, r.Set("object", object))
	assert.Len(test, r.GetAll(), 1)
}

func TestRegistrySets(test *testing.T) {
	var objectA, objectB struct{}

	r := registry.New()

	assert.Same(test, r, r.Sets(registry.Objects{
		"objectA": objectA,
		"objectB": objectB,
	}))

	assert.Len(test, r.GetAll(), 2)
}

func TestRegistryIsExist(test *testing.T) {
	var object struct{}

	r := registry.New()

	assert.False(test, r.IsExist("object"))
	assert.NoError(test, r.Add("object", object))
	assert.True(test, r.IsExist("object"))
}

func TestRegistryIsExists(test *testing.T) {
	var objectA, objectB, objectC struct{}

	r := registry.New()

	assert.False(test, r.IsExists(registry.Names{"objectC", "objectA"}))

	assert.NoError(test, r.Adds(registry.Objects{
		"objectA": objectA,
		"objectB": objectB,
		"objectC": objectC,
	}))

	assert.True(test, r.IsExists(registry.Names{"objectC", "objectA"}))
}

func TestRegistryIsEmpty(test *testing.T) {
	var object struct{}

	r := registry.New()

	assert.True(test, r.IsEmpty())
	assert.NoError(test, r.Add("object", object))
	assert.False(test, r.IsEmpty())
}

func TestRegistrySize(test *testing.T) {
	var object struct{}

	r := registry.New()

	assert.Zero(test, r.Size())
	assert.NoError(test, r.Add("object", object))
	assert.NotZero(test, r.Size())
}

func TestRegistryGet(test *testing.T) {
	var object struct{}

	r := registry.New()

	assert.NoError(test, r.Add("object", object))

	ret, err := r.Get("object")

	assert.NoError(test, err)
	assert.Equal(test, ret, object)
}

func TestRegistryGetError(test *testing.T) {
	object, err := registry.New().Get("object")

	assert.Error(test, err)
	assert.Nil(test, object)
}

func TestRegistryGets(test *testing.T) {
	var objectA, objectB, objectC struct{}

	r := registry.New()

	assert.NoError(test, r.Adds(registry.Objects{
		"objectA": objectA,
		"objectB": objectB,
		"objectC": objectC,
	}))

	objects, err := r.Gets([]string{"objectC", "objectA"})

	assert.NoError(test, err)
	assert.Len(test, objects, 2)
	assert.Equal(test, objectC, objects["objectC"])
	assert.Equal(test, objectA, objects["objectA"])
}

func TestRegistryGetsError(test *testing.T) {
	var objectA, objectB, objectC struct{}

	r := registry.New()

	assert.NoError(test, r.Adds(registry.Objects{
		"objectA": objectA,
		"objectB": objectB,
		"objectC": objectC,
	}))

	objects, err := r.Gets([]string{"objectD", "objectA"})

	assert.Error(test, err)
	assert.Len(test, objects, 1)
	assert.Equal(test, objectA, objects["objectA"])
}

func TestRegistryGetAll(test *testing.T) {
	var objectA, objectB, objectC struct{}

	r := registry.New()

	assert.NoError(test, r.Adds(registry.Objects{
		"objectA": objectA,
		"objectB": objectB,
		"objectC": objectC,
	}))

	objects := r.GetAll()

	assert.Len(test, objects, 3)
	assert.Equal(test, objectA, objects["objectA"])
	assert.Equal(test, objectB, objects["objectB"])
	assert.Equal(test, objectC, objects["objectC"])
}

func TestRegistryRemove(test *testing.T) {
	var objectA, objectB, objectC struct{}

	r := registry.New()

	assert.NoError(test, r.Adds(registry.Objects{
		"objectA": objectA,
		"objectB": objectB,
		"objectC": objectC,
	}))

	assert.Same(test, r, r.Remove("objectB"))

	objects := r.GetAll()

	assert.Len(test, objects, 2)
	assert.Equal(test, objectA, objects["objectA"])
	assert.Equal(test, objectC, objects["objectC"])
}

func TestRegistryRemoveNotExist(test *testing.T) {
	var objectA, objectB, objectC struct{}

	r := registry.New()

	assert.NoError(test, r.Adds(registry.Objects{
		"objectA": objectA,
		"objectB": objectB,
		"objectC": objectC,
	}))

	assert.Same(test, r, r.Remove("objectD"))

	objects := r.GetAll()

	assert.Len(test, objects, 3)
	assert.Equal(test, objectA, objects["objectA"])
	assert.Equal(test, objectB, objects["objectB"])
	assert.Equal(test, objectC, objects["objectC"])
}

func TestRegistryRemoves(test *testing.T) {
	var objectA, objectB, objectC struct{}

	r := registry.New()

	assert.NoError(test, r.Adds(registry.Objects{
		"objectA": objectA,
		"objectB": objectB,
		"objectC": objectC,
	}))

	assert.Same(test, r, r.Removes([]string{"objectA", "objectC"}))

	objects := r.GetAll()

	assert.Len(test, objects, 1)
	assert.Equal(test, objectB, objects["objectB"])
}

func TestRegistryRemoveAll(test *testing.T) {
	var objectA, objectB, objectC struct{}

	r := registry.New()

	assert.NoError(test, r.Adds(registry.Objects{
		"objectA": objectA,
		"objectB": objectB,
		"objectC": objectC,
	}))

	assert.Same(test, r, r.RemoveAll())

	objects := r.GetAll()

	assert.Empty(test, objects)
}
