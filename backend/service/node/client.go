package node

import (
	"SmartLocker/model"
	"SmartLocker/service/cache"
	"github.com/go-playground/log"
	"strconv"
)

func RegisterCabinet(name string, location string, lockerNum int) (int, bool) {
	err := model.AddCabinet(name, location)
	if err != nil {
		log.WithError(err).Info("couldn't add cabinet")
		return -1, false
	}

	c, err := model.GetCabinetByName(name)
	if err != nil {
		log.WithError(err).Info("couldn't get cabinet")
		return -1, false
	}

	cid := strconv.Itoa(c.Id)
	cache.RegisterNode(cid, c.Name)

	// TODO add lockers

	return c.Id, true
}
