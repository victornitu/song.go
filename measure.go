package song

import (
	"fmt"
	"github.com/hajimehoshi/oto"
	"sync"
)

type Measure struct {
	Duration int64
	Notes    []*Note
	options  *options
	context  *oto.Context
	group    *sync.WaitGroup
}

func (m *Measure) play() {
	m.group = &sync.WaitGroup{}
	for _, note := range m.Notes {
		note.timeRef = float64(m.Duration)
		note.options = m.options
		note.player = m.context.NewPlayer()
		m.group.Add(1)
		go note.play(m.group)
	}
	m.group.Wait()
	fmt.Println("Played measure")
	return
}
