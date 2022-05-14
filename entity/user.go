package entity

type User struct {
	User_id    uint64 `gorm:"primary_key:auto_increament" json:"user_id"`
	Name       string `gorm:"type:varchar(255)" json:"name"`
	Email      string `gorm:"uniqueIndex;type:varchar(255)" json:"email"`
	Password   string `gorm:"->;<-;not null" json:"-"`
	Role       string `gorm:"type:varchar(255);default:Customer" json:"role"`
	Usia       uint64 `gorm:"type:varchar(255)" json:"usia"`
	Gender     string `gorm:"type:varchar(255)" json:"gender"`
	No_Telepon string `gorm:"type:varchar(255)" json:"no_telepon"`
	Alamat     string `gorm:"type:varchar(255)" json:"alamat"`
	Token      string `gorm:"-" json:"token,omitempty"`
}
