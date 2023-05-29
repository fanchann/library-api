package environments

import (
	"os"

	"github.com/joho/godotenv"

	"fanchann/library/pkg/utils"
)

var (
	Driver   = os.Getenv("DB_DRIVER")
	Username = os.Getenv("DB_AUTH_USERNAME")
	Password = os.Getenv("DB_AUTH_PASSWORD")
	Db_name  = os.Getenv("DB_NAME")
	Db_url   = os.Getenv("DB_URL")
	Db_port  = os.Getenv("DB_PORT")
)

func LoadEnv() {
	errEnv := godotenv.Load("./pkg/environments/.env")
	utils.LogErrorWithPanic(errEnv)
}
