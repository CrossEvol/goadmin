package add_tags

import "database/sql"

type Category struct {
	ID       int           `db:"id" json:"id"`
	Name     string        `db:"name" json:"name"`
	ParentID sql.NullInt32 `db:"parent_id" json:"parent_id"`
}
