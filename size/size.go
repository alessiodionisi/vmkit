package size

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	numberRegex = regexp.MustCompile("[0-9]+")
	textRegex   = regexp.MustCompile("[a-zA-Z]+")
)

func ToBytes(size string) (uint64, error) {
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

func ParseUint(size uint64, unit string) (string, error) {
	unit = strings.ToLower(unit)

	var number uint64
	switch unit {
	case "kb":
		number = size / 1000
		unit = "kB"
	case "mb":
		number = size / (1000 * 1000)
		unit = "MB"
	case "gb":
		number = size / (1000 * 1000 * 1000)
		unit = "GB"
	case "kib":
		number = size / 1024
		unit = "KiB"
	case "mib":
		number = size / (1024 * 1024)
		unit = "MiB"
	case "gib":
		number = size / (1024 * 1024 * 1024)
		unit = "GiB"
	case "":
		number = size
		unit = "B"
	default:
		return "", fmt.Errorf("unknown unit: %s", unit)
	}

	return fmt.Sprintf("%d%s", number, unit), nil
}
