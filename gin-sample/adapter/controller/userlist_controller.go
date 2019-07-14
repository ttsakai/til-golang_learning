package controller

import (
	"til-golang_learning/gin-sample/usecase/interfaces"
)

type UserListController struct {
	//初期化の時に渡してやる？
	Inputport interfaces.IUserListInputPort
}

// domein や external,repositoryに依存しまくっている。
// func NewUserController(i userlist_inter  ) *UserController {
// 	return &UserController{
// 		Interactor: usecase.UserInteractor{
// 			UserRepository: &gateway.UserRepository{Conn: conn},
// 		},
// 	}
// }
func (controller *UserListController) Execute(username string) {

	controller.Inputport.Handle(username)

	// if err != nil {
	// 	c.JSON(500, NewError(500, err.Error()))
	// 	return
	// }
}
