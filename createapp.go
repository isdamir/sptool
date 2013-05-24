package main

import (
	"fmt"
	"os"
	"path"
	"strings"
)

var cmdCreate = &Command{
	UsageLine: "create [appname]",
	Short:     "create an application base on gospeed framework",
	Long: `
create an application base on gospeed framework

In the current path, will create a folder named [appname]

In the appname folder has the follow struct:

	|- main.go
	|- conf
	    |-  app.json
	|- controllers
	     |- default.go
	|- models
	|- static
	     |- add-on
	     |- web
		 |- js
		 |- css
	         |- img				
	|- views
	    index.html					

`,
}

func init() {
	cmdCreate.Run = createapp
}

func createapp(cmd *Command, args []string) {
	crupath, _ := os.Getwd()
	if len(args) != 1 {
		fmt.Println("error args")
		os.Exit(2)
	}
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		fmt.Println("you should set GOPATH in the env")
		os.Exit(2)
	}
	haspath := false
	if crupath != path.Join(gopath, "src") {
		wgopath := strings.Split(gopath, ";")
		if len(wgopath) >= 1 {
			for _, wg := range wgopath {
				wg = wg + `\src`
				if crupath == wg {
					haspath = true
					break
				}
			}
		}
		if !haspath {
			lgopath := strings.Split(gopath, ":")
			if len(lgopath) >= 1 {
				for _, wg := range lgopath {
					if crupath == path.Join(wg, "src") {
						haspath = true
						break
					}
				}
			}
		}

	} else {
		haspath = true
	}
	if !haspath {
		fmt.Println("can't create application outside of GOPATH")
		fmt.Println("you first should `cd $GOPATH/src` then use create")
		os.Exit(2)
	}
	apppath := path.Join(crupath, args[0])
	os.Mkdir(apppath, 0755)
	fmt.Println("create app folder:", apppath)
	os.Mkdir(path.Join(apppath, "conf"), 0755)
	fmt.Println("create conf:", path.Join(apppath, "conf"))
	os.Mkdir(path.Join(apppath, "controllers"), 0755)
	fmt.Println("create controllers:", path.Join(apppath, "controllers"))
	os.Mkdir(path.Join(apppath, "models"), 0755)
	fmt.Println("create models:", path.Join(apppath, "models"))
	os.Mkdir(path.Join(apppath, "static"), 0755)
	fmt.Println("create static:", path.Join(apppath, "static"))
	os.Mkdir(path.Join(apppath, "static", "add-on"), 0755)
	fmt.Println("create static js:", path.Join(apppath, "static", "add-on"))
	os.Mkdir(path.Join(apppath, "static", "web"), 0755)
	os.Mkdir(path.Join(apppath, "static", "web", "js"), 0755)
	fmt.Println("create static js:", path.Join(apppath, "static", "web", "js"))
	os.Mkdir(path.Join(apppath, "static", "web", "css"), 0755)
	fmt.Println("create static css:", path.Join(apppath, "static", "web", "css"))
	os.Mkdir(path.Join(apppath, "static", "web", "img"), 0755)
	fmt.Println("create static img:", path.Join(apppath, "static", "web", "img"))
	fmt.Println("create views:", path.Join(apppath, "views"))
	os.Mkdir(path.Join(apppath, "views"), 0755)
	fmt.Println("create conf app.json:", path.Join(apppath, "conf", "app.json"))
	writetofile(path.Join(apppath, "conf", "app.json"), strings.Replace(appconf, "{{.Appname}}", args[0], -1))

	fmt.Println("create controllers default.go:", path.Join(apppath, "controllers", "default.go"))
	writetofile(path.Join(apppath, "controllers", "default.go"), controllers)

	fmt.Println("create views index.html:", path.Join(apppath, "views", "index.html"))
	writetofile(path.Join(apppath, "views", "index.html"), indextpl)

	fmt.Println("create main.go:", path.Join(apppath, "main.go"))
	writetofile(path.Join(apppath, "main.go"), strings.Replace(maingo, "{{.Appname}}", args[0], -1))
}

var appconf = `{
	"AppName": "{{.Appname}}",
	"HttpPort": 8080,
	"RunMode": "dev",
	"Custom": {
	}
}
`

var maingo = `package main

import (
	"iyf.cc/gospeed/web"
	_ "testSession/controllers"
)

func main() {
	web.Start()
}
`
var controllers = `package controllers

import (
	"iyf.cc/gospeed/web"
)

type MainController struct {
	web.Controller
}

func init() {
	web.RegisterRouter("/", &MainController{})
}
func (this *MainController) Get() {
	this.Data["Username"] = "YuFeng"
	this.Data["Email"] = "isyufeng@gmail.com"
	this.ServeTpl("index.html")
}
`

var indextpl = `<!DOCTYPE html>
<html>
  <head>
    <title>gospeed welcome template</title>
  </head>
  <body>
    <h1>Hello, world!{{.Username}},{{.Email}}</h1>
  </body>
</html>
`

func writetofile(filename, content string) {
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.WriteString(content)
}
