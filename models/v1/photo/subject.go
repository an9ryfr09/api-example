package photo

type Subject struct {
	Photo
}

func (s *Subject) List() string {
	return "subject list"
}

func (s *Subject) Detail() string {
	return "subject list"
}
