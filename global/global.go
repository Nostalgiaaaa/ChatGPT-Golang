package global

import "github.com/spf13/viper"

var (
	OpenAIKey   string
	Config      System
	ViperConfig *viper.Viper
)

type System struct {
	OpenAIKey      string
	Address        string
	AuthSecretKey  string
	HttpsProxy     string
	ReverseProxy   string
	SocksHost      string
	SocksPort      string
	OpenAPIBaseURL string
}
