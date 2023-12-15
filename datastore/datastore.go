package datastore

import (
	"database/sql"
	"strconv"

	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"

	"books_manager/model"
)

type book struct{}

func New() *book {
	return &book{}
}

func (s *book) GetByID(ctx *gofr.Context, id int) (*model.Book, error) {
	var resp model.Book

	err := ctx.DB().QueryRowContext(ctx, " SELECT id,name,age,class FROM books where id=$1", id).
		Scan(&resp.ID, &resp.Name, &resp.Age, &resp.Class)
	switch err {
	case sql.ErrNoRows:
		return &model.Book{}, errors.EntityNotFound{Entity: "book", ID: strconv.Itoa(id)}
	case nil:
		return &resp, nil
	default:
		return &model.Book{}, err
	}
}

func (s *book) Create(ctx *gofr.Context, book *model.Book) (*model.Book, error) {
	var resp model.Book

	err := ctx.DB().QueryRowContext(ctx, "INSERT INTO books (name, age, class) VALUES($1,$2,$3)"+
		" RETURNING  id,name,age,class", book.ID, book.Name, book.Age, book.Class).Scan(
		&resp.ID, &resp.Name, &resp.Age, &resp.Class)

	if err != nil {
		return &model.Book{}, errors.DB{Err: err}
	}

	return &resp, nil
}

func (s *book) Update(ctx *gofr.Context, book *model.Book) (*model.Book, error) {
	_, err := ctx.DB().ExecContext(ctx, "UPDATE books SET name=$1,age=$2,class=$3 WHERE id=$4",
		book.Name, book.Age, book.Class, book.ID)
	if err != nil {
		return &model.Book{}, errors.DB{Err: err}
	}

	return book, nil
}

func (s *book) Delete(ctx *gofr.Context, id int) error {
	_, err := ctx.DB().ExecContext(ctx, "DELETE FROM books where id=$1", id)
	if err != nil {
		return errors.DB{Err: err}
	}

	return nil
}