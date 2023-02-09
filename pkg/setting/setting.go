package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

var (
	Cfg          *ini.File
	RunMode      string
	HttPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	PageSize  int
	JwtSercet string
)

func init() {
	var err error
	Cfg, err = ini.Load("config/app.ini")
	if err != nil {
		log.Fatalf("解析文件错误-app.ini:%v", err)
	}
	LoadBase()
	LoadServer()
	LoadApp()

}

func LoadBase() {
	//找到run_debug默认值没有则设置默认值
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("app.ini server 模块未找到%v", err)
	}
	//获取服务运行端口号 ip 超时时间
	HttPPort = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

func LoadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("app.ini app 模块未找到%v", err)
	}
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
	JwtSercet = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
}
