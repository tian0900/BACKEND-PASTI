package entity

type Feedback struct {
	Feedback_id uint64 `gorm:"primary_key:auto_increment"`
	Id_customer uint64 `gorm:"type:integer" json:"id_customer"`
	Subjek      string `gorm:"type:varchar(255)" json:"subjek"`
	Deskripsi   string `gorm:"type:varchar(255)" json:"deskripsi"`

	User User `gorm:"foreignkey:id_customer;constraint:onUpdate:CASCADE, onDelete:CASCADE" json:"user"`
}
