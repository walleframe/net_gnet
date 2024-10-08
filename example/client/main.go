package main

import (
	"context"
	"fmt"
	"time"

	"github.com/walleframe/net_gnet"
	"github.com/walleframe/walle/testpkg/wpb"
	"github.com/walleframe/walle/util"
	"github.com/walleframe/walle/zaplog"
)

func main() {
	zaplog.SetFrameLogger(zaplog.NoopLogger)
	zaplog.SetLogicLogger(zaplog.NoopLogger)
	// zaplog.SetFrameLogger(zaplog.GetLogicLogger())

	cli, err := net_gnet.NewClient(
		net_gnet.WithClientOptionAddr(fmt.Sprintf("localhost:%d", 12345)),
	)
	if err != nil {
		util.PanicIfError(err)
	}
	time.Sleep(time.Second)

	wcli := wpb.NewWSvcClient(cli)
	ctx := context.Background()

	rs, err := wcli.Add(ctx, &wpb.AddRq{Params: []int64{100}})
	fmt.Println(rs, err)
}
