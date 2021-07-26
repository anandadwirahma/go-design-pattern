package cor

type Department interface {
	Execute(patient *Patient)
	SetNext(department Department)
}
