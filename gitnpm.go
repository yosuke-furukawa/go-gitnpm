package main

import (
	"encoding/json"
	"fmt"
	"github.com/codegangsta/cli"
	"io/ioutil"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
)

const APPNAME = "gitnpm"
const GNPMRC = ".gnpmrc"
const GITURL = "%s://git@%s:%s.git#%s"

const USAGE = ` 
Usage: gitnpm [subcommand] [options]

subcommand:
  setup    setup command 
  install  install 

Options:
  setup

`

type Gnpmrc struct {
	Protocol string `json:"protocol"`
	Hostname string `json:"hostname"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
func main() {
	app := cli.NewApp()
	app.Name = APPNAME
	app.Usage = USAGE
	app.Commands = []cli.Command{
		{
			Name:      "setup",
			ShortName: "s",
			Usage:     "create setup file and save $HOME/.gnpmrc",
			Flags: []cli.Flag{
				cli.StringFlag{"protocol", "git+ssh", "protocol of url"},
				cli.StringFlag{"hostname", "github.com", "hostname of url"},
			},
			Action: func(c *cli.Context) {
				usr, err := user.Current()
				check(err)
				path := filepath.Join(usr.HomeDir, GNPMRC)
				println(c.String("protocol"))
				println(c.String("hostname"))
				data, err := json.Marshal(&Gnpmrc{c.String("protocol"), c.String("hostname")})
				check(err)
				println(string(data))

				err = ioutil.WriteFile(path, data, 0666)
				check(err)
			},
		},
		{
			Name:      "install",
			ShortName: "i",
			Usage:     "install github module",
			Flags: []cli.Flag{
				cli.StringFlag{"protocol", "git+ssh", "protocol of url"},
				cli.StringFlag{"hostname", "github.com", "hostname of url"},
			},
			Action: func(c *cli.Context) {
				usr, err := user.Current()
				check(err)
				var npmrc *Gnpmrc
				path := filepath.Join(usr.HomeDir, ".gnpmrc")
				if exists(path) {
					data, err := ioutil.ReadFile(path)
					check(err)
					json.Unmarshal(data, &npmrc)
				} else {
					npmrc = &Gnpmrc{c.String("protocol"), c.String("hostname")}
				}

				if len(c.Args()) < 1 {
					println("Usage: gitnpm install <username>/<module>")
					os.Exit(1)
				}
				module := c.Args()[0]

				version := "master"

				if len(c.Args()) > 1 {
					version = c.Args()[1]
				}
				giturl := fmt.Sprintf(GITURL, npmrc.Protocol, npmrc.Hostname, module, version)
				args := []string{"install", giturl}
				if len(c.Args()) > 2 {
					args = append(args, c.Args()[2:]...)
				}
				cmd := exec.Command("npm", args...)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				cmd.Run()
			},
		},
	}
	app.Run(os.Args)
}
