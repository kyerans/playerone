package repos

import "sync"

func New() *Repository {
	return &Repository{
		mapp: make(map[string][]byte),
	}
}

type Repository struct {
	lock sync.Mutex
	mapp map[string][]byte
}

func (r *Repository) Get(k string) []byte {
	r.lock.Lock()
	defer r.lock.Unlock()

	resp, ok := r.mapp[k]
	if ok {
		return resp
	}

	return nil
}

func (r *Repository) Set(k, v string) {
	r.lock.Lock()
	defer r.lock.Unlock()

	r.mapp[k] = []byte(v)
}
