package common

var (
	LogPath  string
	LogLevel int
)

var INI_CONFIG_FILE string = "./conf/conf.ini"

func LoadConfig() {
	cfg, err := ini.LoadConfigFile(ToAbsPath(INI_CONFIG_FILE))
	if err != nil {
		fmt.Printf(err.Error())
		panic(err)
	}
	LogPath = cfg.MustValue("Log", "log_path", "./log/pugserver.log")
	LogLevel = cfg.MustInt("Log", "log_level", 1)
}
