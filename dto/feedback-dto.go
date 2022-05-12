package dto

type FeedbackUpdateDTO struct {
	Feedback_id uint64 `json:"feedback_id"`
	Id_customer uint64 `json:"id_customer,omitempty" form:"id_customer,omitempty"`
	Subjek      uint64 `json:"subjek"`
	Deskripsi   uint64 `json:"deskripsi"`
}

type FeedbackCreateDTO struct {
	Id_customer uint64 `json:"id_customer,omitempty" form:"id_customer,omitempty"`
	Id_produk   uint64 `json:"id_produk"`
	Subjek      uint64 `json:"subjek" form:"subjek" binding:"required"`
	Deskripsi   uint64 `json:"deskripsi" form:"deskripsi" binding:"required"`
}
