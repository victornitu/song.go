package song

import (
	"fmt"
	"github.com/hajimehoshi/oto"
	"io"
	"sync"
	"time"
)

type Pitch float64
type Duration float64
type Dotted float64

const (
	La   Pitch = 440.0
	Sib  Pitch = 466.16
	Si   Pitch = 493.88
	Do   Pitch = 523.25
	Reb  Pitch = 554.37
	Re   Pitch = 587.33
	Mib  Pitch = 622.25
	Mi   Pitch = 659.26
	Fa   Pitch = 698.46
	Solb Pitch = 739.99
	Sol  Pitch = 783.99
	Lab  Pitch = 830.61
)

const (
	Silent    Duration = 0.0
	Whole     Duration = 1.0
	Half      Duration = 0.5
	Quarter   Duration = 0.25
	Eighth    Duration = 0.125
	Sixteenth Duration = 0.0625
)

const (
	NotDotted     Dotted = 1.0
	SingleDotted  Dotted = 1.5
	DoubleDotted  Dotted = 1.75
	TrippleDotted Dotted = 1.875
)

type Note struct {
	Pitch        Pitch
	Duration     Duration
	Offset       Duration
	Dotted       Dotted
	DottedOffset Dotted
	timeRef      float64
	options      *options
	player       *oto.Player
}

func (n *Note) play(group *sync.WaitGroup) {
	if n.Offset != Silent {
		time.Sleep(n.calcOffset())
	}
	fmt.Printf("Play: %v\n", n.Pitch)
	s := n.sineWave()
	if _, err := io.Copy(n.player, s); err != nil {
		panic(err)
	}
	time.Sleep(n.calcDuration())
	if err := n.player.Close(); err != nil {
		panic(err)
	}
	group.Done()
}

func (n *Note) sineWave() *sineWave {
	length := n.options.calculate(n.calcDuration())
	length = length / 4 * 4
	return &sineWave{
		freq:    n.Pitch,
		length:  length,
		options: n.options,
	}
}

func (n *Note) calcOffset() time.Duration {
	if n.DottedOffset == 0 {
		n.DottedOffset = NotDotted
	}
	offset := float64(n.Offset) * float64(n.DottedOffset)
	milliSec := offset * n.timeRef
	return time.Duration(milliSec*1000) * time.Millisecond
}

func (n *Note) calcDuration() time.Duration {
	if n.Dotted == 0 {
		n.Dotted = NotDotted
	}
	duration := float64(n.Duration) * float64(n.Dotted)
	milliSec := duration * n.timeRef
	return time.Duration(milliSec*1000) * time.Millisecond
}
