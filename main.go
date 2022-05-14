package main

import (
	"github.com/NestyTampubolon/golang_gin_gorm_JWT/config"
	"github.com/NestyTampubolon/golang_gin_gorm_JWT/controller"
	"github.com/NestyTampubolon/golang_gin_gorm_JWT/middleware"
	"github.com/NestyTampubolon/golang_gin_gorm_JWT/repository"
	"github.com/NestyTampubolon/golang_gin_gorm_JWT/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db                  *gorm.DB                       = config.SetupDatabaseConnection()
	userRepository      repository.UserRepository      = repository.NewUserRepository(db)
	produkRepository    repository.ProdukRepository    = repository.NewProdukRepository(db)
	keranjangRepository repository.KeranjangRepository = repository.NewKeranjangRepository(db)
	pemesananRepository repository.PemesananRepository = repository.NewPemesananRepository(db)
	feedbackRepository  repository.FeedbackRepository  = repository.NewFeedbackRepository(db)
	jwtService          service.JWTService             = service.NewJWTService()
	userService         service.UserService            = service.NewUserService(userRepository)
	produkService       service.ProdukService          = service.NewProdukService(produkRepository)
	keranjangService    service.KeranjangService       = service.NewKeranjangService(keranjangRepository)
	pemesananService    service.PemesananService       = service.NewPemesananService(pemesananRepository)
	feedbackService     service.FeedbackService        = service.NewFeedbackService(feedbackRepository)
	authService         service.AuthService            = service.NewAuthService(userRepository)
	authController      controller.AuthController      = controller.NewAuthController(authService, jwtService)
	userController      controller.UserController      = controller.NewUserController(userService, jwtService)
	produkController    controller.ProdukController    = controller.NewProdukController(produkService, jwtService)
	keranjangController controller.KeranjangController = controller.NewKeranjangController(keranjangService, jwtService)
	pemesananController controller.PemesananController = controller.NewPemesananController(pemesananService, jwtService)
	feedbackController  controller.FeedbackController  = controller.NewFeedbackController(feedbackService, jwtService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	userRoutes := r.Group("api/user", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/profile", userController.Profile)
		userRoutes.PUT("/profile", userController.Update)
	}

	produkRoutes := r.Group("api/produks")
	{
		produkRoutes.GET("/", produkController.All)
		produkRoutes.POST("/", produkController.Insert)
		produkRoutes.GET("/:id", produkController.FindByID)
		produkRoutes.PUT("/:id", produkController.Update)
		produkRoutes.DELETE("/:id", produkController.Delete)
	}

	keranjangRoutes := r.Group("api/keranjangs", middleware.AuthorizeJWT(jwtService))
	{
		keranjangRoutes.GET("/", keranjangController.All)
		keranjangRoutes.GET("/user", keranjangController.FindUserID)
		keranjangRoutes.POST("/", keranjangController.Insert)
		keranjangRoutes.POST("/insert", keranjangController.InsertPemesanan)
		keranjangRoutes.GET("/:id", keranjangController.FindByID)
		keranjangRoutes.PUT("/:id", keranjangController.Update)
		keranjangRoutes.DELETE("/:id", keranjangController.Delete)
	}

	pemesananRoutes := r.Group("api/pemesanans", middleware.AuthorizeJWT(jwtService))
	{
		pemesananRoutes.GET("/", pemesananController.All)
		pemesananRoutes.GET("/user/:id", pemesananController.FindUserIDD)
		pemesananRoutes.POST("/", pemesananController.Insert)
		pemesananRoutes.GET("/:id", pemesananController.FindByID)
		pemesananRoutes.PUT("/:id", pemesananController.Update)
		pemesananRoutes.DELETE("/:id", pemesananController.Delete)
	}

	feedbackRoutes := r.Group("api/feedbacks")
	{
		feedbackRoutes.GET("/", feedbackController.All)
		feedbackRoutes.POST("/", feedbackController.Insert)
		feedbackRoutes.GET("/:id", feedbackController.FindByID)
		feedbackRoutes.PUT("/:id", feedbackController.Update)
		feedbackRoutes.DELETE("/:id", feedbackController.Delete)
	}
	r.Run()
}
