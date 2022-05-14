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

type ProdukController interface {
	All(context *gin.Context)
	FindByID(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	UpdateStok(context *gin.Context)
	Delete(context *gin.Context)
}

type produkController struct {
	produkService service.ProdukService
	jwtService    service.JWTService
}

func NewProdukController(produkServ service.ProdukService, jwtServ service.JWTService) ProdukController {
	return &produkController{
		produkService: produkServ,
		jwtService:    jwtServ,
	}
}

func (c *produkController) All(context *gin.Context) { //fungsi menampilkan semua data
	var produks []entity.Produk = c.produkService.All()
	// res := helper.BuildResponse(true, "OK", produks)
	context.JSON(http.StatusOK, produks)
}

func (c *produkController) FindByID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
	}
	var produk entity.Produk = c.produkService.FIndById(id)
	if (produk == entity.Produk{}) {
		res := helper.BuildErrorResponse("Data not found", "No Data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		// res := helper.BuildResponse(true, "OK", produk)
		context.JSON(http.StatusOK, produk)
	}
}

func (c *produkController) Insert(context *gin.Context) {
	var produkCreateDTO dto.ProdukCreateDTO
	errDTO := context.ShouldBind(&produkCreateDTO)
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
		result := c.produkService.Insert(produkCreateDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusOK, response)
	}
}

func (c *produkController) Update(context *gin.Context) {
	var produkUpdateDTO dto.ProdukUpdateDTO
	errDTO := context.ShouldBind(&produkUpdateDTO)
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
	produkUpdateDTO.Produk_id = id
	result := c.produkService.Update(produkUpdateDTO)
	response := helper.BuildResponse(true, "OK", result)
	context.JSON(http.StatusOK, response)

}

func (c *produkController) Delete(context *gin.Context) {
	var produk entity.Produk
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to get id", "No param id were found", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}

	produk.Produk_id = id
	c.produkService.Delete(produk)
	res := helper.BuildResponse(true, "Delete", helper.EmptyObj{})
	context.JSON(http.StatusOK, res)

}


func (c *produkController) UpdateStok(context *gin.Context) {
	var produkUpdateDTO dto.ProdukUpdateDTO
	errDTO := context.ShouldBind(&produkUpdateDTO)
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

	var produk entity.Produk = c.produkService.FIndById(id)
	if (produk == entity.Produk{}) {
		res := helper.BuildErrorResponse("Data not found", "No Data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} 
	produkUpdateDTO.Produk_id = id
	produkUpdateDTO.Stok = produk.Stok - produkUpdateDTO.Kuantitas
	produkUpdateDTO.Harga = produk.Harga
	produkUpdateDTO.Gambar = produk.Gambar
	produkUpdateDTO.Kategori = produk.Kategori
	produkUpdateDTO.Nama = produk.Nama
	result := c.produkService.Update(produkUpdateDTO)
	response := helper.BuildResponse(true, "OK", result)
	context.JSON(http.StatusOK, response)

}
