package initializers

import (
	"fmt"

	"github.com/Imran-Sarkar-Sabbir/gin-jwt-auth/models"
)

func SyncDatabase() {
	fmt.Println(DB)
	fmt.Println("Database")
	DB.AutoMigrate(&models.User{})
}
