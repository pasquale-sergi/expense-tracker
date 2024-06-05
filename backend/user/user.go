package user

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pasquale-sergi/expense-tracker/databaseLogic"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       uint   `gorm: "primary key"`
	Username string `gorm: "unique"`
	Email    string `gorm: "unique"`
	Password string
}

var userID uint

var Cookie string

func Signup(c *gin.Context) {
	//Get the email/pass off the req
	var body struct {
		Username string
		Email    string
		Password string
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}
	//Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	//Create the user
	user := User{Username: body.Username, Email: body.Email, Password: string(hash)}
	result := databaseLogic.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed create user",
		})
		return
	}

	//Response
	c.JSON(http.StatusOK, gin.H{})
}

func Login(c *gin.Context) {
	//get email and passw from req body
	var body struct {
		Username string
		Password string
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}
	// look up requested user
	var user User
	databaseLogic.DB.First(&user, "username = ?", body.Username)
	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})

		return
	}
	//compare sent in pass with saved hash pass
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Password Incorrect",
		})
		return
	}
	//generate a jtw token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	//sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SESSION_SECRET_KEY")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed generating token",
		})
		return
	}

	//send it back
	//set the cookie
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("auth-token", tokenString, 3600*24*30, "/", "localhost", false, true)
	Cookie, _ = c.Cookie("auth-token")
	fmt.Print("COOKIE: ", Cookie)

	// Validate(c)
	userID = user.ID
	c.JSON(http.StatusOK, gin.H{"message": "Logged in successfully", "userID": user.ID})

}

// func Validate(c *gin.Context) {
// 	user, _ := c.Get("user")
// 	userID = user.(User).ID
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": userID,
// 	})
// }

func Logout(c *gin.Context) {
	c.SetCookie("Authorization", "", -1, "/", "localhost", false, true)
	c.String(http.StatusOK, "Cookie has been deleted")
}
