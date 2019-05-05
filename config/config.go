package setting

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

const Product = "prod"
const Dev = "dev"

var Config = new(Conf)

type Conf struct {
	RunMode string //运行模式
	Database    struct {
		Dialect      string
		Database     string
		User         string
		Password     string
		Charset      string
		Host         string
		Port         int
		MaxidleConns int
		MaxopenConns int
	}
	RedisConfig struct {
		Url       string //Redis 数据库地址
		Host      string //主机地址
		Port      int    //端口
		Password  string //密码
		MaxIdle   int    //最大空闲数
		MaxActive int    //最大连接数
	}
}

func init() {
	data, err := ioutil.ReadFile("./config.yml")
	CheckError("读取配置文件出错", err)
	err = yaml.Unmarshal(data, Config)
	CheckError("解析配置文件出错", err)
}

func CheckError(text string, err error) {
	if err != nil {
		fmt.Println(text, err.Error())
		return
		//os.Exit(-1)
	}
}
