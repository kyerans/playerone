package repos

import (
	"encoding/base64"
	"encoding/hex"
	"log"
	"sync"
)

func New() *Repository {
	return &Repository{
		mapp: make(map[string]string),
	}
}

type Repository struct {
	lock sync.Mutex
	mapp map[string]string
}

func (r *Repository) Get(k string) string {
	r.lock.Lock()
	defer r.lock.Unlock()

	resp, ok := r.mapp[k]
	if ok {
		return resp
	}

	return ""
}

func (r *Repository) Set(k, v string) {
	r.lock.Lock()
	defer r.lock.Unlock()

	bytes, err := hex.DecodeString(k)
	if err != nil {
		return
	}
	b64k := base64.RawURLEncoding.EncodeToString(bytes)

	bytes, err = hex.DecodeString(v)
	if err != nil {
		return
	}
	b64v := base64.RawURLEncoding.EncodeToString(bytes)

	log.Printf("[repo] save key=%v val=%v", b64k, b64v)
	r.mapp[b64k] = b64v
}
