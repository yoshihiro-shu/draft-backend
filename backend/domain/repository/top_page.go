package repository

import "github.com/yoshihiro-shu/draft-backend/domain/model"

type TopPageRepository interface {
	GetArticles(*[]model.Article, int, int) error
	GetPager(*model.Article) (int, error)
}
