package cache

import (
	"encoding/json"
	"github.com/go-playground/log"
)

func RegisterNode(cid string, name string) bool {
	key := "Node_" + cid
	if exists(key) {
		_, _ = deleteKey(key)
	}
	err := set(key, name, 0)
	if err != nil {
		log.WithError(err).Info("redis error")
		return false
	}
	return true
}

func GenerateToken(name string, token string) bool {
	key := "RegToken_" + name
	if exists(key) {
		_, _ = deleteKey(key)
	}
	err := set(key, token, 0)
	if err != nil {
		log.WithError(err).Info("redis error")
		return false
	}
	return true
}

func CheckToken(name string, token string) bool {
	key := "RegToken_" + name
	if !exists(key) {
		return false
	}

	v, _ := get(key)
	var r string
	_ = json.Unmarshal(v, &r)
	if r != token {
		return false
	}

	return true
}

func NodePingPong(cid string) (string, bool) {
	key := "Node_" + cid
	if exists(key) {
		v, _ := get(key)
		var name string
		_ = json.Unmarshal(v, &name)
		return name, true
	}
	return "", false
}
