package users

import (
	"api-go/pkg/helpers/auth"
	"database/sql"
	"github.com/gin-gonic/gin"
	"regexp"
	"fmt"
)

type Cred struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

func Login(db *sql.DB) func (c *gin.Context) {
	return func(c *gin.Context) {
		var (
			user Cred
			emailValid = regexp.MustCompile(`^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+$`).MatchString
			email string
			id int
			idForPassword int
		)

		_ = c.BindJSON(&user)
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

		_ = db.QueryRow("SELECT id, email FROM users WHERE email = $1", user.Email).Scan(&id, &email)
		if len(email) == 0 {
			c.AbortWithStatusJSON(400, gin.H{
				"error": "User doesn't exists",
			})
			return
		}

		_ = db.QueryRow("SELECT id FROM users WHERE email = $1 AND password = crypt($2, password)", user.Email, user.Password).Scan(&idForPassword)
		fmt.Println(idForPassword)
		if idForPassword < 1 {
			c.AbortWithStatusJSON(400, gin.H{
				"error": "Bad credentials",
			})
			return
		}
		token, _, err := auth.CreateToken(id, email)
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"token": token,
		})
	}
}