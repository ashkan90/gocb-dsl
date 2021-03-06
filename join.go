package main

import (
	"bytes"
)

type JoinClause interface {
	KeysClauses
	As(alias string) KeysClauses
}

type defaultJoinClause struct {
	*defaultKeysClauses
}

func newDefaultJoinClause(parent Statement) *defaultJoinClause {
	return &defaultJoinClause{newDefaultKeysClauses(parent)}
}

func (c *defaultJoinClause) As(alias string) KeysClauses {
	c.setElement(&asElement{alias})
	return newDefaultKeysClauses(c)
}

type joinType string

const (
	defaultJoin joinType = ""
	inner       joinType = "INNER"
	left        joinType = "LEFT"
	leftOuter   joinType = "LEFT OUTER"
	right       joinType = "RIGHT"
	rightOuter  joinType = "RIGHT OUTER"
)

type joinElement struct {
	joinType joinType
	from     string
}

func (e *joinElement) export() string {
	buf := bytes.Buffer{}

	if e.joinType != defaultJoin {
		buf.WriteString(string(e.joinType))
		buf.WriteString(" ")
	}

	buf.WriteString("JOIN ")
	buf.WriteString(e.from)

	return buf.String()
}
