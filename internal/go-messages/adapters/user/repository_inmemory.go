package user

import (
	"context"
	"github.com/kaspermroz/go-message-backend/internal/go-messages/domain"
)

type RepositoryInMemory struct {
	users map[domain.UUID]domain.User
}

func NewRepositoryInMemory() RepositoryInMemory {
	id := domain.MustNewUUID("2137")
	user := domain.MustNewUser(id, domain.MustNewName("test-user"))

	return RepositoryInMemory{
		users: map[domain.UUID]domain.User{
			id: user,
		},
	}
}

func (r RepositoryInMemory) UserByID(_ context.Context, userID domain.UUID) (domain.User, error) {
	user, ok := r.users[userID]
	if !ok {
		return domain.User{}, ErrUserNotFound
	}

	return user, nil
}
