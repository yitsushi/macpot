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
	assert.Equal(t, "invalid OUI: 11:22", err.Error())
}

func TestNew_WithOUI_long(t *testing.T) {
	_, err := macpot.New(
		macpot.WithOUI("11:22:33:44"),
	)

	assert.Error(t, err)
	assert.ErrorAs(t, err, &macpot.OUIError{})
	assert.Equal(t, "invalid OUI: 11:22:33:44", err.Error())
}

func TestNew_WithOUI_invalidCharacter(t *testing.T) {
	_, err := macpot.New(
		macpot.WithOUI("11:2l:33"),
	)

	assert.Error(t, err)
	assert.ErrorAs(t, err, &macpot.OUIError{})
	assert.Equal(t, "strconv.ParseInt: parsing \"2l\": invalid syntax", err.Error())
}
