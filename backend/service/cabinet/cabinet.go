package cabinet

import (
	"SmartLocker/e"
	"SmartLocker/model"
	"github.com/go-playground/log"
)

func GetLocations() ([]string, int) {
	l, err := model.GetCabinetLocations()
	if err != nil {
		log.WithError(err).Warn("Couldn't get locations")
		return nil, e.InternalError
	}
	return l, e.Success
}

func GetCabinets(where string) ([]*model.Cabinet, int) {
	c, err := model.GetCabinetsByLocation(where)
	if err != nil {
		log.WithError(err).Warn("Couldn't get cabinets")
		return nil, e.InternalError
	}

	return c, e.Success
}
