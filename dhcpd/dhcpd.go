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
	name   string
	leases []Lease
}

func (l *Leases) Parse() error {
	file, err := os.Open(l.name)
	if err != nil {
		return err
	}
	defer file.Close()

	var lease Lease
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "{" {
			lease = Lease{}
		} else if line == "}" {
			l.leases = append(l.leases, lease)
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

	return nil
}

func NewLeases(name string) *Leases {
	return &Leases{
		name: name,
	}
}
