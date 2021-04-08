package PugCommon

import (
	"context"
	"fmt"
	log "github.com/cihub/seelog"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/etcdserver/api/v3rpc/rpctypes"
	"time"
)

var (
	dialTimeout    = 5 * time.Second
	requestTimeout = 10 * time.Second
)

type ClientKeepliveChan = <-chan *clientv3.LeaseKeepAliveResponse
type LeaseInfo struct {
	Lease_id        int64
	Ch_alive_status ClientKeepliveChan
}

type EtcdClient struct {
	cli *clientv3.Client
}

func NewEtcdClient(endpoints ...string) (*EtcdClient, error) {
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
			log.Errorf("ctx is canceled by another routine: %v", err)
		case context.DeadlineExceeded:
			log.Errorf("ctx is attached with a deadline is exceeded: %v", err)
		case rpctypes.ErrEmptyKey:
			log.Errorf("client-side error: %v", err)
		default:
			log.Errorf("bad cluster endpoints, which are not etcd servers: %v", err)
		}
		return err
	}
	return nil
}

func (c *EtcdClient) Get(key string) (string, error) {
	if c.cli == nil {
		return "", fmt.Errorf("Client has not connect to the server")
	}

	if len(key) == 0 {
		return "", fmt.Errorf("key is empty")
	}

	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	resp, err := c.cli.Get(ctx, key)
	cancel()

	if err != nil {
		return "", fmt.Errorf("failed to get etcd key %v , info %v", key, err)
	}

	if len(resp.Kvs) == 0 {
		return "", nil
	} else {
		result := resp.Kvs[0].Value
		return string(result), nil
	}
}

func (c *EtcdClient) GetPrefix(prefix string) (map[string]string, error) {
	if c.cli == nil {
		return nil, fmt.Errorf("Client has not connect to the server")
	}

	if len(prefix) == 0 {
		return nil, fmt.Errorf("prefix is empty")
	}

	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	resp, err := c.cli.Get(ctx, prefix, clientv3.WithPrefix(), clientv3.WithSort(clientv3.SortByKey, clientv3.SortDescend))
	cancel()

	if err != nil {
		return nil, err
	}

	if len(resp.Kvs) == 0 {
		return nil, nil
	} else {
		result := make(map[string]string)
		for _, ev := range resp.Kvs {
			log.Debugf("getprefix get %v=>%v reversion %v", ev.Key, ev.Value, resp.Header.Revision)
			result[string(ev.Key)] = string(ev.Value)
		}
		return result, nil
	}
}

func (c *EtcdClient) PutWithLease(key string, value string, ttl int64) (ch_delete_lease chan bool, leaseInfo LeaseInfo, err error) {
	if c.cli == nil {
		return nil, LeaseInfo{}, fmt.Errorf("Client has not connected the server")
	}

	if len(key) == 0 || len(value) == 0 {
		return nil, LeaseInfo{}, fmt.Errorf("key/value is empty")
	}

	if ttl <= 0 {
		return nil, LeaseInfo{}, fmt.Errorf("lease ttl <= 0")
	}

	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	resp, err := c.cli.Grant(ctx, ttl)
	cancel()

	if err != nil {
		return nil, LeaseInfo{}, err
	}

	log.Debugf("begin keepalive")
	keep_ctx, keep_cancel := context.WithCancel(context.Background())
	ch, kaerr := c.cli.KeepAlive(keep_ctx, resp.ID)
	if kaerr != nil {
		return nil, LeaseInfo{}, kaerr
	}

	log.Debugf("begin put")
	if err := c.Put(key, value, resp.ID); err != nil {
		return nil, LeaseInfo{}, err
	}

	ch_close := make(chan bool)

	go func() {
		<-ch_close
		keep_cancel()
		if err := c.deleteLease(resp.ID); err != nil {
			log.Errorf("delete lease %v error", resp.ID)
			return
		}
	}()

	return ch_close, LeaseInfo{Lease_id: int64(resp.ID), Ch_alive_status: ch}, nil
}

func (c *EtcdClient) deleteLease(lease_id clientv3.LeaseID) error {
	if c.cli == nil {
		return fmt.Errorf("Client has not connected to the server")
	}

	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	_, err := c.cli.Revoke(ctx, lease_id)
	cancel()

	if err != nil {
		return err
	}

	return nil
}
