package emailnormalizer

import "strings"

// Rule for Apple
type AppleRule struct {
}

func (rule *AppleRule) processUsername(username string) string {
	return strings.Replace(username, "+", "", -1)
}

func (rule *AppleRule) processDomain(domain string) string {
	return domain
}
