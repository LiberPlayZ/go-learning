package regex

import (
	"fmt"
	"regexp"
	"strings"
)

// IsValidLine checks if the log line starts with a valid log level using regex
func IsValidLine(line string) bool {
	// Ensure the log level is at the start using `^` and grouping the options
	pattern := `^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`
	matched, _ := regexp.MatchString(pattern, line)
	return matched
}

// SplitLogLine splits a log line using the custom separators
func SplitLogLine(line string) []string {
	// Regular expression to match the separators
	pattern := `<[~*=/-]*>` // Matches < followed by ~, *, =, or - any number of times, then >

	// Compile the regex
	re := regexp.MustCompile(pattern)

	// Use regex to split the line
	fields := re.Split(line, -1)

	return fields
}

// CountQuotedPasswords counts occurrences of "password" inside quotes
func CountQuotedPasswords(lines []string) int {
	// Regular expression to match "password" in quotes, case-insensitive
	re := regexp.MustCompile(`"[^"]*password[^"]*"`)

	count := 0
	for _, line := range lines {
		if re.MatchString(strings.ToLower(line)) {
			count++
		}
	}
	return count
}

// RemoveEndOfLineText removes occurrences of "end-of-line" followed by numbers
func RemoveEndOfLineText(line string) string {
	// Regular expression to match "end-of-line" followed by digits
	re := regexp.MustCompile(`end-of-line\d+`)

	// Replace matches with an empty string
	return re.ReplaceAllString(line, "")
}

// TagWithUserName tags lines containing "User <username>"
func TagWithUserName(lines []string) []string {
	// Regular expression to capture "User <username>"
	re := regexp.MustCompile(`User\s+(\S+)`)

	var taggedLines []string

	for _, line := range lines {
		matches := re.FindStringSubmatch(line)
		if len(matches) > 1 {
			username := matches[1]
			taggedLines = append(taggedLines, fmt.Sprintf("[USR] %s %s", username, line))
		} else {
			taggedLines = append(taggedLines, line)
		}
	}

	return taggedLines
}
