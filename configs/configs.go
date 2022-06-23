package configs

import (
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"os"
)

var Cfg = config{}

type config struct {
	DefaultConfig
	Cloudinary
}
type Cloudinary struct {
	CloudName    string `env:"CLOUDINARY_CLOUD_NAME"`
	ApiKey       string `env:"CLOUDINARY_API_KEY"`
	ApiSecret    string `env:"CLOUDINARY_API_SECRET"`
	UploadFolder string `env:"CLOUDINARY_UPLOAD_FOLDER"`
}

type DefaultConfig struct {
	GinMode     string `env:"GIN_MODE" envDefault:"debug"`
	DatabaseUrl string `env:"DATABASE_URL"`
	UploadTo    string `env:"UPLOAD_TO"`
}

func (c *config) show() {
	fmt.Printf("💬 Env runtime set-> %+v\n", c)
}

func init() {
	if err := env.Parse(&Cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}

	if "debug" == Cfg.GinMode {
		err := godotenv.Load(".env." + Cfg.GinMode)
		if err != nil {
		}

		Cfg.DatabaseUrl = os.Getenv("DATABASE_URL")
		Cfg.UploadTo = os.Getenv("UPLOAD_TO")
	}

	Cfg.show()
}
