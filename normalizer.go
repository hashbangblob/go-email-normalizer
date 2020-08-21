package emailNormalizer

import (
	"strings"
)

type Normalizer struct {
	rules map[string]NormalizingRule
}

func NewNormalizer() *Normalizer {
	s := map[string]NormalizingRule{
		"fastmail.com": &FastmailRule{},
		"messagingengine.com": &FastmailRule{},
		"fastmail.fm": &FastmailRule{},
		"gmail.com": &GoogleRule{},
		"googlemail.com": &GoogleRule{},
		"hotmail.com": &MicrosoftRule{},
		"live.com": &MicrosoftRule{},
		"outlook.com": &MicrosoftRule{},
		"rambler.ru": &RamblerRule{},
		"lenta.ru": &RamblerRule{},
		"autorambler.ru": &RamblerRule{},
		"myrambler.ru": &RamblerRule{},
		"ro.ru": &RamblerRule{},
	}
	return &Normalizer{rules:s}
}

func (n *Normalizer) AddRule(domain string, strategy NormalizingRule) {
	n.rules[domain] = strategy
}

func (n *Normalizer) Normalize(email string) string {
	prepared := strings.TrimSpace(email)
	prepared = strings.TrimRight(prepared, ".")
	prepared = strings.ToLower(prepared)

	parts := strings.Split(prepared, "@")

	if len(parts) != 2 {
		return prepared
	}

	username := parts[0]
	domain := parts[1]

	if rule, ok := n.rules[domain]; ok {
		return rule.processUsername(username) + "@" + rule.processDomain(domain)
	}

	return prepared
}