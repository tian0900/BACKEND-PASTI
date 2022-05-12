package middleware

import (
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/NestyTampubolon/golang_gin_gorm_JWT/helper"
	"github.com/NestyTampubolon/golang_gin_gorm_JWT/service"
)

//AuthorizeJWT validates the token user given, return 401 if not valid
func AuthorizeJWT(jwtService service.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response := helper.BuildErrorResponse("Failed to process request", "no token found", nil) //gagal request
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		token, err := jwtService.ValidateToken(authHeader)
		if token.Valid {	//validasi token
			claims := token.Claims.(jwt.MapClaims)
			log.Println("claim[user_id] : ", claims["user_id"])	//mengambil user id
			log.Println("Claim[issuer] :", claims["issuer"])
		} else {
			log.Println(err)
			response := helper.BuildErrorResponse("Token is not valid", err.Error(), nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}
	}
}
