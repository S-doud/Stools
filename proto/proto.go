package proto

import (
	"fmt"
	"os/exec"
)

func CheckExist() error {
	_, err := exec.LookPath("protoc")
	if err != nil {
		fmt.Println("protoc 不存在")
		return err
	}
	return nil
}

func MakeProto(...interface{}) {
	err := CheckExist()
	if err != nil {
		return
	}
	cmd := exec.Command("protoc", "--go_out=./", "*.proto")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("cmd protoc exec error")
	}
	fmt.Printf("combined out:\n%s\n", string(out))
}
