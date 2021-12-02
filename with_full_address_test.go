package macpot_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/macpot"
)

func TestNew_WithFullAddress(t *testing.T) {
	mac, err := macpot.New(
		macpot.WithFullAddress("11:22:33:44:55:66"),
	)

	assert.NoError(t, err)
	assert.Equal(t, "11:22:33:44:55:66", mac.ToString())
}

func TestNew_WithFullAddress_short(t *testing.T) {
	_, err := macpot.New(
		macpot.WithFullAddress("11:22"),
	)

	assert.Error(t, err)
	assert.ErrorAs(t, err, &macpot.AddressError{})
	assert.Equal(t, "invalid MAC address length: 11:22", err.Error())
}

func TestNew_WithFullAddress_long(t *testing.T) {
	_, err := macpot.New(
		macpot.WithFullAddress("11:22:33:44:55:66:77"),
	)

	assert.Error(t, err)
	assert.ErrorAs(t, err, &macpot.AddressError{})
	assert.Equal(t, "invalid MAC address length: 11:22:33:44:55:66:77", err.Error())
}

func TestNew_WithFullAddress_invalidCharacter(t *testing.T) {
	_, err := macpot.New(
		macpot.WithFullAddress("11:2l:33:44:55:66"),
	)

	assert.Error(t, err)
	assert.ErrorAs(t, err, &macpot.AddressError{})
	assert.Equal(t, "strconv.ParseInt: parsing \"2l\": invalid syntax", err.Error())
}
