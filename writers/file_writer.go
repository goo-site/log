package writers

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strconv"
	"time"
)

type FileWriter struct {
	logDir    string
	logSubDir string
	logNumber int   // 当天的日志编号
	logSize   int64 // 单个日志大小
	logTime   time.Time
	logFd     *os.File
}

func (f *FileWriter) Format(prefix string, timestamp int64, filepath string, line int, s string) string {
	f.logTime = time.Unix(timestamp, 0)
	filename := path.Base(filepath)
	return fmt.Sprintf("%s[%s] %s:%d %s\n", prefix, f.logTime.Format("15:04:05"), filename, line, s)
}

func (f *FileWriter) Write(s string) {
	// 检查子文件夹
	f.logSubDir = getSubDir(f.logTime)
	os.MkdirAll(path.Join(f.logDir, f.logSubDir), os.ModePerm)

	// 获取文件信息，没有则创建
	filePath := path.Join(f.logDir, f.logSubDir, fmt.Sprintf("%04d", f.logNumber))
	f.open(filePath)

	// 检查文件大小
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		log.Fatalf("[system] IoWrite stat err: %v\n", err)
	}

	if f.logSize > 0 && fileInfo.Size() > f.logSize {
		f.logNumber++
		f.open(path.Join(f.logDir, f.logSubDir, fmt.Sprintf("%04d", f.logNumber)))
	}

	fmt.Fprintf(f.logFd, "%s", s)
}

// 子路径：202208/29，年+月/日
func getSubDir(now time.Time) string {
	return fmt.Sprintf("%04d%02d/%02d", now.Year(), now.Month(), now.Day())
}

// 按路径打开新文件
func (f *FileWriter) open(path string) {
	fd, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	if err != nil {
		f.logFd = nil
		fmt.Printf("[system] open logfile err: %v", err)
		return
	}
	if f.logFd != nil {
		f.logFd.Sync()
		f.logFd.Close()
	}
	f.logFd = fd
}

// 设置日志输出路径
func (f *FileWriter) SetLogDir(dir string) {
	// 检查dir是否存在
	stat, err := os.Stat(dir)
	if err != nil {
		log.Fatalf("[system] LogDir not exist err: %v\n", err)
	}
	if !stat.IsDir() {
		log.Fatalf("[system] LogDir not exist err: %v\n", err)
	}

	// 查找最新的logNumber
	now := time.Now()
	subDir := getSubDir(now)
	os.MkdirAll(path.Join(dir, subDir), os.ModePerm)
	files, err := ioutil.ReadDir(path.Join(dir, subDir))
	if err != nil {
		log.Fatalf("[system] newLogFile ReadDir err: %v\n", err)
	}

	number := 0
	if len(files) != 0 {
		number, err = strconv.Atoi(files[len(files)-1].Name())
		if err != nil {
			log.Fatalf("[system] newLogFile Atoi err: %v\n", err)
		}
	}

	f.logDir = dir
	f.logNumber = number
}

// 设置单个文件大小
func (f *FileWriter) SetLogSize(size int64) {
	f.logSize = size
}
