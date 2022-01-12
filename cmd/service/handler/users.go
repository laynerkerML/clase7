package handler

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/laynerkerML/clase7/internal/domain"
	"github.com/laynerkerML/clase7/internal/users"
	"github.com/laynerkerML/clase7/pkg/web"
)

type User struct {
	service users.Service
}

func NewUser(e users.Service) *User {
	return &User{
		service: e,
	}
}

// ListUsers godoc
// @Summary List users
// @Tags Users
// @Description get users
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /api/v1/users [get]
func (u *User) GetAll() gin.HandlerFunc {
	type response struct {
		Data []domain.User `json:"data"`
	}

	return func(c *gin.Context) {
		ctx := context.Background()
		users, err := u.service.GetAll(ctx)
		if err != nil {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}

		//c.JSON(http.StatusOK, web.NewResponse(http.StatusOK, &response{users}, ""))
		c.JSON(http.StatusOK, &response{users})
	}
}

func (u *User) Save() gin.HandlerFunc {

	type response struct {
		Data interface{} `json:"data"`
	}
	return func(c *gin.Context) {
		var req domain.User

		err := c.Bind(&req)
		if err != nil {
			c.JSON(422, "json decoding: "+err.Error())
		}
		ctx := context.Background()
		fmt.Println("data ....", req)
		user, _ := u.service.Save(ctx, req)

		c.JSON(http.StatusOK, &response{user})
	}
}

func (u *User) Update() gin.HandlerFunc {
	type response struct {
		Data interface{} `json:"data"`
	}

	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid ID"})
			return
		}

		var req domain.User
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		ctx := context.Background()
		user, err := u.service.Update(ctx, int(id), req)
		if err != nil {
			c.JSON(500, err.Error())
			return
		}

		c.JSON(http.StatusOK, &response{user})
	}
}

func (u *User) Patch() gin.HandlerFunc {
	type response struct {
		Data interface{} `json:"data"`
	}

	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid ID"})
			return
		}

		var req domain.User
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		ctx := context.Background()
		user, err := u.service.FielUpdate(ctx, int(id), req)
		if err != nil {
			c.JSON(500, err.Error())
			return
		}

		c.JSON(http.StatusOK, &response{user})
	}
}

func (u *User) ValidationToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		tokenEnv := os.Getenv("TOKEN")
		if token != tokenEnv {
			c.AbortWithStatusJSON(401, web.NewResponse(401, nil, "token inválido"))
		}
		c.Next()
	}
}

func (u *User) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid ID"})
			return
		}
		err = u.service.Delete(int(id))
		if err != nil {
			c.AbortWithStatusJSON(400, web.NewResponse(400, nil, "error al eliminar el registro"))
		}
		c.JSON(http.StatusOK, web.NewResponse(http.StatusOK, fmt.Sprintf("El usuario %d a sido eliminado", id), ""))
	}
}
