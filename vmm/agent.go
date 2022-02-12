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

package vmm

import (
	"anvirt/vm_process/core/log"
	"anvirt/vm_process/core/protocol/vmm"
	"context"
	"fmt"
	"path/filepath"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type VMMAgent struct {
	token string

	grpc_conn *grpc.ClientConn
	grpc_cli  vmm.VMMClient

	event_delegate EventDelegate
}

func (agent *VMMAgent) InitAgent(token string) (err error) {
	// set token
	agent.token = token

	// init grpc client
	var opts []grpc.DialOption
	{
		// TODO: 未加密的连接
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	socket_file, _ := filepath.Abs("Documents/process/process.sock")
	if agent.grpc_conn, err = grpc.Dial(fmt.Sprintf("unix://%s", socket_file), opts...); err == nil {
		agent.grpc_cli = vmm.NewVMMClient(agent.grpc_conn)
	}
	return
}

func (agent *VMMAgent) HandleEvent(delegate EventDelegate) (err error) {
	agent.event_delegate = delegate
	if cli, grpc_err := agent.grpc_cli.HandleEvent(context.Background(), &vmm.HandleEventRequest{Token: agent.token}); grpc_err != nil {
		err = fmt.Errorf("can not handle event: %v", grpc_err)
	} else {
		go agent.handle_event_internal(cli)
	}
	return
}

func (agent *VMMAgent) Close() {
	agent.grpc_conn.Close()
}

func (agent *VMMAgent) handle_event_internal(cli vmm.VMM_HandleEventClient) {
	// TODO: how to stop/close
	for {
		msg, err := cli.Recv()
		if err != nil {
			log.W2("event stream broken: %v", err)
			return
		}
		agent.event_delegate.OnEvent(int(msg.Type), msg.Payload)
	}
}

func (agent *VMMAgent) ReportStatus(code int, status string) (err error) {
	if _, grpc_err := agent.grpc_cli.ReportStatus(context.Background(), &vmm.ReportStatusRequest{
		Token:      agent.token,
		StatusCode: int32(code),
		Status:     status,
	}); grpc_err != nil {
		err = fmt.Errorf("report status failed: %v", grpc_err)
	}
	return
}

func (agent *VMMAgent) GetStartArgs() (args []string, err error) {
	if reply, ge := agent.grpc_cli.GetStartArgs(context.Background(), &vmm.GetStartArgsRequest{Token: agent.token}); ge != nil {
		err = fmt.Errorf("get args failed: %v", ge)
	} else {
		args = append(args, reply.Args...)
	}
	return
}
