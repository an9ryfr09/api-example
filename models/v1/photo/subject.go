package photo

type Subject struct {
	Photo
}

func (s *Subject) List() string {
	return "abc"
}

func (s *Subject) Detail() string {
	return "bcd"
}
