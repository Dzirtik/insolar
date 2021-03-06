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

package storage

import "github.com/insolar/insolar/ledger/record"

// RecordKey is a composite key for LedgerStore.Get method.
type RecordKey struct {
	Hash     []byte
	TimeSlot uint64
}

// LedgerStorer represents append-only Ladger storage.
type LedgerStorer interface {
	Get(RecordKey) (record.Record, bool)
	Set(record.Record) error
}
