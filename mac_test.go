package macpot_test

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/macpot"
)

func TestNew(t *testing.T) {
	mac, err := macpot.New()

	assert.NoError(t, err)
	assert.Condition(t, func() bool {
		re := regexp.MustCompile("^[a-z0-9]+(:[a-z0-9]{2}){5}$")

		return re.Match([]byte(mac.ToString()))
	}, "random generated MAC address is not valid: %s", mac.ToString())
}

func TestNewFromBytes(t *testing.T) {
	mac := macpot.NewFromBytes([]byte{0x11, 0x22, 0x33, 0x44, 0x55, 0x66})
	assert.Equal(t, "11:22:33:44:55:66", mac.ToString())
}

func TestNewFromBytes_longer(t *testing.T) {
	mac := macpot.NewFromBytes([]byte{0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88})
	assert.Equal(t, "11:22:33:44:55:66", mac.ToString())
}

func TestNewFromBytes_shorter(t *testing.T) {
	mac := macpot.NewFromBytes([]byte{0x11, 0x22, 0x33, 0x44})
	assert.Equal(t, "11:22:33:44:00:00", mac.ToString())
}

func TestNewFromUint64(t *testing.T) {
	mac := macpot.NewFromUint64(18838586676582)
	assert.Equal(t, "11:22:33:44:55:66", mac.ToString())

	assert.Equal(t, uint64(18838586676582), mac.ToUint64())
}

func TestNewFromUint64_max(t *testing.T) {
	mac := macpot.NewFromUint64(281474976710655)
	assert.Equal(t, "ff:ff:ff:ff:ff:ff", mac.ToString())
}

func TestNewFromUint64_overflow(t *testing.T) {
	mac := macpot.NewFromUint64(281474976710665)
	assert.Equal(t, "00:00:00:00:00:09", mac.ToString())
}

func TestMac_Next(t *testing.T) {
	mac := macpot.NewFromBytes([]byte{0x11, 0x22, 0x33, 0x44, 0x55, 0x66})
	assert.Equal(t, "11:22:33:44:55:67", mac.Next().ToString())
}

func TestMac_Next_overflow(t *testing.T) {
	mac := macpot.NewFromBytes([]byte{0x11, 0x22, 0x33, 0x44, 0xff, 0xff})
	assert.Equal(t, "11:22:33:45:00:00", mac.Next().ToString())
}

func TestMac_SetLocal(t *testing.T) {
	mac := macpot.NewFromBytes([]byte{0x11, 0x22, 0x33, 0x44, 0x55, 0x66})
	assert.False(t, mac.IsLocal())

	mac.SetLocal()
	assert.Equal(t, "13:22:33:44:55:66", mac.ToString())
	assert.True(t, mac.IsLocal())
}

func TestMac_SetLocal_noChanges(t *testing.T) {
	mac := macpot.NewFromBytes([]byte{0x13, 0x22, 0x33, 0x44, 0x55, 0x66})
	assert.True(t, mac.IsLocal())

	mac.SetLocal()
	assert.Equal(t, "13:22:33:44:55:66", mac.ToString())
	assert.True(t, mac.IsLocal())
}

func TestMac_SetGlobal(t *testing.T) {
	mac := macpot.NewFromBytes([]byte{0x13, 0x22, 0x33, 0x44, 0x55, 0x66})
	assert.False(t, mac.IsGlobal())

	mac.SetGlobal()
	assert.Equal(t, "11:22:33:44:55:66", mac.ToString())
	assert.True(t, mac.IsGlobal())
}

func TestMac_SetGlobal_noChanges(t *testing.T) {
	mac := macpot.NewFromBytes([]byte{0x11, 0x22, 0x33, 0x44, 0x55, 0x66})
	assert.True(t, mac.IsGlobal())

	mac.SetGlobal()
	assert.Equal(t, "11:22:33:44:55:66", mac.ToString())
	assert.True(t, mac.IsGlobal())
}

func TestMac_SetUnicast(t *testing.T) {
	mac := macpot.NewFromBytes([]byte{0x13, 0x22, 0x33, 0x44, 0x55, 0x66})
	assert.False(t, mac.IsUnicast())

	mac.SetUnicast()
	assert.Equal(t, "12:22:33:44:55:66", mac.ToString())
	assert.True(t, mac.IsUnicast())
}

func TestMac_SetUnicast_noChanges(t *testing.T) {
	mac := macpot.NewFromBytes([]byte{0x12, 0x22, 0x33, 0x44, 0x55, 0x66})
	assert.True(t, mac.IsUnicast())

	mac.SetUnicast()
	assert.Equal(t, "12:22:33:44:55:66", mac.ToString())
	assert.True(t, mac.IsUnicast())
}

func TestMac_SetMulticast(t *testing.T) {
	mac := macpot.NewFromBytes([]byte{0x12, 0x22, 0x33, 0x44, 0x55, 0x66})
	assert.False(t, mac.IsMulticast())

	mac.SetMulticast()
	assert.Equal(t, "13:22:33:44:55:66", mac.ToString())
	assert.True(t, mac.IsMulticast())
}

func TestMac_SetMulticast_noChanges(t *testing.T) {
	mac := macpot.NewFromBytes([]byte{0x13, 0x22, 0x33, 0x44, 0x55, 0x66})
	assert.True(t, mac.IsMulticast())

	mac.SetMulticast()
	assert.Equal(t, "13:22:33:44:55:66", mac.ToString())
	assert.True(t, mac.IsMulticast())
}

func TestMac_SetOctet(t *testing.T) {
	mac := macpot.NewFromBytes([]byte{0x11, 0x22, 0x33, 0x44, 0x55, 0x66})
	err := mac.SetOctet(3, 0x99)

	assert.NoError(t, err)
	assert.Equal(t, "11:22:33:99:55:66", mac.ToString())
}

func TestMac_SetOctet_negativeIndex(t *testing.T) {
	mac := macpot.NewFromBytes([]byte{0x11, 0x22, 0x33, 0x44, 0x55, 0x66})
	err := mac.SetOctet(-3, 0x99)

	assert.Error(t, err)
	assert.ErrorAs(t, err, &macpot.OutOfBoundError{})
	assert.Equal(t, "11:22:33:44:55:66", mac.ToString())
	assert.Equal(t, "unable to set octet -3", err.Error())
}

func TestMac_SetOctet_largeIndex(t *testing.T) {
	mac := macpot.NewFromBytes([]byte{0x11, 0x22, 0x33, 0x44, 0x55, 0x66})
	err := mac.SetOctet(55, 0x99)

	assert.Error(t, err)
	assert.ErrorAs(t, err, &macpot.OutOfBoundError{})
	assert.Equal(t, "11:22:33:44:55:66", mac.ToString())
	assert.Equal(t, "unable to set octet 55", err.Error())
}
