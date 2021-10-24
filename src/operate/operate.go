package operate

import (
    "github.com/zgg2001/StringFinderZ/detect"
    "github.com/zgg2001/StringFinderZ/argument"
    "log"
    "io/ioutil"
    "fmt"
    "strings"
    "sync"
    "os"
    "bufio"
)

//目录操作
func OperateDir(wg *sync.WaitGroup, path string, dirCh chan string, fileCh chan string) {

	result, err := ioutil.ReadDir(path)
	if err != nil {
		log.Printf("\033[31merror : IO error - \033[0m%s", err)
        return
	}

	for _, fi := range result {
		if fi.IsDir() {
			dirCh <- path + "/" + fi.Name()
            wg.Add(1)
		} else {
			fileCh <- path + "/" + fi.Name()
			wg.Add(1)
		}
	}
}

//文件操作 find
func OperateFileF(path string) bool {

    file, err := os.Open(path)
    if err != nil {
        log.Printf("\033[31merror : IO error - \033[0m%s", err)
        return false
    }
    defer file.Close()

    //搜索
    var line_number int = 1
    buf := bufio.NewScanner(file)
    for {
        if !buf.Scan() {
            break
        }
        line := buf.Text()
        if strings.Contains(string(line), argument.Arg.Find) {
            //当此文件为二进制文件时，直接跳过
            if detect.DetectBinary(path) {
                fmt.Printf("Binary file \033[35m%s\033[0m matches\n", path)
                break
            }
            if argument.Arg.Number {
                fmt.Printf("\033[35m%s\033[32m:%d:\033[0m\t%s\n", path, line_number, line)
            } else {
                fmt.Printf("\033[35m%s\033[32m:\033[0m\t%s\n", path, line)
            }
        }
        line_number++
    }
    return false
}

//文件操作 replace
func OperateFileR(path string) bool {

    file, err := os.Open(path)
    if err != nil {
        log.Printf("\033[31merror : IO error - \033[0m%s", err)
        return false
    }
    defer file.Close()

    //替换
    var line_number int = 1
    var state bool = false
    var content string
    //当此文件为二进制文件时，直接跳过
    if detect.DetectBinary(path) {
        fmt.Printf("Skip Binary file \033[35m%s\033[0m\n", path)
        return false
    }

    //开始搜索并替换
    buf := bufio.NewScanner(file)
    for {
        if !buf.Scan() {
            break
        }
        line := buf.Text()
        if strings.Contains(string(line), argument.Arg.Find) {
            if argument.Arg.Number {
                fmt.Printf("\033[35m%s\033[32m:%d:\033[0m\t%s\n", path, line_number, line)
            } else {
                fmt.Printf("\033[35m%s\033[32m:\033[0m\t%s\n", path, line)
            }
            //询问是否替换
            if argument.Arg.AllYes || IsReplace() {
                new_content := strings.ReplaceAll(string(line), argument.Arg.Find, argument.Arg.Replace)
                content = content + new_content + "\n"
                state = true
            } else {
                content = content + line + "\n"
            }
        } else {
            content = content + line + "\n"
        }
        line_number++
    }

    //若改动 则替换原文
    if state == true {
	    err = ioutil.WriteFile(path, []byte(content), 0777)
	    if err != nil {
	        log.Fatal("\033[31merror : IO error - \033[0m", err)
        }
    }

    return false
}

//replace 快速替换
func OperateFileQ(path string) bool {

    content, err := ioutil.ReadFile(path)
    if err != nil {
        log.Printf("\033[31merror : IO error - \033[0m%s", err)
        return false
    }

    //当此文件为二进制文件时，直接跳过
    if detect.DetectBinary(path) {
        fmt.Printf("Skip Binary file \033[35m%s\033[0m\n", path)
        return false
    }

    newReader := strings.ReplaceAll(string(content), argument.Arg.Find, argument.Arg.Replace)
    err = ioutil.WriteFile(path, []byte(newReader), 0777)
    if err != nil {
        log.Fatal("\033[31merror : IO error - \033[0m", err)
    }

    return true
}

//询问是否替换
func IsReplace() bool {
    var input string
    fmt.Printf("是否替换本处? ")
    fmt.Scanf("%s", &input)
    if strings.EqualFold(input, "y") || strings.EqualFold(input, "yes") {
        return true
    }
    return false
}

