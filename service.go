package net_gnet

import (
	"context"
	"net"

	"github.com/walleframe/walle/app"
)

// GNetService implement app.Service interface
type GNetService struct {
	svr  *GNetServer
	name string
	ln   net.Listener
}

func NewService(name string, opt ...ServerOption) app.Service {
	return &GNetService{
		name: name,
		svr:  NewServer(opt...),
	}
}

func (svc *GNetService) Name() string {
	return svc.name
}
func (svc *GNetService) Init(s app.Stoper) (err error) {
	return
}
func (svc *GNetService) Start(s app.Stoper) (err error) {
	svc.svr.initNotify = make(chan error)
	go svc.svr.Run("")
	err = <-svc.svr.initNotify
	close(svc.svr.initNotify)
	svc.svr.initNotify = nil
	return
}
func (svc *GNetService) Stop() {
	svc.svr.Shutdown(context.Background())
	return
}
func (svc *GNetService) Finish() {
	return
}
