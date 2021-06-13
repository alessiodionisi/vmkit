package macaddress

import (
	"crypto/rand"
	"fmt"
)

func NewUnicastLocallyAdministeredMACAddress() (string, error) {
	buf := make([]byte, 6)

	_, err := rand.Read(buf)
	if err != nil {
		return "", err
	}

	// set local bit and ensure unicast address
	buf[0] = (buf[0] | 2) & 0xfe

	return fmt.Sprintf("%02x:%02x:%02x:%02x:%02x:%02x", buf[0], buf[1], buf[2], buf[3], buf[4], buf[5]), nil
}
