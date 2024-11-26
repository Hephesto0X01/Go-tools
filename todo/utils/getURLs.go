package utils

import (
	"strings"
)

func GetURLs(lines []string) []string {
	var urls []string
	for _, line := range lines {
		url := strings.Split(line, "@")
		if len(url) > 1 && len(url[1]) != 0 {
			urls = append(urls, url[1])
		} else {
			urls = append(urls, url[0]+"(null)")
			continue
		}
	}
	return urls
}
