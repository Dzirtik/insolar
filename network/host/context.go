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

package host

import (
	"context"
	"errors"

	"github.com/insolar/insolar/network/host/node"
)

// Context type is localized for future purposes.
// Network Node can have multiple IDs, but each action must be executed with only one ID.
// Context is used in all actions to select specific ID to work with.
type Context context.Context

type ctxKey string

const (
	ctxTableIndex = ctxKey("table_index")
	defaultNodeID = 0
)

// ContextBuilder allows to lazy configure and build new Context.
type ContextBuilder struct {
	dht     *DHT
	actions []func(ctx Context) (Context, error)
}

// NewContextBuilder creates new ContextBuilder.
func NewContextBuilder(dht *DHT) ContextBuilder {
	return ContextBuilder{
		dht: dht,
	}
}

// Build builds and returns new Context.
func (cb ContextBuilder) Build() (ctx Context, err error) {
	ctx = context.Background()
	for _, action := range cb.actions {
		ctx, err = action(ctx)
		if err != nil {
			return
		}
	}
	return
}

// SetNodeByID sets node id in Context.
func (cb ContextBuilder) SetNodeByID(nodeID node.ID) ContextBuilder {
	cb.actions = append(cb.actions, func(ctx Context) (Context, error) {
		for index, id := range cb.dht.origin.IDs {
			if nodeID.Equal(id) {
				return context.WithValue(ctx, ctxTableIndex, index), nil
			}
		}
		return nil, errors.New("node requestID not found")
	})
	return cb
}

// SetDefaultNode sets first node id in Context.
func (cb ContextBuilder) SetDefaultNode() ContextBuilder {
	cb.actions = append(cb.actions, func(ctx Context) (Context, error) {
		return context.WithValue(ctx, ctxTableIndex, defaultNodeID), nil
	})
	return cb
}
