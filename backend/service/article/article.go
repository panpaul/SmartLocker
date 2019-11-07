package article

import (
	"SmartLocker/e"
	"SmartLocker/model"
	"github.com/go-playground/log"
)

type Article struct {
	Id              int    // Locker's id
	Position        int    // Locker's relative position
	UserId          int    // User's id
	CabinetId       int    // Cabinet's id
	CabinetLocation string //Cabinet's location
	CabinetName     string // Cabinet's name
}

func GetArticles(id int) ([]Article, int) {
	a, err := model.GetLockersByUid(id)
	if err != nil {
		log.WithError(err).Warn("Couldn't get lockers")
		return nil, e.InternalError
	}
	var result []Article
	for i := range a {
		r := Article{
			Id:              a[i].Id,
			Position:        a[i].Position,
			UserId:          a[i].Uid,
			CabinetId:       a[i].Cid,
			CabinetLocation: a[i].CabinetInfo.Location,
			CabinetName:     a[i].CabinetInfo.Name,
		}
		result = append(result, r)
	}
	return result, e.Success
}

func (a *Article) RandomOccupy() int { // Param:CabinetId
	lockers, err := model.GetFreeLockers(a.CabinetId)
	if err != nil {
		log.WithError(err).Warn("Couldn't get free lockers")
		return e.InternalError
	}
	if len(lockers) == 0 {
		return e.NoMoreLocker
	}

	a.Id = lockers[0]
	return a.Update(false)
}

func (a *Article) Update(release bool) int { // Param:Id(UserId)
	if a.Id == 0 {
		return e.InvalidParams
	}

	var err error
	if release {
		err = model.ReleaseLockerById(a.Id)
	} else {
		err = model.OccupyLockerById(a.Id, a.UserId)
	}

	if err != nil {
		log.WithError(err).Warn("Couldn't update locker's info")
		return e.InternalError
	}
	return e.Success
}
