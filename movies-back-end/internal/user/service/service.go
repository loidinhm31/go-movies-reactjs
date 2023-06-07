package service

import (
	"context"
	"fmt"
	"log"
	"movies-service/config"
	"movies-service/internal/common/constant"
	"movies-service/internal/common/dto"
	"movies-service/internal/common/entity"
	"movies-service/internal/common/model"
	"movies-service/internal/control"
	"movies-service/internal/errors"
	"movies-service/internal/mail"
	"movies-service/internal/middlewares"
	"movies-service/internal/role"
	"movies-service/internal/user"
	"movies-service/pkg/pagination"
	"strings"
	"time"
)

type userService struct {
	config         *config.Config
	mgmtCtrl       control.Service
	roleRepository role.Repository
	userRepository user.UserRepository
	mailService    mail.Service
}

func NewUserService(config *config.Config, mgmtCtrl control.Service, roleRepository role.Repository, userRepository user.UserRepository, mailService mail.Service) user.Service {
	return &userService{
		config:         config,
		mgmtCtrl:       mgmtCtrl,
		roleRepository: roleRepository,
		userRepository: userRepository,
		mailService:    mailService,
	}
}

func (us *userService) GetUsers(ctx context.Context, pageRequest *pagination.PageRequest, key string, isNew bool) (*pagination.Page[*dto.UserDto], error) {
	log.Println("checking admin privilege...")
	username := fmt.Sprintf("%s", ctx.Value(middlewares.CtxUserKey))
	if !us.mgmtCtrl.CheckAdminPrivilege(username) {
		return nil, errors.ErrUnAuthorized
	}

	page := &pagination.Page[*entity.User]{}

	userResults, err := us.userRepository.FindAllUsers(ctx, pageRequest, page, key, isNew)
	if err != nil {
		log.Println(err)
		return nil, errors.ErrResourceNotFound
	}

	var userDtos []*dto.UserDto
	for _, u := range userResults.Content {
		userDtos = append(userDtos, &dto.UserDto{
			ID:        u.ID,
			Username:  u.Username,
			Email:     u.Email,
			FirstName: u.FirstName,
			LastName:  u.LastName,
			IsNew:     u.IsNew,
			Role: dto.RoleDto{
				ID:       u.Role.ID,
				RoleName: u.Role.RoleName,
				RoleCode: u.Role.RoleCode,
			},
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
		})
	}
	return &pagination.Page[*dto.UserDto]{
		PageSize:      pageRequest.PageSize,
		PageNumber:    pageRequest.PageNumber,
		Sort:          pageRequest.Sort,
		TotalElements: userResults.TotalElements,
		TotalPages:    userResults.TotalPages,
		Content:       userDtos,
	}, nil
}

func (us *userService) UpdateUserRole(ctx context.Context, userDto *dto.UserDto) error {
	log.Println("checking admin privilege...")
	username := fmt.Sprintf("%s", ctx.Value(middlewares.CtxUserKey))
	if !us.mgmtCtrl.CheckAdminPrivilege(username) {
		return errors.ErrUnAuthorized
	}

	theUser, err := us.userRepository.FindUserByID(ctx, userDto.ID)
	if err != nil {
		return err
	}

	theRole, err := us.roleRepository.FindRoleByRoleCode(ctx, userDto.Role.RoleCode)
	if err != nil {
		return err
	}

	err = us.userRepository.UpdateUserRole(ctx, theUser.ID, theRole.ID)
	if err != nil {
		return err
	}

	// Send email to user
	go func() {
		htmlMessage := fmt.Sprintf(`
		<strong>Updated Role</strong><br>
		Dear %s %s, <br>
		This is notification for your updated role to <strong>%s</strong>.`,
			theUser.FirstName, theUser.LastName, strings.ToUpper(theRole.RoleName))

		err := us.mailService.SendMessage(ctx, &model.MailData{
			To:           theUser.Email,
			From:         us.config.Mail.From,
			Subject:      constant.UpdateRoleSubject,
			Content:      htmlMessage,
			TemplateMail: "basic.html",
		})
		if err != nil {
			log.Println(err)
		}

	}()

	return nil
}

func (us *userService) AddOidcUser(ctx context.Context, userDto *dto.UserDto) (*dto.UserDto, error) {
	log.Println("Checking user...")
	fmtUsername := strings.ToLower(userDto.Username)
	euser, _ := us.userRepository.FindUserByUsernameAndIsNew(ctx, fmtUsername, false)
	if euser != nil {
		return nil, errors.ErrUserExisted
	}

	log.Println("checking admin privilege...")
	author := fmt.Sprintf("%s", ctx.Value(middlewares.CtxUserKey))
	if !us.mgmtCtrl.CheckAdminPrivilege(author) {
		return nil, errors.ErrUnAuthorized
	}

	getRole, err := us.roleRepository.FindRoleByRoleCode(ctx, userDto.Role.RoleCode)

	err = us.userRepository.InsertUser(ctx, &entity.User{
		Username:  userDto.Username,
		Email:     userDto.Email,
		FirstName: userDto.FirstName,
		LastName:  userDto.LastName,
		Role:      getRole,
		IsNew:     false,
		CreatedAt: time.Now(),
		CreatedBy: author,
		UpdatedAt: time.Now(),
		UpdatedBy: author,
	})
	if err != nil {
		return nil, err
	}
	return userDto, nil
}
