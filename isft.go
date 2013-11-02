package wav

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/campoy/riff"
)

func init() {
	funcs[riff.NewID("ISFT")] = isftDec
}

type ISFTContent string

func (c *ISFTContent) String() string {
	return fmt.Sprintf("%s", c)
}

func isftDec(r io.Reader) (interface{}, error) {
	p, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return ISFTContent(p), nil
}
