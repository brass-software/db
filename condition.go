package db

import "fmt"

type Condition struct {
	Op  string
	LHS string
	RHS string
}

func (c *Condition) String() string {
	return fmt.Sprintf("%s %s %s", c.LHS, c.Op, c.RHS)
}
