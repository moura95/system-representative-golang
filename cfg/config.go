package cfg

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	AWSAccessKeyID       string        `mapstructure:"AWS_ACCESS_KEY_ID"`
	AWSSecretAccessKey   string        `mapstructure:"AWS_SECRET_ACCESS_KEY"`
	AWSBucketName        string        `mapstructure:"AWS_BUCKET_NAME"`
	AWSRegion            string        `mapstructure:"AWS_REGION"`
	DBSource             string        `mapstructure:"DB_SOURCE"`
	HTTPServerAddress    string        `mapstructure:"HTTP_SERVER_ADDRESS"`
	TokenSymmetricKey    string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
	StripeKey            string        `mapstructure:"STRIPE_KEY"`
	SecretKey            string        `mapstructure:"SECRET_KEY"`
	StripeWebhookSecret  string        `mapstructure:"STRIPE_WEBHOOK_SECRET"`
	GinMode              string        `mapstructure:"GIN_MODE"`
	NewRelicKey          string        `mapstructure:"NEW_RELIC_KEY"`
	NewRelicAppName      string        `mapstructure:"NEW_RELIC_APP_NAME"`
	NewRelicLogForward   bool          `mapstructure:"NEW_RELIC_LOG_FORWARD"`
	SilverPlanID         string        `mapstructure:"SILVER_PLAN_ID"`
	GoldPlanID           string        `mapstructure:"GOLD_PLAN_ID"`
	PdfTurtleClient      string        `mapstructure:"PDF_TURTLE_CLIENT_URL"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigFile(".env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
