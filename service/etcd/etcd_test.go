// ***********************************************************************************************
// ***                                     G O L A N D                                         ***
// ***********************************************************************************************
// * Auth: ColeCai
// * Date: 2023/10/09 17:37:28
// * Proj: public
// * Pack: etcd
// * File: etcd_test.go
// *----------------------------------------------------------------------------------------------
// * Overviews:
// *----------------------------------------------------------------------------------------------
// * Functions:
// * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

package etcd

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestEtcdTest(t *testing.T) {
	client, err := NewEtcdClient([]string{"192.168.31.28:2379"}, 5*time.Second, nil)
	if err != nil {
		return
	}
	//_, err = client.client.Put(context.Background(), "/no", "xxxx", clientv3.WithPrevKV())
	//fmt.Printf("err:%+v", err)

	res, err := client.client.Get(context.Background(), "/no")
	fmt.Printf("key:%s, val:%s, err:%+v\n", res.Kvs[0].Key, res.Kvs[0].Value, err)
}
