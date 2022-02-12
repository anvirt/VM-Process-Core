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

package process

import (
	"anvirt/vm_process/core/log"
	"anvirt/vm_process/core/qemu"
	"anvirt/vm_process/core/vmm"
	"flag"
	"os"
)

var process_name string

func init() {
	flag.StringVar(&process_name, "name", "", "")
}

type RequestCloseFunc = func()

type VMProcess struct {
	// VMM agent (grpc client)
	agent vmm.VMMAgent
	// QEMU agent
	qemu_agent qemu.QEMUAgent
	// name from vmm
	name string

	request_close RequestCloseFunc
}

func (proc *VMProcess) Initialize() (err error) {
	flag.Parse()

	proc.name = process_name

	token := os.Getenv("__VM_PROCESS_TOKEN__")
	err = proc.agent.InitAgent(token)
	return
}

func (proc *VMProcess) GetName() string { return proc.name }

func (proc *VMProcess) CheckReady() (err error) {
	if err = proc.agent.ReportStatus(0, "i am ready"); err == nil {
		proc.agent.HandleEvent(proc)
	}
	return
}

func (proc *VMProcess) OnEvent(t int, payload string) {
	switch t {
	case 19999:
		if proc.request_close != nil {
			proc.request_close()
		}
	default:
		log.W2("TODO: recv event: [%d] %s", t, payload)
	}
}

func (proc *VMProcess) OnExit() {
	_ = proc.agent.ReportStatus(9999, "bye")
}

func (proc *VMProcess) StartVM(request_close_fn RequestCloseFunc) (err error) {
	var args []string
	if args, err = proc.agent.GetStartArgs(); err != nil {
		return
	}
	proc.request_close = request_close_fn
	return proc.qemu_agent.Start(args)
}

//////////////////////////////////////////////////////////////////////////////////////////
// singleton
//////////////////////////////////////////////////////////////////////////////////////////

var global_process VMProcess

func Global() *VMProcess { return &global_process }
