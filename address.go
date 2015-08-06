package smtpd

import (
	"fmt"
	"strings"
)

func parseAddress(src string) (string, error) {
	src = strings.Trim(src, " \t")
	if src == "" {
		return "", fmt.Errorf("No address given%s", "")
	}
	posBegin := strings.Index(src, "<")
	posEnd := strings.Index(src, ">")
	if posBegin > posEnd || posBegin < 0 || posEnd < 0 {
		return "", fmt.Errorf("Ill-formatted e-mail address: %s", src)
	}
	src = src[posBegin : posEnd+1]

	if strings.Count(src, "@") > 1 {
		return "", fmt.Errorf("Ill-formatted e-mail address: %s", src)
	}

	return src[1 : len(src)-1], nil
}
