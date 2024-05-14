package app

import "github.com/google/uuid"

func (a *App) Generate() uuid.UUID {
	var u uuid.UUID

	if len(a.uuids) > 0 {
		u = a.uuids[0]
		a.uuids = a.uuids[1:]
	} else {
		u = uuid.New()
	}

	return u
}
