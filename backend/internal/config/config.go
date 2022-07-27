package config

type Config struct {
	Server  Server
	MongoDB MongoDB
}

type Server struct {
	Port string `default:":8080"`
}

type MongoDB struct {
	URI      string `default:"mongodb://localhost:27017"`
	User     string `default:"admin"`
	Password string `default:"admin"`
	DB       string `default:"emails"`
}
