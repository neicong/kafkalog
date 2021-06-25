package server

import (
	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
	config "kafkalog/logagent/config"
)

var (
	Client sarama.SyncProducer
	MsgChan chan *sarama.ProducerMessage
)
func Init(configObj *config.Config)(err error) {
	kafkaConf := sarama.NewConfig()
	kafkaConf.Producer.RequiredAcks = sarama.WaitForAll
	kafkaConf.Producer.Partitioner = sarama.NewRandomPartitioner
	kafkaConf.Producer.Return.Successes = true
	Client ,err = sarama.NewSyncProducer([]string{configObj.KafkaConfig.Addrress},kafkaConf)
	if err != nil{
		logrus.Error("连接失败err:",err)
		return
	}
	MsgChan = make(chan *sarama.ProducerMessage , configObj.KafkaConfig.MsgChanSize)
	go sendMsg()
	return
}

func sendMsg()  {
	for{
		select{
		case msg := <-MsgChan:
		pid , offset , ok := Client.SendMessage(msg)
		if ok !=nil{
			logrus.Error("提交kafka失败err:",ok)

		}
		logrus.Infof("kafka提交成功，pid：%v offset:%v",pid,offset)

		}
	}
	
}
