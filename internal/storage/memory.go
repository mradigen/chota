package storage

type MemoryStorage struct {
	data map[string]string
}

func NewMemory() *MemoryStorage {
	return &MemoryStorage{data: make(map[string]string)}
}

func (m* MemoryStorage) Save(slug string, u string) (string, error) {
	if _, ok := m.data[slug]; ok {
		return "", ErrExists
	}

	m.data[slug] = u
	return slug, nil
}

func (m* MemoryStorage) Get(slug string) (string, error) {
	if u, ok := m.data[slug]; ok {
		return u, nil
	}

	return "", ErrNotFound
}
