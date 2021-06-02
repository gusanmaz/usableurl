package usableurl

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"regexp"
	"strings"
	"time"
)

const curlTimeout = 2 * time.Minute

func Sanitize(url string) string {
	ctx, cancel := context.WithTimeout(context.Background(), curlTimeout)
	defer cancel()

	url = strings.TrimSpace(url)
	cmd := exec.CommandContext(ctx, "curl", "-sLI", url)
	bs, _ := cmd.Output()

	if ctx.Err() == context.DeadlineExceeded {
		fmt.Printf("Command timed out. We will still attempt to expand the %v. "+
			"If unssuccesful url will be expanded to empty string\n", url)
	}

	buf := bytes.Buffer{}
	buf.Write(bs)

	re := regexp.MustCompile(`[Ll]ocation:\s*(.*)\s`)
	res := re.FindAllStringSubmatch(buf.String(), -1)
	bestURL := ""
	for _, matches := range res {
		match := matches[1]
		if strings.HasPrefix(match, "http") {
			if len(bestURL) < len(match) {
				bestURL = match
			}
		}
	}
	return bestURL
}
