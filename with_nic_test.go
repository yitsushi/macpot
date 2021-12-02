package macpot_test

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/macpot"
)

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
	assert.Equal(t, "invalid NIC: 11:22", err.Error())
}

func TestNew_WithNIC_long(t *testing.T) {
	_, err := macpot.New(
		macpot.WithNIC("11:22:33:44"),
	)

	assert.Error(t, err)
	assert.ErrorAs(t, err, &macpot.NICError{})
	assert.Equal(t, "invalid NIC: 11:22:33:44", err.Error())
}

func TestNew_WithNIC_invalidCharacter(t *testing.T) {
	_, err := macpot.New(
		macpot.WithNIC("11:2l:33"),
	)

	assert.Error(t, err)
	assert.ErrorAs(t, err, &macpot.NICError{})
	assert.Equal(t, "strconv.ParseInt: parsing \"2l\": invalid syntax", err.Error())
}
