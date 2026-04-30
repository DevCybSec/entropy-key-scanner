package analyzer

import "regexp"

type Rule struct {
	ID          string
	Description string
	Pattern     *regexp.Regexp
	Severity    string
}

// ScannerRules contiene la lista de patrones conocidos
var ScannerRules = []Rule{
	{
		ID:          "SEC-001",
		Description: "AWS Access Key ID detectada",
		Pattern:     regexp.MustCompile(`(A3T[A-Z0-9]|AKIA|AGPA|AIDA|AROA|AIPA|ANPA|ANVA|ASIA)[A-Z0-9]{16}`),
		Severity:    "CRITICAL",
	},
	{
		ID:          "SEC-002",
		Description: "GitHub Personal Access Token detectado",
		Pattern:     regexp.MustCompile(`ghp_[a-zA-Z0-9]{36}`),
		Severity:    "CRITICAL",
	},
	{
		ID:          "SEC-003",
		Description: "Private Key de SSH/RSA detectada",
		Pattern:     regexp.MustCompile(`-----BEGIN [A-Z ]+ PRIVATE KEY-----`),
		Severity:    "HIGH",
	},
	{
		ID:          "SEC-004",
		Description: "Posible Slack Webhook detectado",
		Pattern:     regexp.MustCompile(`https://hooks\.slack\.com/services/T[a-zA-Z0-9_]+/B[a-zA-Z0-9_]+/[a-zA-Z0-9_]+`),
		Severity:    "HIGH",
	},
}
