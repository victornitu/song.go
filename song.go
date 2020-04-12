package song

import (
	"github.com/hajimehoshi/oto"
	"time"
)

type song struct {
	Measures []*Measure
	options  *options
	context  *oto.Context
}

type options struct {
	sampleRate      int
	channelNum      int
	bitDepthInBytes int
}

func (o *options) calculate(duration time.Duration) int64 {
	return int64(o.channelNum) * int64(o.bitDepthInBytes) * int64(o.sampleRate) * int64(duration) / int64(time.Second)
}

func New(sampleRate int, channelNum int, bitDepthInBytes int) (m *song, err error) {
	ctx, err := oto.NewContext(sampleRate, channelNum, bitDepthInBytes, 4096)
	if err != nil {
		return
	}
	m = &song{
		options: &options{
			sampleRate:      sampleRate,
			channelNum:      channelNum,
			bitDepthInBytes: bitDepthInBytes,
		},
		context: ctx,
	}
	return
}

func (s *song) Play() {
	for _, measure := range s.Measures {
		measure.context = s.context
		measure.options = s.options
		measure.play()
	}
}
