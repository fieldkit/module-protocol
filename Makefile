all: fk-module.proto.json fk-module.pb.go fk-module.pb.c fk-module.pb.h

node_modules/.bin/pbjs:
	npm install

fk-module.proto.json: node_modules/.bin/pbjs fk-module.proto
	pbjs fk-module.proto -t json -o fk-module.proto.json

fk-module.pb.c fk-module.pb.h: fk-module.proto
	protoc --nanopb_out=./ fk-module.proto

fk-module.pb.go: fk-module.proto
	protoc --go_out=./ fk-module.proto
