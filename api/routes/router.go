package routes

import (
	"fmt"
	"kkl-v2/api/controllers"
	"kkl-v2/data/repositories"
	"kkl-v2/data/services"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func NewRouter(router fiber.Router, db *gorm.DB, rd *redis.Client) {
	userRepo := repositories.NewUserRepository(db)
	userSvc := services.NewUserService(userRepo)
	userCtrl := controllers.NewUserCtrl(userSvc)
	fmt.Print(userCtrl)
	// Implement routes
}
