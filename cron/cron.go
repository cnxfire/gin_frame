package cron

import (
	"github.com/robfig/cron"
	"go_frame/config"
)
var cronMap = map[string]func(){}
func init()  {
	if setting.Config.RunMode != setting.Dev {
		cronMap["1 * * * * *"] = yesterDaySale
	}
}

func New() *cron.Cron{
	c:= cron.New()
	for spec,cmd :=range cronMap{
		c.AddFunc(spec,cmd)
	}
	return c
}
