package models

import "gorm.io/gorm"

type Repository struct {
	userRepository UserRepository
	postRepository PostRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		UserRepository{
			db: db,
		},
		PostRepository{
			db: db,
		},
	}
}

func (r *Repository) GetUserRepository() *UserRepository {
	return &r.userRepository
}

func (r *Repository) GetPostRepository() *PostRepository {
	return &r.postRepository
}
