package routes

import (
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/Real-Dev-Squad/tiny-site-backend/models"
	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

func tinyRoutes(rg *gin.RouterGroup, db *bun.DB) {
	tiny_url := rg.Group("/Tinyurl")

	tiny_url.GET("", func(ctx *gin.Context) {

		var users []models.Tinyurl
		err := db.NewSelect().Model(&users).OrderExpr("id ASC").Limit(10).Scan(ctx)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": err,
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": users,
		})
	})

	tiny_url.POST("", func(ctx *gin.Context) {
		var body models.Tinyurl

		err := ctx.BindJSON(&body)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err,
			})
			return
		}
		body.Short_url = encryptURL(body.Org_url)
		body.Created_at = time.Now()
		body.Created_by = "System"
		body.Valid_up = time.Now().AddDate(0, 0, 7)
		//	fmt.Println(decryptURL(body.Short_url))
		_, err = db.NewInsert().Model(&body).Exec(ctx)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "error",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "user created successfully",
			"data":    "",
		})
	})

}
func encryptURL(url string) string {
	// Generate a random 6-character string.
	characters := "abcdefghijklmnopqrstuvwxyz0123456789"
	encryptedURL := make([]byte, 6)
	for i := range encryptedURL {
		encryptedURL[i] = characters[rand.Intn(len(characters))]
	}
	return string(encryptedURL)
}

func decryptURL(encryptedURL string) string {
	// Reverse the encryption process.
	characters := "abcdefghijklmnopqrstuvwxyz0123456789"
	decryptedURL := ""
	for _, char := range encryptedURL {
		decryptedURL += string(characters[len(characters)-strings.Index(string(char), characters)-1])
	}
	return decryptedURL
}

// func md5_url(url string) string {
// 	url = url + time.Nanosecond.String()
// 	hash := md5.New()
// 	hash.Write([]byte(url))
// 	return hex.EncodeToString(hash.Sum(nil))[:8]
// }
