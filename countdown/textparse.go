package countdown

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

func parseTime(s string) (time.Duration, error) {
	// Parses the time in the text input
	var minutes int
	var seconds int

	// Try to parse an integer
	regex, _ := regexp.Compile(`^\d+`)
	match := regex.FindStringIndex(s)
	if match == nil {
		return 0, fmt.Errorf("expected integer at start")
	}

	val, err := strconv.Atoi(s[0:match[1]])
	if err != nil {
		return 0, fmt.Errorf("failed to parse integer")
	}

	s = s[match[1]:]
	if len(s) == 0 {
		return 0, fmt.Errorf("expected m or s after integer")
	}

	parsedMinutes := false

	switch s[0] {
	case 'm':
		minutes = val
		parsedMinutes = true
	case 's':
		seconds = val
	default:
		return 0, fmt.Errorf("expected m or s after integer")
	}

	s = s[1:]

	if len(s) > 0 {
		// If we parsed minutes, we're not done
		if !parsedMinutes {
			return 0, fmt.Errorf("unexpected extra content after seconds")
		}

		// Parse another integer
		match := regex.FindStringIndex(s)
		if match == nil {
			return 0, fmt.Errorf("expected integer at start")
		}

		val, err := strconv.Atoi(s[0:match[1]])
		if err != nil {
			return 0, fmt.Errorf("failed to parse integer")
		}

		s = s[match[1]:]
		if len(s) == 0 || s[0] != 's' {
			return 0, fmt.Errorf("expected s after integer")
		}

		seconds = val
	}

	return time.Duration(minutes)*time.Minute + time.Duration(seconds)*time.Second, nil
}
