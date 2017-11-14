package configuration
import (
	"encoding/json"
	"os"
	"fmt"
)

type Configuration struct {
	Host	string
	Username	string
	Password	string
	Database	string
}
var Config Configuration

func ReadConfig(path string)  {
	file, _ := os.Open(path)
	decoder := json.NewDecoder(file)
	Config = Configuration{}
	err := decoder.Decode(&Config)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(Config.Host) // output: [UserA, UserB]
}