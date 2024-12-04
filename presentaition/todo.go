package presentation

import (
	"fmt"
	"himaplus-api/application"

	"github.com/gin-gonic/gin"
)

type ToDoHandler struct {
	s *application.ToDoService
}

// ファクトリー関数
func NewToDoHandler(s *application.ToDoService) *ToDoHandler {
	return &ToDoHandler{
		s: s,
	}
}

// todo登録
func (h *ToDoHandler) RegisterToDoHandler(ctx *gin.Context) {

	fmt.Println("todoはんどらーです")
		// 構造体にマッピング
		// var bToDo requests.ToDo // 構造体のインスタンス
		// if err := ctx.ShouldBindJSON(&bReq); err != nil {
		// 	responder.SendFailedBindJSON(ctx, err)
		// 	return
		// }

}
