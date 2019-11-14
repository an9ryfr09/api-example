package model

type Modeler interface {
	List() string
	Detail() string
}

type Model struct{}

func (model Model) List(m Modeler) string {
	return m.List()
}

func (model *Model) Detail(m Modeler) string {
	return m.Detail()
}
