// ***********************************************************************************************
// ***                                     G O L A N D                                         ***
// ***********************************************************************************************
// * Auth: ColeCai
// * Date: 2023/10/12 01:00:53
// * Proj: public
// * Pack: redis
// * File: redis_test.go
// *----------------------------------------------------------------------------------------------
// * Overviews:
// *----------------------------------------------------------------------------------------------
// * Functions:
// * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

package redis

import (
	"context"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
	"os"
	"testing"
	"time"
)

func TestConnredis(t *testing.T) {
	options := &redis.RingOptions{
		Addrs: map[string]string{
			"redis-4501": "192.168.31.54:4501",
			"redis-4502": "192.168.31.54:4502",
			"redis-4503": "192.168.31.54:4503",
		},
	}
	ring := redis.NewRing(options)
	ring.Set(context.Background(), "aaaaa", "192.168.31.54:4501", 100*time.Second)
	ring.Set(context.Background(), "bbbbb", "192.168.31.54:4501", 100*time.Second)
	ring.Set(context.Background(), "ccc", "192.168.31.54:4501", 100*time.Second)
	ring.Set(context.Background(), "eee", "192.168.31.54:4501", 100*time.Second)
	clusterOpt := &redis.ClusterOptions{
		Addrs:    []string{"192.168.31.28:6000", "192.168.31.28:6001", "192.168.31.28:6002"},
		Username: "admin",
		Password: "19931207cjj@",
	}
	cluster := redis.NewClusterClient(clusterOpt)
	ret := cluster.Set(context.Background(), "clustertestr", "testtt", 1000*time.Second)
	s, e := ret.Result()
	fmt.Printf("ret:%s, err:%+v\n", s, e)
}

func TestSentinel(t *testing.T) {
	opt := &redis.FailoverOptions{
		MasterName: "mymaster",
		SentinelAddrs: []string{
			"192.168.31.28:8000",
			"192.168.31.28:8001",
			"192.168.31.28:8002",
		},
		SentinelUsername: "admin",
		SentinelPassword: "19931207cjj@",
		Username:         "admin",
		Password:         "19931207cjj@",
	}
	sen := redis.NewFailoverClient(opt)
	ret := sen.Set(context.Background(), "test", "balue", 1000*time.Second)
	s, e := ret.Result()
	fmt.Printf("ret:%s, err:%+v\n", s, e)
}

func TestReadClusterConf(t *testing.T) {
	f, err := os.ReadFile("./redis.yaml")
	if err != nil {
		t.Logf("readfile failed:%+v", err)
		return
	}
	conf := new(redis.ClusterOptions)
	err = yaml.Unmarshal(f, conf)
	if err != nil {
		t.Logf("unmarshal failed:%+v", err)
		return
	}
	t.Logf("conf:%+v", conf)
}
func TestWriteRedisCOnf(t *testing.T) {
	type Conf struct {
		Fn   func() `yaml:f,omitempty`
		Name string
	}
}

func TestViper(t *testing.T) {
	v := viper.New()
	v.SetConfigFile("./redis.yaml")
	v.SetConfigName("redis")
	v.AddConfigPath(".")
	v.SetDefault("clientname", "caijunjun")
	err := v.ReadInConfig()
	if err != nil {
		t.Logf("read failed:%+v", err)
		return
	}
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		t.Logf("name:%s", in.Name)
		t.Logf("op:%s", in.String())
		conf := new(redis.ClusterOptions)
		v.Unmarshal(conf)
		t.Logf("E1:%+v", conf)
	})

	conf := new(redis.ClusterOptions)
	v.Unmarshal(conf)
	t.Logf("E:%+v", conf)

	select {}
}
