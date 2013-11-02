package wav

import (
	"io"

	"github.com/campoy/riff"
)

var funcs = map[riff.ID]riff.DecoderFunc{}

type Stream struct {
	c *riff.Chunk
}

func NewStream(r io.Reader) (*Stream, error) {
	d := riff.NewDecoder(r)

	for id, f := range funcs {
		if err := d.Map(id, f); err != nil {
			return nil, err
		}
	}

	c, err := d.Decode()
	if err != nil {
		return nil, err
	}

	return &Stream{c}, nil
}
