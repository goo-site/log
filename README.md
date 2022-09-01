# log
## 特点
* 插入式的日志输出接口
* 自带两个输出接口的实现

## 示例
```
import "log"
import "log/writes"

func Init(){
    consoleWriter := ConsoleWriter{}
    
    fileWriter := FileWriter{}
    fileWriter.SetLogDir("logdir")
    fileWriter.SetLogSize(1024*1024)

    log.SetLevel("Info")
    log.AddWriter(consoleWriter, fileWriter)
}

func main(){
    Init()
    s := "hello world"
    log.Info("this is Info: %s", s)
    log.Error("this is Error: %s", s)
}
```
