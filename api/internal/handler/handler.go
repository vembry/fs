package handler

import "api/pkg/db"

type handler struct {
	q db.Querier
}

func NewHandler(q db.Querier) *handler {
	return &handler{
		q: q,
	}
}
