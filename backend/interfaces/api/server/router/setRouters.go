package router

import (
	"net/http"

	"github.com/yoshihiro-shu/draft-backend/interfaces/api/server/auth"
	"github.com/yoshihiro-shu/draft-backend/interfaces/api/server/handler"
	"github.com/yoshihiro-shu/draft-backend/interfaces/api/server/middleware"
	"github.com/yoshihiro-shu/draft-backend/interfaces/api/server/request"
	"github.com/yoshihiro-shu/draft-backend/registory"
)

func (r Router) ApplyRouters() {
	ctx := request.NewContext(r.Config)
	r.Use(middleware.CorsMiddleware)

	h := handler.Handler{
		Context: ctx,
	}

	{
		r.AppHandle("/", h.Index).Methods(http.MethodGet)
	}
	{
		// th := api.NewTopPageHandler(ctx)
		topPageHandler := registory.NewTopPageRegistory(ctx)
		r.AppHandle("/top", topPageHandler.Get).Methods(http.MethodGet)
	}
	{
		twitterHandler := registory.NewTwitterRegistory(ctx)
		twitter := r.Group("/twitter")
		twitter.AppHandle("/timeline", twitterHandler.GetTimeLine).Methods(http.MethodGet)
	}
	{
		// Grouping
		t := r.Group("/test")
		t.AppHandle("", h.TestHandler).Methods(http.MethodGet)
		t.AppHandle("/redis", h.TestSetRedis).Methods(http.MethodPost)
		t.AppHandle("/redis/{key}", h.TestGetRedis).Methods(http.MethodGet)
		t.AppHandle("/v2", h.Index).Methods(http.MethodGet)
	}
	{
		c := r.Group("/cmd")
		c.AppHandle("", h.Command).Methods(http.MethodGet)
	}
	{
		user := r.Group("/users")
		userHandler := registory.NewUserRegistory(ctx)
		user.AppHandle("/login", userHandler.Login).Methods(http.MethodPost)
		user.AppHandle("/signup", userHandler.SignUp).Methods(http.MethodPost)
		// user.HandleFunc("/register", h.RegisterAccount).Methods(http.MethodPost)
	}
	{
		articleHandler := registory.NewArticleRegistory(ctx)
		article := r.Group("/articles")
		// article.AppHandle("", h.GetArticles).Methods(http.MethodGet)
		article.AppHandle("/{id}", articleHandler.Get).Methods(http.MethodGet)
	}
	{
		a := r.Group("/auth")
		a.Use(auth.AuthMiddleware)
		a.AppHandle("/index", h.AuthIndex).Methods(http.MethodGet)
	}
}
