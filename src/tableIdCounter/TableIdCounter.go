package tableIdCounter

type TableIdCounter struct {
	id int
}

func New() *TableIdCounter {
	return &TableIdCounter{
		id: 0,
	}
}

func (counter *TableIdCounter) getIdRef() *int {
	return &counter.id
}

func (counter *TableIdCounter) Get() int {
	idRef := counter.getIdRef()
	defer func() {
		*idRef++
	}()

	return *idRef
}
