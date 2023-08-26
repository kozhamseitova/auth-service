package service

import (
	"context"

	"github.com/kozhamseitova/auth-service/internal/entity"
	"github.com/kozhamseitova/auth-service/utils"
)

func (m *Manager) Create(ctx context.Context) (string, error) {
	id, err := m.repository.Create(ctx)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (m *Manager) Login(ctx context.Context, id string) (*entity.Token, error) {
	_, err := m.repository.GetUserById(ctx, id) 
	if err != nil {
		return nil, err
	}

	accesToken, err := m.token.CreateToken(id, m.config.TimeToLiveAccess)
	if err != nil {
		return nil, err
	}

	refreshToken, err := utils.GenerateRefreshToken()
	if err != nil {
		return nil, err
	}

	hashedRefreshToken, err := utils.HashPassword(refreshToken)
	if err != nil {
		return nil, err
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

func (m *Manager) Refresh(ctx context.Context, refreshToken string) (*entity.Token, error) {
	user, err := m.repository.GetByRefreshToken(ctx, refreshToken) 
	if err != nil {
		return nil, err
	}

	err = utils.CheckPassword(refreshToken, user.RefreshToken)
	if err != nil {
		return nil, err
	}

	accesToken, err := m.token.CreateToken(user.ID, m.config.TimeToLiveAccess)
	if err != nil {
		return nil, err
	}

	newRefreshToken, err := utils.GenerateRefreshToken()
	if err != nil {
		return nil, err
	}

	hashedRefreshToken, err := utils.HashPassword(newRefreshToken)
	if err != nil {
		return nil, err
	}

	err = m.repository.UpdateRefreshToken(ctx, user.ID, hashedRefreshToken)
	if err != nil {
		return nil, err
	}

	tokens := &entity.Token{
		AccessToken: accesToken,
		RefreshToken: refreshToken,
	}

	return tokens, nil

}

func(m *Manager) VerifyToken(ctx context.Context, token string) (string, error) {
	payload, err := m.token.ValidateToken(token)
	if err != nil {
		return "", utils.ErrInternalError
	}

	return payload.UserId, nil
}
