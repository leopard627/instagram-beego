package models

import (
	"time"
)

type Photo struct {
	Id       int
	Url      string
	Created  time.Time  `orm:"auto_now_add;type(datetime)"`
	User     *User      `orm:"rel(fk)"`
	Hashtag  *Hashtag   `orm:"rel(fk);null"`
	Comments []*Comment `orm:"reverse(many)"`
}
