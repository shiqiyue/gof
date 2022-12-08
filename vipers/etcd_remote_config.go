package vipers

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/shiqiyue/gof/loggers"
	"go.uber.org/zap"
	"io"
	"io/ioutil"
	"strings"
	"sync"
	"time"

	"github.com/spf13/viper"
)

type EtcdRemoteConfig struct {
	Mutex       sync.Mutex
	ClientCache map[string]*clientv3.Client
}

// 获取配置信息
func (c *EtcdRemoteConfig) Get(rp viper.RemoteProvider) (io.Reader, error) {
	client, err := c.newClient(rp)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	paths := rp.Path()
	ps := strings.Split(paths, ";")
	bf := bytes.Buffer{}
	for _, p := range ps {
		resp, err := client.Get(ctx, p)
		if err != nil {
			return nil, err
		}
		bf.Write(resp.Kvs[0].Value)
	}
	return bytes.NewReader(bf.Bytes()), nil
}

// 监听配置信息
func (c *EtcdRemoteConfig) Watch(rp viper.RemoteProvider) (io.Reader, error) {
	client, err := c.newClient(rp)
	if err != nil {
		return nil, err
	}
	paths := rp.Path()
	ps := strings.Split(paths, ";")
	var watchCh = make(chan bool, 1)
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()
	for _, p := range ps {
		go func(p string) {
			<-client.Watch(ctx, p)
			watchCh <- true
		}(p)
	}
	<-watchCh
	return c.Get(rp)
}

// 监听配置信息
func (c *EtcdRemoteConfig) WatchChannel(rp viper.RemoteProvider) (<-chan *viper.RemoteResponse, chan bool) {
	rr := make(chan *viper.RemoteResponse)
	stop := make(chan bool)
	go func() {
		client, err := c.newClient(rp)
		if err != nil {
			loggers.Error(context.Background(), "获取etcd客户端异常", zap.Error(err))
			panic(err)
		}
		paths := rp.Path()
		ps := strings.Split(paths, ";")
		ch := make(chan bool, 1)
		ctx, cancelFunc := context.WithCancel(context.Background())
		defer cancelFunc()
		for _, p := range ps {
			go func(p string) {
				<-client.Watch(ctx, p)
				ch <- true
			}(p)
		}
		for {
			select {
			case <-stop:
				return
			case <-ch:
				reader, err := c.Get(rp)
				if err != nil {
					loggers.Error(context.Background(), "获取数据异常", zap.Error(err))
					panic(err)
				}
				bs, err := ioutil.ReadAll(reader)
				if err != nil {
					loggers.Error(context.Background(), "获取数据异常", zap.Error(err))
					panic(err)
				}
				rr <- &viper.RemoteResponse{
					Value: bs,
				}

			}
		}
	}()

	return rr, stop
}

// 新建etcd客户端
func (c *EtcdRemoteConfig) newClient(rp viper.RemoteProvider) (*clientv3.Client, error) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	if c.ClientCache == nil {
		c.ClientCache = make(map[string]*clientv3.Client, 0)
	}
	client := c.ClientCache[rp.Endpoint()]
	if client != nil {
		return client, nil
	}
	sec := rp.SecretKeyring()
	etcdCert := &etcdCertificate{}
	err := json.Unmarshal([]byte(sec), etcdCert)
	if err != nil {
		return nil, err
	}
	client, err = clientv3.New(clientv3.Config{
		Endpoints:   strings.Split(rp.Endpoint(), ","),
		DialTimeout: 2 * time.Second,
		Username:    etcdCert.Username,
		Password:    etcdCert.Password,
	})
	if err != nil {
		return nil, err
	}
	c.ClientCache[rp.Endpoint()] = client

	return client, nil
}
