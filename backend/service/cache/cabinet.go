package cache

import (
	"SmartLocker/e"
	"encoding/json"
	"github.com/go-playground/log"
)

func GetCabinetLocations() ([]string, int) {
	if exists("CabinetLocations") {
		l, err := get("CabinetLocations")
		if err != nil {
			log.WithError(err).Info("redis get error")
			return nil, e.RedisError
		}
		var r []string
		err = json.Unmarshal(l, &r)
		if err != nil {
			log.WithError(err).Info("couldn't unmarshal")
			return nil, e.JsonUnmarshalError
		}
		return r, e.Success
	}
	return nil, e.CacheNotFound
}

func SetCabinetLocations(s []string) {
	err := set("CabinetLocations", s, 3600)
	if err != nil {
		log.WithError(err).Info("couldn't set cache")
	}
}
