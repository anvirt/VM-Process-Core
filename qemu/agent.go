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

package qemu

/*
#include <stdlib.h>
typedef char** qemu_args;
typedef int (*anv_qemu_main)(int argc, char** argv);

static inline int start_qemu_with_args(anv_qemu_main qemu_main, int argc, qemu_args argv) {
	return qemu_main(argc, argv);
}

static inline qemu_args make_args(int argc) {
	return calloc(argc, sizeof(char *));
}

static inline void put_arg(qemu_args argv, int i, char *arg) {
	argv[i] = arg;
}

static inline void free_args(int argc, qemu_args argv) {
	for (int i = 0; i < argc; ++i) if (argv[i]) free(argv[i]);
	free(argv);
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

type QEMUAgent struct {
}

func (agent *QEMUAgent) Start(args []string) (err error) {
	argc := len(args)
	if argc == 0 {
		return fmt.Errorf("no args to run qemu")
	}

	argv := C.make_args(C.int(argc))
	for i, arg := range args {
		C.put_arg(argv, C.int(i), C.CString(arg))
	}
	defer C.free_args(C.int(argc), argv)

	if rc := C.start_qemu_with_args(qemu_main, C.int(argc), argv); rc != 0 {
		return fmt.Errorf("recv error from qemu: %d", rc)
	}
	return
}

var qemu_main C.anv_qemu_main

func SetMain(fn unsafe.Pointer) { qemu_main = (C.anv_qemu_main)(fn) }
