package user

import (
	"SmartLocker/e"
	"SmartLocker/model"
	"SmartLocker/service/article"
	"SmartLocker/service/cache"
	"github.com/go-playground/log"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"unicode"
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
	var errInt int

	if u.Id != 0 {
		user, errInt = cache.GetUserInfo(strconv.Itoa(u.Id))
	} else if u.Username != "" {
		user, errInt = cache.GetUserInfo(u.Username)
	} else {
		return e.InvalidParams
	}

	if errInt == e.Success {
		u.Id = user.Id
		u.Username = user.Username
		u.Password = user.Password
		u.Role = user.Role

		return e.Success
	}

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
	u.Password = user.Password
	u.Role = user.Role

	if errInt == e.CacheNotFound {
		cache.SetUserInfo(strconv.Itoa(user.Id), user)
		cache.SetUserInfo(user.Username, user)
	}

	return e.Success
}

// get the user's articles
func (u *User) GetArticles() int { // Param:Id
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
	if unicode.IsDigit(rune(u.Username[0])) {
		return e.UsernameInvalid
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

// Verify password
func (u *User) Verify() (bool, int) {
	user := User{Username: u.Username}
	user.Get()

	err := bcrypt.CompareHashAndPassword(user.Password, u.Password)

	if err == nil {
		u.Id = user.Id
		u.Role = user.Role
		return true, e.Success
	} else {
		return false, e.Success
	}

}
