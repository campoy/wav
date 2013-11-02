package wav

import (
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/campoy/riff"
)

func init() {
	funcs[riff.NewID("fmt ")] = fmtDec
}

type FmtContent struct {
	CompressionCode           uint16
	NumChans                  uint16
	SampleRate                uint32
	AvgBytesPerSecond         uint32
	BlockAlign                uint16
	SignificantBytesPerSample uint16
	ExtraFormatBytes          []byte
}

func (c *FmtContent) String() string {
	return fmt.Sprintf("cc:%d|nc:%d|sr:%d|abps:%d|ba:%d|sbps:%d",
		c.CompressionCode,
		c.NumChans,
		c.SampleRate,
		c.AvgBytesPerSecond,
		c.BlockAlign,
		c.SignificantBytesPerSample)
}

func fmtDec(r io.Reader) (interface{}, error) {
	fc := new(FmtContent)

	vs := []interface{}{
		&fc.CompressionCode,
		&fc.NumChans,
		&fc.SampleRate,
		&fc.AvgBytesPerSecond,
		&fc.BlockAlign,
		&fc.SignificantBytesPerSample,
	}

	for _, v := range vs {
		err := binary.Read(r, binary.LittleEndian, v)
		if err != nil {
			return nil, err
		}
	}

	p, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	fc.ExtraFormatBytes = p

	return fc, nil
}
