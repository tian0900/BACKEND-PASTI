package controller

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strconv"
	"time"

	"github.com/NestyTampubolon/golang_gin_gorm_JWT/dto"
	"github.com/NestyTampubolon/golang_gin_gorm_JWT/entity"
	"github.com/NestyTampubolon/golang_gin_gorm_JWT/helper"
	"github.com/NestyTampubolon/golang_gin_gorm_JWT/service"
	"github.com/gin-gonic/gin"
)

type KeranjangController interface {
	All(context *gin.Context)
	FindByID(context *gin.Context)
	Insert(context *gin.Context)
	InsertPemesanan(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
	FindUserID(context *gin.Context)
}

type keranjangController struct {
	keranjangService       service.KeranjangService
	jwtService             service.JWTService
	pemesananService       service.PemesananService
	pemesanandetailService service.PemesananDetailService
}

func NewKeranjangController(keranjangServ service.KeranjangService, jwtServ service.JWTService) KeranjangController {
	return &keranjangController{
		keranjangService: keranjangServ,
		jwtService:       jwtServ,
	}
}

type pemesanan1Controller struct {
	pemesananService       service.PemesananService
	keranjangController    service.KeranjangService
	pemesanandetailService service.PemesananDetailService
	jwtService             service.JWTService
}

func NewPemesanan1Controller(pemesananServ service.PemesananService, jwtServ service.JWTService) PemesananController {
	return &pemesananController{
		pemesananService: pemesananServ,
		jwtService:       jwtServ,
	}
}
func (c *keranjangController) All(context *gin.Context) { //fungsi menampilkan semua data
	var keranjangs []entity.Keranjang = c.keranjangService.All()
	res := helper.BuildResponse(true, "OK", keranjangs)
	context.JSON(http.StatusOK, res)
}

func (c *keranjangController) FindByID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
	}
	var keranjang entity.Keranjang = c.keranjangService.FIndById(id)
	if (keranjang == entity.Keranjang{}) {
		res := helper.BuildErrorResponse("Data not found", "No Data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK", keranjang)
		context.JSON(http.StatusOK, res)
	}
}

func (c *keranjangController) Insert(context *gin.Context) {
	var keranjangCreateDTO dto.KeranjangCreateDTO
	errDTO := context.ShouldBind(&keranjangCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {

		keranjangCreateDTO.Kuantitas = keranjangCreateDTO.Stok
		keranjangCreateDTO.Total = keranjangCreateDTO.Harga * keranjangCreateDTO.Stok
		
		result := c.keranjangService.Insert(keranjangCreateDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusOK, response)
	}
}

func (c *keranjangController) Update(context *gin.Context) {
	var keranjangUpdateDTO dto.KeranjangUpdateDTO
	errDTO := context.ShouldBind(&keranjangUpdateDTO)
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
	keranjangUpdateDTO.Id_keranjang = id
	result := c.keranjangService.Update(keranjangUpdateDTO)
	response := helper.BuildResponse(true, "OK", result)
	context.JSON(http.StatusOK, response)

}

func (c *keranjangController) Delete(context *gin.Context) {
	var keranjang entity.Keranjang
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to get id", "No param id were found", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}

	keranjang.Id_keranjang = id
	c.keranjangService.Delete(keranjang)
	res := helper.BuildResponse(true, "Delete", helper.EmptyObj{})
	context.JSON(http.StatusOK, res)

}

func (c *keranjangController) getUserIDByToken(token string) string { //fungsi menampilkan data sesuai ID
	aToken, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := aToken.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id
}

func (c *keranjangController) FindUserID(context *gin.Context) {
	authHeader := context.GetHeader("Authorization")
	id_customer := c.getUserIDByToken(authHeader)

	convertedUserID, err := strconv.ParseUint(id_customer, 10, 64)

	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
	}
	var pemesanan []entity.Keranjang = c.keranjangService.FIndByUserId(convertedUserID)
	res := helper.BuildResponse(true, "OK", pemesanan)
	context.JSON(http.StatusOK, res)

}

func (c *keranjangController) InsertPemesanan(context *gin.Context) {
	var pemesananCreateDTO dto.PemesananCreateDTO

	errDTO := context.ShouldBind(&pemesananCreateDTO)

	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		authHeader := context.GetHeader("Authorization")
		id_customer := c.getUserIDByToken(authHeader)

		convertedUserID, err := strconv.ParseUint(id_customer, 10, 64)

		if err == nil {
			pemesananCreateDTO.Id_customer = convertedUserID
			pemesananCreateDTO.Tanggal_pemesanan = time.Now()
			pemesananCreateDTO.Status = "Verifikasi"
		}

		var pemesananDetailCreateDTO dto.PemesananDetailCreateDTO
		errDTO1 := context.ShouldBind(&pemesananDetailCreateDTO)
		if errDTO1 != nil {
			res := helper.BuildErrorResponse("Failed to process request", errDTO1.Error(), helper.EmptyObj{})
			context.JSON(http.StatusBadRequest, res)
		} else {
			var keranjang []entity.Keranjang = c.keranjangService.FIndByUserId(convertedUserID)
			res := helper.BuildResponse(true, "OK", keranjang)
			context.JSON(http.StatusOK, res)
			result := c.pemesananService.Insert(pemesananCreateDTO)
			helper.BuildResponse(true, "OK", keranjang)
			// for i := 0; i < len(keranjang); i++ {
			// pemesananDetailCreateDTO.Id_pemesanan = result.Id_pemesanan
			// pemesananDetailCreateDTO.Id_produk = keranjang[i].Id_produk
			// pemesananDetailCreateDTO.Kuantitas_pesan = keranjang[i].Kuantitas
			// pemesananDetailCreateDTO.Total_harga = keranjang[i].Total
			// c.pemesanandetailService.Insert(pemesananDetailCreateDTO)
			// 	fmt.Print(keranjang[].Total)
			// }
			response := helper.BuildResponse(true, "OK", result)
			context.JSON(http.StatusOK, response)
		}
	}
}
