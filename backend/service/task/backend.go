package task

import (
	"SmartLocker/model"
	"SmartLocker/service/cache"
)

func AddBackendTask(cid string, taskType int, work int, uid int) bool {
	u, err := model.GetUserInfoById(uid)
	if err != nil {
		return false
	}
	cache.AddBackendTask(cid, work, u.Username)
	return true
}

func ConsumeBackendTask(u string) bool {
	cache.DeleteBackendTask(u)
	return true
}

func GetBackendTask(u string) ([]string, bool) {
	t := cache.GetBackendTask(u)
	if t == nil {
		return nil, false
	} else if len(t) == 0 {
		return nil, false
	}
	return t, true
}
