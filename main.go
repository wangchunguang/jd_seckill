package main

import (
	"bufio"
	"flag"
	"jd_seckill/global"
	logs "jd_seckill/log"
	"jd_seckill/scekill"
	"os"
	"strings"
	"time"
)

var (
	skuId       = flag.String("sku", "100012043978", "商品id")
	num         = flag.Int("num", 2, "商品数量")
	works       = flag.Int("works", 6, "并发数量")
	start       = flag.String("time", "09:59:59", "开始时间")
	brwoserPath = flag.String("execPath", "", "浏览器执行路径")
	err         error
)

func init() {
	flag.Parse()
}

func main() {
	execPath := ""
	if *brwoserPath != "" {
		execPath = *brwoserPath
	}
RE:
	jdSecKill := scekill.NewJdSecKill(execPath, *skuId, *num, *works)
	jdSecKill.StartTime, err = global.Hour2Unix(*start)
	if err != nil {
		logs.Fatal("开始时间初始化失败", err)
	}
	if jdSecKill.StartTime.Unix() < time.Now().Unix() {
		jdSecKill.StartTime = jdSecKill.StartTime.AddDate(0, 0, 1)
	}

	jdSecKill.SyncJdTime()
	logs.PrintlnInfo("开始执行时间为：", jdSecKill.StartTime.Format(global.DateTimeFormatStr))
	err = jdSecKill.Run()
	if err != nil {
		if strings.Contains(err.Error(), "exec") {
			logs.PrintlnInfo("默认浏览器执行路径未找到，" + execPath + "  请重新输入：")
			scanner := bufio.NewScanner(os.Stdin)
			for scanner.Scan() {
				execPath = scanner.Text()
				if execPath != "" {
					break
				}
			}
			goto RE
		}
		logs.Fatal(err)
	}
	goto RE

}
