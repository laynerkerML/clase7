package users

import (
	"context"

	"github.com/laynerkerML/clase7/internal/domain"
)

type Service interface {
	GetAll(ctx context.Context) ([]domain.User, error)
	Save(ctx context.Context, user domain.User) (domain.User, error)
	Update(ctx context.Context, id int, user domain.User) (domain.User, error)
	FielUpdate(ctx context.Context, id int, user domain.User) (domain.User, error)
	Delete(id int) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}
func (s *service) GetAll(ctx context.Context) ([]domain.User, error) {
	return s.repository.GetAll(ctx)
}

func (s *service) Save(ctx context.Context, user domain.User) (domain.User, error) {
	lastId, err := s.repository.LastId()
	if err != nil {
		return domain.User{}, err
	}
	user.Id = lastId + 1
	_, err = s.repository.Save(ctx, user)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (s *service) Update(ctx context.Context, id int, user domain.User) (domain.User, error) {
	err := s.repository.Update(ctx, id, user)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (s *service) FielUpdate(ctx context.Context, id int, user domain.User) (domain.User, error) {
	allUsers, err := s.repository.GetAll(ctx)
	lastUser := domain.User{}
	for k, u := range allUsers {
		if u.Id == id {
			lastUser = allUsers[k]
		}
	}
	updatedFields := updateFields(lastUser, user)

	err = s.repository.Update(ctx, id, updatedFields)
	if err != nil {
		return domain.User{}, err
	}
	return updatedFields, nil
}

func updateFields(lastUser domain.User, newUser domain.User) domain.User {
	if newUser.Id != lastUser.Id && newUser.Id != 0 {
		lastUser.Id = newUser.Id
	}
	if newUser.Nombre != lastUser.Nombre && newUser.Nombre != "" {
		lastUser.Nombre = newUser.Nombre
	}
	if newUser.Apellido != lastUser.Apellido && newUser.Apellido != "" {
		lastUser.Apellido = newUser.Apellido
	}
	if newUser.Email != lastUser.Email && newUser.Email != "" {
		lastUser.Email = newUser.Email
	}
	if newUser.Edad != lastUser.Edad && newUser.Edad != 0 {
		lastUser.Edad = newUser.Edad
	}
	if newUser.Altura != lastUser.Altura && newUser.Altura != 0 {
		lastUser.Altura = newUser.Altura
	}
	if newUser.Activo != lastUser.Activo {
		lastUser.Activo = newUser.Activo
	}
	if newUser.FechaCreacion != lastUser.FechaCreacion && newUser.FechaCreacion != "" {
		lastUser.FechaCreacion = newUser.FechaCreacion
	}
	return lastUser
}

func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}
