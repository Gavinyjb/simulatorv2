package shell

import (
	"fmt"
	"os/exec"
)

func GetCacaheInfo(command string)string{
	fmt.Println("进入shell函数")
	//command :="/home/gavin/trace/tools/05-analyer -input /mnt/hgfs/FB_2009_0_DNlog\\(DEBUG\\)/jsonDir/slave4rKV.json"
	cmd := exec.Command("/bin/bash", "-c", command)

	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Execute Shell:%s failed with error:%s", command, err.Error())
		return err.Error()
	}
	return string(output)
}
