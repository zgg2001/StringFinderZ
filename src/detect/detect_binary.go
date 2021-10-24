package detect

import (
    "os"
    "log"
    "bufio"
)

func DetectBinary(path string) bool {

    file, err := os.Open(path)
    if err != nil {
        log.Printf("\033[31merror : IO error - \033[0m%s", err)
        return false
    }
    defer file.Close()

    r := bufio.NewReader(file)
    buf := make([]byte, 1024)
    n, err := r.Read(buf)

    var white_byte int = 0
    for i := 0; i < n; i++ {
        if (buf[i] >= 0x20 && buf[i] <= 0xff) ||
            buf[i] == 9  ||
            buf[i] == 10 ||
            buf[i] == 13 {
            white_byte++
        } else if buf[i] <= 6 || (buf[i] >= 14 && buf[i] <= 31) {
            return true
        }
    }

    if white_byte >= 1 {
        return false
    }
    return true
}
