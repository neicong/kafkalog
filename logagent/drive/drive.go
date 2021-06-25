package drive

import (
	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
	server "kafkalog/logagent/server"
	"strings"
	"time"
	config "kafkalog/logagent/config"
)
func Run(configObj *config.Config) {

	for  {
		line , err :=<-server.Tailas.Lines
		if !err {
			logrus.Error("数据读取失败！err:",err)
			time.Sleep(time.Second)
			continue
		}
		if len(strings.Trim(line.Text,"\r")) == 0{
			continue
		}
		msg :=&sarama.ProducerMessage{}
		msg.Topic = configObj.KafkaConfig.Topic
		msg.Value = sarama.StringEncoder(line.Text)
		server.MsgChan <- msg
	}

}
