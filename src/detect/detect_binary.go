package detect

import (
    "os"
    "log"
    "bufio"
)

func DetectFile(path string) bool {

    file, err := os.Open(path)
    if err != nil {
        log.Printf("\033[31merror : IO error - \033[0m%s", err)
        return false
    }
    defer file.Close()

    r := bufio.NewReader(file)
    buf := make([]byte, 100)
    n, err := r.Read(buf)

    //var bad_byte int = 0
    for i := 0; i < n; i++ {
        if (buf[i] >= 0x20 && buf[i] <= 0xff) ||
            buf[i] == '\n' ||
            buf[i] == '\t' {
            continue
        }
        return true
    }

    return false
}
