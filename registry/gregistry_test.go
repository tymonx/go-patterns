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

func TestGlobalRegistryAdd(test *testing.T) {
	defer registry.RemoveAll()

	var object struct{}

	assert.NoError(test, registry.Add("object", object))
	assert.Len(test, registry.GetAll(), 1)
}

func TestGlobalRegistryAdds(test *testing.T) {
	defer registry.RemoveAll()

	var objectA, objectB struct{}

	assert.NoError(test, registry.Adds(registry.Objects{
		"objectA": objectA,
		"objectB": objectB,
	}))

	assert.Len(test, registry.GetAll(), 2)
}

func TestGlobalRegistryAddsError(test *testing.T) {
	defer registry.RemoveAll()

	var object struct{}

	assert.NoError(test, registry.Add("object", object))
	assert.Error(test, registry.Adds(registry.Objects{
		"object": object,
	}))
}

func TestGlobalRegistrySet(test *testing.T) {
	defer registry.RemoveAll()

	var object struct{}

	registry.Set("object", object)
	registry.Set("object", object)
	assert.Len(test, registry.GetAll(), 1)
}

func TestGlobalRegistrySets(test *testing.T) {
	defer registry.RemoveAll()

	var objectA, objectB struct{}

	registry.Sets(registry.Objects{
		"objectA": objectA,
		"objectB": objectB,
	})

	assert.Len(test, registry.GetAll(), 2)
}

func TestGlobalRegistryIsExist(test *testing.T) {
	defer registry.RemoveAll()

	var object struct{}

	assert.False(test, registry.IsExist("object"))
	assert.NoError(test, registry.Add("object", object))
	assert.True(test, registry.IsExist("object"))
}

func TestGlobalRegistryIsExists(test *testing.T) {
	defer registry.RemoveAll()

	var objectA, objectB, objectC struct{}

	assert.False(test, registry.IsExists(registry.Names{"objectC", "objectA"}))

	assert.NoError(test, registry.Adds(registry.Objects{
		"objectA": objectA,
		"objectB": objectB,
		"objectC": objectC,
	}))

	assert.True(test, registry.IsExists(registry.Names{"objectC", "objectA"}))
}

func TestGlobalRegistryIsEmpty(test *testing.T) {
	defer registry.RemoveAll()

	var object struct{}

	assert.True(test, registry.IsEmpty())
	assert.NoError(test, registry.Add("object", object))
	assert.False(test, registry.IsEmpty())
}

func TestGlobalRegistrySize(test *testing.T) {
	defer registry.RemoveAll()

	var object struct{}

	assert.Zero(test, registry.Size())
	assert.NoError(test, registry.Add("object", object))
	assert.NotZero(test, registry.Size())
}

func TestGlobalRegistryGet(test *testing.T) {
	defer registry.RemoveAll()

	var object struct{}

	assert.NoError(test, registry.Add("object", object))

	ret, err := registry.Get("object")

	assert.NoError(test, err)
	assert.Equal(test, ret, object)
}

func TestGlobalRegistryGetError(test *testing.T) {
	defer registry.RemoveAll()

	object, err := registry.New().Get("object")

	assert.Error(test, err)
	assert.Nil(test, object)
}

func TestGlobalRegistryGets(test *testing.T) {
	defer registry.RemoveAll()

	var objectA, objectB, objectC struct{}

	assert.NoError(test, registry.Adds(registry.Objects{
		"objectA": objectA,
		"objectB": objectB,
		"objectC": objectC,
	}))

	objects, err := registry.Gets([]string{"objectC", "objectA"})

	assert.NoError(test, err)
	assert.Len(test, objects, 2)
	assert.Equal(test, objectC, objects["objectC"])
	assert.Equal(test, objectA, objects["objectA"])
}

func TestGlobalRegistryGetsError(test *testing.T) {
	defer registry.RemoveAll()

	var objectA, objectB, objectC struct{}

	assert.NoError(test, registry.Adds(registry.Objects{
		"objectA": objectA,
		"objectB": objectB,
		"objectC": objectC,
	}))

	objects, err := registry.Gets([]string{"objectD", "objectA"})

	assert.Error(test, err)
	assert.Len(test, objects, 1)
	assert.Equal(test, objectA, objects["objectA"])
}

func TestGlobalRegistryGetAll(test *testing.T) {
	defer registry.RemoveAll()

	var objectA, objectB, objectC struct{}

	assert.NoError(test, registry.Adds(registry.Objects{
		"objectA": objectA,
		"objectB": objectB,
		"objectC": objectC,
	}))

	objects := registry.GetAll()

	assert.Len(test, objects, 3)
	assert.Equal(test, objectA, objects["objectA"])
	assert.Equal(test, objectB, objects["objectB"])
	assert.Equal(test, objectC, objects["objectC"])
}

func TestGlobalRegistryRemove(test *testing.T) {
	defer registry.RemoveAll()

	var objectA, objectB, objectC struct{}

	assert.NoError(test, registry.Adds(registry.Objects{
		"objectA": objectA,
		"objectB": objectB,
		"objectC": objectC,
	}))

	registry.Remove("objectB")

	objects := registry.GetAll()

	assert.Len(test, objects, 2)
	assert.Equal(test, objectA, objects["objectA"])
	assert.Equal(test, objectC, objects["objectC"])
}

func TestGlobalRegistryRemoveNotExist(test *testing.T) {
	defer registry.RemoveAll()

	var objectA, objectB, objectC struct{}

	assert.NoError(test, registry.Adds(registry.Objects{
		"objectA": objectA,
		"objectB": objectB,
		"objectC": objectC,
	}))

	registry.Remove("objectD")

	objects := registry.GetAll()

	assert.Len(test, objects, 3)
	assert.Equal(test, objectA, objects["objectA"])
	assert.Equal(test, objectB, objects["objectB"])
	assert.Equal(test, objectC, objects["objectC"])
}

func TestGlobalRegistryRemoves(test *testing.T) {
	defer registry.RemoveAll()

	var objectA, objectB, objectC struct{}

	assert.NoError(test, registry.Adds(registry.Objects{
		"objectA": objectA,
		"objectB": objectB,
		"objectC": objectC,
	}))

	registry.Removes([]string{"objectA", "objectC"})

	objects := registry.GetAll()

	assert.Len(test, objects, 1)
	assert.Equal(test, objectB, objects["objectB"])
}

func TestGlobalRegistryRemoveAll(test *testing.T) {
	defer registry.RemoveAll()

	var objectA, objectB, objectC struct{}

	assert.NoError(test, registry.Adds(registry.Objects{
		"objectA": objectA,
		"objectB": objectB,
		"objectC": objectC,
	}))

	registry.RemoveAll()

	objects := registry.GetAll()

	assert.Empty(test, objects)
}
