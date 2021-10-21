package main

import (
    "flag"
    "github.com/zgg2001/StringFinderZ/argument"
    "github.com/zgg2001/StringFinderZ/operate"
    "log"
    "os"
    "sync"
)

func main() {

    flag.StringVar(&argument.Arg.Path, "p", ".", "path ") //路径 默认为当前路径
    flag.StringVar(&argument.Arg.Find, "f", "", "want to find") //搜索的内容
    flag.StringVar(&argument.Arg.Replace, "r", "", "want to replace") //替换的内容
    flag.BoolVar(&argument.Arg.Number, "n", false, "show line number") //显示行号
    flag.BoolVar(&argument.Arg.AllYes, "y", false, "Allow all replace") //允许所有替换
    flag.BoolVar(&argument.Arg.Quick, "q", false, "quick replace") //快速替换

    flag.Parse();

    //搜索内容为空
    if argument.Arg.Find == "" {
        log.Fatal("\033[31merror : Missing argument - \033[0m'find' no content")
    }

    //路径判定
    file, err := os.Stat(argument.Arg.Path)
    if err != nil {
        log.Fatal("\033[31merror : Path error - \033[0m", err)
    }

    //判断模式
    if argument.Arg.Replace == "" {
        argument.Arg.IsFind = true
    } else {
        argument.Arg.IsFind = false
    }
    if argument.Arg.IsFind && argument.Arg.Quick {
        log.Fatal("\033[31merror : Missing argument - \033[0m'replace' no content")
    }

	dirCh := make(chan string, 1000)
	fileCh := make(chan string, 1000)

	var wg sync.WaitGroup

    log.Printf("Start.")

	if file.IsDir() {
		operate.OperateDir(&wg, argument.Arg.Path, dirCh, fileCh)
	} else {
        if argument.Arg.IsFind {
            operate.OperateFileF(argument.Arg.Path)
        } else {
            operate.OperateFileR(argument.Arg.Path)
        }
	}

    //线程1 遍历目录
	go func() {
		for c := range dirCh {
			operate.OperateDir(&wg, c, dirCh, fileCh)
			wg.Done()
		}
	}()

    //线程2 遍历文件
	go func() {
        if argument.Arg.IsFind {
		    for c := range fileCh {
		        operate.OperateFileF(c)
			    wg.Done()
		    }
        } else if argument.Arg.Quick {
		    for c := range fileCh {
		        operate.OperateFileQ(c)
			    wg.Done()
		    }
        } else {
		    for c := range fileCh {
		        operate.OperateFileR(c)
			    wg.Done()
		    }
        }
    }()

	wg.Wait()
    log.Printf("Done.")
}
