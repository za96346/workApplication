package logger

import (
	"fmt"
	_ "net/http/pprof"
	"os"
	"time"
    "github.com/sirupsen/logrus"
    "path"

	// "time"

	// "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Logger() *logrus.Logger {
    now := time.Now()
    logFilePath := ""
    if dir, err := os.Getwd(); err == nil {
        logFilePath = dir + "/logs/"
    }
    if err := os.MkdirAll(logFilePath, 0777); err != nil {
        fmt.Println(err.Error())
    }
    logFileName := now.Format("2006-01-02") + ".log"
    //日誌檔案
    fileName := path.Join(logFilePath, logFileName)
    if _, err := os.Stat(fileName); err != nil {
        if _, err := os.Create(fileName); err != nil {
            fmt.Println(err.Error())
        }
    }
    //寫入檔案
    src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
    if err != nil {
        fmt.Println("err", err)
    }

    //例項化
    logger := logrus.New()

    //設定輸出
    logger.Out = src

    //設定日誌級別
    logger.SetLevel(logrus.DebugLevel)

    //設定日誌格式
    logger.SetFormatter(&logrus.TextFormatter{
        TimestampFormat: "2006-01-02 15:04:05",
    })
    return logger
}

func LoggerToFile() gin.HandlerFunc {
    logger := Logger()
    return func(c *gin.Context) {
        // 開始時間
        startTime := time.Now()

        // 處理請求
        c.Next()

        // 結束時間
        endTime := time.Now()

        // 執行時間
        latencyTime := endTime.Sub(startTime)

        // 請求方式
        reqMethod := c.Request.Method

        // 請求路由
        reqUri := c.Request.RequestURI

        // 狀態碼
        statusCode := c.Writer.Status()

        // 請求IP
        clientIP := c.ClientIP()

        //日誌格式
        logger.Infof("| %3d | %13v | %15s | %s | %s |",
            statusCode,
            latencyTime,
            clientIP,
            reqMethod,
            reqUri,
        )
    }
}