package global

// Config 所有配置
type Config struct {
	AesKey    string `mapstructure:"aeskey" yaml:"aeskey"`
	JWTKey    string `mapstructure:"jwtkey" yaml:"jwtkey"`
	Port      int    `mapstructure:"port"   yaml:"port"`
	Mysql     Mysql  `mapstructure:"mysql"  yaml:"mysql"`
	SSOServer string `mapstructure:"sso"    yaml:"sso"`
	Admin     Admin  `mapstructure:"admin"  yaml:"admin"`
	Static    string `mapstructure:"static"`
}

// Mysql 连接配置
type Mysql struct {
	Username string `mapstructure:"username" yaml:"username"`
	Password string `mapstructure:"password" yaml:"password"`
	Path     string `mapstructure:"path"     yaml:"path"`
	Dbname   string `mapstructure:"dbname"   yaml:"dbname"`
	Parm     string `mapstructure:"parm"     yaml:"parm"`
}

// Admin 默认管理员账号密码
type Admin struct {
	Username string `mapstructure:"username" yaml:"username"`
	Password string `mapstructure:"password" yaml:"password"`
}
