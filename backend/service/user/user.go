package user

import (
	"SmartLocker/e"
	"SmartLocker/model"
	"SmartLocker/service/article"
	"github.com/go-playground/log"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       int
	Username string
	Password []byte
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
	u.Username = user.Username
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
	if u.Username == "" || string(u.Password) == "" {
		return e.InvalidParams
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(u.Password, bcrypt.DefaultCost)
	if err != nil {
		log.WithError(err).Warn("couldn't generate password")
		return e.InternalError
	}

	err = model.AddUser(u.Username, hashedPassword, model.USER, "")
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

	err = bcrypt.CompareHashAndPassword(user.Password, u.Password)

	if err == nil {
		u.Id = user.Id
		u.Role = user.Role
		return true, e.Success
	} else {
		return false, e.Success
	}

}
