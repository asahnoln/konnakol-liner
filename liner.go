// Package konnakolliner presents a couple of methods to format given strings
// into separated lines by count of given thalam
package konnakolliner

import (
	"errors"
	"io"
	"regexp"
	"strconv"
	"strings"
)

// Thalams, available to work with
const (
	ThalamAdi     = 8
	ThalamRoopaga = 3
)

// Gathi, available to work with
const (
	GathiTisram     = 3
	GathiChatushram = 4
	GathiKanda      = 5
	GathiMisram     = 7
	GathiSangeerna  = 9
)

var (
	// ErrGathiLessThanOne signals that given gathi is less than one
	ErrGathiLessThanOne = errors.New("gathi must be greater than zero")
	// ErrThalamLessThanOne signals that given thalam count is less than one
	ErrThalamLessThanOne = errors.New("thalam must be greater than zero")
)

// Line formats given konnakol line according to given thalam and gathi into separate lines for each count
func Line(line string, thalam, gathi int) (string, error) {
	if thalam < 1 {
		return "", ErrThalamLessThanOne
	}
	if gathi < 1 {
		return "", ErrGathiLessThanOne
	}

	var result strings.Builder
	thalamIndex := 0
	sols := regexp.MustCompile(`\s+`).Split(line, -1)

	for i, s := range sols {
		if i%gathi == 0 {
			thalamIndex = thalamIndex%thalam + 1
			if i > 0 {
				result.WriteByte('\n')
			}

			result.WriteString("|" + strconv.Itoa(thalamIndex) + "|")
		}
		result.WriteString(" " + s)
	}
	return result.String(), nil
}

// Out does the same thing as Line/Highlight, but it reads from a Reader and prints to Writer
func Out(r io.Reader, w io.Writer, f func(string, int, int) (string, error), thalam, gathi int) error {
	s, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	line, err := f(string(s), thalam, gathi)
	if err != nil {
		return err
	}
	_, err = w.Write([]byte(line))
	if err != nil {
		return err
	}

	return nil
}
