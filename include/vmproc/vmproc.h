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
 * 
 *  vmproc.h
 *  Created by Chen Zhen <cz.worker@gmail.com> on 2021/12/23.
 */

#ifndef __anvirt_vm_proc_h__
#define __anvirt_vm_proc_h__

#include "def.h"

__BEGIN_DECLS

typedef int (*anv_gui_main_loop_func)(void);
typedef int (*anv_qemu_main)(int argc, char** argv);
typedef void (*anv_vmproc_request_close)(void);

#ifndef GOLANG_API
ANV_VMPROC_API int anv_vmproc_main(anv_gui_main_loop_func main_loop, anv_qemu_main qemu_main);
ANV_VMPROC_API int anv_vmproc_check_ready(void);
ANV_VMPROC_API int anv_vmproc_boot_complete(void);
ANV_VMPROC_API char *anv_vmproc_get_name(void);
ANV_VMPROC_API void anv_vmproc_on_exit(void);
ANV_VMPROC_API int anv_vmproc_start_vm(anv_vmproc_request_close request_close);
#endif

__END_DECLS

#endif // __anvirt_vm_proc_h__
