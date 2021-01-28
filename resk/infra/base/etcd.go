package base

import (
	"context"
	"time"

	"go.etcd.io/etcd/clientv3"
	"resk.micro/infra"
)

var clt *clientv3.Client.Client

type EtcdStarter struct {
	starter infra.BaseStarter
	cli *clientv3.Client
}

func (etcdclt *EtcdStarter) Init(ctx infra.StarterContext) {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{""},
		DialTimeout: 2 * time.Second,
	})

	if err != nil {
		Logger().Fatal("etcd connect failed")
	}
	clt = client
}

/*func (etcdclt *EtcdStarter) Setup(ctx infra.StarterContext) {
	ticker := time.NewTicker(time.Second * time.Duration(11))
	kv := clientv3.NewKV(clt)
	lease := clientv3.NewLease(clt)
	var curleaseid clientv3.LeaseID = 0
	for  {
		if curleaseid == 0 {
			grantResponse, err := lease.Grant(context.TODO(), 10)

			key := grantResponse.ID
		}
	}
}*/

func (etcdclt *EtcdStarter) Start(ctx infra.StarterContext) {
	panic("implement me")
}

func (etcdclt *EtcdStarter) StartBlocking() bool {
	panic("implement me")
}

func (etcdclt *EtcdStarter) Stop(ctx infra.StarterContext) {
	panic("implement me")
}

func GetEtcdClt(ctx context.Context) *clientv3.Client {

	return clt
}

