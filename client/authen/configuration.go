package authen

// Config describes a 2-legged OAuth2 flow, with both the
// client application information and the server's endpoint URLs.
type Configuration struct {
	// ClientID is the application's ID.
	ClientID string

	// ClientSecret is the application's secret.
	ClientSecret string

	// TokenURL is the resource server's token endpoint
	// URL. This is a constant specific to each server.
	TokenURL string
}

func NewConfiguration(ClientID string, ClientSecret string, TokenURL string) *Configuration {
	config := &Configuration{
		ClientID:     ClientID,
		ClientSecret: ClientSecret,
		TokenURL:     TokenURL,
	}
	return config
}
