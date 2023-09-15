package routes

import (
	"net/http"

	"github.com/Real-Dev-Squad/tiny-site-backend/models"
	"github.com/Real-Dev-Squad/tiny-site-backend/services"
	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

func userRoutes(rg *gin.RouterGroup, db *bun.DB) {
	users := rg.Group("/users")

	users.GET("", func(ctx *gin.Context) {

		var users []models.User
		err := db.NewSelect().Model(&users).OrderExpr("id ASC").Limit(10).Scan(ctx)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "error",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": users,
		})
	})

	users.POST("", func(ctx *gin.Context) {
		var user models.User
		err := ctx.BindJSON(&user)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "error",
			})
			return
		}
		//user.Hashed_passkey = hashFunc(user.Hashed_passkey, generateSalt(16))
		user.Hashed_passkey = services.HashFunc(user.Hashed_passkey, services.GenerateSalt(0))
		_, err = db.NewInsert().Model(&user).Exec(ctx)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": err,
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "user created successfully",
			"data":    "",
		})
	})
}

// // Define salt size
// const saltSize = 16

// // Generate 16 bytes randomly
// func generateSalt(saltSize int) []byte {
// 	var salt = make([]byte, saltSize)
// 	_, err := rand.Read(salt[:])

// 	if err != nil {
// 		panic(err)
// 	}

// 	return salt
// }

// // Combine password and salt then hash them using the SHA-512
// func hashFunc(password string, salt []byte) string {
// 	// Convert password string to byte slice
// 	var pwdByte = []byte(password)

// 	// Create sha-512 hasher
// 	var sha512 = sha512.New()

// 	pwdByte = append(pwdByte, salt...)

// 	sha512.Write(pwdByte)

// 	// Get the SHA-512 hashed password
// 	var hashedPassword = sha512.Sum(nil)

// 	// Convert the hashed to hex string
// 	var hashedPasswordHex = hex.EncodeToString(hashedPassword)
// 	return hashedPasswordHex
// }
