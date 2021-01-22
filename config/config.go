package config

// Config 所有配置
type Config struct {
	AesKey    string `json:"aeskey" yaml:"aeskey"`
	JWTKey    string `json:"jwtkey" yaml:"jwtkey"`
	Port      int    `json:"port"   yaml:"port"`
	Mysql     Mysql  `json:"mysql"  yaml:"mysql"`
	SSOServer string `json:"sso"    yaml:"sso"`
	Admin     Admin  `json:"admin"  yaml:"admin"`
}

// Mysql 连接配置
type Mysql struct {
	Username string `json:"username" yaml:"username"`
	Password string `json:"password" yaml:"password"`
	Path     string `json:"path"     yaml:"path"`
	Dbname   string `json:"dbname"   yaml:"dbname"`
	Parm     string `json:"parm"     yaml:"parm"`
}

// Admin 默认管理员账号密码
type Admin struct {
	Username string `json:"username" yaml:"username"`
	Password string `json:"password" yaml:"password"`
}
