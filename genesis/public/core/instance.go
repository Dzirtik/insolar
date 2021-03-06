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

package core

import (
	"fmt"

	"github.com/insolar/insolar/genesis/model/class"
	"github.com/insolar/insolar/genesis/model/domain"
	"github.com/insolar/insolar/genesis/model/factory"
	"github.com/insolar/insolar/genesis/model/object"
)

// InstanceDomainName is a name for instance domain.
const InstanceDomainName = "InstanceDomain"

// InstanceDomain is a contract that stores instances of other domains
type InstanceDomain interface {
	// Base domain implementation.
	domain.Domain
	// CreateInstance is used to create new instance as a child to domain storage.
	CreateInstance(*factory.Factory) (string, error)
	// GetInstance returns instance from its record in domain storage.
	GetInstance(string) (*factory.Factory, error)
}

type instanceDomain struct {
	domain.BaseDomain
}

// newInstanceDomain creates new instance of InstanceDomain
func newInstanceDomain(parent object.Parent) (*instanceDomain, error) {
	if parent == nil {
		return nil, fmt.Errorf("parent must not be nil")
	}

	instDomain := &instanceDomain{
		BaseDomain: *domain.NewBaseDomain(parent, InstanceDomainName),
	}
	return instDomain, nil
}

// GetClassID return string representation of InstanceDomain's class.
func (instDom *instanceDomain) GetClassID() string {
	return class.InstanceDomainID
}

// CreateInstance creates new instance as a child to domain storage.
func (instDom *instanceDomain) CreateInstance(fc factory.Factory) (string, error) {
	instance := fc.Create(instDom)
	if instance == nil {
		return "", fmt.Errorf("factory returns nil")
	}

	record, err := instDom.ChildStorage.Set(instance)
	if err != nil {
		return "", err
	}

	return record, nil
}

// GetInstance returns instance from its record in domain storage.
func (instDom *instanceDomain) GetInstance(record string) (object.Proxy, error) {
	instance, err := instDom.ChildStorage.Get(record)
	if err != nil {
		return nil, err
	}

	result, ok := instance.(object.Proxy)
	if !ok {
		return nil, fmt.Errorf("object with record `%s` is not `Proxy` instance", record)
	}

	return result, nil
}

type instanceDomainProxy struct {
	instance *instanceDomain
}

// newInstanceDomainProxy creates new proxy and associate it with new instance of InstanceDomain.
func newInstanceDomainProxy(parent object.Parent) (*instanceDomainProxy, error) {
	instance, err := newInstanceDomain(parent)
	if err != nil {
		return nil, err
	}
	return &instanceDomainProxy{
		instance: instance,
	}, nil
}

// CreateInstance proxy call for instance method.
func (idp *instanceDomainProxy) CreateInstance(fc factory.Factory) (string, error) {
	return idp.instance.CreateInstance(fc)
}

// GetInstance proxy call for instance method.
func (idp *instanceDomainProxy) GetInstance(record string) (object.Proxy, error) {
	return idp.instance.GetInstance(record)
}

// GetReference proxy call for instance method.
func (idp *instanceDomainProxy) GetReference() *object.Reference {
	return idp.instance.GetReference()
}

// GetParent proxy call for instance method.
func (idp *instanceDomainProxy) GetParent() object.Parent {
	return idp.instance.GetParent()
}

// GetClassID proxy call for instance method.
func (idp *instanceDomainProxy) GetClassID() string {
	return class.InstanceDomainID
}

type instanceDomainFactory struct{}

// NewInstanceDomainFactory creates new factory for InstanceDomain.
func NewInstanceDomainFactory() factory.Factory {
	return &instanceDomainFactory{}
}

// GetClassID return string representation of InstanceDomain's class.
func (idf *instanceDomainFactory) GetClassID() string {
	return class.InstanceDomainID
}

// GetReference returns nil for not published factory
func (idf *instanceDomainFactory) GetReference() *object.Reference {
	return nil
}

// Create is factory method that used to create new InstanceDomain instances.
func (idf *instanceDomainFactory) Create(parent object.Parent) object.Proxy {
	proxy, err := newInstanceDomainProxy(parent)
	if err != nil {
		return nil
	}

	_, err = parent.AddChild(proxy)
	if err != nil {
		return nil
	}
	return proxy
}
