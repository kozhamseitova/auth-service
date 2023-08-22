package service

import "context"

func (m *Manager) Create(ctx context.Context) (string, error) {
	return m.repository.Create(ctx)
}