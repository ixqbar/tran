package server

import (
	"common"
	"context"
	"github.com/liuzl/gocc"
	"github.com/tidwall/redcon"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func RunRedisServer(ctx context.Context, opcc *gocc.OpenCC) {
	rs := redcon.NewServerNetwork("tcp", common.Config.Address,
		func(conn redcon.Conn, cmd redcon.Command) {
			switch strings.ToLower(string(cmd.Args[0])) {
			default:
				conn.WriteError("ERR unknown command '" + string(cmd.Args[0]) + "'")
			case "select":
				conn.WriteBulkString("OK")
			case "version":
				conn.WriteBulkString(common.Version)
			case "get":
				if len(cmd.Args) != 2 {
					conn.WriteError("ERR wrong number of arguments for '" + string(cmd.Args[0]) + "' command")
					return
				}
				out, err := opcc.Convert(string(cmd.Args[1]))
				if err != nil {
					conn.WriteRaw(cmd.Args[1])
					return
				}
				conn.WriteBulkString(out)
			}
		},
		func(conn redcon.Conn) bool {
			log.Printf("redis server accept client %s", conn.RemoteAddr())
			return true
		},
		func(conn redcon.Conn, err error) {
			if err != nil {
				log.Printf("redis server closed client %s, err: %v", conn.RemoteAddr(), err)
			}
		},
	)

	go func() {
		if err := rs.ListenAndServe(); err != nil {
			log.Printf("redis server listen fail %v", err)
		}
	}()

	select {
	case <-ctx.Done():
		log.Printf("redis server catch exit signal")
		rs.Close()
		return
	}
}


func Run() error {
	gocc.Dir = &common.Config.DataPath
	opcc, err := gocc.New("s2t")
	if err != nil {
		return err
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())

	go RunRedisServer(ctx, opcc)

	select {
	case <-quit:
		log.Printf("Shutdown Server ...")
		cancel()
	}

	<-time.After(3 * time.Second)
	log.Printf("Server exiting")

	return nil
}