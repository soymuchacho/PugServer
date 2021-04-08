package PugCommon

import (
	"fmt"
	"sync"
)

type Registry interface {
	RegisterWithKeep() error
	UnRegister() error
}

type ServiceInfo struct {
	ServiceName string `json:"service_name"`
	ServiceType string `json:"service_type"`
	ServiceAddr string `json:"service_addr"`
	Version     string `json:"version"`
	Load        int64  `json:"Load"`
}
