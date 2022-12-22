// Code generated by github.com/v1def/go-polybase. DO NOT EDIT.

package generated

import (
	"context"

	"github.com/v1def/go-polybase"
)

type City struct {
	Id      string  `json:"id"`
	Name    string  `json:"name"`
	Country *string `json:"country,omitempty"`
}
type ICity interface {
	Constructor(ctx context.Context, input *CityConstructorInput) (*polybase.SingleResponse[City], error)
	SetCountry(ctx context.Context, id string, input *CitySetCountryInput) (*polybase.SingleResponse[City], error)
	UpdateName(ctx context.Context, id string, input *CityUpdateNameInput) (*polybase.SingleResponse[City], error)
}

type city struct{ coll polybase.Collection }

func NewCity(db polybase.Polybase) ICity {
	return &city{coll: db.Collection("V-City/City")}
}

type CityConstructorInput struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func (c *city) Constructor(ctx context.Context, input *CityConstructorInput) (*polybase.SingleResponse[City], error) {
	var response polybase.SingleResponse[City]
	return &response, c.coll.Create(ctx, polybase.ParseInput(input), &response)
}

type CitySetCountryInput struct {
	Country string `json:"country"`
}

func (c *city) SetCountry(ctx context.Context, id string, input *CitySetCountryInput) (*polybase.SingleResponse[City], error) {
	var response polybase.SingleResponse[City]
	return &response, c.coll.Record(id).Call(ctx, "setCountry", polybase.ParseInput(input), &response)
}

type CityUpdateNameInput struct {
	Name string `json:"name"`
}

func (c *city) UpdateName(ctx context.Context, id string, input *CityUpdateNameInput) (*polybase.SingleResponse[City], error) {
	var response polybase.SingleResponse[City]
	return &response, c.coll.Record(id).Call(ctx, "updateName", polybase.ParseInput(input), &response)
}