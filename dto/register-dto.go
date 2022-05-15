package dto


type RegisterDTO struct {
	Name     string `json:"name" form:"name" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required"`
	Role       string `json:"role" form:"role"`
	Usia       uint64 `json:"usia,string" form:"usia" binding:"required"`
	Gender     string `json:"gender" form:"gender" binding:"required"`
	No_Telepon string `json:"no_telepon" form:"no_telepon" binding:"required"`
	Alamat     string `json:"alamat" form:"alamat" binding:"required"`
}


