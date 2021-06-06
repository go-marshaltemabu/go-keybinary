package keybinary

import (
	"strconv"
)

// ErrIncorrectDataSize indicate size of given data is not correct.
type ErrIncorrectDataSize struct {
	ExpectSize   int
	ReceivedSize int
}

func (e *ErrIncorrectDataSize) Error() string {
	return "[ErrIncorrectDataSize: expect=" + strconv.FormatInt(int64(e.ExpectSize), 10) + ", received=" + strconv.FormatInt(int64(e.ReceivedSize), 10) + "]"
}
