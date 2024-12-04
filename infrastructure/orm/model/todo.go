package model

import "time"

// TODOテーブル
type ToDo struct {
	UserUuid     string        `xorm:"varchar(36) pk" json:"userUUID"`                           // ユーザーUUID
	ToDoUuid     string        `xorm:"varchar(36) pk" json:"todoUUID"`                           // todoUUID
	ToDoTitile   string        `xorm:"varchar(36) not null" json:"todotitel" binding:"required"` // todoタイトル
	Importance   int           `xorm:"int" json:"importance not null" binding:"required"`        // 重要度	1:高、2:中、3:低
	RequiredTime time.Duration `xorm:"DATETIME not null" json:"requiredTime"`                    // 必要時間
	ToDoMemo     string        `xorm:"text" json:"todoMemo"`                                     // memo
	ToDoDate     time.Time     `xorm:"DATETIME not null" json:"ToDoDate"`                        // 登録した時間
	ParentUuid   *string        `xorm:"varchar(36) pk" json:"parentUUID"`                         // 親要素のUUID
}