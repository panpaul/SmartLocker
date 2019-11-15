package cabinet

import (
	"SmartLocker/e"
	"SmartLocker/model"
	"SmartLocker/service/cache"
	"github.com/go-playground/log"
)

func GetLocations() ([]string, int) {
	l, errInt := cache.GetCabinetLocations()
	if errInt == e.Success {
		//log.Info("using cache")
		return l, e.Success
	}

	l, err := model.GetCabinetLocations()
	if err != nil {
		log.WithError(err).Warn("Couldn't get locations")
		return nil, e.InternalError
	}

	if errInt == e.CacheNotFound {
		//log.Info("write cache")
		cache.SetCabinetLocations(l)
	}
	return l, e.Success
}

func GetCabinets(where string) ([]*model.Cabinet, int) {
	c, errInt := cache.GetCabinetsByLocation(where)
	if errInt == e.Success {
		//log.Info("using cache")
		return c, e.Success
	}

	c, err := model.GetCabinetsByLocation(where)
	if err != nil {
		log.WithError(err).Warn("Couldn't get cabinets")
		return nil, e.InternalError
	}

	if errInt == e.CacheNotFound && len(c) != 0 {
		//log.Info("write cache")
		cache.SetCabinets(c, where)
	}

	return c, e.Success
}
