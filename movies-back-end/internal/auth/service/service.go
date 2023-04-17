package service

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"movies-service/internal/auth"
	"movies-service/internal/dto"
	"movies-service/internal/errors"
	"movies-service/internal/models"
	"strings"
	"time"
)

type AuthClaims struct {
	jwt.StandardClaims
	Username string `json:"username"`
	UserId   string `json:"userId"`
}

type authService struct {
	userRepository auth.UserRepository
	signingKey     []byte
	expireDuration time.Duration
	clientId       string
	clientSecret   string
}

func NewAuthService(
	userRepository auth.UserRepository,
	signingKey []byte,
	tokenTTL int64,
	clientId string,
	clientSecret string,
) auth.Service {
	return &authService{
		userRepository: userRepository,
		signingKey:     signingKey,
		expireDuration: time.Second * time.Duration(tokenTTL),
		clientId:       clientId,
		clientSecret:   clientSecret,
	}
}

func (a *authService) SignUp(ctx context.Context, userDto *dto.UserDto) error {
	fmtUsername := strings.ToLower(userDto.Username)
	euser, _ := a.userRepository.FindUserByUsername(ctx, fmtUsername)

	if euser != nil {
		return errors.ErrUserExisted
	}
	user := &models.User{
		UserID:    uuid.New().String(),
		Username:  fmtUsername,
		Password:  userDto.Password,
		FirstName: userDto.FirstName,
		LastName:  userDto.LastName,
	}
	user.HashPassword()
	err := a.userRepository.InsertUser(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (a *authService) SignIn(ctx context.Context, username, password string) (string, error) {
	user, _ := a.userRepository.FindUserByUsername(ctx, username)
	if user == nil {
		return "", errors.ErrNotFound
	}

	if !user.ComparePassword(password) {
		return "", errors.ErrWrongPassword
	}
	claims := AuthClaims{
		Username: user.Username,
		UserId:   user.UserID,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			Issuer:    "qushift",
			ExpiresAt: time.Now().Add(a.expireDuration).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(a.signingKey)
}

func (a *authService) ParseToken(ctx context.Context, accessToken string) (string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return a.signingKey, nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(*AuthClaims); ok && token.Valid {
		return claims.UserId, nil
	}

	return "", errors.ErrInvalidAccessToken
}

func (a *authService) VerifyToken(ctx context.Context, c *gin.Context, introspectForm *dto.Introspect) (*models.User, error) {
	token := introspectForm.Token

	if !(introspectForm.ClientId == a.clientId &&
		introspectForm.ClientSecret == a.clientSecret) {
		return nil, errors.ErrInvalidClient
	}

	userId, err := a.ParseToken(ctx, token)
	if err != nil {
		return nil, errors.ErrInvalidAccessToken
	}

	user, err := a.userRepository.FindUserById(ctx, fmt.Sprintf("%s", userId))
	if err != nil {
		return nil, errors.ErrNotFound
	}
	return user, nil
}
