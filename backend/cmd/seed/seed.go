package main

import (
	"SmartLocker/config"
	"SmartLocker/logger"
	"SmartLocker/model"
	"SmartLocker/service/article"
	"github.com/go-playground/log"
	"math/rand"
	"runtime"
	"strconv"
	"time"
)

func main() {
	// init the helpers
	logger.Setup()
	config.Setup("config.yaml")
	model.Setup()

	// check if data exist
	cl, err := model.GetCabinetLocations()
	if err != nil {
		log.WithError(err)
		return
	}
	if len(cl) != 0 {
		log.Info("Data exist")
		return
	}

	// fill cabinet
	var c []*model.Cabinet
	for i := 0; i < 50; i++ {
		c = append(c, &model.Cabinet{Name: "C" + strconv.Itoa(i), Location: "L" + strconv.Itoa(int(i/4))})
	}
	err = model.AddCabinets(c)
	if err != nil {
		log.WithError(err)
		err = nil
	}

	runtime.GC()

	// fill user
	var u []*model.User
	for i := 0; i < 10000; i++ {
		u = append(u, &model.User{Username: "test" + strconv.Itoa(i), Password: "test", Role: 0, Connect: ""})
	}
	err = model.AddUsers(u)
	if err != nil {
		log.WithError(err)
		err = nil
	}

	runtime.GC()

	// fill lockers

	// get cabinets' ids
	var cid []int
	for i := 0; i < 13; i++ {
		j, err := model.GetCabinetIdsByLocation("L" + strconv.Itoa(i))
		if err != nil {
			log.WithError(err)
			err = nil
		} else {
			for k := range j {
				cid = append(cid, j[k])
			}
		}
	}

	log.Info(cid)

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
		err := model.AddLockers(l)
		if err != nil {
			log.WithError(err)
		}
		l = nil
		runtime.GC()
	}

	runtime.GC()

	// Randomly make occupations
	ri := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))
	for i := 0; i < 100; i++ { //随机找100个人
		art := article.Article{UserId: ri.Intn(10000)}
		for j := 0; j < 5; j++ { //随机找5个柜子
			art.CabinetId = cid[ri.Intn(len(cid))]
			for k := 0; k < 5; k++ { // 每个柜子申请5格
				art.RandomOccupy()
			}
		}
	}

}
