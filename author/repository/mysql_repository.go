package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/bxcodec/go-clean-arch/author"
	"github.com/bxcodec/go-clean-arch/models"
)

type mysqlAuthorRepo struct {
	DB *mongo.Database
}

// NewMysqlAuthorRepository will create an implementation of author.Repository
func NewMysqlAuthorRepository(db *mongo.Database) author.Repository {
	return &mysqlAuthorRepo{
		DB: db,
	}
}

func (m *mysqlAuthorRepo) getOne(ctx context.Context, query string, args ...interface{}) (*models.Author, error) {
	filter := bson.D{{"name", "Ash"}}
	a := &models.Author{}
	_ = m.DB.Collection("author").FindOne(ctx, filter).Decode(&a)

	return a, nil
}

func (m *mysqlAuthorRepo) GetByID(ctx context.Context, id int64) (*models.Author, error) {
	query := `SELECT id, name, created_at, updated_at FROM author WHERE id=?`
	return m.getOne(ctx, query, id)
}
