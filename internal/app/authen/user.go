package authen

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	db "github.com/minhtri6179/tinyface-users-service/internal/gorm"
	util "github.com/minhtri6179/tinyface-users-service/utils"
)

func (server *Server) createUser(ctx *gin.Context) {
	var req db.RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	newUser := db.User{
		HashedPassword: hashedPassword,
		FirstName:      info.GetFirstName(),
		LastName:       info.GetLastName(),
		DateOfBirth:    info.Dob.AsTime(),
		Email:          info.GetEmail(),
		UserName:       info.GetUserName(),
	}
	a.db.Create(&newUser)
	return &authen_and_post.UserResult{
		Status: authen_and_post.UserStatus_OK,
		Info: &authen_and_post.UserDetailInfo{
			UserId:   int64(newUser.ID),
			UserName: newUser.UserName,
		},
	}, nil
	user, err := server.store.CreateUser(ctx, newUser)
	if err != nil {
		if db.ErrorCode(err) == db.UniqueViolation {
			ctx.JSON(http.StatusForbidden, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := newUserResponse(user)
	ctx.JSON(http.StatusOK, rsp)
}

type Server struct {
	redisDatabase *redis.Client
	router        *gin.Engine
}
