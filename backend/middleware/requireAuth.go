package middleware

import (
	"fmt"

	"net/http"
	"os"
	"time"

	"github.com/pasquale-sergi/expense-tracker/databaseLogic"
	"github.com/pasquale-sergi/expense-tracker/user"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *gin.Context) {
	//get the cookie off req
	tokenString := user.Cookie
	fmt.Print("TOKEN: ", tokenString)
	// if err != nil {
	// 	c.AbortWithStatus(http.StatusUnauthorized)
	// 	c.Abort()
	// 	return
	// }
	//decode/validate it
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SESSION_SECREY_KEY")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		//check the exp

		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token expired"})
			c.Abort()
			return
		}

		//find the user with token sub
		var user user.User
		databaseLogic.DB.First(&user, claims["sub"])

		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		//attach to req
		c.Set("user", user)

		//continue
		c.Next()

	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.Abort()
	}
}
