package objects

type ModelObject interface {
	ToDTO() *DTObject
}

type DTObject interface {
}
