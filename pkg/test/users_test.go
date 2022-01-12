package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/laynerkerML/clase7/cmd/service/handler"
	"github.com/laynerkerML/clase7/internal/domain"
	"github.com/laynerkerML/clase7/internal/users"
	"github.com/laynerkerML/clase7/pkg/store"
	"github.com/stretchr/testify/assert"
)

func createService() *gin.Engine {
	_ = os.Setenv("TOKEN", "123456")
	db := store.New(store.FileType, "../../users.json")
	repo := users.NewRepository(db)
	service := users.NewService(repo)
	p := handler.NewUser(service)
	r := gin.Default()

	pr := r.Group("api/v1/users")
	{
		pr.GET("/", p.ValidationToken(), p.GetAll())
	}
	return r
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "123456")
	return req, httptest.NewRecorder()
}

func Test_GetUsers_OK(t *testing.T) {
	type response struct {
		Data []domain.User
	}
	r := createService()
	req, rr := createRequestTest(http.MethodGet, "/api/v1/users/", "")
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	var objRes response
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	assert.Nil(t, err)
	assert.True(t, len(objRes.Data) > 0)
}
