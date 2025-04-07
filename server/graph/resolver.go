package graph

import (
	"github.com/Pawan2061/timeline_grpc_go/server/store"
)

type Resolver struct {
	store *store.Store
}

func NewResolver(store *store.Store) *Resolver {
	return &Resolver{store: store}
}
