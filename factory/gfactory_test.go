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

package factory_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/tymonx/go-patterns/factory"
)

func TestGlobalFactoryCreate(test *testing.T) {
	defer factory.RemoveAll()

	assert.NoError(test, factory.Add("constructor", Constructor))

	object, err := factory.Create("constructor")

	assert.NoError(test, err)
	assert.NotNil(test, object)
}

func TestGlobalFactoryCreateNoExist(test *testing.T) {
	defer factory.RemoveAll()

	object, err := factory.Create("constructor")

	assert.Error(test, err)
	assert.Nil(test, object)
}

func TestGlobalFactoryCreateError(test *testing.T) {
	defer factory.RemoveAll()

	assert.NoError(test, factory.Add("constructor", ConstructorError))

	object, err := factory.Create("constructor")

	assert.Error(test, err)
	assert.Nil(test, object)
}

func TestGlobalFactoryCreateNil(test *testing.T) {
	defer factory.RemoveAll()

	assert.NoError(test, factory.Add("constructor", ConstructorNil))

	object, err := factory.Create("constructor")

	assert.Error(test, err)
	assert.Nil(test, object)
}

func TestGlobalFactoryCreates(test *testing.T) {
	defer factory.RemoveAll()

	assert.NoError(test, factory.Adds(factory.Constructors{
		"constructorA": Constructor,
		"constructorB": Constructor,
		"constructorC": Constructor,
	}))

	objects, err := factory.Creates([]string{"constructorC", "constructorA"})

	assert.NoError(test, err)
	assert.Len(test, objects, 2)
	assert.NotNil(test, objects[0])
	assert.NotNil(test, objects[1])
}

func TestGlobalFactoryCreatesError(test *testing.T) {
	defer factory.RemoveAll()

	assert.NoError(test, factory.Adds(factory.Constructors{
		"constructorA": Constructor,
		"constructorB": Constructor,
		"constructorC": ConstructorError,
	}))

	objects, err := factory.Creates([]string{"constructorC", "constructorA"})

	assert.Error(test, err)
	assert.Len(test, objects, 1)
	assert.NotNil(test, objects[0])
}

func TestGlobalFactoryAdd(test *testing.T) {
	defer factory.RemoveAll()

	assert.NoError(test, factory.Add("constructor", Constructor))
	assert.Len(test, factory.GetAll(), 1)
}

func TestGlobalFactoryAddError(test *testing.T) {
	defer factory.RemoveAll()

	assert.NoError(test, factory.Add("constructor", Constructor))
	assert.Error(test, factory.Add("constructor", Constructor))
	assert.Len(test, factory.GetAll(), 1)
}

func TestGlobalFactoryAdds(test *testing.T) {
	defer factory.RemoveAll()

	assert.NoError(test, factory.Adds(factory.Constructors{
		"constructorA": Constructor,
		"constructorB": Constructor,
	}))

	assert.Len(test, factory.GetAll(), 2)
}

func TestGlobalFactoryAddsError(test *testing.T) {
	defer factory.RemoveAll()

	assert.NoError(test, factory.Add("constructor", Constructor))
	assert.Error(test, factory.Adds(factory.Constructors{
		"constructor": Constructor,
	}))
}

func TestGlobalFactorySet(test *testing.T) {
	defer factory.RemoveAll()

	factory.Set("constructor", Constructor)
	factory.Set("constructor", Constructor)
	assert.Len(test, factory.GetAll(), 1)
}

func TestGlobalFactorySets(test *testing.T) {
	defer factory.RemoveAll()

	factory.Sets(factory.Constructors{
		"constructorA": Constructor,
		"constructorB": Constructor,
	})

	assert.Len(test, factory.GetAll(), 2)
}

func TestGlobalFactoryIsExist(test *testing.T) {
	defer factory.RemoveAll()

	assert.False(test, factory.IsExist("constructor"))
	assert.NoError(test, factory.Add("constructor", Constructor))
	assert.True(test, factory.IsExist("constructor"))
}

func TestGlobalFactoryIsExists(test *testing.T) {
	defer factory.RemoveAll()

	assert.False(test, factory.IsExists(factory.Names{"constructorC", "constructorA"}))

	assert.NoError(test, factory.Adds(factory.Constructors{
		"constructorA": Constructor,
		"constructorB": Constructor,
		"constructorC": Constructor,
	}))

	assert.True(test, factory.IsExists(factory.Names{"constructorC", "constructorA"}))
}

func TestGlobalFactoryIsEmpty(test *testing.T) {
	defer factory.RemoveAll()

	assert.True(test, factory.IsEmpty())
	assert.NoError(test, factory.Add("constructor", Constructor))
	assert.False(test, factory.IsEmpty())
}

func TestGlobalFactorySize(test *testing.T) {
	defer factory.RemoveAll()

	assert.Zero(test, factory.Size())
	assert.NoError(test, factory.Add("constructor", Constructor))
	assert.NotZero(test, factory.Size())
}

func TestGlobalFactoryGet(test *testing.T) {
	defer factory.RemoveAll()

	assert.NoError(test, factory.Add("constructor", Constructor))

	ret, err := factory.Get("constructor")

	assert.NoError(test, err)
	assert.NotNil(test, ret)
}

func TestGlobalFactoryGetError(test *testing.T) {
	defer factory.RemoveAll()

	constructor, err := factory.Get("constructor")

	assert.Error(test, err)
	assert.Nil(test, constructor)
}

func TestGlobalFactoryGets(test *testing.T) {
	defer factory.RemoveAll()

	assert.NoError(test, factory.Adds(factory.Constructors{
		"constructorA": Constructor,
		"constructorB": Constructor,
		"constructorC": Constructor,
	}))

	constructors, err := factory.Gets([]string{"constructorC", "constructorA"})

	assert.NoError(test, err)
	assert.Len(test, constructors, 2)
	assert.Contains(test, constructors, "constructorA")
	assert.Contains(test, constructors, "constructorC")
}

func TestGlobalFactoryGetsError(test *testing.T) {
	defer factory.RemoveAll()

	assert.NoError(test, factory.Adds(factory.Constructors{
		"constructorA": Constructor,
		"constructorB": Constructor,
		"constructorC": Constructor,
	}))

	constructors, err := factory.Gets([]string{"constructorD", "constructorA"})

	assert.Error(test, err)
	assert.Len(test, constructors, 1)
	assert.Contains(test, constructors, "constructorA")
}

func TestGlobalFactoryGetAll(test *testing.T) {
	defer factory.RemoveAll()

	assert.NoError(test, factory.Adds(factory.Constructors{
		"constructorA": Constructor,
		"constructorB": Constructor,
		"constructorC": Constructor,
	}))

	constructors := factory.GetAll()

	assert.Len(test, constructors, 3)
	assert.Contains(test, constructors, "constructorA")
	assert.Contains(test, constructors, "constructorB")
	assert.Contains(test, constructors, "constructorC")
}

func TestGlobalFactoryRemove(test *testing.T) {
	defer factory.RemoveAll()

	assert.NoError(test, factory.Adds(factory.Constructors{
		"constructorA": Constructor,
		"constructorB": Constructor,
		"constructorC": Constructor,
	}))

	factory.Remove("constructorB")

	constructors := factory.GetAll()

	assert.Len(test, constructors, 2)
	assert.Contains(test, constructors, "constructorA")
	assert.Contains(test, constructors, "constructorC")
}

func TestGlobalFactoryRemoveNotExist(test *testing.T) {
	defer factory.RemoveAll()

	assert.NoError(test, factory.Adds(factory.Constructors{
		"constructorA": Constructor,
		"constructorB": Constructor,
		"constructorC": Constructor,
	}))

	factory.Remove("constructorD")

	constructors := factory.GetAll()

	assert.Len(test, constructors, 3)
	assert.Contains(test, constructors, "constructorA")
	assert.Contains(test, constructors, "constructorB")
	assert.Contains(test, constructors, "constructorC")
}

func TestGlobalFactoryRemoves(test *testing.T) {
	defer factory.RemoveAll()

	assert.NoError(test, factory.Adds(factory.Constructors{
		"constructorA": Constructor,
		"constructorB": Constructor,
		"constructorC": Constructor,
	}))

	factory.Removes([]string{"constructorA", "constructorC"})

	constructors := factory.GetAll()

	assert.Len(test, constructors, 1)
	assert.Contains(test, constructors, "constructorB")
}

func TestGlobalFactoryRemoveAll(test *testing.T) {
	defer factory.RemoveAll()

	assert.NoError(test, factory.Adds(factory.Constructors{
		"constructorA": Constructor,
		"constructorB": Constructor,
		"constructorC": Constructor,
	}))

	factory.RemoveAll()

	constructors := factory.GetAll()

	assert.Empty(test, constructors)
}
