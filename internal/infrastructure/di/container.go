package di

import (
	"pet-manage-be/internal/infrastructure/config"
	"pet-manage-be/internal/infrastructure/database"
	"pet-manage-be/internal/infrastructure/migrations"
	"pet-manage-be/internal/infrastructure/repository"
	"pet-manage-be/internal/interface/http/handlers"
	"pet-manage-be/internal/interface/http/routes"
	"pet-manage-be/internal/usecase/meal"
	"pet-manage-be/internal/usecase/owner"
	"pet-manage-be/internal/usecase/pet"

	"github.com/gin-gonic/gin"
)

// Container 의존성 주입 컨테이너
type Container struct {
	Config *config.Config
	Router *gin.Engine
}

// NewContainer 새로운 컨테이너 생성
func NewContainer() *Container {
	// 설정 로드
	cfg := config.Load()

	// 데이터베이스 연결
	if err := database.Connect(&cfg.Database); err != nil {
		panic("데이터베이스 연결 실패: " + err.Error())
	}

	// 마이그레이션 실행
	if err := migrations.RunMigrations(database.GetDB()); err != nil {
		panic("마이그레이션 실패: " + err.Error())
	}

	// Gin 라우터 생성
	r := gin.Default()

	// 미들웨어 설정
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(corsMiddleware())

	// 리포지토리 생성
	ownerRepo := repository.NewOwnerRepository()
	petRepo := repository.NewPetRepository()
	mealRepo := repository.NewMealRepository()
	mealRepo.SetDB(database.GetDB())

	// 유스케이스 생성
	ownerUsecase := owner.NewOwnerUsecase(ownerRepo)
	petUsecase := pet.NewPetUsecase(petRepo, ownerRepo)
	mealUsecase := meal.NewMealUsecase(mealRepo)

	// 핸들러 생성
	ownerHandler := handlers.NewOwnerHandler(ownerUsecase)
	petHandler := handlers.NewPetHandler(petUsecase)
	mealHandler := handlers.NewMealHandler(mealUsecase)

	// 라우트 설정
	routes.SetupRoutes(r, ownerHandler, petHandler, mealHandler)

	return &Container{
		Config: cfg,
		Router: r,
	}
}

// CORS 미들웨어
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
