package __test__

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/crossevol/goadmin/internal/utils"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
	"time"
)

type User struct {
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
}

func TestNull(t *testing.T) {
	var user User
	err := json.Unmarshal([]byte("{}"), &user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(utils.ToJson(user))

	err = json.Unmarshal([]byte(`{"name":"test","age":18}`), &user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(utils.ToJson(user))
}

type NullString struct {
	ID sql.NullInt32 `json:"id"`
}

func TestSqlNull(t *testing.T) {
	nullString := NullString{ID: sql.NullInt32{Int32: 18, Valid: true}}
	fmt.Println(utils.ToJson(nullString))

	err := json.Unmarshal([]byte(`{"id":18}`), &nullString)
	require.NotNil(t, err)
}

func TestNullTime(t *testing.T) {
	type D struct {
		Birth time.Time `json:"birth,omitempty"`
	}
	d := D{}
	fmt.Println(utils.ToJson(d))
}
