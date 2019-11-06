package main

import (
	"SmartLocker/config"
	"SmartLocker/logger"
	"SmartLocker/model"
	"github.com/go-playground/log"
	"runtime"
	"strconv"
)

func main() {
	// init the helpers
	logger.Setup()
	config.Setup()
	model.Setup()

	var c []*model.Cabinet
	for i := 0; i < 50; i++ {
		c = append(c, &model.Cabinet{Name: "C" + strconv.Itoa(i), Location: "L" + strconv.Itoa(int(i/4))})
	}
	err := model.AddCabinets(c)
	if err != nil {
		log.WithError(err)
		err = nil
	}

	runtime.GC()

	var u []*model.User
	for i := 0; i < 100000; i++ {
		u = append(u, &model.User{Username: "test" + strconv.Itoa(i), Password: "test", Role: 0, Connect: ""})
	}
	err = model.AddUsers(u)
	if err != nil {
		log.WithError(err)
		err = nil
	}

	runtime.GC()

	var cid []int
	for i := 0; i < 14; i++ {
		j, err := model.GetCabinetsByLocation("L" + strconv.Itoa(i))
		if err != nil {
			log.WithError(err)
			err = nil
		} else {
			for k := range j {
				cid = append(cid, j[k])
			}
		}
	}
	var l []*model.Locker
	for i := 0; i < len(cid); i++ {
		for j := 0; j < 1000; j++ {
			l = append(l, &model.Locker{
				Position:     j,
				Availability: true,
				Cid:          cid[i],
			})
		}

		log.Info(i)
		go tx(l)
		l = nil
		runtime.GC()
	}

	runtime.GC()
}

func tx(l []*model.Locker) {
	err := model.AddLockers(l)
	if err != nil {
		log.WithError(err)
	}
}
