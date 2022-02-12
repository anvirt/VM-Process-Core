//go:build debug

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
	log "github.com/sirupsen/logrus"
)

var (
	log_level = log.DebugLevel
)

// 日志，级别：调试。只在测试环境输出
func D(v ...interface{}) {
	// 只在测试环境生效
	if _log != nil {
		// debug.Println(v...)
		// debug.Output(LOG_DEPTH, fmt.Sprintln(v...))
		_log.Debugln(v...)
	}
}

// 日志，级别：调试。只在测试环境输出
func D2(format string, v ...interface{}) {
	// 只在测试环境生效
	if _log != nil {
		// debug.Println(v...)
		// debug.Output(LOG_DEPTH, fmt.Sprintf(format, v...))
		_log.Debugf(format, v...)
	}
}
