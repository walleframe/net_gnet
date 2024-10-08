// Code generated by "gogen option"; DO NOT EDIT.
// Exec: "gogen option -n ServerOption -o option.server.go"
// Version: 0.0.4

package net_gnet

import (
	"net"
	"time"

	gnet "github.com/panjf2000/gnet/v2"
	"github.com/walleframe/walle/process"
	"github.com/walleframe/walle/zaplog"
)

var _ = walleServer()

// ServerOption
type ServerOptions struct {
	// Addr Server Addr
	Addr string
	// NetOption modify raw options
	NetConnOption func(net.Conn)
	// accepted load limit
	AcceptLoadLimit func(sess Session, cnt int64) bool
	// Process Options
	ProcessOptions []process.ProcessOption
	// process router
	Router Router
	// SessionRouter custom session router
	SessionRouter func(sess Session, global Router) (r Router)
	// frame log
	FrameLogger *zaplog.Logger
	// SessionLogger custom session logger
	SessionLogger func(sess Session, global *zaplog.Logger) (r *zaplog.Logger)
	// NewSession custom session
	NewSession func(in Session) (Session, error)
	// StopImmediately when session finish,business finish immediately.
	StopImmediately bool
	// Heartbeat use websocket ping/pong.
	Heartbeat time.Duration
	// WithMulticore sets up multi-cores in gnet server.
	Multicore bool
	// WithLockOSThread sets up LockOSThread mode for I/O event-loops.
	LockOSThread bool
	// WithLoadBalancing sets up the load-balancing algorithm in gnet server.
	LoadBalancing gnet.LoadBalancing
	// WithNumEventLoop sets up NumEventLoop in gnet server.
	NumEventLoop int
	// WithReusePort sets up SO_REUSEPORT socket option.
	ReusePort bool
	// WithTCPKeepAlive sets up the SO_KEEPALIVE socket option with duration.
	TCPKeepAlive time.Duration
	// WithTCPNoDelay enable/disable the TCP_NODELAY socket option.
	TCPNoDelay gnet.TCPSocketOpt
	// WithReadBufferCap sets up ReadBufferCap for reading bytes.
	ReadBufferCap int
	// WithSocketRecvBuffer sets the maximum socket receive buffer in bytes.
	SocketRecvBuffer int
	// WithSocketSendBuffer sets the maximum socket send buffer in bytes.
	SocketSendBuffer int
	// WithTicker indicates that a ticker is set.
	Ticker time.Duration
	// ReuseReadBuffer 复用read缓存区。影响Process.DispatchFilter.
	// 如果此选项设置为true，在DispatchFilter内如果开启协程，需要手动复制内存。
	// 如果在DispatchFilter内不开启协程，设置为true可以减少内存分配。
	// 默认为false,是为了防止错误的配置导致bug。
	ReuseReadBuffer bool
}

// Addr Server Addr
func WithAddr(v string) ServerOption {
	return func(cc *ServerOptions) ServerOption {
		previous := cc.Addr
		cc.Addr = v
		return WithAddr(previous)
	}
}

// NetOption modify raw options
func WithNetConnOption(v func(net.Conn)) ServerOption {
	return func(cc *ServerOptions) ServerOption {
		previous := cc.NetConnOption
		cc.NetConnOption = v
		return WithNetConnOption(previous)
	}
}

// accepted load limit
func WithAcceptLoadLimit(v func(sess Session, cnt int64) bool) ServerOption {
	return func(cc *ServerOptions) ServerOption {
		previous := cc.AcceptLoadLimit
		cc.AcceptLoadLimit = v
		return WithAcceptLoadLimit(previous)
	}
}

// Process Options
func WithProcessOptions(v ...process.ProcessOption) ServerOption {
	return func(cc *ServerOptions) ServerOption {
		previous := cc.ProcessOptions
		cc.ProcessOptions = v
		return WithProcessOptions(previous...)
	}
}

// process router
func WithRouter(v Router) ServerOption {
	return func(cc *ServerOptions) ServerOption {
		previous := cc.Router
		cc.Router = v
		return WithRouter(previous)
	}
}

// SessionRouter custom session router
func WithSessionRouter(v func(sess Session, global Router) (r Router)) ServerOption {
	return func(cc *ServerOptions) ServerOption {
		previous := cc.SessionRouter
		cc.SessionRouter = v
		return WithSessionRouter(previous)
	}
}

// frame log
func WithFrameLogger(v *zaplog.Logger) ServerOption {
	return func(cc *ServerOptions) ServerOption {
		previous := cc.FrameLogger
		cc.FrameLogger = v
		return WithFrameLogger(previous)
	}
}

// SessionLogger custom session logger
func WithSessionLogger(v func(sess Session, global *zaplog.Logger) (r *zaplog.Logger)) ServerOption {
	return func(cc *ServerOptions) ServerOption {
		previous := cc.SessionLogger
		cc.SessionLogger = v
		return WithSessionLogger(previous)
	}
}

// NewSession custom session
func WithNewSession(v func(in Session) (Session, error)) ServerOption {
	return func(cc *ServerOptions) ServerOption {
		previous := cc.NewSession
		cc.NewSession = v
		return WithNewSession(previous)
	}
}

// StopImmediately when session finish,business finish immediately.
func WithStopImmediately(v bool) ServerOption {
	return func(cc *ServerOptions) ServerOption {
		previous := cc.StopImmediately
		cc.StopImmediately = v
		return WithStopImmediately(previous)
	}
}

// Heartbeat use websocket ping/pong.
func WithHeartbeat(v time.Duration) ServerOption {
	return func(cc *ServerOptions) ServerOption {
		previous := cc.Heartbeat
		cc.Heartbeat = v
		return WithHeartbeat(previous)
	}
}

// WithMulticore sets up multi-cores in gnet server.
func WithMulticore(v bool) ServerOption {
	return func(cc *ServerOptions) ServerOption {
		previous := cc.Multicore
		cc.Multicore = v
		return WithMulticore(previous)
	}
}

// WithLockOSThread sets up LockOSThread mode for I/O event-loops.
func WithLockOSThread(v bool) ServerOption {
	return func(cc *ServerOptions) ServerOption {
		previous := cc.LockOSThread
		cc.LockOSThread = v
		return WithLockOSThread(previous)
	}
}

// WithLoadBalancing sets up the load-balancing algorithm in gnet server.
func WithLoadBalancing(v gnet.LoadBalancing) ServerOption {
	return func(cc *ServerOptions) ServerOption {
		previous := cc.LoadBalancing
		cc.LoadBalancing = v
		return WithLoadBalancing(previous)
	}
}

// WithNumEventLoop sets up NumEventLoop in gnet server.
func WithNumEventLoop(v int) ServerOption {
	return func(cc *ServerOptions) ServerOption {
		previous := cc.NumEventLoop
		cc.NumEventLoop = v
		return WithNumEventLoop(previous)
	}
}

// WithReusePort sets up SO_REUSEPORT socket option.
func WithReusePort(v bool) ServerOption {
	return func(cc *ServerOptions) ServerOption {
		previous := cc.ReusePort
		cc.ReusePort = v
		return WithReusePort(previous)
	}
}

// WithTCPKeepAlive sets up the SO_KEEPALIVE socket option with duration.
func WithTCPKeepAlive(v time.Duration) ServerOption {
	return func(cc *ServerOptions) ServerOption {
		previous := cc.TCPKeepAlive
		cc.TCPKeepAlive = v
		return WithTCPKeepAlive(previous)
	}
}

// WithTCPNoDelay enable/disable the TCP_NODELAY socket option.
func WithTCPNoDelay(v gnet.TCPSocketOpt) ServerOption {
	return func(cc *ServerOptions) ServerOption {
		previous := cc.TCPNoDelay
		cc.TCPNoDelay = v
		return WithTCPNoDelay(previous)
	}
}

// WithReadBufferCap sets up ReadBufferCap for reading bytes.
func WithReadBufferCap(v int) ServerOption {
	return func(cc *ServerOptions) ServerOption {
		previous := cc.ReadBufferCap
		cc.ReadBufferCap = v
		return WithReadBufferCap(previous)
	}
}

// WithSocketRecvBuffer sets the maximum socket receive buffer in bytes.
func WithSocketRecvBuffer(v int) ServerOption {
	return func(cc *ServerOptions) ServerOption {
		previous := cc.SocketRecvBuffer
		cc.SocketRecvBuffer = v
		return WithSocketRecvBuffer(previous)
	}
}

// WithSocketSendBuffer sets the maximum socket send buffer in bytes.
func WithSocketSendBuffer(v int) ServerOption {
	return func(cc *ServerOptions) ServerOption {
		previous := cc.SocketSendBuffer
		cc.SocketSendBuffer = v
		return WithSocketSendBuffer(previous)
	}
}

// WithTicker indicates that a ticker is set.
func WithTicker(v time.Duration) ServerOption {
	return func(cc *ServerOptions) ServerOption {
		previous := cc.Ticker
		cc.Ticker = v
		return WithTicker(previous)
	}
}

// ReuseReadBuffer 复用read缓存区。影响Process.DispatchFilter.
// 如果此选项设置为true，在DispatchFilter内如果开启协程，需要手动复制内存。
// 如果在DispatchFilter内不开启协程，设置为true可以减少内存分配。
// 默认为false,是为了防止错误的配置导致bug。
func WithReuseReadBuffer(v bool) ServerOption {
	return func(cc *ServerOptions) ServerOption {
		previous := cc.ReuseReadBuffer
		cc.ReuseReadBuffer = v
		return WithReuseReadBuffer(previous)
	}
}

// SetOption modify options
func (cc *ServerOptions) SetOption(opt ServerOption) {
	_ = opt(cc)
}

// ApplyOption modify options
func (cc *ServerOptions) ApplyOption(opts ...ServerOption) {
	for _, opt := range opts {
		_ = opt(cc)
	}
}

// GetSetOption modify and get last option
func (cc *ServerOptions) GetSetOption(opt ServerOption) ServerOption {
	return opt(cc)
}

// ServerOption option define
type ServerOption func(cc *ServerOptions) ServerOption

// NewServerOptions create options instance.
func NewServerOptions(opts ...ServerOption) *ServerOptions {
	cc := newDefaultServerOptions()
	for _, opt := range opts {
		_ = opt(cc)
	}
	if watchDogServerOptions != nil {
		watchDogServerOptions(cc)
	}
	return cc
}

// InstallServerOptionsWatchDog install watch dog
func InstallServerOptionsWatchDog(dog func(cc *ServerOptions)) {
	watchDogServerOptions = dog
}

var watchDogServerOptions func(cc *ServerOptions)

// newDefaultServerOptions new option with default value
func newDefaultServerOptions() *ServerOptions {
	cc := &ServerOptions{
		Addr: "tcp://0.0.0.0:8080",
		NetConnOption: func(net.Conn) {
		},
		AcceptLoadLimit: func(sess Session, cnt int64) bool {
			return false
		},
		ProcessOptions: nil,
		Router:         process.GetRouter(),
		SessionRouter: func(sess Session, global Router) (r Router) {
			return global
		},
		FrameLogger: zaplog.GetFrameLogger(),
		SessionLogger: func(sess Session, global *zaplog.Logger) (r *zaplog.Logger) {
			return global
		},
		NewSession: func(in Session) (Session, error) {
			return in, nil
		},
		StopImmediately:  false,
		Heartbeat:        0,
		Multicore:        false,
		LockOSThread:     false,
		LoadBalancing:    gnet.SourceAddrHash,
		NumEventLoop:     0,
		ReusePort:        false,
		TCPKeepAlive:     0,
		TCPNoDelay:       gnet.TCPNoDelay,
		ReadBufferCap:    0,
		SocketRecvBuffer: 0,
		SocketSendBuffer: 0,
		Ticker:           0,
		ReuseReadBuffer:  false,
	}
	return cc
}
