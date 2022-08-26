package interactor

import "go-ecommerce/contract"

type Interactor struct {
	store contract.Repository
}

func New(store contract.Repository) Interactor {

	return Interactor{store: store}
}
