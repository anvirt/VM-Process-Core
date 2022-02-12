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

package main

// #cgo CFLAGS: -Iinclude
/*
#include "vmproc.h"
static inline int do_gui_main_loop(anv_gui_main_loop_func main_loop) {
  return main_loop();
}

static inline void do_request_close(anv_vmproc_request_close fn) {
	fn();
}

*/
import "C"
import (
	"anvirt/vm_process/core/log"
	"anvirt/vm_process/core/process"
	"anvirt/vm_process/core/qemu"
	"unsafe"
)

//export anv_vmproc_main
func anv_vmproc_main(main_loop C.anv_gui_main_loop_func, qemu_main C.anv_qemu_main) C.int {
	log.I("[vm process core] init")
	if err := process.Global().Initialize(); err != nil {
		log.W2("[vm process core] init failed: %v", err)
		return -1
	}
	qemu.SetMain(unsafe.Pointer(qemu_main))
	return C.do_gui_main_loop(main_loop)
}

//export anv_vmproc_check_ready
func anv_vmproc_check_ready() int {
	if err := process.Global().CheckReady(); err != nil {
		log.W2("check ready failed: %v", err)
		return -1
	} else {
		return 0
	}
}

//export anv_vmproc_get_name
func anv_vmproc_get_name() *C.char {
	return C.CString(process.Global().GetName())
}

//export anv_vmproc_on_exit
func anv_vmproc_on_exit() {
	process.Global().OnExit()
}

//export anv_vmproc_start_vm
func anv_vmproc_start_vm(request_close_fn C.anv_vmproc_request_close) int {
	reqClose := func() {
		C.do_request_close(request_close_fn)
	}
	if err := process.Global().StartVM(reqClose); err != nil {
		return -1
	}
	return 0
}
