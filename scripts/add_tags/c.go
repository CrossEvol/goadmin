package add_tags

import "time"

type TodoTag struct {
	ID        int       `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	TodoID    int       `db:"todo_id" json:"todo_id"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}
