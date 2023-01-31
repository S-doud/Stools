package main

import "github.com/S-doud/tools/proto"

var toolMap map[string]func(...interface{})

func init() {
	toolMap = make(map[string]func(...interface{}))

	toolMap["protoc"] = proto.MakeProto
}
