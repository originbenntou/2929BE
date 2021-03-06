// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type ChildSuggest struct {
	Word   string   `json:"word"`
	Growth *Growth  `json:"growth"`
	Graphs []*Graph `json:"graphs"`
}

type Graph struct {
	Date  string  `json:"date"`
	Value float64 `json:"value"`
}

type Growth struct {
	Short  Arrow `json:"short"`
	Midium Arrow `json:"midium"`
	Long   Arrow `json:"long"`
}

type History struct {
	SuggestID int      `json:"suggestId"`
	Status    Progress `json:"status"`
}

type Suggest struct {
	Keyword       string          `json:"keyword"`
	ChildSuggests []*ChildSuggest `json:"childSuggests"`
}

type Arrow string

const (
	ArrowUp   Arrow = "Up"
	ArrowFlat Arrow = "Flat"
	ArrowDown Arrow = "Down"
)

var AllArrow = []Arrow{
	ArrowUp,
	ArrowFlat,
	ArrowDown,
}

func (e Arrow) IsValid() bool {
	switch e {
	case ArrowUp, ArrowFlat, ArrowDown:
		return true
	}
	return false
}

func (e Arrow) String() string {
	return string(e)
}

func (e *Arrow) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Arrow(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Arrow", str)
	}
	return nil
}

func (e Arrow) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Progress string

const (
	ProgressInprogress Progress = "INPROGRESS"
	ProgressCompleted  Progress = "COMPLETED"
)

var AllProgress = []Progress{
	ProgressInprogress,
	ProgressCompleted,
}

func (e Progress) IsValid() bool {
	switch e {
	case ProgressInprogress, ProgressCompleted:
		return true
	}
	return false
}

func (e Progress) String() string {
	return string(e)
}

func (e *Progress) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Progress(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Progress", str)
	}
	return nil
}

func (e Progress) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
