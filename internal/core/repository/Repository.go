package repository

type Repository struct {
	Auth AuthRepository
	User UserRepository
}
