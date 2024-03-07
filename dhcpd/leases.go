package dhcpd

import (
	"bufio"
	"os"
	"strings"
)

type Lease struct {
	IPAddress string
	HWAddress string
}

type Leases struct {
	leases []Lease
}

func (l *Leases) FindByHWAddress(hwAddress string) *Lease {
	for _, lease := range l.leases {
		if lease.HWAddress == hwAddress {
			return &lease
		}
	}

	return nil
}

func ParseLeases(path string) (*Leases, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var leases []Lease
	var lease Lease

	for scanner.Scan() {
		line := scanner.Text()

		if line == "{" {
			lease = Lease{}
		} else if line == "}" {
			leases = append(leases, lease)
		} else {
			parts := strings.SplitN(line, "=", 2)

			if len(parts) == 2 {
				key := strings.TrimSpace(parts[0])
				value := strings.TrimSpace(parts[1])

				switch key {
				case "ip_address":
					lease.IPAddress = value
				case "hw_address":
					lease.HWAddress = strings.Split(value, ",")[1]
				}
			}
		}
	}

	return &Leases{
		leases: leases,
	}, nil
}
