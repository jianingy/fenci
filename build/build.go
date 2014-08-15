/*
 * filename   : build.go
 * created at : 2014-08-15 12:03:10
 * author     : Jianing Yang <jianingy.yang@gmail.com>
 */

package main

import (
    "bytes"
    "fmt"
    "io/ioutil"
    "os"
    "regexp"
    "strings"
    "strconv"

    "github.com/codegangsta/cli"
    "github.com/torbit/cdb"

    . "github.com/jianingy/fenci/constants"
    . "github.com/jianingy/fenci/utils"
)

func main() {
    app := cli.NewApp()
    app.Name = "fenci-build"
    app.Flags = []cli.Flag {
        cli.StringFlag{
            Name: "in",
            Value: "people-daily.txt",
            Usage: "manually segmented dataset ",
        },
        cli.StringFlag{
            Name: "out",
            Value: "dict.cdb",
            Usage: "output file ",
        },
    }
    app.Action = func(c *cli.Context) {
        build(c)
    }
    app.Run(os.Args)
}

func build(c *cli.Context) {
    var text []byte
    var err error

    Log("INFO", "开始处理语料库 %s\n", c.String("in"))
    if text, err = ioutil.ReadFile(c.String("in")); err != nil {
        panic(err)
    }

    out, err := os.OpenFile(c.String("out"),
        os.O_CREATE | os.O_WRONLY |os.O_TRUNC, 0644)
    if err != nil {
        panic(err)
    }
    defer func() { out.Close() }()

    // 匹配语料库里 “文字/词性”这种正则。
    // 为了简化，这里不处理复合词，也就是 “[xxxx]yy”这种格式我们不处理
    re := regexp.MustCompile("[^\\x00-\\xff]+?/[a-z]+")
    hash := make(map[string]int)

    for i, term := range re.FindAllString(string(text), -1) {
        pair := strings.Split(term, "/")
        switch pair[1] {
        case "t":     // 不要时间词
            fallthrough
        case "m":     // 不要数字词
            fallthrough
        case "w":     // 不要标点符号
            continue
        }
        hash[pair[0]]++
        if i % 500 == 0 {
            Log("INFO", "已经识别 %d 个单词\r", i)
        }
    }
    Log("INFO", "共识别 %d 个单词\n", len(hash))

    records := bytes.NewBuffer(nil)
    var total, maxlength int
    for word, count := range hash {
        total = total + count
        if count > len(word) {
            maxlength = len(word)
        }
        cnt := strconv.Itoa(count)
        fmt.Fprintf(records, "+%d,%d:%s->%s\n", len(word), len(cnt), word, cnt)
        if total % 500 == 0 {
            Log("INFO", "已经插入 %d 个单词\r", total)
        }
    }
    Log("INFO", "共插入 %d 个单词\n", total)

    // 存储字典统计信息
    stats := make(map[string]string)
    stats[KEY_TOTAL] = strconv.Itoa(total)
    stats[KEY_MAXLENGTH] = strconv.Itoa(maxlength)
    stats[KEY_NUMTERMS] = strconv.Itoa(len(hash))

    for name, value := range stats {
        fmt.Fprintf(records, "+%d,%d:%s->%s\n",
            len(name), len(value), name, value)
    }
    // 最后要多一个 “\n” 不然会 panic: EOF
    fmt.Fprintf(records, "\n")

    Log("INFO", "写入字典文件 %s ... \n", c.String("out"))
    if err = cdb.Make(out, records); err != nil {
        panic(err)
    }
    Log("INFO", "完成\n")
}
