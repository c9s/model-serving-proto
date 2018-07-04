PROTOC = protoc
GRPC_CPP_PLUGIN = grpc_cpp_plugin
GRPC_CPP_PLUGIN_PATH ?= `which $(GRPC_CPP_PLUGIN)`


PROTOS := $(wildcard *.proto)
PY_PROTOS := $(PROTOS:.proto=_pb2.py)
GO_PROTOS := $(PROTOS:.proto=.pb.go)
CPP_PROTOS := $(PROTOS:.proto=.pb.cc)
CPP_GRPC_PROTOS := $(PROTOS:.proto=.grpc.pb.cc)

PROTO_GEN_FILES := $(PY_PROTOS) $(GO_PROTOS) $(CPP_PROTOS) $(CPP_GRPC_PROTOS)

all: $(PROTO_GEN_FILES)

deps:
	go get github.com/golang/protobuf/protoc-gen-go

vpath %.proto

%_pb2.py: %.proto deps
	python -m grpc_tools.protoc -I. --python_out=. --grpc_python_out=. $<

%.pb.go: %.proto deps
	$(PROTOC) $< --go_out=plugins=grpc:.

%.grpc.pb.cc: %.proto
	$(PROTOC) -I. --grpc_out=. --plugin=protoc-gen-grpc=$(GRPC_CPP_PLUGIN_PATH) $<

%.pb.cc: %.proto
	$(PROTOC) -I. --cpp_out=. $<

clean:
	rm -rf $(PROTO_GEN_FILES)
