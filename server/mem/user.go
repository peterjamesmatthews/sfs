package mem

import "pjm.dev/sfs/graph"

func (m *Database) GetUserByID(id string) (graph.User, error) {
	users := m.Users

	for _, user := range users {
		if user.ID == id {
			return *user, nil
		}
	}

	return graph.User{}, graph.ErrNotFound
}
