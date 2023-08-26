package service

import (
	"context"
	"errors"

	"github.com/kozhamseitova/auth-service/internal/entity"
	"github.com/kozhamseitova/auth-service/utils"
	"golang.org/x/crypto/bcrypt"
)

func (m *Manager) Create(ctx context.Context) (string, error) {
	return m.repository.Create(ctx)
}

func (m *Manager) Login(ctx context.Context, id string) (*entity.Token, error) {
	_, err := m.repository.GetUserById(ctx, id) 
	if err != nil {
		return nil, err
	}

	accesToken, err := m.token.CreateToken(id, m.config.TimeToLiveAccess)
	if err != nil {
		m.logger.Errorf(ctx, "[CreateToken] err: %v", err)
		return nil, utils.ErrInternalError
	}

	refreshToken, err := utils.GenerateRefreshToken()
	if err != nil {
		m.logger.Errorf(ctx, "[GenerateRefreshToken] err: %v", err)
		return nil, utils.ErrInternalError
	}

	hashedRefreshToken, err := utils.HashPassword(refreshToken)
	if err != nil {
		m.logger.Errorf(ctx, "[HashPassword] err: %v", err)
		return nil, utils.ErrInternalError
	}

	err = m.repository.UpdateRefreshToken(ctx, id, hashedRefreshToken)
	if err != nil {
		return nil, err
	}

	token := &entity.Token{
		AccessToken: accesToken,
		RefreshToken: refreshToken,
	}

	return token, nil

}

func (m *Manager) Refresh(ctx context.Context, reqUser entity.User) (*entity.Token, error) {
	user, err := m.repository.GetUserById(ctx, reqUser.ID) 
	if err != nil {
		return nil, err
	}

	err = utils.CheckPassword(reqUser.RefreshToken, user.RefreshToken)
	if err != nil {
		m.logger.Errorf(ctx, "[CheckPassword] err: %v", err)
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return nil, utils.ErrInvalidCredentials
		}
		return nil, utils.ErrInternalError
	}

	accesToken, err := m.token.CreateToken(user.ID, m.config.TimeToLiveAccess)
	if err != nil {
		m.logger.Errorf(ctx, "[CreateToken] err: %v", err)
		return nil, utils.ErrInternalError
	}

	newRefreshToken, err := utils.GenerateRefreshToken()
	if err != nil {
		m.logger.Errorf(ctx, "[GenerateRefreshToken] err: %v", err)
		return nil, utils.ErrInternalError
	}

	hashedRefreshToken, err := utils.HashPassword(newRefreshToken)
	if err != nil {
		m.logger.Errorf(ctx, "[HashPassword] err: %v", err)
		return nil, utils.ErrInternalError
	}

	err = m.repository.UpdateRefreshToken(ctx, user.ID, hashedRefreshToken)
	if err != nil {
		return nil, err
	}

	tokens := &entity.Token{
		AccessToken: accesToken,
		RefreshToken: newRefreshToken,
	}

	return tokens, nil

}

func(m *Manager) VerifyToken(ctx context.Context, token string) (string, error) {
	payload, err := m.token.ValidateToken(token)
	if err != nil {
		m.logger.Errorf(ctx, "[ValidateToken] err: %v", err)
		return "", err
	}

	return payload.UserId, nil
}
