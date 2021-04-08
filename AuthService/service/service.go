package service

import (
	"PugCommon"
	"encoding/json"
	"fmt"
	log "github.com/cihub/seelog"
	"sync"
)

type ServiceRegistry struct {
	Cli      *PugCommon.EtcdClient
	Lock     sync.Mutex
	Ch_del   chan bool
	Lease    *PugCommon.LeaseInfo
	EndPoint []string
}

var SerReg *ServiceRegistry

func NewService(endpoints ...string) (*ServiceRegistry, error) {
	if len(endpoints) == 0 {
		return nil, fmt.Errorf("endpoints is empty")
	}

	c, err := PugCommon.NewEtcdClient(endpoints...)
	if err != nil {
		return nil, err
	}

	return &ServiceRegistry{
		Cli:      c,
		EndPoint: endpoints,
	}, nil
}

func (r *ServiceRegistry) RegisterWithKeep(serinfo PugCommon.ServiceInfo, ttl int64) error {
	r.Lock.Lock()

	if r.Cli == nil {
		r.Lock.Unlock()
		return fmt.Errorf("etcd client is nil")
	}

	key := serinfo.ServiceType + "/" + serinfo.ServiceName
	serinfo_str, err := json.Marshal(serinfo)
	if err != nil {
		r.Lock.Unlock()
		return err
	}
	ch_del, leaseinfo, err := r.Cli.PutWithLease(key, string(serinfo_str[:]), ttl)
	if err != nil {
		r.Lock.Unlock()
		return err
	}

	r.Ch_del = ch_del
	r.Lease = &leaseinfo
	r.Lock.Unlock()

	log.Debugf("register success")
	for {
		select {
		case msg, ok := <-r.Lease.Ch_alive_status:
			if !ok {
				log.Debugf("etcd keepalive was closed")
				return nil
			} else {
				log.Debugf("recv etcd keep service alive message: %v", msg)
			}
		}
	}
	return nil
}

func (r *ServiceRegistry) UnRegister() error {
	r.Lock.Lock()
	defer r.Lock.Unlock()

	if r.Lease != nil {
		r.Ch_del <- true
	}
	return nil
}
