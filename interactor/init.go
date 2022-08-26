package interactor

type Interactor struct {
	store interface{}
}

func New(store interface{}) Interactor {

	return Interactor{store: store}
}
