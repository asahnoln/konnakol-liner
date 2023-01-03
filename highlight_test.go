package konnakolliner_test

import (
	"testing"

	liner "github.com/asahnoln/konnakol-liner"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHighlight(t *testing.T) {
	str, err := liner.Highlight(
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
	assert.Equal(t, `|1|(Tha) Ka Dhi Mi
|2|(Tha) Ka Dhi
	Mi |3|(Tha) Ka Dhi    Mi |1|(Tha)
Ka 						Dhi
		Mi
|2|(Tha)
	Ka
		Dhi
	Mi`,
	str)
}
