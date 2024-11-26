package utils

import (
	"strings"
)

func GetBriefLines(lines []string) []string {
	var briefLines []string
	for _, line := range lines {
		briefLine := strings.Split(line, "@")
		if len(briefLine) > 1 && len(briefLine[1]) != 0 {
			briefLines = append(briefLines, briefLine[0]+"(@)")
		} else {
			briefLines = append(briefLines, briefLine[0])
		}
	}
	return briefLines
}
