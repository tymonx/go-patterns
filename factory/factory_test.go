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

func TestFactoryNew(test *testing.T) {
	f := factory.New()

	assert.NotNil(test, f)
	assert.Empty(test, f.GetAll())
}

func TestFactoryCreate(test *testing.T) {
	f := factory.New()

	assert.NoError(test, f.Add("constructor", Constructor))

	object, err := f.Create("constructor")

	assert.NoError(test, err)
	assert.NotNil(test, object)
}

func TestFactoryCreateNoExist(test *testing.T) {
	f := factory.New()

	object, err := f.Create("constructor")

	assert.Error(test, err)
	assert.Nil(test, object)
}

func TestFactoryCreateError(test *testing.T) {
	f := factory.New()

	assert.NoError(test, f.Add("constructor", ConstructorError))

	object, err := f.Create("constructor")

	assert.Error(test, err)
	assert.Nil(test, object)
}

func TestFactoryCreateNil(test *testing.T) {
	f := factory.New()

	assert.NoError(test, f.Add("constructor", ConstructorNil))

	object, err := f.Create("constructor")

	assert.Error(test, err)
	assert.Nil(test, object)
}

func TestFactoryCreates(test *testing.T) {
	f := factory.New()

	assert.NoError(test, f.Adds(factory.Constructors{
		"constructorA": Constructor,
		"constructorB": Constructor,
		"constructorC": Constructor,
	}))

	objects, err := f.Creates([]string{"constructorC", "constructorA"})

	assert.NoError(test, err)
	assert.Len(test, objects, 2)
	assert.NotNil(test, objects[0])
	assert.NotNil(test, objects[1])
}

func TestFactoryCreatesError(test *testing.T) {
	f := factory.New()

	assert.NoError(test, f.Adds(factory.Constructors{
		"constructorA": Constructor,
		"constructorB": Constructor,
		"constructorC": ConstructorError,
	}))

	objects, err := f.Creates([]string{"constructorC", "constructorA"})

	assert.Error(test, err)
	assert.Len(test, objects, 1)
	assert.NotNil(test, objects[0])
}

func TestFactoryAdd(test *testing.T) {
	f := factory.New()

	assert.NoError(test, f.Add("constructor", Constructor))
	assert.Len(test, f.GetAll(), 1)
}

func TestFactoryAddError(test *testing.T) {
	f := factory.New()

	assert.Error(test, f.Add("constructor", nil))
	assert.Empty(test, f.GetAll())
}

func TestFactoryAdds(test *testing.T) {
	f := factory.New()

	assert.NoError(test, f.Adds(factory.Constructors{
		"constructorA": Constructor,
		"constructorB": Constructor,
	}))

	assert.Len(test, f.GetAll(), 2)
}

func TestFactoryAddsError(test *testing.T) {
	f := factory.New()

	assert.NoError(test, f.Add("constructor", Constructor))
	assert.Error(test, f.Adds(factory.Constructors{
		"constructor": Constructor,
	}))
}

func TestFactorySet(test *testing.T) {
	f := factory.New()

	assert.NoError(test, f.Set("constructor", Constructor))
	assert.NoError(test, f.Set("constructor", Constructor))
	assert.Len(test, f.GetAll(), 1)
}

func TestFactorySetError(test *testing.T) {
	f := factory.New()

	assert.Error(test, f.Set("constructor", nil))
	assert.Empty(test, f.GetAll())
}

func TestFactorySets(test *testing.T) {
	f := factory.New()

	assert.NoError(test, f.Sets(factory.Constructors{
		"constructorA": Constructor,
		"constructorB": Constructor,
	}))

	assert.Len(test, f.GetAll(), 2)
}

func TestFactorySetsError(test *testing.T) {
	f := factory.New()

	assert.Error(test, f.Sets(factory.Constructors{
		"constructorA": nil,
		"constructorB": Constructor,
	}))

	assert.Len(test, f.GetAll(), 1)
}

func TestFactoryIsExist(test *testing.T) {
	f := factory.New()

	assert.False(test, f.IsExist("constructor"))
	assert.NoError(test, f.Add("constructor", Constructor))
	assert.True(test, f.IsExist("constructor"))
}

func TestFactoryIsExists(test *testing.T) {
	f := factory.New()

	assert.False(test, f.IsExists(factory.Names{"constructorC", "constructorA"}))

	assert.NoError(test, f.Adds(factory.Constructors{
		"constructorA": Constructor,
		"constructorB": Constructor,
		"constructorC": Constructor,
	}))

	assert.True(test, f.IsExists(factory.Names{"constructorC", "constructorA"}))
}

func TestFactoryIsEmpty(test *testing.T) {
	f := factory.New()

	assert.True(test, f.IsEmpty())
	assert.NoError(test, f.Add("constructor", Constructor))
	assert.False(test, f.IsEmpty())
}

func TestFactorySize(test *testing.T) {
	f := factory.New()

	assert.Zero(test, f.Size())
	assert.NoError(test, f.Add("constructor", Constructor))
	assert.NotZero(test, f.Size())
}

func TestFactoryGet(test *testing.T) {
	f := factory.New()

	assert.NoError(test, f.Add("constructor", Constructor))

	ret, err := f.Get("constructor")

	assert.NoError(test, err)
	assert.NotNil(test, ret)
}

func TestFactoryGetError(test *testing.T) {
	constructor, err := factory.New().Get("constructor")

	assert.Error(test, err)
	assert.Nil(test, constructor)
}

func TestFactoryGets(test *testing.T) {
	f := factory.New()

	assert.NoError(test, f.Adds(factory.Constructors{
		"constructorA": Constructor,
		"constructorB": Constructor,
		"constructorC": Constructor,
	}))

	constructors, err := f.Gets([]string{"constructorC", "constructorA"})

	assert.NoError(test, err)
	assert.Len(test, constructors, 2)
	assert.Contains(test, constructors, "constructorA")
	assert.Contains(test, constructors, "constructorC")
}

func TestFactoryGetsError(test *testing.T) {
	f := factory.New()

	assert.NoError(test, f.Adds(factory.Constructors{
		"constructorA": Constructor,
		"constructorB": Constructor,
		"constructorC": Constructor,
	}))

	constructors, err := f.Gets([]string{"constructorD", "constructorA"})

	assert.Error(test, err)
	assert.Len(test, constructors, 1)
	assert.Contains(test, constructors, "constructorA")
}

func TestFactoryGetAll(test *testing.T) {
	f := factory.New()

	assert.NoError(test, f.Adds(factory.Constructors{
		"constructorA": Constructor,
		"constructorB": Constructor,
		"constructorC": Constructor,
	}))

	constructors := f.GetAll()

	assert.Len(test, constructors, 3)
	assert.Contains(test, constructors, "constructorA")
	assert.Contains(test, constructors, "constructorB")
	assert.Contains(test, constructors, "constructorC")
}

func TestFactoryRemove(test *testing.T) {
	f := factory.New()

	assert.NoError(test, f.Adds(factory.Constructors{
		"constructorA": Constructor,
		"constructorB": Constructor,
		"constructorC": Constructor,
	}))

	assert.Same(test, f, f.Remove("constructorB"))

	constructors := f.GetAll()

	assert.Len(test, constructors, 2)
	assert.Contains(test, constructors, "constructorA")
	assert.Contains(test, constructors, "constructorC")
}

func TestFactoryRemoveNotExist(test *testing.T) {
	f := factory.New()

	assert.NoError(test, f.Adds(factory.Constructors{
		"constructorA": Constructor,
		"constructorB": Constructor,
		"constructorC": Constructor,
	}))

	assert.Same(test, f, f.Remove("constructorD"))

	constructors := f.GetAll()

	assert.Len(test, constructors, 3)
	assert.Contains(test, constructors, "constructorA")
	assert.Contains(test, constructors, "constructorB")
	assert.Contains(test, constructors, "constructorC")
}

func TestFactoryRemoves(test *testing.T) {
	f := factory.New()

	assert.NoError(test, f.Adds(factory.Constructors{
		"constructorA": Constructor,
		"constructorB": Constructor,
		"constructorC": Constructor,
	}))

	assert.Same(test, f, f.Removes([]string{"constructorA", "constructorC"}))

	constructors := f.GetAll()

	assert.Len(test, constructors, 1)
	assert.Contains(test, constructors, "constructorB")
}

func TestFactoryRemoveAll(test *testing.T) {
	f := factory.New()

	assert.NoError(test, f.Adds(factory.Constructors{
		"constructorA": Constructor,
		"constructorB": Constructor,
		"constructorC": Constructor,
	}))

	assert.Same(test, f, f.RemoveAll())

	constructors := f.GetAll()

	assert.Empty(test, constructors)
}
