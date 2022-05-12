package controller

import (
	"net/http"
	"strconv"

	"github.com/NestyTampubolon/golang_gin_gorm_JWT/dto"
	"github.com/NestyTampubolon/golang_gin_gorm_JWT/entity"
	"github.com/NestyTampubolon/golang_gin_gorm_JWT/helper"
	"github.com/NestyTampubolon/golang_gin_gorm_JWT/service"
	"github.com/gin-gonic/gin"
)

//Feedback
type FeedbackController interface {
	All(context *gin.Context)
	FindByID(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type feedbackController struct {
	feedbackService service.FeedbackService
	jwtService      service.JWTService
}

func NewFeedbackController(feedbackServ service.FeedbackService, jwtServ service.JWTService) FeedbackController {
	return &feedbackController{
		feedbackService: feedbackServ,
		jwtService:      jwtServ,
	}
}

func (c *feedbackController) All(context *gin.Context) { //fungsi menampilkan semua data
	var feedbacks []entity.Feedback = c.feedbackService.All()
	// res := helper.BuildResponse(true, "OK", produks)
	context.JSON(http.StatusOK, feedbacks)
}

func (c *feedbackController) FindByID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
	}
	var feedback entity.Feedback = c.feedbackService.FIndById(id)
	if (feedback == entity.Feedback{}) {
		res := helper.BuildErrorResponse("Data not found", "No Data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		// res := helper.BuildResponse(true, "OK", produk)
		context.JSON(http.StatusOK, feedback)
	}
}

func (c *feedbackController) Insert(context *gin.Context) {
	var feedbackCreateDTO dto.FeedbackCreateDTO
	errDTO := context.ShouldBind(&feedbackCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		// authHeader := context.GetHeader("Authorization")
		// userID := c.getUserIDByToken(authHeader)
		// convertedUserID, err := strconv.ParseUint(userID, 10, 64)
		// if err == nil {
		// 	produkCreateDTO.UserID = convertedUserID
		// }
		result := c.feedbackService.Insert(feedbackCreateDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusOK, response)
	}
}

func (c *feedbackController) Update(context *gin.Context) {
	var feedbackUpdateDTO dto.FeedbackUpdateDTO
	errDTO := context.ShouldBind(&feedbackUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}

	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to get id", "No param id were found", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}
	feedbackUpdateDTO.Feedback_id = id
	result := c.feedbackService.Update(feedbackUpdateDTO)
	response := helper.BuildResponse(true, "OK", result)
	context.JSON(http.StatusOK, response)

}

func (c *feedbackController) Delete(context *gin.Context) {
	var feedback entity.Feedback
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to get id", "No param id were found", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}

	feedback.Feedback_id = id
	c.feedbackService.Delete(feedback)
	res := helper.BuildResponse(true, "Delete", helper.EmptyObj{})
	context.JSON(http.StatusOK, res)

}
