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
	if src[0] != '<' || src[len(src)-1] != '>' {
		return "", fmt.Errorf("Ill-formatted e-mail address: %s", src)
	}

	if strings.Count(src, "@") > 1 {
		return "", fmt.Errorf("Ill-formatted e-mail address: %s", src)
	}

	return src[1 : len(src)-1], nil
}
