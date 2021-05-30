package usableurl

import (
	"bytes"
	"os/exec"
	"regexp"
	"strings"
)

func Sanitize(url string) string{
	url = strings.TrimSpace(url)
	cmd := exec.Command("curl", "-sLI", url)
	buf := bytes.Buffer{}
	cmd.Stdout = &buf
	cmd.Run()
	re := regexp.MustCompile(`[Ll]ocation:\s*(.*)\s`)
	res := re.FindAllStringSubmatch(buf.String(), -1)
	bestURL := ""
	for _, matches := range res{
		match := matches[1]
		if strings.HasPrefix(match, "http"){
			if len(bestURL) < len(match){
				bestURL = match
			}
		}
	}
	return bestURL
}
