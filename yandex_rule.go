package emailNormalizer

import "strings"

type YandexRule struct {
}

func (rule *YandexRule) processUsername(username string) string {
	return strings.Replace(username, "+", "", -1)
}

func (rule *YandexRule) processDomain(domain string) string {
	return domain
}