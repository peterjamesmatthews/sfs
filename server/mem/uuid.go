package mem

import "github.com/google/uuid"

func (m *Database) Generate() uuid.UUID {
	if len(m.UUIDs) == 0 {
		return uuid.New()
	}

	u := m.UUIDs[0]
	m.UUIDs = m.UUIDs[1:]
	return u
}
