package db

import (
	"strings"

	"github.com/mikerybka/schema.cafe/pkg/util"
)

type Where struct {
	Conditions []*Condition
}

func (w *Where) String() string {
	if len(w.Conditions) == 0 {
		return ""
	}
	return " WHERE " + strings.Join(
		util.Map(
			w.Conditions,
			func(c *Condition) string {
				return c.String()
			},
		),
		" AND ",
	)
}
