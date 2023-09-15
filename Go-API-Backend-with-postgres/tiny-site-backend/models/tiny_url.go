package models

import (
	"time"

	"github.com/uptrace/bun"
)

type Tinyurl struct {
	bun.BaseModel `bun:"table:tiny_url"`
	Org_url       string
	Short_url     string
	Comment       string
	User_id       string
	Valid_up      time.Time
	Created_at    time.Time
	Created_by    string
}
