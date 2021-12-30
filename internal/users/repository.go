package users

import (
	"context"
	"errors"
	"fmt"

	"github.com/laynerkerML/clase7/internal/domain"
	"github.com/laynerkerML/clase7/pkg/store"
)

var accesos []domain.User
var lastID int

type Repository interface {
	GetAll(ctx context.Context) ([]domain.User, error)
	Save(ctx context.Context, user domain.User) (int, error)
	Update(ctx context.Context, id int, user domain.User) error
	LastId() (int, error)
	Delete(id int) error
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll(ctx context.Context) ([]domain.User, error) {
	var ass []domain.User
	r.db.Read(&ass)
	return ass, nil
}
func (r *repository) Save(ctx context.Context, user domain.User) (int, error) {
	var ass []domain.User
	r.db.Read(&ass)
	ass = append(ass, user)
	if err := r.db.Write(ass); err != nil {
		return 0, err
	}
	return user.Id, nil
}

func (r *repository) Update(ctx context.Context, id int, user domain.User) error {
	updated := false
	var ass []domain.User
	r.db.Read(&ass)
	for k, u := range ass {
		if u.Id == id {
			ass[k] = user
			updated = true
		}
	}
	if err := r.db.Write(ass); err != nil {
		return err
	}
	if !updated {
		return errors.New("No Se modificio el registor")
	}
	return nil
}

func (r *repository) LastId() (int, error) {
	var acc []domain.User
	if err := r.db.Read(&acc); err != nil {
		return 0, err
	}
	return acc[len(acc)-1].Id, nil
}

func (r *repository) Delete(id int) error {
	deleted := false
	var index int
	var acc []domain.User
	r.db.Read(&acc)
	for i := range acc {
		if acc[i].Id == id {
			index = i
			deleted = true
		}
	}
	if !deleted {
		return fmt.Errorf("producto %d no encontrado", id)
	}
	acc = append(acc[:index], acc[index+1:]...)
	if err := r.db.Write(acc); err != nil {
		return err
	}
	return nil
}
