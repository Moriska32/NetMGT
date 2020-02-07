package services

import "github.com/gin-gonic/gin"

// SayHello Сервис теста
func SayHello() gin.HandlerFunc {
	fn := func(ctx *gin.Context) {
		ctx.SecureJSON(200, gin.H{"Message": "Hello, world!"})
	}
	return gin.HandlerFunc(fn)
}
