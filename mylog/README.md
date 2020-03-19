#### mylog
***

传入生成的日志文件名称作为参数

```
import (
	"github.com/my-golang-lib/mylog"
	"time"
)

func main() {
	log := mylog.NewMyLog("running.log")
	log.Debug("TEST Debug")
	log.Error("test error log")

	time.Sleep(2 * time.Second)
}
```