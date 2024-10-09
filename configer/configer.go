// *********************************************************************************************************************
// ***                                        CONFIDENTIAL --- CUSTOM STUDIOS                                        ***
// *********************************************************************************************************************
// * Auth: ColeCai                                                                                                     *
// * Date: 2024/03/03 16:04:03                                                                                         *
// * Proj: work                                                                                                        *
// * Pack: configer                                                                                                    *
// * File: configer.go                                                                                                 *
// *-------------------------------------------------------------------------------------------------------------------*
// * Overviews:                                                                                                        *
// *-------------------------------------------------------------------------------------------------------------------*
// * Functions:                                                                                                        *
// * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - *

package configer

import (
	"context"
	"github.com/pkg/errors"
	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
	"time"
)

type IConfiger interface {
	GetKeys() []string
	SetConfig(key, value []byte)
}

type Configer struct {
	client  *clientv3.Client
	modules []IConfiger
}

func MustInit(addrs []string, module ...IConfiger) {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   addrs,
		DialTimeout: 5 * time.Second,
		DialOptions: []grpc.DialOption{
			grpc.WithBlock(),
		},
	})
	if err != nil {
		panic(err)
	}
	configer := &Configer{client: client, modules: module}
	configer.onInit()
}

func (c *Configer) watcher(module IConfiger, key string) {
	recvs := c.client.Watch(context.Background(), key)
	for recv := range recvs {
		for _, ev := range recv.Events {
			switch ev.Type {
			case mvccpb.PUT:
				module.SetConfig(ev.Kv.Key, ev.Kv.Value)
			}
		}
	}
}

func (c *Configer) get(module IConfiger, key string) {
	recv, err := c.client.Get(context.Background(), key)
	if err != nil {
		panic(err)
	}
	if len(recv.Kvs) != 1 {
		panic(errors.Errorf("not found %v config.", key))
	}
	module.SetConfig(recv.Kvs[0].Key, recv.Kvs[0].Value)
}

func (c *Configer) onInit() {
	for _, v := range c.modules {
		keys := v.GetKeys()
		for _, key := range keys {
			c.get(v, key)
			go c.watcher(v, key)
		}
	}
}
