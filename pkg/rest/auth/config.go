package auth

type Config struct {
	IssueUrl  string   `mapstructure:"ISSUE_URL"` // This requires a trailing slash
	Audiences []string `mapstructure:"AUDIENCES"`
}
