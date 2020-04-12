package main

import (
	"flag"
	"fmt"
	"log"
	"song"
	"time"
)

var (
	sampleRate      = flag.Int("samplerate", 44100, "sample rate")
	channelNum      = flag.Int("channelnum", 2, "number of channel")
	bitDepthInBytes = flag.Int("bitdepthinbytes", 2, "bit depth in bytes")
)

func main() {
	flag.Parse()
	s, err := song.New(*sampleRate, *channelNum, *bitDepthInBytes)
	if err != nil {
		log.Fatalf("new song failed: %v", err)
	}
	s.Measures = []*song.Measure{
		{	// Measure #1
			Duration: 4,
			Notes: []*song.Note{
				{
					Pitch:    song.Do,
					Duration: song.Quarter,
				},
				{
					Pitch:    song.Sol,
					Duration: song.Eighth,
					Offset:   song.Eighth,
				},
				{
					Pitch:    song.Fa,
					Duration: song.Eighth,
					Offset:   song.Quarter,
				},
				{
					Pitch:        song.Mi,
					Duration:     song.Eighth,
					Dotted:       song.SingleDotted,
					Offset:       song.Quarter,
					DottedOffset: song.SingleDotted,
				},
				{
					Pitch:        song.Do,
					Duration:     song.Sixteenth,
					Offset:       song.Half,
					DottedOffset: song.DoubleDotted,
				},
				{
					Pitch:        song.Re,
					Duration:     song.Quarter,
					Offset:       song.Half,
					DottedOffset: song.SingleDotted,
				},
			},
		},
		{	// Measure #2
			Duration: 4,
			Notes: []*song.Note{
				{
					Pitch:    song.Do,
					Duration: song.Quarter,
				},
				{
					Pitch:    song.Sol,
					Duration: song.Eighth,
					Offset:   song.Eighth,
				},
				{
					Pitch:    song.Fa,
					Duration: song.Eighth,
					Offset:   song.Quarter,
				},
				{
					Pitch:        song.La,
					Duration:     song.Quarter,
					Dotted:       song.SingleDotted,
					Offset:       song.Quarter,
					DottedOffset: song.SingleDotted,
				},
				{
					Pitch:        song.Do,
					Duration:     song.Sixteenth,
					Offset:       song.Half,
					DottedOffset: song.DoubleDotted,
				},
				{
					Pitch:        song.Re,
					Duration:     song.Quarter,
					Offset:       song.Half,
					DottedOffset: song.SingleDotted,
				},
			},
		},
		{	// Measure #3
			Duration: 4,
			Notes: []*song.Note{
				{
					Pitch:    song.Do,
					Duration: song.Quarter,
				},
				{
					Pitch:    song.Sol,
					Duration: song.Eighth,
					Offset:   song.Eighth,
				},
				{
					Pitch:    2 * song.La,
					Duration: song.Eighth,
					Offset:   song.Quarter,
				},
				{
					Pitch:        song.Mi,
					Duration:     song.Eighth,
					Dotted:       song.SingleDotted,
					Offset:       song.Quarter,
					DottedOffset: song.SingleDotted,
				},
				{
					Pitch:        song.Do,
					Duration:     song.Sixteenth,
					Offset:       song.Half,
					DottedOffset: song.DoubleDotted,
				},
				{
					Pitch:        song.Re,
					Duration:     song.Quarter,
					Offset:       song.Half,
					DottedOffset: song.SingleDotted,
				},
			},
		},
		{	// Measure #4
			Duration: 4,
			Notes: []*song.Note{
				{
					Pitch:    song.Do,
					Duration: song.Quarter,
				},
				{
					Pitch:    song.Sol,
					Duration: song.Eighth,
					Offset:   song.Eighth,
				},
				{
					Pitch:    song.Mi,
					Duration: song.Eighth,
					Offset:   song.Quarter,
				},
				{
					Pitch:        song.La,
					Duration:     song.Quarter,
					Dotted:       song.SingleDotted,
					Offset:       song.Quarter,
					DottedOffset: song.SingleDotted,
				},
				{
					Pitch:        song.Do,
					Duration:     song.Sixteenth,
					Offset:       song.Half,
					DottedOffset: song.DoubleDotted,
				},
				{
					Pitch:        song.Re,
					Duration:     song.Quarter,
					Offset:       song.Half,
					DottedOffset: song.SingleDotted,
				},
			},
		},
	}
	for {
		s.Play()
		fmt.Println("Song played")
		time.Sleep(0 * time.Second)
	}
}
