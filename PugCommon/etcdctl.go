package PugCommon

import (
	"context"
	"fmt"
	log "github.com/cihub/seelog"
	"github.com/coreos/etcd/clientv3"
)

var (
	dialTimeout    = 5 * time.Second
	requestTimeout = 10 * time.Second
)

type EtcdClient struct {
	cli *clientv3.Client
}

func New(endpoints ...string) (*EtcdClient, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: dialTimeout,
	})

	if err != nil {
		return nil, err
	}

	c := &EtcdClient{}
	c.cli = cli

	return c, nil
}

func (c *EtcdClient) Put(key string, value string, lease_id ...clientv3.LeaseID) error {
	if c.cli == nil {
		return fmt.Errorf("client has not connect to the etcd server")
	}

	if len(key) == 0 {
		return fmt.Errorf("key is empty")
	}

	var err error
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	if len(lease_id) > 0 && int64(lease_id[0]) > 0 {
		_, err = c.cli.Put(ctx, key, value, clientv3.WithLease(lease_id[0]))
	} else {
		_, err = c.cli.Put(ctx, key, value)
	}

	cancel()

	if err != nil {
		switch err {
		case context.Canceled:
			log.Error("ctx is canceled by another routine: %v", err)
		case context.DealineExceeded:
			log.Error("ctx is attached with a deadline is exceeded: %v", err)
		case rpctypes.ErrEmptyKey:
			log.Error("client-side error: %v", err)
		default:
			log.Error("bad cluster endpoints, which are not etcd servers: %v", err)
		}
		return err
	}
	return nil
}

func (c *EtcdClient) Get(key string) (string, error) {
	if c.cli == nil {
		return "", fmt.Errorf("Client has not connect to the server")
	}
}
