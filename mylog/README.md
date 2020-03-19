#### mylog
***

传入生成的日志文件名称作为参数

```
import "github.com/xiejia1992/my-golang-lib/mylog"

func main() {
	log := mylog.NewMyLog("running.log")
	log.Debug("TEST Debug")
	log.Error("TEST Error")
	log.Info("TEST Info")
	log.Warning("TEST Warning")
}
```