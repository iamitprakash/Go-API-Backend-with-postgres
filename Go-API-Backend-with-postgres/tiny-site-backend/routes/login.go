package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Real-Dev-Squad/tiny-site-backend/models"
	"github.com/Real-Dev-Squad/tiny-site-backend/services"
	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

func LoginRoutes(rg *gin.RouterGroup, db *bun.DB) {
	login := rg.Group("/login")

	login.GET("", func(ctx *gin.Context) {
		var tuser models.User
		decoder := json.NewDecoder(ctx.Request.Body)
		err1 := decoder.Decode(&tuser)
		if err1 != nil {

			fmt.Println(err1)
		}
		user := new(models.User)
		userpassword := services.HashFunc(tuser.Hashed_passkey, services.GenerateSalt(0))
		err := db.NewSelect().Model(user).Where("? = ?", bun.Ident("username"), tuser.Username).Where("? =?", bun.Ident("hashed_passkey"), userpassword).Scan(ctx)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": err,
			})
			log.Println("rrrr", err)

		}
		claims1 := models.Claims{
			Username:  user.Username,
			Role:      user.Roles,
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
		}
		token, err2 := services.GenerateBearerToken(services.Claims(claims1))
		if err != nil {
			//w.WriteHeader(http.StatusInternalServerError)
			fmt.Println(err2)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			//"message": user,
			"Auth token": token,
		})
	})
}
