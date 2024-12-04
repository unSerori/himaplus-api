package route

import (
	"himaplus-api/view"

	"github.com/gin-gonic/gin"
	
)

// エンドポイントのルーティング
func routing(engine *gin.Engine, handlers Handlers) {

	// ver1グループ
	v1 := engine.Group("/v1")
	{

		// todosグループ
		todos := v1.Group("/todos")
		{
			// todo新規登録
			todos.POST("/register", handlers.ToDoHandler.RegisterToDoHandler) //
		}
	}
}

// エンジンを作成して返す
func SetupRouter(handlers Handlers) (*gin.Engine, error) {
	// エンジンを作成
	engine := gin.Default()

	// 静的ファイル設定
	err := view.LoadingStaticFile(engine)
	if err != nil {
		return nil, err
	}

	// マルチパートフォームのメモリ使用制限を設定
	engine.MaxMultipartMemory = 8 << 20 // 20bit左シフトで8MiB

	// ルーティング
	routing(engine, handlers)

	// router設定されたengineを返す。
	return engine, nil
}
