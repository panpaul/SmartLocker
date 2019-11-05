package user

import (
	"SmartLocker/e"
	"SmartLocker/model"
	"SmartLocker/service/article"
	"github.com/go-playground/log"
)

type User struct {
	Id       int
	Username string
	Password string
	Role     int
	Articles []article.Article
}

// fill the blank
func (u *User) Get() int {
	var user *model.User
	var err error

	if u.Username != "" {
		user, err = model.GetUserInfoByName(u.Username)
	} else {
		user, err = model.GetUserInfoById(u.Id)
	}

	if err != nil {
		log.WithError(err).Warn("Couldn't get user info")
		return e.InternalError
	}

	u.Id = user.Id
	u.Role = user.Role

	return u.getArticles()
}

// get the user's articles
func (u *User) getArticles() int {
	a, err := article.GetArticles(u.Id)
	if err != e.Success {
		return err
	}
	u.Articles = a
	return e.Success
}
