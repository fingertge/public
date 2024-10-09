// ***********************************************************************************************
// ***                                     G O L A N D                                         ***
// ***********************************************************************************************
// * Auth: ColeCai
// * Date: 2023/10/09 17:25:42
// * Proj: public
// * Pack: service
// * File: etcd.go
// *----------------------------------------------------------------------------------------------
// * Overviews:
// *----------------------------------------------------------------------------------------------
// * Functions:
// * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

package etcd

import (
	"context"
	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
	"time"
)

type IService interface {
	AddService(key, val []byte)
	DelService(key, val []byte)
}

type EtcdClient struct {
	servicHandler IService
	client        *clientv3.Client
	keepAliveChan <-chan *clientv3.LeaseKeepAliveResponse
	leaseID       clientv3.LeaseID
	opTimeout     time.Duration
}

func NewEtcdClient(endpoints []string, opTimeOut time.Duration, handler IService) (*EtcdClient, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: opTimeOut,
		DialOptions: []grpc.DialOption{grpc.WithBlock()},
	})
	if err != nil {
		return nil, err
	}
	ser := &EtcdClient{
		client:        cli,
		servicHandler: handler,
		opTimeout:     opTimeOut,
	}
	return ser, nil
}

func (s *EtcdClient) PutKeyWithLease(key, val string, lease int64) error {
	//设置租约时间
	gCtx, gCancle := context.WithTimeout(context.Background(), s.opTimeout)
	defer gCancle()
	resp, err := s.client.Grant(gCtx, lease)
	if err != nil {
		return err
	}
	//注册服务并绑定租约
	pCtx, pCancle := context.WithTimeout(context.Background(), s.opTimeout)
	defer pCancle()
	_, err = s.client.Put(pCtx, key, val, clientv3.WithLease(resp.ID))
	if err != nil {
		return err
	}
	//设置续租 定期发送需求请求
	kCtx, kCancle := context.WithTimeout(context.Background(), s.opTimeout)
	defer kCancle()
	leaseRespChan, err := s.client.KeepAlive(kCtx, resp.ID)

	if err != nil {
		return err
	}
	s.leaseID = resp.ID
	s.keepAliveChan = leaseRespChan
	go s.lisitenLeaseRespChan()
	return nil
}

func (s *EtcdClient) GetKeyWithPrefix(key string) error {
	gCtx, gCancle := context.WithTimeout(context.Background(), s.opTimeout)
	defer gCancle()
	rsp, err := s.client.Get(gCtx, key, clientv3.WithPrefix())
	if err != nil {
		return err
	}
	for _, v := range rsp.Kvs {
		s.servicHandler.AddService(v.Key, v.Value)
	}
	return nil
}

func (s *EtcdClient) GetKey(key string) error {
	gCtx, gCancle := context.WithTimeout(context.Background(), s.opTimeout)
	defer gCancle()
	rsp, err := s.client.Get(gCtx, key)
	if err != nil {
		return err
	}
	for _, v := range rsp.Kvs {
		s.servicHandler.AddService(v.Key, v.Value)
	}
	return nil
}

func (s *EtcdClient) PutKey(key, val string) error {
	gCtx, gCancle := context.WithTimeout(context.Background(), s.opTimeout)
	defer gCancle()
	_, err := s.client.Put(gCtx, key, val)
	return err
}

func (s *EtcdClient) lisitenLeaseRespChan() {
	for range s.keepAliveChan {

	}
}

func (s *EtcdClient) Close() error {
	//撤销租约
	ctx, cancle := context.WithTimeout(context.Background(), s.opTimeout)
	defer cancle()
	if _, err := s.client.Revoke(ctx, s.leaseID); err != nil {
		return err
	}
	return s.client.Close()
}

func (s *EtcdClient) Watcher(prefix string) {
	rch := s.client.Watch(context.Background(), prefix, clientv3.WithPrefix(), clientv3.WithPrevKV())
	for wresp := range rch {
		for _, ev := range wresp.Events {
			switch ev.Type {
			case mvccpb.PUT: //修改或者新增
				s.servicHandler.AddService(ev.Kv.Key, ev.Kv.Value)
			case mvccpb.DELETE: //删除
				s.servicHandler.DelService(ev.PrevKv.Key, ev.PrevKv.Value)
			}
		}
	}
}
