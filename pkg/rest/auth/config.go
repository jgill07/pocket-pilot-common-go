package auth

type Config struct {
	IssuerUrl string   `mapstructure:"ISSUER_URL"` // This requires a trailing slash
	Audiences []string `mapstructure:"AUDIENCES"`
}
