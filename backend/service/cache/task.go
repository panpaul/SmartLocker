package cache

import (
	"encoding/json"
	"github.com/go-playground/log"
	"strconv"
)

func AddClientTask(cid string, task int) bool {
	key := "NodeTask_" + cid
	var tasks []int
	if exists(key) {
		v, _ := get(key)
		_ = json.Unmarshal(v, &tasks)
	}
	tasks = append(tasks, task)
	err := set(key, tasks, 0)
	if err != nil {
		log.WithError(err).Info("redis error")
		return false
	}
	return true

}

func DeleteClientTask(cid string) {
	key := "NodeTask_" + cid
	_, err := deleteKey(key)
	if err != nil {
		log.WithError(err).Info("redis error")
	}
}

func GetClientTask(cid string) []int {
	key := "NodeTask_" + cid
	var tasks []int
	if exists(key) {
		v, _ := get(key)
		_ = json.Unmarshal(v, &tasks)
		return tasks
	}
	return nil
}

func AddBackendTask(cid string, task int, user string) bool {
	key := "BackendTask_" + user
	var tasks []string
	if exists(key) {
		v, _ := get(key)
		_ = json.Unmarshal(v, &tasks)
	}
	tasks = append(tasks, strconv.Itoa(task)+"-"+cid)
	err := set(key, tasks, 0)
	if err != nil {
		log.WithError(err).Info("redis error")
		return false
	}
	return true
}

func DeleteBackendTask(user string) {
	key := "BackendTask_" + user
	_, err := deleteKey(key)
	if err != nil {
		log.WithError(err).Info("redis error")
	}
}

func GetBackendTask(user string) []string {
	key := "BackendTask_" + user
	var tasks []string
	if exists(key) {
		v, _ := get(key)
		_ = json.Unmarshal(v, &tasks)
		return tasks
	}
	return nil
}
