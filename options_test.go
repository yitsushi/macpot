package macpot_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/macpot"
)

func TestNew_AsMulticast(t *testing.T) {
	mac, err := macpot.New(
		macpot.WithOUI("12:22:33"),
		macpot.WithNICFromIPv4("192.168.31.7"),
		macpot.AsMulticast(),
	)

	assert.Condition(t, mac.IsMulticast)
	assert.NoError(t, err)
	assert.Equal(t, "13:22:33:a8:1f:07", mac.ToString())
}

func TestNew_AsUnicast(t *testing.T) {
	mac, err := macpot.New(
		macpot.WithOUI("13:22:33"),
		macpot.WithNICFromIPv4("192.168.31.7"),
		macpot.AsUnicast(),
	)

	assert.Condition(t, mac.IsUnicast)
	assert.NoError(t, err)
	assert.Equal(t, "12:22:33:a8:1f:07", mac.ToString())
}

func TestNew_AsLocal(t *testing.T) {
	mac, err := macpot.New(
		macpot.WithOUI("14:22:33"),
		macpot.WithNICFromIPv4("192.168.31.7"),
		macpot.AsLocal(),
	)

	assert.Condition(t, mac.IsLocal)
	assert.NoError(t, err)
	assert.Equal(t, "16:22:33:a8:1f:07", mac.ToString())
}

func TestNew_AsGlobal(t *testing.T) {
	mac, err := macpot.New(
		macpot.WithOUI("16:22:33"),
		macpot.WithNICFromIPv4("192.168.31.7"),
		macpot.AsGlobal(),
	)

	assert.Condition(t, mac.IsGlobal)
	assert.NoError(t, err)
	assert.Equal(t, "14:22:33:a8:1f:07", mac.ToString())
}
