package general

type AppService struct {
	App           AppAccount   `json:",omitempty"`
	Route         RouteAccount `json:",omitempty"`
	Database      Database     `json:",omitempty"`
	Authorization AuthAccount  `json:",omitempty"`
	KeyData       KeyAccount   `json:",omitempty"`
}

type Database struct {
	Read  DBDetail `json:",omitempty"`
	Write DBDetail `json:",omitempty"`
}

type DBDetail struct {
	Username     string `json:",omitempty"`
	Password     string `json:",omitempty"`
	URL          string `json:",omitempty"`
	Port         string `json:",omitempty"`
	DBName       string `json:",omitempty"`
	MaxIdleConns int    `json:",omitempty"`
	MaxOpenConns int    `json:",omitempty"`
	MaxLifeTime  int    `json:",omitempty"`
	Timeout      string `json:",omitempty"`
	SSLMode      string `json:",omitempty"`
}

type AppAccount struct {
	Name         string `json:",omitempty"`
	Environtment string `json:",omitempty"`
	URL          string `json:",omitempty"`
	Port         string `json:",omitempty"`
	SecretKey    string `json:",omitempty"`
}

type AuthAccount struct {
	JWT    JWTCredential    `json:",omitempty"`
	Public PublicCredential `json:",omitempty"`
}

type JWTCredential struct {
	IsActive              bool   `json:",omitempty"`
	AccessTokenSecretKey  string `json:",omitempty"`
	AccessTokenDuration   int    `json:",omitempty"`
	RefreshTokenSecretKey string `json:",omitempty"`
	RefreshTokenDuration  int    `json:",omitempty"`
}

type PublicCredential struct {
	SecretKey string `json:",omitempty"`
}

type RouteAccount struct {
	Methods []string `json:",omitempty"`
	Headers []string `json:",omitempty"`
	Origins []string `json:",omitempty"`
}

type KeyAccount struct {
	User    string `json:",omitempty"`
	Company string `json:",omitempty"`
}
