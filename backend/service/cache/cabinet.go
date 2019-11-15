package cache

import (
	"SmartLocker/e"
	"SmartLocker/model"
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

func GetCabinetsByLocation(where string) ([]*model.Cabinet, int) {
	if exists("Cabinet_" + where) {
		l, err := get("Cabinet_" + where)
		if err != nil {
			log.WithError(err).Info("redis get error")
			return nil, e.RedisError
		}
		var r []*model.Cabinet
		err = json.Unmarshal(l, &r)
		if err != nil {
			log.WithError(err).Info("couldn't unmarshal")
			return nil, e.JsonUnmarshalError
		}
		return r, e.Success
	}
	return nil, e.CacheNotFound
}

func SetCabinets(s []*model.Cabinet, where string) {
	err := set("Cabinet_"+where, s, 3600)
	if err != nil {
		log.WithError(err).Info("couldn't set cache")
	}
}
