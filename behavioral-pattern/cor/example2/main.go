package main

import (
	"fmt"
	"strings"
)

type StringHandler interface {
	SetNext(h StringHandler)
	Process(s string) string
}

type LowerCaseHandler struct {
	next StringHandler
}

func (l *LowerCaseHandler) SetNext(h StringHandler) {
	l.next = h
}

func (l *LowerCaseHandler) Process(s string) string {
	s = strings.ToLower(s)
	if l.next != nil {
		s = l.next.Process(s)
	}

	return s
}

type SpaceRemoval struct {
	next StringHandler
}

func (s *SpaceRemoval) SetNext(h StringHandler) {
	s.next = h
}

func (s *SpaceRemoval) Process(st string) string {
	st = strings.Replace(st, " ", "", -1)
	if s.next != nil {
		st = s.next.Process(st)
	}

	return st
}

func main() {
	sr := SpaceRemoval{}
	lc := LowerCaseHandler{}
	lc.SetNext(&sr)

	st := lc.Process("THE titanic")
	fmt.Println(st)
}
