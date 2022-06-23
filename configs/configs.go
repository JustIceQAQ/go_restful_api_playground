package configs

import (
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"os"
)

var Cfg = config{}

type config struct {
	GinMode     string `env:"GIN_MODE" envDefault:"debug"`
	DatabaseUrl string `env:"DATABASE_URL"`
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
	}

	Cfg.show()
}
