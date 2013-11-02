package wav

import (
	"fmt"
	"os"
	"testing"
)

func TestNewStream(t *testing.T) {
	f, err := os.Open("data/test.wav")
	if err != nil {
		t.Fatalf("open test file: %v", err)
	}
	s, err := NewStream(f)
	if err != nil {
		t.Errorf("NewStream: %v", err)
	}
	fmt.Printf("%v\n", s.c)
}

func TestWrite(t *testing.T) {
	f, err := os.Open("data/hand.wav")
	if err != nil {
		t.Fatalf("open test file: %v", err)
	}
	s, err := NewStream(f)
	if err != nil {
		t.Errorf("NewStream: %v", err)
	}
	c := s.c.Chunks[2]

	orig := c.Data
	c.Data = orig[:len(orig)/2]
	c.Len = uint32(len(c.Data))

	fw, err := os.Create("data/first.wav")
	if err != nil {
		t.Fatalf("open write file: %v", err)
	}
	_, err = c.WriteTo(fw)
	if err != nil {
		t.Fatalf("writing: %v", err)
	}

	c.Data = orig[len(orig)/2:]
	c.Len = uint32(len(c.Data))

	fw, err = os.Create("data/second.wav")
	if err != nil {
		t.Fatalf("open write file: %v", err)
	}
	_, err = c.WriteTo(fw)
	if err != nil {
		t.Fatalf("writing: %v", err)
	}
}

/*
func TestNew(t *testing.T) {
	fw, err := os.Create("data/extra.wav")
	if err != nil {
		t.Fatalf("open write file: %v", err)
	}

}*/
