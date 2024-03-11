package memdb

import "github.com/google/uuid"

func (m *MemDatabase) Generate() uuid.UUID {
	if len(m.UUIDs) == 0 {
		return uuid.New()
	}

	u := m.UUIDs[0]
	m.UUIDs = m.UUIDs[1:]
	return u
}
