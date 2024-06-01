package add_tags

type Group struct {
	ID   int    `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
	Desc string `db:"desc" json:"desc"`
}
