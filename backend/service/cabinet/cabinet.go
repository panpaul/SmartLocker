package cabinet

import (
	"SmartLocker/e"
	"SmartLocker/model"
	"github.com/go-playground/log"
)

type Cabinet struct {
	Id       int
	Name     string
	Location string
}

func GetLocations() ([]string, int) {
	l, err := model.GetCabinetLocations()
	if err != nil {
		log.WithError(err).Warn("Couldn't get locations")
		return nil, e.InternalError
	}
	return l, e.Success
}

func GetCabinets(where string) ([]Cabinet, int) {
	index, err := model.GetCabinetsByLocation(where)
	if err != nil {
		log.WithError(err).Warn("Couldn't get cabinets")
		return nil, e.InternalError
	}

	var r []Cabinet
	for i := range index {
		r = append(r, Cabinet{Id: index[i]})
	}
	return r, e.Success
}