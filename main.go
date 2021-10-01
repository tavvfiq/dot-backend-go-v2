package main

import (
	"context"
	"fmt"
	_articleHttpDelivery "go-backend-v2/internal/article/delivery/http"
	_articleRepoCache "go-backend-v2/internal/article/repository/cache"
	_articleRepo "go-backend-v2/internal/article/repository/postgres"
	_articleUsecase "go-backend-v2/internal/article/usecase"
	_commentHttpDelivery "go-backend-v2/internal/comment/delivery/http"
	_commentRepo "go-backend-v2/internal/comment/repository/postgres"
	_commentUsecase "go-backend-v2/internal/comment/usecase"
	_db "go-backend-v2/internal/infrastructure/database/postgres"
	"go-backend-v2/internal/infrastructure/database/redis"
	"go-backend-v2/internal/interface/middleware"
	_userHttpDelivery "go-backend-v2/internal/user/delivery/http"
	_userRepo "go-backend-v2/internal/user/repository/postgres"
	_userUsecase "go-backend-v2/internal/user/usecase"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func init() {
	// get root path
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)
	// get server type
	server := os.Getenv("ENV")
	// load env based on server type
	viper.SetConfigFile(fmt.Sprintf("%s/.%s.env", basepath, server))
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("error on reading config file. got: %v", err)
	}
}

func main() {
	// initialize database
	dbHost := viper.GetString("HOST_POSTGRE")
	dbPort := viper.GetString("PORT_POSTGRE")
	dbUser := viper.GetString("USER_POSTGRE")
	dbPass := viper.GetString("PASSWORD_POSTGRE")
	dbName := viper.GetString("DBNAME_POSTGRE")
	var dsn string
	dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPass, dbName, dbPort)
	pg := _db.NewPostgresDb(dsn)

	// create cache provider
	redisAddr := viper.GetString("REDIS_ADDR")
	redisPass := viper.GetString("REDIS_PASS")
	redisDB := viper.GetInt("REDIS_DB")
	rd := redis.NewRedis(redisAddr, redisPass, redisDB)

	// create new repository
	userRepo := _userRepo.NewUserPostgresRepository(pg)
	articleRepo := _articleRepo.NewArticlePostgresRepository(pg)
	// inject article repo with cache provider
	articleRepo = _articleRepoCache.NewArticleRepositoryWithCache(rd, articleRepo)
	commentRepo := _commentRepo.NewCommentPostgresRepository(pg)

	// create new usecase
	userUsecase := _userUsecase.NewUserUsecase(userRepo)
	articleUsecase := _articleUsecase.NewArticleUsecase(articleRepo, commentRepo)
	commentUsecase := _commentUsecase.NewCommentUsecase(commentRepo)

	// new http handler
	e := echo.New()
	// error handler middleware
	e.Use(middleware.ErrorHandler)
	// grouping by version
	v1 := e.Group("v1")
	_userHttpDelivery.NewUserHttpHandler(v1, userUsecase)
	_articleHttpDelivery.NewArticleHttpHandler(v1, articleUsecase)
	_commentHttpDelivery.NewCommentHttpHandler(v1, commentUsecase)

	// set and Run the server
	// server port
	port := os.Getenv("PORT")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		oscall := <-c
		log.Printf("system call: %+v", oscall)
		cancel()
	}()

	errServer := make(chan error)

	go func() {
		log.Printf("server running on port: %s", port)
		errServer <- e.Start(port)
	}()

	select {
	case <-ctx.Done():
		ctxShutDown, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := e.Shutdown(ctxShutDown); err != nil {
			log.Fatalf("server shutdown failed: %+s", err)
		}
		log.Printf("server exited properly")
	case err := <-errServer:
		log.Fatalf("server error. got: %v", err)
	}
}
