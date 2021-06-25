package main

import (
	"github.com/sirupsen/logrus"
	config "kafkalog/logagent/config"
	server "kafkalog/logagent/server"
	drive "kafkalog/logagent/drive"
)

func main()  {
	configObj :=config.GetConfig()
	err := server.Init(configObj)
	if err != nil{
		logrus.Error("kafka连接失败err:",err)
		return
	}

	err = server.InitTail(configObj)
	if err != nil{
		logrus.Error("读取日志失败err:",err)
		return
	}
	drive.Run(configObj)

	//ctx,cancel := server.IniEtcd(configObj)

	//consumer, err := sarama.NewConsumer([]string{configObj.KafkaConfig.Addrress},nil)
	//if err != nil{
	//	logrus.Error("消费者连接失败err:",err)
	//	return
	//}
	//partitionList ,err := consumer.Partitions(configObj.KafkaConfig.Topic)
	//if err != nil{
	//	logrus.Error("获取topic列表失败err:",err)
	//	return
	//}
	//
	//var wg sync.WaitGroup
	//for partition := range partitionList{
	//	pc , err := consumer.ConsumePartition(configObj.KafkaConfig.Topic,int32(partition),sarama.OffsetNewest)
	//	if err != nil{
	//		logrus.Error("分区读取失败失败err:",err)
	//		return
	//	}
	//	wg.Add(1)
	//	defer pc.AsyncClose()
	//	go func(sarama.PartitionConsumer) {
	//		for msg := range pc.Messages(){
	//			fmt.Printf("获取数据Partition:%d-----Offset:%d-----Key:%s-----Value:%s\n",msg.Partition,msg.Offset,msg.Key,msg.Value)
	//		}
	//	}(pc)
	//}
	//
	//wg.Wait()

}
