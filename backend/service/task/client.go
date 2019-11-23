package task

import "SmartLocker/service/cache"

func AddClientTask(cid string, taskType int, work int) {
	cache.AddClientTask(cid, work)
}

func ConsumeClientTask(cid string) {
	cache.DeleteClientTask(cid)
}

func GetClientTask(cid string) (bool, []int) {
	t := cache.GetClientTask(cid)
	if t == nil {
		return false, nil
	} else if len(t) == 0 {
		return false, nil
	}
	return true, t
}
