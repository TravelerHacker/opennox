package blobs

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFormatAccesses(t *testing.T) {
	if os.Getenv("OPENNOX_BLOB_FORMAT") != "1" {
		t.Skip("tree-wide blob formatter mutates sources; set OPENNOX_BLOB_FORMAT=1 to run")
	}
	SetPath("../../")
	err := FormatAccesses()
	require.NoError(t, err)
}
