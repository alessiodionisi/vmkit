package bytesize

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	binaryUnits  = []string{"B", "KiB", "MiB", "GiB", "TiB", "PiB", "EiB", "ZiB", "YiB"}
	decimalUnits = []string{"B", "kB", "MB", "GB", "TB", "PB", "EB", "ZB", "YB"}
	numberRegex  = regexp.MustCompile("[0-9]+")
	textRegex    = regexp.MustCompile("[a-zA-Z]+")
)

func Parse(size string) (uint64, error) {
	size = strings.ToLower(size)

	numberString := numberRegex.FindString(size)
	unit := textRegex.FindString(size)

	number, err := strconv.ParseUint(numberString, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("error parsing number: %w", err)
	}

	if unit == "" {
		return number, nil
	}

	var bytes uint64
	switch unit {
	case "kb":
		bytes = number * 1000
	case "mb":
		bytes = number * 1000 * 1000
	case "gb":
		bytes = number * 1000 * 1000 * 1000
	case "kib":
		bytes = number * 1024
	case "mib":
		bytes = number * 1024 * 1024
	case "gib":
		bytes = number * 1024 * 1024 * 1024
	default:
		return 0, fmt.Errorf("unknown unit: %s", unit)
	}

	return bytes, nil
}

// FormatDecimal returns a string representation of the size in bytes using decimal units.
// The size is formatted to 4 significant figures.
// The units are chosen based on the size, and the largest unit that can represent the size is used.
func FormatDecimal(size uint64) string {
	return format(size, false)
}

// FormatBinary returns a string representation of the size in bytes using binary units.
// The size is formatted to 4 significant figures.
// The units are chosen based on the size, and the largest unit that can represent the size is used.
func FormatBinary(size uint64) string {
	return format(size, true)
}

// format returns a string representation of the size in bytes using the specified units.
// If binary is true, it uses binary units, otherwise it uses decimal units.
// The size is formatted to 4 significant figures.
// The units are chosen based on the size, and the largest unit that can represent the size is used.
func format(size uint64, binary bool) string {
	base := float64(1000)
	units := decimalUnits

	if binary {
		base = 1024
		units = binaryUnits
	}

	fsize := float64(size)
	unitsLimit := len(units) - 1
	i := 0

	for fsize >= base && i < unitsLimit {
		fsize = fsize / base
		i++
	}

	return fmt.Sprintf("%.4g %s", fsize, units[i])
}
