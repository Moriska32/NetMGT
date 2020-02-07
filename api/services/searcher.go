package services

import (
	"time"

	"github.com/gin-gonic/gin"
)

var (
	database = map[string]string{
		"privet":   "kak dela",
		"kak dela": "norm",
		"lala":     "lele",
	}
)

type SearchAnswer struct {
	Data *string   `json:"data,omitmepty"`
	Date time.Time `json:"request_date,omitempty"`
}

// SearchData Сервис поиска версия 1
func SearchData() gin.HandlerFunc {
	fn := func(ctx *gin.Context) {
		requestedText := ctx.Param("search_field")

		answer := SearchAnswer{
			Data: nil,
			Date: time.Now(),
		}
		for k, v := range database {
			if k == requestedText {
				answer.Data = &v
			}
		}
		ctx.SecureJSON(200, gin.H{"data": answer})

	}
	return gin.HandlerFunc(fn)
}

// SearchDataV2 Сервис поиска версия 2
func SearchDataV2() gin.HandlerFunc {
	fn := func(ctx *gin.Context) {
		query := ctx.Request.URL.Query()
		request := []string{}
		ok := false
		if request, ok = query["req"]; !ok {
			ctx.AbortWithStatusJSON(400, gin.H{"Error": "Provide field \"req\" please"})
			return
		}

		if len(request) == 0 {
			ctx.AbortWithStatusJSON(400, gin.H{"Error": "Field \"req\" should not be empty"})
			return
		}

		requestedText := request[0]
		answer := SearchAnswer{
			Data: nil,
			Date: time.Now(),
		}
		for k, v := range database {
			if k == requestedText {
				answer.Data = &v
			}
		}
		ctx.SecureJSON(200, gin.H{"data": answer})
	}
	return gin.HandlerFunc(fn)
}
