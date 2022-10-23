package service

import (
	"fiber-go-exercise/pkg/models"

	"go.uber.org/zap"
)

type PostService struct {
	postRepository *models.PostRepository
}

func (ps *PostService) FindAllPosts() (*[]models.Post, error) {
	return ps.postRepository.FindAll()
}

func (ps *PostService) CreatePost(p *models.Post) error {
	err := ps.postRepository.Save(p)
	if err != nil {
		zap.S().Warn("Unable to create a new post, error: ", err)
		return err
	}
	return nil
}

func (ps *PostService) FindPostByID(id uint32) (*models.Post, error) {
	u, err := ps.postRepository.FindOne(id)
	if err != nil {
		zap.S().Warn("Unable to find post, error: ", err)
		return nil, err
	}
	return u, nil
}

func (ps *PostService) DeletePost(id uint32) error {
	_, err := ps.postRepository.DeleteOne(id)
	if err != nil {
		zap.S().Warn("Unable to delete post, error: ", err)
		return err
	}
	return nil
}

func (ps *PostService) UpdatePost(id uint32, p *models.Post) (*models.Post, error) {
	err := ps.postRepository.UpdateOne(id, p)

	if err != nil {
		zap.S().Warn("Unable to update post, error: ", err)
		return nil, err
	}

	post, err := ps.postRepository.FindOne(id)
	if err != nil {
		zap.S().Warn("Unable to find updated post data, error: ", err)
		return nil, err
	}
	return post, nil
}
