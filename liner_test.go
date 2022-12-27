package konnakolliner_test

import (
	"bytes"
	"errors"
	"strings"
	"testing"
	"testing/iotest"

	liner "github.com/asahnoln/konnakol-liner"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMessyInputString(t *testing.T) {
	str, err := liner.Line(
		`Tha Ka Dhi Mi
Tha Ka Dhi
	Mi Tha Ka Dhi    Mi Tha
Ka 						Dhi
		Mi
Tha
	Ka
		Dhi
	Mi`,
		liner.ThalamRoopaga,
		liner.GathiChatushram)
	require.NoError(t, err)
	assert.Equal(t, `|1| Tha Ka Dhi Mi
|2| Tha Ka Dhi Mi
|3| Tha Ka Dhi Mi
|1| Tha Ka Dhi Mi
|2| Tha Ka Dhi Mi`, str)
}

func TestGetLineOut(t *testing.T) {
	b := &bytes.Buffer{}
	s := strings.NewReader("Thaam . Dheem . Thaam . Dheem . Thaam . Dheem .")

	err := liner.LineOut(s, b, liner.ThalamAdi, liner.GathiTisram)
	require.NoError(t, err)

	assert.Equal(t, `|1| Thaam . Dheem
|2| . Thaam .
|3| Dheem . Thaam
|4| . Dheem .`, b.String())
}

func TestZeroCounts(t *testing.T) {
	_, err := liner.Line("", 1, 0)
	require.ErrorIs(t, err, liner.ErrGathiLessThanOne)

	_, err = liner.Line("", 1, -5)
	require.ErrorIs(t, err, liner.ErrGathiLessThanOne)

	_, err = liner.Line("", 0, 1)
	require.ErrorIs(t, err, liner.ErrThalamLessThanOne)

	_, err = liner.Line("", -10, 1)
	require.ErrorIs(t, err, liner.ErrThalamLessThanOne)
}

type ErrWriter struct {
}

func (*ErrWriter) Write([]byte) (int, error) {
	return 0, iotest.ErrTimeout
}

func TestErrWritingLineOut(t *testing.T) {
	r := strings.NewReader("")
	w := &ErrWriter{}

	err := liner.LineOut(r, w, 1, 1)

	require.Error(t, err)
}

func TestErrLineOutWrongCounts(t *testing.T) {
	r := strings.NewReader("")
	w := &bytes.Buffer{}

	err := liner.LineOut(r, w, 0, 0)

	require.Error(t, err)
}

func TestErrReadingLineOut(t *testing.T) {
	r := iotest.ErrReader(errors.New("error reading"))
	w := &bytes.Buffer{}

	err := liner.LineOut(r, w, 0, 0)

	require.Error(t, err)
}
