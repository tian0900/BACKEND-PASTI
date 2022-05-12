package dto

type FeedbackUpdateDTO struct {
	Feedback_id uint64 `json:"feedback_id"`
	Id_customer uint64 `json:"id_customer,omitempty" form:"id_customer,omitempty"`
	Subjek      string `json:"subjek"`
	Deskripsi   string `json:"deskripsi"`
}

type FeedbackCreateDTO struct {
	Id_customer uint64 `json:"id_customer,omitempty" form:"id_customer,omitempty"`
	Subjek      string `json:"subjek" form:"subjek" binding:"required"`
	Deskripsi   string `json:"deskripsi" form:"deskripsi" binding:"required"`
}
