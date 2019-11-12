package photo

type Building struct {
	Photo
}

func (b *Building) List() string {
	return "building list"
}

func (b *Building) Detail() string {
	return "building detail"
}
