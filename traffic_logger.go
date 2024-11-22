package traffic

import (
    "log"
    "os"
)

func LogTraffic(data string) {
    file, err := os.OpenFile("traffic.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Printf("Failed to log traffic: %s", err)
        return
    }
    defer file.Close()
    logger := log.New(file, "TRAFFIC: ", log.LstdFlags)
    logger.Println(data)
}
