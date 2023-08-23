package service

import (
	"context"

	"github.com/kozhamseitova/auth-service/internal/entity"
	"github.com/kozhamseitova/auth-service/utils"
)

func (m *Manager) Create(ctx context.Context, user *entity.User) (string, error) {
	user, err := m.repository.GetUserByName(ctx, user.Name)
	if err != nil {
		return "", err
	}

	if user != nil {
		return "", utils.ErrUserAlreadyExists
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return "", utils.ErrInternalError
	}

	user.Password = hashedPassword

	id, err := m.repository.Create(ctx, user)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (m *Manager) Login(ctx context.Context, name, password string) (string, string, error) {
	user, err := m.repository.GetUserByName(ctx, name) 
	if err != nil {
		return "", "", err
	}

	err = utils.CheckPassword(password, user.Password)
	if err != nil {
		return "", "", err
	}

	accesToken, err := m.token.CreateToken(user.ID, m.config.TimeToLiveAccess)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := m.token.CreateToken(user.ID, m.config.TimeToLiveRefresh)
	if err != nil {
		return "", "", err
	}

	hashedRefreshToken, err := utils.HashPassword(refreshToken)
	if err != nil {
		return "", "", err
	}

	err = m.repository.UpdateRefreshToken(ctx, user.ID, hashedRefreshToken)
	if err != nil {
		return "", "", err
	}

	return accesToken, refreshToken, nil

}

func (m *Manager) Refresh(ctx context.Context, id, refreshToken string) (string, string, error) {
	user, err := m.repository.GetUserById(ctx, id) 
	if err != nil {
		return "", "", err
	}

	err = utils.CheckPassword(refreshToken, user.Password)
	if err != nil {
		return "", "", err
	}

	accesToken, err := m.token.CreateToken(user.ID, m.config.TimeToLiveAccess)
	if err != nil {
		return "", "", err
	}

	newRefreshToken, err := m.token.CreateToken(user.ID, m.config.TimeToLiveRefresh)
	if err != nil {
		return "", "", err
	}

	hashedRefreshToken, err := utils.HashPassword(newRefreshToken)
	if err != nil {
		return "", "", err
	}

	err = m.repository.UpdateRefreshToken(ctx, user.ID, hashedRefreshToken)
	if err != nil {
		return "", "", err
	}

	return accesToken, newRefreshToken, nil

}
