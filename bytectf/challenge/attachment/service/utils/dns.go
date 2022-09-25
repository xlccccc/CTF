package utils

import (
	"errors"
	"net"
	"sync"
	"time"
)

type dnsItem struct {
	UpdatedAt time.Time `json:"updateAt"`
	Hosts     []string  `json:"hosts"`
}

var (
	dnsMutex sync.RWMutex
	dnscache = make(map[string]dnsItem)
	errEmpty = errors.New("returns empty host")
)

func LookupIPAddr(name string, cacheTime time.Duration) ([]string, error) {
	dnsMutex.Lock()
	defer dnsMutex.Unlock()
	v := dnscache[name]
	cachehosts := v.Hosts
	if time.Since(v.UpdatedAt) < cacheTime && len(cachehosts) > 0 {
		return cachehosts, nil
	}

	hosts, err := net.LookupHost(name)
	if err != nil {
		if len(cachehosts) > 0 {
			return cachehosts, nil
		}
		return nil, err
	}
	item := dnsItem{Hosts: hosts, UpdatedAt: time.Now()}
	dnscache[name] = item
	if len(hosts) == 0 {
		return nil, errEmpty
	}
	return hosts, nil
}

func DeleteInvalidIpAddr(name string, ipAddr string) {
	dnsMutex.RLock()
	defer dnsMutex.RUnlock()
	hosts := dnscache[name].Hosts
	for index, host := range hosts {
		if host == ipAddr {
			hosts = append(hosts[:index], hosts[index+1:]...)
			item := dnsItem{Hosts: hosts, UpdatedAt: time.Now()}
			dnscache[name] = item
			return
		}
	}
}
