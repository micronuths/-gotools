package xlog

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/micronuths/gotools/xlog/config"
)

const (
	//DEBUG is a constant of string type
	DEBUG = "DEBUG"
	INFO  = "INFO"
	WARN  = "WARN"
	ERROR = "ERROR"
	FATAL = "FATAL"
)

// Config is a struct which stores details for maintaining logs
type Config struct {
	LoggerLevel    string
	LoggerFile     string
	Writers        []string
	EnableRsyslog  bool
	RsyslogNetwork string
	RsyslogAddr    string

	LogFormatText bool
}

var configvar = DefaultConfig()
var m sync.RWMutex

// Writers is a map
var Writers = make(map[string]io.Writer)

// RegisterWriter is used to register a io writer
func RegisterWriter(name string, writer io.Writer) {
	m.Lock()
	Writers[name] = writer
	m.Unlock()
}

// DefaultConfig is a function which retuns config object with default configuration
func DefaultConfig() *Config {
	return &Config{
		LoggerLevel:    INFO,
		LoggerFile:     "",
		EnableRsyslog:  false,
		RsyslogNetwork: "udp",
		RsyslogAddr:    "127.0.0.1:5140",
		LogFormatText:  false,
	}
}

// Init is a function which initializes all config struct variables
func LagerInit(c Config) {
	if c.LoggerLevel != "" {
		configvar.LoggerLevel = c.LoggerLevel
	}

	if c.LoggerFile != "" {
		configvar.LoggerFile = c.LoggerFile
		configvar.Writers = append(configvar.Writers, "file")
	}

	if c.EnableRsyslog {
		configvar.EnableRsyslog = c.EnableRsyslog
	}

	if c.RsyslogNetwork != "" {
		configvar.RsyslogNetwork = c.RsyslogNetwork
	}

	if c.RsyslogAddr != "" {
		configvar.RsyslogAddr = c.RsyslogAddr
	}
	if len(c.Writers) == 0 {
		configvar.Writers = append(configvar.Writers, "stdout")

	} else {
		configvar.Writers = c.Writers
	}
	configvar.LogFormatText = c.LogFormatText
	RegisterWriter("stdout", os.Stdout)
	var file io.Writer
	var err error
	if configvar.LoggerFile != "" {
		file, err = os.OpenFile(configvar.LoggerFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			panic(err)
		}

	}
	for _, sink := range configvar.Writers {
		if sink == "file" {
			if file == nil {
				log.Panic("Must set file path")
			}
			RegisterWriter("file", file)
		}
	}
}

// NewLogger is a function
func NewLogger(component string) config.Logger {
	return NewLoggerExt(component, component)
}

// NewLoggerExt is a function which is used to write new logs
func NewLoggerExt(component string, appGUID string) config.Logger {
	var lagerLogLevel config.LogLevel
	switch strings.ToUpper(configvar.LoggerLevel) {
	case DEBUG:
		lagerLogLevel = config.DEBUG
	case INFO:
		lagerLogLevel = config.INFO
	case WARN:
		lagerLogLevel = config.WARN
	case ERROR:
		lagerLogLevel = config.ERROR
	case FATAL:
		lagerLogLevel = config.FATAL
	default:
		panic(fmt.Errorf("unknown logger level: %s", configvar.LoggerLevel))
	}
	logger := config.NewLoggerExt(component, configvar.LogFormatText)
	for _, sink := range configvar.Writers {

		writer, ok := Writers[sink]
		if !ok {
			log.Panic("Unknow writer: ", sink)
		}
		sink := config.NewReconfigurableSink(config.NewWriterSink(sink, writer, config.DEBUG), lagerLogLevel)
		logger.RegisterSink(sink)
	}

	return logger
}

func Debug(action string, data ...config.Data) {
	Logger.Debug(action, data...)
}

func Info(action string, data ...config.Data) {
	Logger.Info(action, data...)
}

func Warn(action string, data ...config.Data) {
	Logger.Warn(action, data...)
}

func Error(action string, err error, data ...config.Data) {
	Logger.Error(action, err, data...)
}

func Fatal(action string, err error, data ...config.Data) {
	Logger.Fatal(action, err, data...)
}

func Debugf(format string, args ...interface{}) {
	Logger.Debugf(format, args...)
}

func Infof(format string, args ...interface{}) {
	Logger.Infof(format, args...)
}

func Warnf(format string, args ...interface{}) {
	Logger.Warnf(format, args...)
}

func Errorf(err error, format string, args ...interface{}) {
	Logger.Errorf(err, format, args...)
}

func Fatalf(err error, format string, args ...interface{}) {
	Logger.Fatalf(err, format, args...)
}
