package main

import (
	"flag"
	"ivan_search/api"
	"ivan_search/configuration"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	confName := flag.String("conf", "server.cfg", "Config file path")
	flag.Parse()
	log.Println("Starting server...")

	// Парсинг конфигурационного файла
	mainCfg := configuration.Configuration{}
	mainCfg.SetParams(*confName)

	allConnections := configuration.Connections{}
	// Установка коннекта с Postgres (опционально)
	// allConnections.EstablishPostgresConnections(&mainCfg.PostgresDatabaseCfg)
	// for _, v := range allConnections {
	// 	defer v.Close()
	// }
	// Установка коннекта с Redis (опционально)
	// configuration.EstablishRedisConnection(&mainCfg.RedisCfg)

	// Запуск http-сервера
	gin.SetMode(gin.DebugMode)
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(DataBaseMiddleware(&allConnections))

	// Инициализация "групп" и главного API
	mainGroup := r.Group("/api")
	mainGroupV001 := mainGroup.Group("/v0.0.1")
	api.InitAPI(mainGroupV001, &allConnections)
	// Инициализация доступа к FE
	r.Use(static.Serve("/", static.LocalFile("./fe_static", true)))

	// Старт сервера
	r.Run(":" + mainCfg.ServerCfg.Port)
	s := &http.Server{
		Addr:         ":" + mainCfg.ServerCfg.Port,
		Handler:      r,
		ReadTimeout:  180 * time.Second,
		WriteTimeout: 180 * time.Second,
	}

	// NO SSL
	s.ListenAndServe()

	// SSL
	// go s.ListenAndServe()
	// http.ListenAndServeTLS(":"+mainCfg.ServerCfg.PortSSL, "server.crt", "server.key", r)
}

// DataBaseMiddleware Middleware для Postgres
func DataBaseMiddleware(databases *configuration.Connections) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		for k, v := range *databases {
			ctx.Set(k, v)
		}
		ctx.Next()
	}
}
