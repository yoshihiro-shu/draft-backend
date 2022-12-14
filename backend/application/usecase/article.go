package usecase

import (
	"github.com/yoshihiro-shu/draft-backend/domain/model"
	"github.com/yoshihiro-shu/draft-backend/domain/repository"
)

type ArticleUseCase interface {
	Create(title, content string, userId, categoryId int) (*model.Article, error)
	FindByID(id int) (*model.Article, error)
	Update(id int, title, content string) (*model.Article, error)
	Delete(id int) error
}

type articleUseCase struct {
	articleRepo repository.ArticleRepository
}

func NewArticleUseCase(articleRepo repository.ArticleRepository) ArticleUseCase {
	return &articleUseCase{articleRepo: articleRepo}
}

func (au *articleUseCase) Create(title, content string, userId, categoryId int) (*model.Article, error) {
	article := &model.Article{
		Title:      title,
		Content:    content,
		UserId:     userId,
		CategoryId: categoryId,
	}

	createdArticle, err := au.articleRepo.Create(article)
	if err != nil {
		return nil, err
	}

	return createdArticle, nil
}

func (au *articleUseCase) FindByID(id int) (*model.Article, error) {
	foundArticle, err := au.articleRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return foundArticle, nil
}

func (au *articleUseCase) Update(id int, title, content string) (*model.Article, error) {
	return &model.Article{}, nil
}

func (au *articleUseCase) Delete(id int) error {
	return nil
}
