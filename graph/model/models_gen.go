// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type Mutation struct {
}

type NewStaffInput struct {
	FirstName  string  `json:"firstName"`
	LastName   string  `json:"lastName"`
	Email      string  `json:"email"`
	WorkNumber *string `json:"workNumber,omitempty"`
	Role       string  `json:"role"`
	Department string  `json:"department"`
	Salary     float64 `json:"salary"`
}

type Pagination struct {
	Page  *int `json:"page,omitempty"`
	Limit *int `json:"limit,omitempty"`
}

type Project struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	Department  string  `json:"department"`
	Budget      float64 `json:"budget"`
}

type Query struct {
}

type Staff struct {
	ID         string     `json:"id"`
	FirstName  string     `json:"firstName"`
	LastName   string     `json:"lastName"`
	Email      string     `json:"email"`
	WorkNumber *string    `json:"workNumber,omitempty"`
	Role       string     `json:"role"`
	Department string     `json:"department"`
	Salary     float64    `json:"salary"`
	Projects   []*Project `json:"projects,omitempty"`
}

type StaffFilter struct {
	Role       *string  `json:"role,omitempty"`
	Department *string  `json:"department,omitempty"`
	SalaryMin  *float64 `json:"salaryMin,omitempty"`
	SalaryMax  *float64 `json:"salaryMax,omitempty"`
}

type Order string

const (
	OrderAsc  Order = "ASC"
	OrderDesc Order = "DESC"
)

var AllOrder = []Order{
	OrderAsc,
	OrderDesc,
}

func (e Order) IsValid() bool {
	switch e {
	case OrderAsc, OrderDesc:
		return true
	}
	return false
}

func (e Order) String() string {
	return string(e)
}

func (e *Order) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Order(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Order", str)
	}
	return nil
}

func (e Order) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type SortBy string

const (
	SortByFirstName SortBy = "FIRST_NAME"
	SortByLastName  SortBy = "LAST_NAME"
	SortByRole      SortBy = "ROLE"
	SortBySalary    SortBy = "SALARY"
)

var AllSortBy = []SortBy{
	SortByFirstName,
	SortByLastName,
	SortByRole,
	SortBySalary,
}

func (e SortBy) IsValid() bool {
	switch e {
	case SortByFirstName, SortByLastName, SortByRole, SortBySalary:
		return true
	}
	return false
}

func (e SortBy) String() string {
	return string(e)
}

func (e *SortBy) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SortBy(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SortBy", str)
	}
	return nil
}

func (e SortBy) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
