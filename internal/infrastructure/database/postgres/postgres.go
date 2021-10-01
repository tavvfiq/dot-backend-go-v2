package postgres

import (
	"context"
	"go-backend-v2/internal/model"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresInterface interface {
	Db(ctx context.Context) *gorm.DB
}

type postgresDb struct {
	Conn *gorm.DB
}

func NewPostgresDb(dsn string) PostgresInterface {
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Panic(err)
	}
	err = db.AutoMigrate(&model.User{}, &model.Article{}, &model.Comment{})
	if err != nil {
		log.Printf("error on auto migrating. got: %v", err)
	}
	log.Printf("successfully connect db with dsn: %s", dsn)
	return &postgresDb{Conn: db}
}

func (p *postgresDb) Db(ctx context.Context) *gorm.DB {
	tx := ctx.Value("txCtx")
	if tx != nil {
		return tx.(*gorm.DB)
	}
	return p.Conn
}
