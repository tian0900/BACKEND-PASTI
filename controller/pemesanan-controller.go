package controller

import (
	"fmt"
	"net/http"
	"strconv"


	"github.com/NestyTampubolon/golang_gin_gorm_JWT/dto"
	"github.com/NestyTampubolon/golang_gin_gorm_JWT/entity"
	"github.com/NestyTampubolon/golang_gin_gorm_JWT/helper"
	"github.com/NestyTampubolon/golang_gin_gorm_JWT/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type PemesananController interface {
	All(context *gin.Context)
	FindByID(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
	FindUserIDD(context *gin.Context)
}

type pemesananController struct {
	pemesananService       service.PemesananService
	keranjangController    service.KeranjangService
	pemesanandetailService service.PemesananDetailService
	jwtService             service.JWTService
}

func NewPemesananController(pemesananServ service.PemesananService, jwtServ service.JWTService) PemesananController {
	return &pemesananController{
		pemesananService: pemesananServ,
		jwtService:       jwtServ,
	}
}

func (c *pemesananController) All(context *gin.Context) { //fungsi menampilkan semua data
	var pemesanans []entity.Pemesanan = c.pemesananService.All()
	res := helper.BuildResponse(true, "OK", pemesanans)
	context.JSON(http.StatusOK, res)
}

func (c *pemesananController) FindByID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
	}
	var pemesanan entity.Pemesanan = c.pemesananService.FIndById(id)
	if (pemesanan == entity.Pemesanan{}) {
		res := helper.BuildErrorResponse("Data not found", "No Data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK", pemesanan)
		context.JSON(http.StatusOK, res)
	}
}

func (c *pemesananController) Insert(context *gin.Context) {
	var pemesananCreateDTO dto.PemesananCreateDTO
	// var pemesananDetailCreateDTO dto.PemesananDetailCreateDTO
	errDTO := context.ShouldBind(&pemesananCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		authHeader := context.GetHeader("Authorization")
		id_customer := c.getUserIDToken(authHeader)

		convertedUserID, err := strconv.ParseUint(id_customer, 10, 64)

		if err == nil {
			pemesananCreateDTO.Id_customer = convertedUserID
			pemesananCreateDTO.Status = "Verifikasi"
		}
		result := c.pemesananService.Insert(pemesananCreateDTO)
		var keranjang []entity.Keranjang = c.keranjangController.FIndByUserId(convertedUserID)
		for i := 0; i < len(keranjang); i++ {
			// pemesananDetailCreateDTO.Id_pemesanan = result.Id_pemesanan
			// pemesananDetailCreateDTO.Id_produk = keranjang[i].Id_produk
			// pemesananDetailCreateDTO.Kuantitas_pesan = keranjang[i].Kuantitas
			// pemesananDetailCreateDTO.Total_harga = keranjang[i].Total
			// c.pemesanandetailService.Insert(pemesananDetailCreateDTO)
			fmt.Print(keranjang[i].Kuantitas)
		}

		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusOK, response)
	}
}

func (c *pemesananController) Update(context *gin.Context) {
	var pemesananUpdateDTO dto.PemesananUpdateDTO
	errDTO := context.ShouldBind(&pemesananUpdateDTO)
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
	pemesananUpdateDTO.Id_pemesanan = id
	result := c.pemesananService.Update(pemesananUpdateDTO)
	response := helper.BuildResponse(true, "OK", result)
	context.JSON(http.StatusOK, response)

}

func (c *pemesananController) Delete(context *gin.Context) {
	var pemesanan entity.Pemesanan
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to get id", "No param id were found", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}

	pemesanan.Id_pemesanan = id
	c.pemesananService.Delete(pemesanan)
	res := helper.BuildResponse(true, "Delete", helper.EmptyObj{})
	context.JSON(http.StatusOK, res)

}

func (c *pemesananController) getUserIDToken(token string) string { //fungsi menampilkan data sesuai ID
	aToken, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := aToken.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id
}

func (c *pemesananController) FindUserIDD(context *gin.Context) {
	authHeader := context.GetHeader("Authorization")
	id_customer := c.getUserIDToken(authHeader)

	convertedUserID, err := strconv.ParseUint(id_customer, 10, 64)

	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
	}
	var pemesanan []entity.Keranjang = c.keranjangController.FIndByUserId(convertedUserID)
	res := helper.BuildResponse(true, "OK", pemesanan)
	context.JSON(http.StatusOK, res)
}
