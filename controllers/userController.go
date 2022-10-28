package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/afrimadhia/task-5-vix-btpns-afrima_dhia_defara/config"
	"github.com/afrimadhia/task-5-vix-btpns-afrima_dhia_defara/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func Index(c *gin.Context) {
	user, _ := c.Get("user")

	if user.(models.Users).ID == 0 {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{})
	} else {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	}
}

// LOGIN
func LoginGet(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

func LoginPost(c *gin.Context) {
	//catch form
	username := c.PostForm("username")
	password := c.PostForm("password")

	//look up user
	var user models.Users

	config.DB.First(&user, "username = ?", username)

	if user.ID == 0 {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"message": "Invalid Username or password",
		})
		return
	}

	//compare sent with pass database
	errPassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if errPassword != nil {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"message": "Invalid Username or password",
		})
		return
	}

	//generate jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token",
		})
		return
	}

	//send it back
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Auth", tokenString, 3600, "", "", false, true)

	c.HTML(http.StatusOK, "index.html", gin.H{})
}

// LOGOUT
func Logout(c *gin.Context) {
	c.SetCookie("Auth", "", -1, "", "", false, true)
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

// SIGNUP
func SignupGet(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{})
}
func SignupPost(c *gin.Context) {
	//Catch form
	email := c.PostForm("email")
	password := c.PostForm("password")
	username := c.PostForm("username")
	nama_lengkap := c.PostForm("nama_lengkap")

	//Hashing password
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to Hash Password",
		})
		return
	}

	//Create user
	user := models.Users{Email: email, Password: string(hash), Username: username, NamaLengkap: nama_lengkap}
	result := config.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to Create User",
		})

		return
	}

	//Respond
	c.Redirect(http.StatusFound, "/login")
}

// CREATE FILES

func CreateFilesGet(c *gin.Context) {
	user, _ := c.Get("user")

	if user.(models.Users).ID == 0 {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{})
	} else {
		c.HTML(http.StatusOK, "upload.html", gin.H{})
	}

}

func CreateFilesPost(c *gin.Context) {

	// single file
	file, err := c.FormFile("image")

	if err != nil {
		c.HTML(http.StatusBadRequest, "upload.html", gin.H{
			"error": "Failed to upload image",
		})
		return
	}

	//save file
	err = c.SaveUploadedFile(file, "assets/uploads/"+file.Filename)

	if err != nil {
		c.HTML(http.StatusBadRequest, "upload.html", gin.H{
			"error": "Failed to upload image",
		})
		return
	}

	//respond
	c.HTML(http.StatusOK, "index.html", gin.H{
		"image":   "assets/uploads/" + file.Filename,
		"message": "Upload Berhasil",
	})
}
