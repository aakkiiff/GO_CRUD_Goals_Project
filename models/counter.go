package models

import "github.com/kamva/mgm/v3"

type Counter struct {
	mgm.DefaultModel `bson:"inline"`
	Count            int64 `json:"count" bson:"count"`
}

func NewCounter(count int64) *Counter {
	return &Counter{
		Count: count,
	}
}
