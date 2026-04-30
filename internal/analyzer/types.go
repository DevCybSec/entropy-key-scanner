package analyzer

// Job It's a file which needs to be scanned
type Job struct {
	Path string
}

// Result represents the findings found in a file
type Result struct {
	FilePath string
	Findings []string // Por ahora strings, luego usaremos nuestro SecurityFinding
	Error    error
}
