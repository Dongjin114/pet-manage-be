package main

import (
	"log"
	"net/http"

	"pet-manage-be/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	// Gin 라우터 생성
	r := gin.Default()

	// 미들웨어 설정
	r.Use(gin.Logger())     // 요청 로깅
	r.Use(gin.Recovery())   // 패닉 복구
	r.Use(corsMiddleware()) // CORS 처리

	// 라우트 설정
	setupRoutes(r)

	// 서버 시작
	port := ":8080"
	log.Printf("서버가 포트 %s에서 시작됩니다", port)
	if err := r.Run(port); err != nil {
		log.Fatal("서버 시작 실패:", err)
	}
}

// CORS 미들웨어 - 크로스 오리진 요청 허용
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 모든 오리진에서의 요청 허용
		c.Header("Access-Control-Allow-Origin", "*")
		// 허용할 HTTP 메서드들
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		// 허용할 헤더들
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		// OPTIONS 요청 처리 (preflight 요청)
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// 모든 라우트 설정
func setupRoutes(r *gin.Engine) {
	// 헬스체크 엔드포인트 - 서비스 상태 확인
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "Pet Management API is running",
		})
	})

	// API v1 라우트 그룹
	v1 := r.Group("/api/v1")
	{
		// 펫 관련 라우트
		pets := v1.Group("/pets")
		{
			pets.GET("", handlers.GetPets)          // 펫 목록 조회
			pets.GET("/:id", handlers.GetPet)       // 특정 펫 조회
			pets.POST("", handlers.CreatePet)       // 펫 생성
			pets.PUT("/:id", handlers.UpdatePet)    // 펫 수정
			pets.DELETE("/:id", handlers.DeletePet) // 펫 삭제
		}

		// 소유자 관련 라우트
		owners := v1.Group("/owners")
		{
			owners.GET("", handlers.GetOwners)          // 소유자 목록 조회
			owners.GET("/:id", handlers.GetOwner)       // 특정 소유자 조회
			owners.POST("", handlers.CreateOwner)       // 소유자 생성
			owners.PUT("/:id", handlers.UpdateOwner)    // 소유자 수정
			owners.DELETE("/:id", handlers.DeleteOwner) // 소유자 삭제
		}

		// 급식 관련 라우트
		meals := v1.Group("/meals")
		{
			meals.GET("/types", handlers.GetMealTypes)          // 급식 타입 조회
			meals.GET("/items", handlers.GetMealItems)          // 급식 아이템 목록 조회
			meals.GET("/items/:id", handlers.GetMealItem)       // 특정 급식 아이템 조회
			meals.POST("/items", handlers.CreateMealItem)       // 급식 아이템 생성
			meals.PUT("/items/:id", handlers.UpdateMealItem)    // 급식 아이템 수정
			meals.DELETE("/items/:id", handlers.DeleteMealItem) // 급식 아이템 삭제
		}
	}
}
