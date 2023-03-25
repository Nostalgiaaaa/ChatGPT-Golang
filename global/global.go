package global

var (
	OpenAIKey string
	Config    SystemConfig
)

type SystemConfig struct {
	System struct {
		OpenAIKey      string
		Address        string
		AuthSecretKey  string
		HttpsProxy     string
		ReverseProxy   string
		SocksHost      string
		SocksPort      string
		OpenAPIBaseURL string
		DatabasePath   string
	}
}
