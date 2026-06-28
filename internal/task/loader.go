package task

type Loader struct{}

func NewLoader() *Loader {
	return &Loader{}
}

func (l *Loader) Load(path string) (*Task, error) {
	return Load(path)
}
