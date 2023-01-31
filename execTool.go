package main

import "github.com/S-doud/Stools/proto"

var toolMap map[string]func(...interface{})

func init() {
	toolMap = make(map[string]func(...interface{}))

	toolMap["-p"] = proto.MakeProto
}
