PROTOS := $(wildcard *.proto)
PY_PROTOS := $(PROTOS:.proto=_pb2.py)
GO_PROTOS := $(PROTOS:.proto=.pb.go)

%_pb2.py: %.proto
	python -m grpc_tools.protoc -I. --python_out=. --grpc_python_out=. $<

%.pb.go: %.proto
	protoc $< --go_out=plugins=grpc:.

all: $(PY_PROTOS) $(GO_PROTOS)

clean:
	rm -rf $(PY_PROTOS) $(GO_PROTOS)
