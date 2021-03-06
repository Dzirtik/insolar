/*
 *    Copyright 2018 INS Ecosystem
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

package resolver

import (
	"testing"

	"github.com/insolar/insolar/genesis/mock/storage"
	"github.com/insolar/insolar/genesis/model/object"
	"github.com/stretchr/testify/assert"
)

func TestNewResolverHandler(t *testing.T) {
	mockParent := &mockParent{}
	handler := NewResolverHandler(mockParent)

	assert.Equal(t, &resolverHandler{
		globalResolver: GlobalResolver,
		childResolver: &childResolver{
			parent: mockParent,
		},
		contextResolver: &contextResolver{
			parent: mockParent,
		},
	}, handler)
}

func TestResolverHandler_GetObject_GlobalScope(t *testing.T) {
	mockParent := &mockParent{}
	resolverHandler := NewResolverHandler(nil)
	ref, _ := object.NewReference("1", "1", object.GlobalScope)
	(*GlobalResolver.globalInstanceMap)[ref] = mockParent

	obj, err := resolverHandler.GetObject(ref, "mockParent")

	assert.NoError(t, err)
	assert.Equal(t, mockParent, obj)
}

func TestResolverHandler_GetObject_ChildScope(t *testing.T) {
	mockParent := &mockParent{}
	resolverHandler := NewResolverHandler(mockParent)
	ref, _ := object.NewReference("1", "1", object.ChildScope)

	obj, err := resolverHandler.GetObject(ref, "mockChild")

	assert.NoError(t, err)
	assert.Equal(t, child, obj)
}

func TestResolverHandler_GetObject_ContextScope(t *testing.T) {
	contextStorage := storage.NewMapStorage()
	record, _ := contextStorage.Set(child)
	mockParent := &mockParent{
		ContextStorage: contextStorage,
	}
	resolverHandler := NewResolverHandler(mockParent)
	ref, _ := object.NewReference(record, "1", object.ContextScope)

	obj, err := resolverHandler.GetObject(ref, "mockChild")

	assert.NoError(t, err)
	assert.Equal(t, child, obj)
}

func TestResolverHandler_GetObject_default(t *testing.T) {
	mockParent := &mockParent{}
	resolverHandler := NewResolverHandler(mockParent)
	ref := &object.Reference{
		Scope: object.ScopeType(10000),
	}

	obj, err := resolverHandler.GetObject(ref, "mockChild")

	assert.EqualError(t, err, "unknown scope type: 10000")
	assert.Nil(t, obj)
}
