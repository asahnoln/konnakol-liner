package konnakolliner

import (
	"errors"
	"io"
	"regexp"
	"strconv"
	"strings"
)

const (
	ThalamAdi     = 8
	ThalamRoopaga = 3
)

const (
	GathiTisram     = 3
	GathiChatushram = 4
	GathiKanda      = 5
	GathiMisram     = 7
	GathiSangeerna  = 9
)

var (
	ErrGathiLessThanOne  = errors.New("gathi must be greater than zero")
	ErrThalamLessThanOne = errors.New("thalam must be greater than zero")
)

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

func LineOut(r io.Reader, w io.Writer, thalam, gathi int) error {
	s, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	line, err := Line(string(s), thalam, gathi)
	if err != nil {
		return err
	}
	_, err = w.Write([]byte(line))
	if err != nil {
		return err
	}

	return nil
}
