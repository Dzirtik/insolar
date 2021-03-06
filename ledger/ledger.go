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

package ledger

import (
	"github.com/insolar/insolar/ledger/record"
	"github.com/insolar/insolar/ledger/storage"
)

// Ledgerer is high level Ledger interface
// TODO: signature probably will change
type Ledgerer interface {
	Get(id record.Hash) (bool, record.Record)
	Set(record record.Record) error
}

// Ledger defines parameters for running ledger storer
// TODO: should implements Ledgerer interface
type Ledger struct {
	Store storage.LedgerStorer
}

// NewLedger creates new Ledger
func NewLedger() (Ledger, error) {
	panic("implement me")
}
