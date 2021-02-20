package base

import (
	"context"
	"go.uber.org/zap"
	"strings"
	"time"

	"go.etcd.io/etcd/clientv3"
	"resk.micro/infra"
)

var clt *clientv3.Client

//EtcdStarter
type EtcdStarter struct {
	starter infra.BaseStarter
	cli     *clientv3.Client
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

func (etcdclt *EtcdStarter) Setup(ctx infra.StarterContext) {

}

func (etcdclt *EtcdStarter) Start(ctx infra.StarterContext) {
	panic("implement me")
}

func (etcdclt *EtcdStarter) StartBlocking() bool {
	panic("implement me")
}

func (etcdclt *EtcdStarter) Stop(ctx infra.StarterContext) {
	panic("implement me")
}

func RegisterToETCD(name string, addr string, ttl int64) *clientv3.Client {
	ticker := time.NewTicker(time.Second * time.Duration(ttl))

	if clt == nil {
		client, _ := clientv3.New(clientv3.Config{
			Endpoints:   []string{""},
			DialTimeout: 2 * time.Second,
		})
		clt = client
	}

	var builder strings.Builder
	builder.WriteString(addr)
	builder.WriteString(name)
	key := builder.String()
	go func() {
		for {
			getResponse, err := clt.Get(context.Background(), key)
			if err != nil {
				logger.Fatal("register service error", zap.String("error", err.Error()))
			} else if getResponse.Count == 0 {
				withAlive(key, addr, ttl)
			} else {

			}

			<-ticker.C
		}
	}()

	return clt
}

func withAlive(key, addr string, ttl int64) error {
	grantResponse, err := clt.Grant(context.Background(), ttl)
	if err != nil {

		return err
	}

	_, err = clt.Put(context.Background(), key, addr, clientv3.WithLease(grantResponse.ID))
	if err != nil {
		return err
	}

	_, err = clt.KeepAlive(context.Background(), grantResponse.ID)
	if err != nil {
		return err
	}
	return nil
}
