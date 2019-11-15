package cache

import (
	"SmartLocker/e"
	"SmartLocker/model"
	"encoding/json"
	"github.com/go-playground/log"
)

func GetUserInfo(index string) (*model.User, int) {
	if exists("UserInfo_" + index) {
		l, err := get("UserInfo_" + index)
		if err != nil {
			log.WithError(err).Info("redis get error")
			return nil, e.RedisError
		}
		var r *model.User
		err = json.Unmarshal(l, &r)
		if err != nil {
			log.WithError(err).Info("couldn't unmarshal")
			return nil, e.JsonUnmarshalError
		}
		return r, e.Success
	}
	return nil, e.CacheNotFound
}

func SetUserInfo(index string, info *model.User) {
	err := set("UserInfo_"+index, info, 3600)
	if err != nil {
		log.WithError(err).Info("couldn't set cache")
	}
}
