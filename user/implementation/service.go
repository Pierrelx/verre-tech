package implementation 

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"

	user "github.com/Pierrelx/verre-tech/user"
	userSvc "github.com/Pierrelx/verre-tech/user"
)

type service struct{
	repository userSvc.Repository
	logger log.Logger
}

//NewService créé et retourne une nouvelle instance du userService
func NewService(repo userSvc.Service, logger log.Logger) userSvc.Service {
	return &service{
		repository: repo,
		logger: logger,
	}
}

//CreateUser permet d'ajouter un utiisateur
func (s *service) CreateUser(ctx context.Context, user user.User)(int64, error ){
	logger := log.With(s.logger, "method", "CreateUser", "mail", user.Email)
	user.CreatedOn = time.Now().Unix()
	user.UpdatedOn = time.Now().Unix()
	var id int64
	if id, err := s.repository.CreateUser(ctx, user); err != nil{
		level.Error(logger).Log("err", err)
		return id, userSvc.ErrCmdRepository
	}
	return id, nil
}

//UpdateUser mets à jour un utlisateur
func (s *service) UpdateUser(ctx context.Context, user user.User)(error){
	logger := log.With(s.logger, "method", "UpdateUser", "id", user.ID, "mail", user.Email)
	user.UpdatedOn = time.Now().Unix()
	if err := s.repository.UpdateUser(ctx, user); err != nil{
		level.Error(logger).Log("err", err)
		return userSvc.ErrCmdRepository
	} 
	return nil
}

//GetUserById retourne un utilisateur grâce à son id
func (s *service) GetUserById(ctx context.Context, ID int)(user.User, error){
	logger := log.With(s.logger, "method", "GetUserById", "id", ID)
	var userToReturn user.User
	if userToReturn, err := s.repository.GetUserById(ctx, ID); err != nil{
		level.Error(logger).Log("err", err)
		return userToReturn, err
	}
	return userToReturn, nil
}

//GetUserByMail retourne un utilisateur grâce à son mail
func (s *service) GetUserByMail(ctx context.Context, mail string)(user.User, error){
	logger := log.With(s.logger, "method", "GetUserByMail", "mail", mail)
	var userToReturn user.User
	if userToReturn, err := s.repository.GetUserByMail(ctx, mail); err != nil{
		level.Error(logger).Log("err", err)
		return userToReturn, err
	}
	return userToReturn, nil
}

//CreateUserPassword permet de créer un mot de passe utilisateur
func (s *service) CreateUserPassword(ctx context.Context, password string, user user.User)(error){
	logger := log.With(s.logger, "method", "CreateUserPassword", "UserId", user.ID)
	if err := s.repository.CreateUserPassword(ctx, password, user); err != nil{
		level.Error(logger).Log("err", err)
		return err
	}
	return nil
}

//UpdatePassword permet de mettre à jour un mot de passe utilisateur
func (s *service) UpdateUserPassword(ctx context.Context, oldPassword string, newPassword string, user user.User) (error){
	logger := log.With(s.logger, "method", "UpdateUserPassword", "UserId", user.ID)
	if err := s.repository.UpdateUserPassword(ctx, oldPassword, newPassword, user); err != nil{
		level.Error(logger).Log("err", err)
		return err
	}
	return nil
}

//LoginUser implique la création d'un cookie utilisateur
func (s *service) LoginUser(ctx context.Context, email string, password string)(error){
	logger := log.With(s.logger, "method", "LoginUser", "Email", email)
	if err := s.repository.LoginUser(ctx, email, password); err != nil{
		level.Error(logger).Log("err", err)
		return err
	}
	return nil
}

//LogoutUser implique la suppression d'un cookie utilisateur
func (s *service) LogoutUser(ctx context.Context, user user.User)(error){
	logger := log.With(s.logger, "method", "LogoutUser", "UserId", user.ID)
	if err := s.repository.LogoutUser(ctx, user); err != nil{
		level.Error(logger).Log("err", err)
		return err
	}
	return nil
}