package user

import (
	"SmartLocker/e"
	"SmartLocker/model"
	"SmartLocker/service/article"
	"SmartLocker/util"
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
func (u *User) Get() int { //Param:Id/Username
	var user *model.User
	var err error

	if u.Id != 0 {
		user, err = model.GetUserInfoById(u.Id)
	} else if u.Username != "" {
		user, err = model.GetUserInfoByName(u.Username)
	} else {
		return e.InvalidParams
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
func (u *User) getArticles() int { // Param:Id
	a, err := article.GetArticles(u.Id)
	if err != e.Success {
		return err
	}
	u.Articles = a
	return e.Success
}

// register
func (u *User) Register() int {
	if u.Username == "" || u.Password == "" {
		return e.InvalidParams
	}
	err := model.AddUser(u.Username, util.EncodeSha256(u.Password), model.USER, "")
	if err != nil {
		log.WithError(err).Warn("Couldn't register")
		return e.InternalError
	}
	return e.Success
}

func (u *User) Verify() (bool, int) {
	if u.Username == "" {
		return false, e.InvalidParams
	}

	user, err := model.GetUserInfoByName(u.Username)

	if err != nil {
		log.WithError(err).Warn("couldn't get user info")
		return false, e.InternalError
	}

	if user.Password == util.EncodeSha256(u.Password) {
		u.Id = user.Id
		u.Role = user.Role
		return true, e.Success
	} else {
		return false, e.Success
	}
}
