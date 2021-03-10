package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

var (
	Conf *ini.File

	// 运行模式
	RunMode string
	// 端口
	HTTPPort int
	// 读取超时时间
	ReadTimeout time.Duration
	// 写超时时间
	WriteTimeout time.Duration
	// 条数
	PageSize int
	// JWT
	JwtSecret string
)

func init() {
	// 错误类型变量
	var err error
	// 读取ini文件
	Conf, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("解析ini文件:%v", err)
	}
	// 加载运行模式
	LoadBase()
	// 加载服务
	LoadServer()
	// 加载app配置
	LoadApp()
}
func LoadBase() {
	RunMode = Conf.Section("").Key("RUN_MODE").MustString("debug")
}
func LoadServer() {
	sec, err := Conf.GetSection("server")
	if err != nil {
		log.Fatalf("获取'server':%v", err)
	}
	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}
func LoadApp() {
	sec, err := Conf.GetSection("app")
	if err != nil {
		log.Fatalf("获取'app': %v", err)
	}

	JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
}
