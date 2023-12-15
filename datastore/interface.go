package datastore

import (
	"simple-rest-api/model"
	"gofr.dev/pkg/gofr"
)

type Books interface {
	GetByID(ctx *gofr.Context, id int) (*model.Books, error)
	Create(ctx *gofr.Context, model *model.Books) (*model.Books, error)
	Update(ctx *gofr.Context, model *model.Books) (*model.Books, error)
	Delete(ctx *gofr.Context, id int) error
}