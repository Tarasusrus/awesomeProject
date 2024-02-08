package pattern

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

type State interface {
	open() error
	close() error
	save() error
}

type Document struct {
	state State
}

func (d *Document) SetState(s State) {
	d.state = s
}

func (d *Document) Open() error {
	return d.state.open()
}

func (d *Document) Close() error {
	return d.state.close()
}

func (d *Document) Save() error {
	return d.state.save()
}

type StateOpened struct{}

func (s *StateOpened) open() error {
	// implementation for opened state
}

func (s *StateOpened) close() error {
	// implementation for opened state
}

func (s *StateOpened) save() error {
	// implementation for opened state
}
