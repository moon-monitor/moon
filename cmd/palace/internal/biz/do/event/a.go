package event

import (
	"sync"
)

func Models() []any {
	return []any{
		&Realtime{},
	}
}

var hasTable func(tableName string) bool = nil
var setTableFun func(tableName string) = nil

var setHasTableFunOnce sync.Once
var setTableFunOnce sync.Once

func SetHasTableFun(f func(tableName string) bool) {
	setHasTableFunOnce.Do(func() {
		hasTable = f
	})
}

func SetTableFun(f func(tableName string)) {
	setTableFunOnce.Do(func() {
		setTableFun = f
	})
}
