package mem

import (
	"errors"
	"fmt"

	"pjm.dev/sfs/graph"
	"pjm.dev/sfs/graph/model"
)

func (m *Database) owns(user model.User, node model.Node) bool {
	return user.ID == node.GetOwner().ID
}

func (m *Database) has(user model.User, accessType model.AccessType, node model.Node) (bool, error) {
	_, err := m.getAccess(user, accessType, node)
	if err == nil {
		return true, nil
	}

	if !errors.Is(err, graph.ErrNotFound) {
		return false, fmt.Errorf("failed to check user %s has access %s to node %s: %w", user.ID, accessType.String(), node.GetID(), err)
	}

	if node.GetParent() == nil {
		return false, nil
	}

	return m.has(user, accessType, node.GetParent())
}

func (m *Database) getAccess(user model.User, accessType model.AccessType, node model.Node) (*model.Access, error) {
	for _, a := range m.Access {
		if a.User.ID != user.ID {
			continue
		}
		if a.Type != accessType {
			continue
		}
		if a.Target.GetID() != node.GetID() {
			continue
		}

		return a, nil
	}

	return nil, graph.ErrNotFound
}
