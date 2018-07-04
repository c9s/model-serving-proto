PROTOC = protoc
GRPC_CPP_PLUGIN = grpc_cpp_plugin
GRPC_CPP_PLUGIN_PATH ?= `which $(GRPC_CPP_PLUGIN)`

PROTOS := $(wildcard *.proto)
PY_PROTO_FILES := $(PROTOS:.proto=_pb2.py)
GO_PROTO_FILES := $(PROTOS:.proto=.pb.go)
CPP_PROTO_FILES := $(PROTOS:.proto=.pb.cc)
CPP_GRPC_PROTO_FILES := $(PROTOS:.proto=.grpc.pb.cc)

ALL_PROTO_FILES := $(PY_PROTO_FILES) $(GO_PROTO_FILES) $(CPP_PROTO_FILES) $(CPP_GRPC_PROTO_FILES)

all: $(ALL_PROTO_FILES)

py: $(PY_PROTO_FILES)

go: $(GO_PROTO_FILES)

cpp: $(CPP_PROTO_FILES)

protoc-gen-go:
	go get github.com/golang/protobuf/protoc-gen-go

vpath %.proto

%_pb2.py: %.proto
	python -m grpc_tools.protoc -I. --python_out=. --grpc_python_out=. $<

%.pb.go: %.proto protoc-gen-go
	$(PROTOC) $< --go_out=plugins=grpc:.

%.grpc.pb.cc: %.proto
	$(PROTOC) -I. --grpc_out=. --plugin=protoc-gen-grpc=$(GRPC_CPP_PLUGIN_PATH) $<

%.pb.cc: %.proto
	$(PROTOC) -I. --cpp_out=. $<

clean:
	rm -rf $(ALL_PROTO_FILES)

.phony: py go cpp
