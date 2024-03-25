package easyget

import (
	"fmt"
	"sync"
)

var dfmgr = NewKVManager()

func NewKVManager() *KVManager {
	return &KVManager{
		m: map[string]Getter{},
	}
}

type KVManager struct {
	mu sync.Mutex
	m  map[string]Getter
}

func (kvm *KVManager) SetGetter(key string, g Getter) {
	kvm.mu.Lock()
	defer kvm.mu.Unlock()

	kvm.m[key] = g
}

func (kvm *KVManager) SetDefault(key string, v string) {
	kvm.mu.Lock()
	defer kvm.mu.Unlock()

	g, ok := kvm.m[key]
	if !ok {
		kvm.m[key] = NewStaticGetter(v)
		return
	}

	kvm.m[key] = WithDefault(g, v)
}

func (kvm *KVManager) Get(key string) (string, bool) {
	kvm.mu.Lock()
	defer kvm.mu.Unlock()

	g, ok := kvm.m[key]
	if !ok {
		return "", false
	}

	return g.Get(key)
}
func (kvm *KVManager) MustGet(key string) string {
	v, ok := kvm.Get(key)
	if !ok {
		panic(fmt.Errorf("[easyget] MustGet %s fail", key))
	}

	return v
}

func SetGetter(key string, g Getter) {
	dfmgr.SetGetter(key, g)
}

func SetDefault(key string, v string) {
	dfmgr.SetDefault(key, v)
}

func Get(key string) (string, bool) {
	return dfmgr.Get(key)
}

func MustGet(key string) string {
	return dfmgr.MustGet(key)
}
