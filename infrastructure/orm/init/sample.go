// サンプルデータ
package orm

import (
	model "himaplus-api/infrastructure/orm/model"
	"time"

	"xorm.io/xorm"
)

// テストデータ
// 外部キーの参照先テーブルを先に登録する必要がある。
func sampleData() []interface{} {
	return []interface{}{ // サンプル
		// ToDo
		&model.ToDo{
			UserUuid:     "de8fbfdb-c7a0-43c1-8df3-e5d741fd0e92",
			ToDoUuid:     "097dd3e7-59b5-411b-9c96-71d494499c4c",
			ToDoTitile:   "書類提出",
			Importance:   1,
			RequiredTime: 1 * time.Hour,
			ToDoMemo:     "",
			ToDoDate:     time.Now().Add(24 * time.Hour),
		},
	}
}

// サンプルデータ作成
func RegisterSample(db *xorm.Engine) {
	// テスト用データ作成
	for _, sample := range sampleData() {
		db.Insert(sample)
	}
}
