package schema

import (
	"time"

	"hiyoko-fiber/internal/pkg/ent/util"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Mixin of the User.
// updateやcreateをmixinするとカラムの順序を制御出来ない
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			MinLen(4).
			MaxLen(255).
			Unique(),
		field.String("sub").
			GoType(util.ULID("")).
			DefaultFunc(func() util.ULID {
				return util.NewULID()
			}).
			Immutable().
			Unique(),
		field.String("email").
			MinLen(4).
			MaxLen(255).
			NotEmpty(),
		field.String("password").Sensitive().NotEmpty(),
		field.Time("created_at").Immutable().Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{}
}
