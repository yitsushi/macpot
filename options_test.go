package macpot_test

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/macpot"
)

func TestNew_WithOUI(t *testing.T) {
	mac, err := macpot.New(
		macpot.WithOUI("11:22:33"),
	)

	assert.NoError(t, err)
	assert.Condition(t, func() bool {
		re := regexp.MustCompile("^11:22:33(:[a-z0-9]{2}){3}$")

		return re.Match([]byte(mac.ToString()))
	}, "random generated MAC address is not valid: %s", mac.ToString())
}

func TestNew_WithOUI_short(t *testing.T) {
	_, err := macpot.New(
		macpot.WithOUI("11:22"),
	)

	assert.Error(t, err)
	assert.ErrorAs(t, err, &macpot.OUIError{})
}

func TestNew_WithOUI_long(t *testing.T) {
	_, err := macpot.New(
		macpot.WithOUI("11:22:33:44"),
	)

	assert.Error(t, err)
	assert.ErrorAs(t, err, &macpot.OUIError{})
}

func TestNew_WithOUI_invalidCharacter(t *testing.T) {
	_, err := macpot.New(
		macpot.WithOUI("11:2l:33"),
	)

	assert.Error(t, err)
	assert.ErrorAs(t, err, &macpot.OUIError{})
}

func TestNew_WithNIC(t *testing.T) {
	mac, err := macpot.New(
		macpot.WithNIC("11:22:33"),
	)

	assert.NoError(t, err)
	assert.Condition(t, func() bool {
		re := regexp.MustCompile("^[a-z0-9]{2}(:[a-z0-9]{2}){2}:11:22:33$")

		return re.Match([]byte(mac.ToString()))
	}, "random generated MAC address is not valid: %s", mac.ToString())
}

func TestNew_WithNIC_short(t *testing.T) {
	_, err := macpot.New(
		macpot.WithNIC("11:22"),
	)

	assert.Error(t, err)
	assert.ErrorAs(t, err, &macpot.NICError{})
}

func TestNew_WithNIC_long(t *testing.T) {
	_, err := macpot.New(
		macpot.WithNIC("11:22:33:44"),
	)

	assert.Error(t, err)
	assert.ErrorAs(t, err, &macpot.NICError{})
}

func TestNew_WithNIC_invalidCharacter(t *testing.T) {
	_, err := macpot.New(
		macpot.WithNIC("11:2l:33"),
	)

	assert.Error(t, err)
	assert.ErrorAs(t, err, &macpot.NICError{})
}

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
	assert.ErrorAs(t, err, &macpot.NICError{})
}

func TestNew_WithNICFromIPv4_long(t *testing.T) {
	_, err := macpot.New(
		macpot.WithNICFromIPv4("192.168.31.7.33"),
	)

	assert.Error(t, err)
	assert.ErrorAs(t, err, &macpot.NICError{})
}

func TestNew_WithNICFromIPv4_invalidCharacter(t *testing.T) {
	_, err := macpot.New(
		macpot.WithNICFromIPv4("192.168.p1.7"),
	)

	assert.Error(t, err)
	assert.ErrorAs(t, err, &macpot.NICError{})
}

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
