package service

import "fiber-go-exercise/pkg/models"

type Service struct {
	userService UserService
	postService PostService
}

func New(ur *models.UserRepository, pr *models.PostRepository) *Service {
	return &Service{
		userService: UserService{
			userRepository: ur,
		},
		postService: PostService{
			postRepository: pr,
		},
	}
}

func (s *Service) GetUserService() *UserService {
	return &s.userService
}

func (s *Service) GetPostService() *PostService {
	return &s.postService
}
