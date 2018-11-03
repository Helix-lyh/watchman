package utils

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func Run(strCmd string) (info string) {
	runResult := exec.Command("/bin/bash", "-c", strCmd)
	stdout, err := runResult.StdoutPipe()
	if err != nil {
		fmt.Printf("Error:can not obtain stdout pipe for command:%s\n", err)
		return
	}

	//执行命令
	if err := runResult.Start(); err != nil {
		fmt.Println("Error:The command is err,", err)
		return
	}

	//读取所有输出
	byteInfo, err := ioutil.ReadAll(stdout)
	if err != nil {
		fmt.Println("ReadAll Stdout:", err.Error())
		return
	}

	if err := runResult.Wait(); err != nil {
		fmt.Println("wait:", err.Error())
		return
	}
	return string(byteInfo)
}
