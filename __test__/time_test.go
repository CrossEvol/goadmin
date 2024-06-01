package __test__

import (
	"encoding/json"
	"fmt"
	"github.com/crossevol/goadmin/internal/utils"
	"log"
	"testing"
	"time"
)

type Date struct {
	Birth time.Time
}

func TestName(t *testing.T) {
	date := Date{Birth: time.Now()}
	dateJson := utils.ToJson(date)
	fmt.Println(dateJson)

	var d2 Date
	const json2 = `{"Birth":"2024-06-01T16:00:00Z"}`
	err := json.Unmarshal([]byte(json2), &d2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(d2)
}
