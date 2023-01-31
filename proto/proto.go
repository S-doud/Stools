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

//MakeProto 生成proto文件
func MakeProto(file ...interface{}) {

	switch file[0].([]string)[0] {

	case "init":
		ProInit()

	default:
		MakeProtoFile(file[0].([]string)[0])

	}

}

//ProInit 生成proto文件
func ProInit() {

	cmd := exec.Command("go", "install", "google.golang.org/protobuf/cmd/protoc-gen-go@latest")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("cmd go exec error")
	}
	fmt.Println(string(out))

	cmd = exec.Command("go", "install", "google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest")
	out, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Println("cmd go exec error")
	}
	fmt.Println(string(out))
}

//MakeProtoFile 生成proto文件
func MakeProtoFile(file string) {
	err := CheckExist()
	if err != nil {
		return
	}
	cmd := exec.Command("protoc", "--go_out=.", "--go-grpc_out=.", file)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("cmd protoc exec error")
	}
	fmt.Printf("combined out:\n%s\n", string(out))
}
