package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/afrimadhia/task-5-vix-btpns-afrima_dhia_defara/config"
	"github.com/afrimadhia/task-5-vix-btpns-afrima_dhia_defara/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func Authorization(c *gin.Context) {
	//Get cookie
	tokenString, err := c.Cookie("Auth")

	if err != nil {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"message": "Silahkan Login Terlebih Dahulu",
		})
		return
	}
	//Decode / validate

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		//check exp date

		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.HTML(http.StatusUnauthorized, "login.html", gin.H{
				"message": "Silahkan Login Terlebih Dahulu",
			})
			return
		}

		//find user with token
		var user models.Users
		config.DB.First(&user, claims["sub"])

		if user.ID == 0 {
			c.HTML(http.StatusUnauthorized, "login.html", gin.H{
				"message": "Silahkan Login Terlebih Dahulu",
			})
			return
		}
		//attach to req
		c.Set("user", user)

		//continue
		c.Next()

	} else {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{})
	}

}
