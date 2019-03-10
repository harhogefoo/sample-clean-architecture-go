package usecase

import "github.com/harhogefoo/sample-clean-architecture-go/src/app/domain"

type UserRepository interface {
	Store(domain.User) (int, error)
	FindById(int) (domain.User, error)
	FindAll() (domain.Users, error)
}
