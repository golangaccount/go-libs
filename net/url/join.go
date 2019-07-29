package url

import (
	"strings"

	gstrings "github.com/golangaccount/go-libs/strings"
)

func Join(sub ...string) string {
	for i, item := range sub {
		if !gstrings.IsEmptyOrWhite(item) {
			sub[i] = strings.TrimSuffix(strings.TrimPrefix(item, "/"), "/")
		}
	}
	return strings.Join(sub, "/")
}
