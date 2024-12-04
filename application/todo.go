package application

import "himaplus-api/infrastructure/orm"

// サービスの構造体
type ToDoService struct {
	i *orm.ToDoInfrastruture
}

func NewToDoService(i *orm.ToDoInfrastruture) *ToDoService{
	return &ToDoService{
		i: i,
	}
}

