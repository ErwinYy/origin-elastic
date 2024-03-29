// +build linux

package linux

import (
	"bytes"
	"encoding/binary"
	"testing"

	"github.com/ErwinYy/origin-elastic/gosigar/sys"
	"github.com/stretchr/testify/assert"
)

func TestParseNetlinkErrorDataTooShort(t *testing.T) {
	assert.Error(t, ParseNetlinkError(nil), "too short")
}

func TestParseNetlinkErrorErrno(t *testing.T) {
	buf := new(bytes.Buffer)
	binary.Write(buf, sys.GetEndian(), -1*int32(NLE_MSG_TOOSHORT))
	assert.Equal(t, ParseNetlinkError(buf.Bytes()), NLE_MSG_TOOSHORT)
}
