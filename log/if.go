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

// 外部日志接口，提供自定义日志接口的能力

package log

type ILogger interface {
	E(format string, v ...interface{})
	W(format string, v ...interface{})
	I(format string, v ...interface{})
	D(format string, v ...interface{})
	T(format string, v ...interface{})
}

type ILoggerFactory = func(tag string) ILogger

var logger_factory ILoggerFactory = nil

func RegisterLogger(factory ILoggerFactory) {
	logger_factory = factory
	if _log_if == nil {
		_log_if = logger_factory("default")
	}
}
