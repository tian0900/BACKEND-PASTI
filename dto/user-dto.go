package dto

type UserUpdateDTO struct {
	User_id  uint64 `json:"user_id" form:"user_id"`                                          //id wajib diisi
	Name     string `json:"name" form:"name" binding:"required"`                             //name wajib diisi
	Email    string `json:"email" form:"email" binding:"required,email"`                     //email wajib diisi
	Password string `json:"password,omitempty" form:"password,omitempty" binding:"required"` //password wajib diisi
}

//UserCreateDTO is used by client when create user
// type UserCreateDTO struct {
// 	Name     string `json:"name" form:"name" binding:"required"`
// 	Email    string `json:"email" form:"email" binding:"required" validate:"email"`
// 	Password string `json:"password, omitempty" form:"password, omitempty" validate:"min:6" binding:"required" `
// }
