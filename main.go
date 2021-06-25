package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)
func main()  {


	//
	//filename := "F:\\www\\api_photo_frame\\storage\\logs\\laravel.log"
	//config := tail.Config{
	//	ReOpen: true,
	//	Follow: true,
	//	Location: &tail.SeekInfo{Offset: 0,Whence: 2},
	//	MustExist: false,
	//	Poll: true,
	//}
	//
	//tailas , err := tail.TailFile(filename,config)
	//
	//if err !=nil{
	//	fmt.Println("失败！err:",err)
	//}
	//
	//var(
	//	msg *tail.Line
	//	ok bool
	//)
	//
	//
	//for{
	//	msg,ok = <-tailas.Lines
	//	if !ok{
	//		fmt.Println("数据读取失败！")
	//		time.Sleep(time.Second)
	//	}
	//
	//	fmt.Println("msg:",msg.Text)
	//
	//
	//}
	//
	//
	//










	// 1.生产者配置
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	// 2.连接kafka
	Client ,err := sarama.NewSyncProducer([]string{"192.168.137.10:9092"},config)
	if err != nil{
		fmt.Println("连接失败err:",err)
		return
	}
	defer Client.Close()
	// 3.封装消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = "web_logs"
	msg.Value = sarama.StringEncoder("2021.6.24 kafka 消息-1")
	// 4.发送消息
	pid, offset, err :=Client.SendMessage(msg)
	if err != nil {
		fmt.Println("发送失败err:",err)
		return
	}
	fmt.Println("发送成功%v\n",pid,offset)





}


