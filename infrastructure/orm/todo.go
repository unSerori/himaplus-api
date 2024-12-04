package orm

import "xorm.io/xorm"

// サービスの構造体
type ToDoInfrastruture struct {
	db *xorm.Engine
}

// ファクトリー関数
func NewToDoInfrastruture(db *xorm.Engine) *ToDoInfrastruture {
	return &ToDoInfrastruture{db: db}
}
