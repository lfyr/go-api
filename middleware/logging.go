package middleware

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"runtime"
	"time"
)

var (
	dunno     = []byte("???")
	centerDot = []byte("·")
	dot       = []byte(".")
	slash     = []byte("/")
)

// 日志中间件
//是否需要打印body
func LoggerWithWriter(logBody bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery
		c.Next()
		end := time.Now()
		latency := end.Sub(start)
		statusCode, exit := c.Get("statusCode")
		if !exit {
			statusCode = c.Writer.Status()
		}
		//uid := utils.GetUid(c)
		if raw != "" {
			path = path + "?" + raw
		}
		fieldMap := logrus.Fields{
			"method":     c.Request.Method,
			"path":       path,
			"statusCode": statusCode,
			"referer":    c.Request.Referer(),
			"latency":    fmt.Sprintf("%13v", latency),
			"ip":         c.Request.Header.Get("X-Forwarded-For"),
			"remoteAddr": c.Request.RemoteAddr,
			"user-Agent": c.Request.Header.Get("User-Agent"),
			//"uid":        uid,
			//"token":      getTokenByContext(c),
			//"loginType":  gin_utils.GetLoginType(c),
		}
		if logBody {
			body, err := ioutil.ReadAll(c.Request.Body)
			if err != nil {
				return
			}
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
			fieldMap["body"] = string(body)
		}
		AppVersion := c.Request.Header.Get("App-Version")
		appInfo := c.Request.Header.Get("app-info")
		if AppVersion != "" {
			fieldMap["AppVersion"] = AppVersion
		}
		if appInfo != "" {
			fieldMap["appInfo"] = appInfo
		}
		logrus.WithFields(fieldMap).Info("access log")
	}
}

func Recovery() gin.HandlerFunc {
	return RecoveryWithWriter()
}

// RecoveryWithWriter returns a middleware for a given writer that recovers from any panics and writes a 500 if there was one.
func RecoveryWithWriter() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				field := make(logrus.Fields)
				stack := stack(3)
				httprequest, _ := httputil.DumpRequest(c.Request, true)
				field["request"] = string(httprequest)
				field["stack"] = string(stack)
				field["error"] = err
				logrus.WithFields(field).Error("panic recover")
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}

// stack returns a nicely formatted stack frame, skipping skip frames.
func stack(skip int) []byte {
	buf := new(bytes.Buffer) // the returned data
	// As we loop, we open files and read them. These variables record the currently
	// loaded file.
	var lines [][]byte
	var lastFile string
	for i := skip; ; i++ { // Skip the expected number of frames
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		// Print this much at least.  If we can't find the source, it won't show.
		fmt.Fprintf(buf, "%s:%d (0x%x)\n", file, line, pc)
		if file != lastFile {
			data, err := ioutil.ReadFile(file)
			if err != nil {
				continue
			}
			lines = bytes.Split(data, []byte{'\n'})
			lastFile = file
		}
		fmt.Fprintf(buf, "\t%s: %s\n", function(pc), source(lines, line))
	}
	return buf.Bytes()
}

// source returns a space-trimmed slice of the n'th line.
func source(lines [][]byte, n int) []byte {
	n-- // in stack trace, lines are 1-indexed but our array is 0-indexed
	if n < 0 || n >= len(lines) {
		return dunno
	}
	return bytes.TrimSpace(lines[n])
}

// function returns, if possible, the name of the function containing the PC.
func function(pc uintptr) []byte {
	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return dunno
	}
	name := []byte(fn.Name())
	// The name includes the path name to the package, which is unnecessary
	// since the file name is already included.  Plus, it has center dots.
	// That is, we see
	//	runtime/debug.*T·ptrmethod
	// and want
	//	*T.ptrmethod
	// Also the package path might contains dot (e.g. code.google.com/...),
	// so first eliminate the path prefix
	if lastslash := bytes.LastIndex(name, slash); lastslash >= 0 {
		name = name[lastslash+1:]
	}
	if period := bytes.Index(name, dot); period >= 0 {
		name = name[period+1:]
	}
	name = bytes.Replace(name, centerDot, dot, -1)
	return name
}
