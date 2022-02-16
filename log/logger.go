/*
 *  VM Process Core Library
 *  Copyright (C) 2021 AnVirt
 *
 *  This program is free software; you can redistribute it and/or modify
 *  it under the terms of the GNU General Public License as published by
 *  the Free Software Foundation; either version 2 of the License, or
 *  (at your option) any later version.
 *
 *  This program is distributed in the hope that it will be useful,
 *  but WITHOUT ANY WARRANTY; without even the implied warranty of
 *  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *  GNU General Public License for more details.
 *
 *  You should have received a copy of the GNU General Public
 *  License along with this library; if not, see <http://www.gnu.org/licenses/>.
 */

package log

import (
	"os"
	"runtime"

	log "github.com/sirupsen/logrus"
)

const (
	LOG_DEPTH int = 2
)

var (
	// info  *log.Logger
	// warn  *log.Logger
	// err   *log.Logger
	// debug *log.Logger

	_log *log.Logger

	_log_if ILogger
)

type LogFile struct {
	f *os.File
}

func (file *LogFile) Write(p []byte) (n int, err error) {
	return file.f.Write(p)
}

// 初始化日志
func init() {
	// TODO: 日志文件路径应该在某处设定？
	Log2File(os.Stdout)
}

func Log2File(outFile *os.File) {
	// flag := log.LstdFlags | log.Lmicroseconds
	// info = log.New(outFile, "[INFO] ", flag|log.Lshortfile)
	// warn = log.New(outFile, "[WARN] ", flag)
	// err = log.New(outFile, "[ERROR] ", flag|log.Lshortfile)
	// debug = log.New(outFile, "[DEBUG] ", flag|log.Lshortfile)
	_log = log.New()
	// logrus内部判断输出文件是否为terminal时，在intel的mac上，会导致一个栈对齐的问题，暂时不知道怎么解决。包装一层Writer可以跳过这个判断
	if runtime.GOOS == "darwin" && runtime.GOARCH == "amd64" {
		_log.SetOutput(&LogFile{f: outFile})
	} else {
		_log.SetOutput(outFile)
	}
	_log.SetLevel(log_level)
}

// 日志，级别：跟踪
func T(v ...interface{}) {
	if _log != nil {
		_log.Traceln(v...)
	}
}

// 日志，级别：跟踪
func T2(format string, v ...interface{}) {
	if _log != nil {
		_log.Tracef(format, v...)
	}
}

// 日志，级别：信息
func I(v ...interface{}) {
	if _log != nil {
		_log.Infoln(v...)
	}
}

// 日志，级别：信息
func I2(format string, v ...interface{}) {
	if _log_if != nil {
		_log_if.I(format, v...)
	} else if _log != nil {
		_log.Infof(format, v...)
	}
}

// 日志，级别：警告
func W(v ...interface{}) {
	if _log != nil {
		_log.Warnln(v...)
	}
}

// 日志，级别：警告
func W2(format string, v ...interface{}) {
	if _log_if != nil {
		_log_if.W(format, v...)
	} else if _log != nil {
		_log.Warnf(format, v...)
	}
}

// 日志，级别：错误。输出并退出程序
func E(v ...interface{}) {
	if _log != nil {
		// err.Fatalln(v...)
		_log.Fatalln(v...)
		//os.Exit(1)
	}
}

// 日志，级别：错误。输出并退出程序
func E2(format string, v ...interface{}) {
	if _log_if != nil {
		_log_if.E(format, v...)
	} else if _log != nil {
		// err.Fatalln(v...)
		_log.Fatalf(format, v...)
		//os.Exit(1)
	}
}

// 日志，级别：信息。带有字段信息
func F(fields map[string]interface{}, v ...interface{}) {
	_log.WithFields(fields).Infoln(v...)
}

func Logger() *log.Logger {
	return _log
}
