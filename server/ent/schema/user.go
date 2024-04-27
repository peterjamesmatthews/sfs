package schema

import (
	"regexp"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()).
			Unique().
			Default(uuid.New).
			Annotations(entsql.DefaultExprs(map[string]string{
				dialect.Postgres: "gen_random_uuid()",
			})),
		field.String("name").
			Unique().
			Match(regexp.MustCompile("^[a-zA-Z]+$")).
			MinLen(1).
			MaxLen(64),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
