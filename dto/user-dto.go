package dto

type UserUpdateDTO struct {
	User_id    uint64 `json:"user_id" form:"user_id"`
	Name       string `json:"name" form:"name" binding:"required"`
	Email      string `json:"email" form:"email" binding:"required,email"`
	Password   string `json:"password,omitempty" form:"password,omitempty" binding:"required"`
	Role       string `json:"role" form:"role" binding:"required"`
	Usia       uint64 `json:"usia" form:"usia" binding:"required"`
	Gender     string `json:"gender" form:"gender" binding:"required"`
	No_Telepon string `json:"no_telepon" form:"no_telepon" binding:"required"`
	Alamat     string `json:"alamat" form:"alamat" binding:"required"`
}

//UserCreateDTO is used by client when create user
// type UserCreateDTO struct {
// 	Name     string `json:"name" form:"name" binding:"required"`
// 	Email    string `json:"email" form:"email" binding:"required" validate:"email"`
// 	Password string `json:"password, omitempty" form:"password, omitempty" validate:"min:6" binding:"required" `
// }
