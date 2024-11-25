package utils

import (
	"strings"
)

func GetBriefLines(lines []string) []string {
	var briefLines []string
	for _, line := range lines {
		briefLine := strings.Replace(line, "https://", "", 1)
		// briefLine = strings.Replace(briefLine, "www.", "", 1)
		domains := strings.Split(briefLine, "/")
		domain := domains[0]
		briefLines = append(briefLines, domain)
	}
	return briefLines
}
