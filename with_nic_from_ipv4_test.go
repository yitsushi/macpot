package macpot_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/macpot"
)

func TestNew_WithNICFromIPv4(t *testing.T) {
	mac, err := macpot.New(
		macpot.WithOUI("11:22:33"),
		macpot.WithNICFromIPv4("192.168.31.7"),
	)

	assert.NoError(t, err)
	assert.Equal(t, "11:22:33:a8:1f:07", mac.ToString())
}

func TestNew_WithNICFromIPv4_short(t *testing.T) {
	_, err := macpot.New(
		macpot.WithNICFromIPv4("168.31.7"),
	)

	assert.Error(t, err)
	assert.ErrorAs(t, err, &macpot.IPv4Error{})
	assert.Equal(t, "invalid IPv4 address: 168.31.7", err.Error())
}

func TestNew_WithNICFromIPv4_long(t *testing.T) {
	_, err := macpot.New(
		macpot.WithNICFromIPv4("192.168.31.7.33"),
	)

	assert.Error(t, err)
	assert.ErrorAs(t, err, &macpot.IPv4Error{})
	assert.Equal(t, "invalid IPv4 address: 192.168.31.7.33", err.Error())
}

func TestNew_WithNICFromIPv4_invalidCharacter(t *testing.T) {
	_, err := macpot.New(
		macpot.WithNICFromIPv4("192.168.p1.7"),
	)

	assert.Error(t, err)
	assert.ErrorAs(t, err, &macpot.IPv4Error{})
	assert.Equal(t, "strconv.ParseInt: parsing \"p1\": invalid syntax", err.Error())
}
