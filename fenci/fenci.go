/*
 * filename   : build.go
 * created at : 2014-08-15 12:03:10
 * author     : Jianing Yang <jianingy.yang@gmail.com>
 */

package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	. "github.com/jianingy/fenci/core"
)


func main() {
	app := cli.NewApp()
	app.Name = "fenci"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "db",
			Value: "dict.cdb",
			Usage: "字典文件",
		},
		cli.StringFlag{
			Name:  "text",
			Value: "",
			Usage: "需要分词的句子",
		},
	}
	app.Action = func(c *cli.Context) {
		var err error
		seg, err := NewSegmentor(c.String("db"))
		if err != nil {
			panic(err)
		}
		result, err := seg.DoText(c.String("text"))
		if err != nil {
			panic(err)
		}
		fmt.Println(result)
	}
	app.Run(os.Args)
}
