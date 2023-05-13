package service

import (
	"hacktiv8-msib-final-project-3/entity"
	"hacktiv8-msib-final-project-3/repository/userrepository"

	"github.com/gin-gonic/gin"
)

type AuthService interface {
	Authentication() gin.HandlerFunc
}

type authService struct {
	userRepo userrepository.UserRepository
}

func NewAuthService(userRepo userrepository.UserRepository) AuthService {
	return &authService{userRepo}
}

func (a *authService) Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bearerToken := ctx.GetHeader("Authorization")

		var user entity.User

		if err := user.ValidateToken(bearerToken); err != nil {
			ctx.AbortWithStatusJSON(err.StatusCode(), err)
			return
		}

		result, err := a.userRepo.GetUserByID(user.ID)
		if err != nil {
			ctx.AbortWithStatusJSON(err.StatusCode(), err)
			return
		}

		ctx.Set("userData", result)
		ctx.Next()
	}
}
