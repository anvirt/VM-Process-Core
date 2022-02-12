#
#  VM Process Core Library
#  Copyright (C) 2021 AnVirt
#
#  This program is free software; you can redistribute it and/or modify
#  it under the terms of the GNU General Public License as published by
#  the Free Software Foundation; either version 2 of the License, or
#  (at your option) any later version.
#
#  This program is distributed in the hope that it will be useful,
#  but WITHOUT ANY WARRANTY; without even the implied warranty of
#  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
#  GNU General Public License for more details.
#
#  You should have received a copy of the GNU General Public
#  License along with this library; if not, see <http://www.gnu.org/licenses/>.
#

PROJECT_TOP ?= $(shell dirname `pwd`)
OUT_LIB := ${PROJECT_TOP}/output/lib

TAGS := ""
ifeq (${DEBUG},1)
	TAGS := "${TAGS} debug trace"
endif

.PHONY: build
build: proto
	@mkdir -p ${OUT_LIB}
	@go clean -cache
	@go build -a \
		-tags ${TAGS} \
		-buildmode=c-archive \
		-o ${OUT_LIB}/libvm-process.a \
		.

VM_PROCESS_PROTOCOL := vm-process-protocol

.PHONY: proto
proto: ${VM_PROCESS_PROTOCOL}
	@echo "> generate protobuf & gRPC src"
	@protoc \
		--go_out=. \
		--go-grpc_out=. \
		--proto_path=. \
		--proto_path=${VM_PROCESS_PROTOCOL}/ \
		${VM_PROCESS_PROTOCOL}/vmm.proto

${VM_PROCESS_PROTOCOL}:
	@git submodule update $@
