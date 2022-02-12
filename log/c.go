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

// #cgo CFLAGS: -I../include
// #cgo CFLAGS: -D__GOLANG_LOG__
// #include <base/log.h>
import "C"
import "fmt"

//export anv_log_log
func anv_log_log(level C.int, tag, msg *C.char) {
	if level == C.AnvLogLevelError {
		E(fmt.Sprintf("[%s] %s", C.GoString(tag), C.GoString(msg)))
	} else if level == C.AnvLogLevelWarn {
		W(fmt.Sprintf("[%s] %s", C.GoString(tag), C.GoString(msg)))
	} else if level == C.AnvLogLevelInfo {
		I(fmt.Sprintf("[%s] %s", C.GoString(tag), C.GoString(msg)))
	} else if level == C.AnvLogLevelDebug {
		D(fmt.Sprintf("[%s] %s", C.GoString(tag), C.GoString(msg)))
	} else if level == C.AnvLogLevelTrace {
		T(fmt.Sprintf("[%s] %s", C.GoString(tag), C.GoString(msg)))
	}
}
