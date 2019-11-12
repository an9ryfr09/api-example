package photo

type Designer struct {
	Photo
}

func (d *Designer) List() string {
	return "designer list"
}

func (d *Designer) Detail() string {
	return "designer detail"
}
