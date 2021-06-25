package server

import (
	"context"
	"github.com/sirupsen/logrus"
	"go.etcd.io/etcd/clientv3"
	"time"
	config "kafkalog/logagent/config"
)

func IniEtcd(configObj *config.Config) (context.Context,context.CancelFunc) {


	etcdConfig := clientv3.Config{
		Endpoints:   []string{configObj.EtcdConfig.EtcdAddrress},
		DialTimeout: 10 * time.Second,
	}
	client, err := clientv3.New(etcdConfig)
	defer client.Close()
	if err != nil{
		logrus.Error("etcd连接失败err:",err)
	}
	return context.WithTimeout(context.Background() , time.Second)
	//_ , err =client.Put(CtxEtcd , "foo","123456789")
	//if err != nil{
	//	logrus.Error("etcd写入失败err:",err)
	//}
	//CancelEtcd()


}
