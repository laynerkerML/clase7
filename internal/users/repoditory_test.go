package users

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/laynerkerML/clase7/internal/domain"
	"github.com/stretchr/testify/assert"
)

type storeTest struct{}

func (st *storeTest) Read(data interface{}) error {
	datosJson := `[{"id": 1,"nombre": "Laynerker","apellido": "Guerrero","email": "lay@gmail.com","edad": 31,"altura": 2,"activo": true,"fechaCreacion": "2012"}]`
	return json.Unmarshal([]byte(datosJson), &data)
}
func (st *storeTest) Write(data interface{}) error {
	return nil
}

func TestGetAll(t *testing.T) {
	resultadoEsperado := []domain.User{}
	resultadoEsperado = append(resultadoEsperado, domain.User{
		Id:            1,
		Nombre:        "Laynerker",
		Apellido:      "Guerrero",
		Email:         "lay@gmail.com",
		Edad:          31,
		Altura:        2,
		Activo:        true,
		FechaCreacion: "2012",
	})
	storeT := &storeTest{}
	repositorioT := NewRepository(storeT)
	ctx := context.Background()
	resul, _ := repositorioT.GetAll(ctx)

	assert.Equal(t, resultadoEsperado, resul)
}
