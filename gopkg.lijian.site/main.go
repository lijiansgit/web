package main

import (
	"github.com/astaxie/beego"
	_ "myweb/routers"
)

func main() {
	//pkgdoc
	beego.SetStaticPath("/pkg/", "views/pkg/")

	beego.Run()
}
