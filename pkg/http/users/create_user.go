package users

import "C"
import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"regexp"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}


func CreateOneUser(db *sql.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var (
			user Person
			userID int
			emailValid = regexp.MustCompile(`^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+$`).MatchString
		)
		_ = c.BindJSON(&user)
		_ = db.QueryRow("SELECT id FROM users WHERE email = $1", user.Email).Scan(&userID)
		if userID >= 1 {
			c.AbortWithStatusJSON(400, gin.H{
				"error": "user already exists",
			})
			return
		}

		if len(user.Password) < 5 {
			c.AbortWithStatusJSON(400, gin.H{
				"error": "minimal password 6 chars",
			})
			return
		}

		if !emailValid(user.Email) {
			c.AbortWithStatusJSON(400, gin.H{
				"error": "invalid email",
			})
			return
		}

		stmt, err := db.Prepare("INSERT INTO users (name,email,password) VALUES ($1, $2, crypt($3,gen_salt('bf')) ) RETURNING id")
		
		if err != nil {
			fmt.Println("na salah disi")
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		err = stmt.QueryRow(user.Name, user.Email, user.Password).Scan(&userID)
		if err != nil {
			fmt.Println("heem disini")
			c.AbortWithStatusJSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		defer stmt.Close()
		c.JSON(200, gin.H{
			"id": userID,
		})
	}
}