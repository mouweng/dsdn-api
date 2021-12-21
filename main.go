package main

import (
	"math/rand"
	"net/http"
	"os"
	"runtime/debug"
	"time"

	"ginTemplate/config"
	"ginTemplate/api"
	"github.com/cihub/seelog"
	"github.com/facebookgo/grace/gracehttp"
	"github.com/facebookgo/pidfile"
	"github.com/urfave/cli"

	_ "ginTemplate/api/test"
	_ "ginTemplate/api/message"
)

var buildstamp = ""
var githash = ""

func main() {
	rand.Seed(time.Now().UnixNano())
	debug.SetTraceback("crash")
	app := cli.NewApp()
	app.Name = "gin_template"
	app.Usage = "gin框架web server项目脚手架"
	app.Flags = config.Flags
	app.Version = githash + " " + buildstamp
	app.Action = action
	app.Run(os.Args)
}

func action(c *cli.Context) error {
	logger, _ := seelog.LoggerFromConfigAsBytes([]byte(logtoconsoleconf))
		seelog.ReplaceLogger(logger)
		defer seelog.Flush()
		config.Initialize(c)
		pidfile.SetPidfilePath(os.Args[0] + ".pid")
		pidfile.Write()
		srv := api.NewServer(c)
		host := c.String("host")
		port := c.String("port")
		listenAddress := host + ":" + port
		seelog.Info("Serve on ", listenAddress)
		return gracehttp.Serve(&http.Server{
			Addr:         listenAddress,
			Handler:      srv,
			ReadTimeout:  100 * time.Second,
			WriteTimeout: 100 * time.Second,
		})
}

const (
	logtoconsoleconf = `
	<seelog>
		<outputs>
			<console formatid="out"/>
		</outputs>
		<formats>
		    <format id="out" format="[%Level] %File:%Line %Func %Msg%n"/>
		</formats>
	</seelog>
	`
)