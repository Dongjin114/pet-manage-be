package routes

import (
	"net/http"

	"pet-manage-be/internal/interface/http/handlers"

	"github.com/gin-gonic/gin"
)

// SetupRoutes 모든 라우트 설정
func SetupRoutes(r *gin.Engine, ownerHandler *handlers.OwnerHandler, petHandler *handlers.PetHandler, mealHandler *handlers.MealHandler) {
	// 헬스체크 엔드포인트
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "Pet Management API is running",
		})
	})

	// API v1 라우트 그룹
	v1 := r.Group("/api/v1")
	{
		// 소유자 관련 라우트
		owners := v1.Group("/owners")
		{
			owners.GET("", ownerHandler.GetAllOwners)       // 소유자 목록 조회
			owners.GET("/:id", ownerHandler.GetOwnerByID)   // 특정 소유자 조회
			owners.POST("", ownerHandler.CreateOwner)       // 소유자 생성
			owners.PUT("/:id", ownerHandler.UpdateOwner)    // 소유자 수정
			owners.DELETE("/:id", ownerHandler.DeleteOwner) // 소유자 삭제
		}

		// 펫 관련 라우트
		pets := v1.Group("/pets")
		{
			pets.GET("", petHandler.GetAllPets)                       // 펫 목록 조회
			pets.GET("/:id", petHandler.GetPetByID)                   // 특정 펫 조회
			pets.POST("", petHandler.CreatePet)                       // 펫 생성
			pets.PUT("/:id", petHandler.UpdatePet)                    // 펫 수정
			pets.DELETE("/:id", petHandler.DeletePet)                 // 펫 삭제
			pets.GET("/owner/:owner_id", petHandler.GetPetsByOwnerID) // 소유자별 펫 목록 조회
		}

		// 급식 관련 라우트
		meals := v1.Group("/meals")
		{
			meals.GET("/types", mealHandler.GetMealTypes)               // 급식 타입 조회
			meals.GET("/items", mealHandler.GetAllMealss)               // 급식 아이템 목록 조회
			meals.GET("/items/:id", mealHandler.GetMealsByID)           // 특정 급식 아이템 조회
			meals.POST("/items", mealHandler.CreateMeals)               // 급식 아이템 생성
			meals.PUT("/items/:id", mealHandler.UpdateMeals)            // 급식 아이템 수정
			meals.DELETE("/items/:id", mealHandler.DeleteMeals)         // 급식 아이템 삭제
			meals.GET("/items/type/:type", mealHandler.GetMealssByType) // 타입별 급식 아이템 조회
		}
	}
}
