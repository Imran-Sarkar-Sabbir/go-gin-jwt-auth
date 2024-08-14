package middlewares

import (
	"net/http"
	"os"
	"time"

	"github.com/Imran-Sarkar-Sabbir/gin-jwt-auth/initializers"
	"github.com/Imran-Sarkar-Sabbir/gin-jwt-auth/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid authorization token",
		})
		return
	}

	token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil {
		println("error parsing token")
	}

	claims := token.Claims.(jwt.MapClaims)
	if float64(time.Now().Unix()) > claims["exp"].(float64) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid authorization token time",
		})
		return
	}

	var user models.User
	initializers.DB.First(&user, claims["sub"])

	if user.ID == 0 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid authorization token time",
		})
		return
	}
	c.Set("user", user)
	c.Next()
}
