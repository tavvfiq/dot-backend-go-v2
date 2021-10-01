package it

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
	"go-backend-v2/internal/model"
	_userHttpDelivery "go-backend-v2/internal/user/delivery/http"
	_userRepo "go-backend-v2/internal/user/repository/postgres"
	_userUsecase "go-backend-v2/internal/user/usecase"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
)

type e2eTestSuite struct {
	suite.Suite
	baseUrl    string
	randSource rand.Source
	r          *rand.Rand
}

func TestE2ETestSuite(t *testing.T) {
	suite.Run(t, new(e2eTestSuite))
}

func (s *e2eTestSuite) SetupTest() {
	s.randSource = rand.NewSource(time.Now().UnixNano())
	s.r = rand.New(s.randSource)
}

func (s *e2eTestSuite) SetupSuite() {
	// initialize database
	dbHost := os.Getenv("HOST_POSTGRE")
	dbPort := os.Getenv("PORT_POSTGRE")
	dbUser := os.Getenv("USER_POSTGRE")
	dbPass := os.Getenv("PASSWORD_POSTGRE")
	dbName := os.Getenv("DBNAME_POSTGRE")
	var dsn string
	dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPass, dbName, dbPort)
	pg := _db.NewPostgresDb(dsn)

	// create user with name "taufiq" and ID = 20 (for testing purpose) and sample article
	pg.Db(context.Background()).Model(&model.User{}).Create(&model.User{ID: 20, Name: "taufiq"})
	// article for update
	pg.Db(context.Background()).Model(&model.Article{}).Create(&model.Article{
		ID:       1,
		Title:    "e2e test",
		Subtitle: "e2e test for update",
		AuthorId: 20,
		Content:  "e2e test content",
	})
	// article for delete
	pg.Db(context.Background()).Model(&model.Article{}).Create(&model.Article{
		ID:       2,
		Title:    "e2e test",
		Subtitle: "e2e test for delete",
		AuthorId: 20,
		Content:  "e2e test content",
	})
	// comment for delete
	pg.Db(context.Background()).Model(&model.Comment{}).Create(&model.Comment{
		ID:        1,
		UserId:    20,
		ArticleId: 1,
		Content:   "test e2e delete comment",
	})

	// create cache provider
	redisAddr := os.Getenv("REDIS_ADDR")
	redisPass := os.Getenv("REDIS_PASS")
	redisDB, _ := strconv.ParseInt(os.Getenv("REDIS_DB"), 10, 64)
	rd := redis.NewRedis(redisAddr, redisPass, int(redisDB))

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
	s.baseUrl = fmt.Sprintf("http://localhost%s/v1", port)
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

	go func() {
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
	}()
}

func (s *e2eTestSuite) TearDownSuite() {
	p, _ := os.FindProcess(syscall.Getpid())
	p.Signal(os.Interrupt)
}
