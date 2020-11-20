// description: log 
// 
// @author: xwc1125
// @date: 2019/9/18
package tests

import (
	"github.com/chain5j/log15"
	"os"
	"testing"
)

var (
	ostream log.Handler
	glogger *log.GlogHandler
)

func TestLog(t *testing.T) {
	// 颜色
	//usecolor := (isatty.IsTerminal(os.Stderr.Fd()) || isatty.IsCygwinTerminal(os.Stderr.Fd())) && os.Getenv("TERM") != "dumb"
	//output := io.Writer(os.Stderr)
	//if usecolor {
	//	output = colorable.NewColorableStderr()
	//}
	//ostream = log.StreamHandler(output, log.TerminalFormat(usecolor))// usecolor表示日志是否颜色显示
	ostream = log.StreamHandler(os.Stderr, log.TerminalFormat(true)) // usecolor表示日志是否颜色显示
	glogger = log.NewGlogHandler(ostream)
	glogger.Verbosity(log.Lvl(5)) // 日志级别

	//glogger.Vmodule(*vmodule)
	log.Root().SetHandler(glogger)

	// all loggers can have key/value context
	srvlog := log.New("app/server")

	// all log messages can have key/value context
	//srvlog.Warn("abnormal conn rate", "rate", "curRate", "low", "lowRate", "high", "highRate")

	srvlog.Print("aaaaaa")
	srvlog.Print(111, 222, 333, "12345")

	srvlog.Fatal("111")

	// child loggers with inherited context
	connlog := srvlog.New("srvlog", "raddr", "127.0.0.1")
	connlog.Info("connection open")

	// lazy evaluation
	connlog.Debug("ping remote", "latency", log.Lazy{"pingRemote"})

	// flexible configuration
	srvlog.SetHandler(log.MultiHandler(
		log.StreamHandler(os.Stderr, log.LogfmtFormat()),
		log.LvlFilterHandler(
			log.LvlError,
			log.Must.FileHandler("errors.json", log.JSONFormat()))))
}