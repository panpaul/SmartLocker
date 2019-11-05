package article

import (
	"SmartLocker/e"
	"SmartLocker/model"
	"github.com/go-playground/log"
)

type Article struct {
	Id              int    // Locker's id
	Position        int    // Locker's relative position
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
	for _, v := range a {
		r := Article{
			Id:              v.Id,
			Position:        v.Position,
			CabinetId:       v.Cid,
			CabinetLocation: v.CabinetInfo.Location,
			CabinetName:     v.CabinetInfo.Name,
		}
		result = append(result, r)
	}
	return result, e.Success
}
