package config

type Config struct {
	Server  Server
	MongoDB MongoDB
}

type Server struct {
	Port string `default:":8080"`
}

type MongoDB struct {
	URI      string `default:"mongodb://localhost:27017" json:"uri,omitempty"`
	User     string `default:"admin" json:"user,omitempty"`
	Password string `default:"admin" json:"password,omitempty"`
	DB       string `default:"emails" json:"db,omitempty"`
}
