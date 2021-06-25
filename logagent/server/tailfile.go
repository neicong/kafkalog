package server

import (
	"github.com/hpcloud/tail"
	"github.com/sirupsen/logrus"
	config "kafkalog/logagent/config"
)

var(
	Tailas *tail.Tail
)
func InitTail(configObj *config.Config)(err error)  {


	tailConfig := tail.Config{
		ReOpen: true,
		Follow: true,
		Location: &tail.SeekInfo{Offset: 0,Whence: 2},
		MustExist: false,
		Poll: true,
	}

	Tailas , err = tail.TailFile(configObj.CollectConfig.LogfilePath,tailConfig)
	if err !=nil{
		logrus.Error("读取日志失败！err:",err)
		return
	}
	return


}