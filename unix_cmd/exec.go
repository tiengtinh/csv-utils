package unix_cmd

import (
	"errors"
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	"github.com/golang/glog"
)

func Exec(cmd string) (string, error) {
	// splitting head => g++ parts => rest of the command
	parts := strings.Fields(cmd)
	head := parts[0]
	parts = parts[1:len(parts)]

	out, err := exec.Command(head, parts...).CombinedOutput()
	if err != nil {
		return "", errors.New(fmt.Sprintf("Error[%s] while executing cmd[%s]", err, cmd))
	}

	return string(out), nil
}

func DelFirstLine(filepath string) error {
	cmdStr := fmt.Sprintf(`sed -i 1d %s`, filepath)
	if _, err := Exec(cmdStr); err != nil {
		return errors.New(fmt.Sprintf("Error[%s] while removing headline", err))
	}
	return nil
}

func Sort(
	srcFilepath, outFilepath string,
	comma rune,
	keys []int,
) error {
	//http://stackoverflow.com/questions/1037365/unix-sort-with-tab-delimiter
	if comma == "\t" {
		comma = "$'\t'"
	}

	cmdStr := fmt.Sprintf(`sort --ignore-leading-blanks -t%s -k%s -o %s %s`,
		comma, keysCmd(keys), outFilepath, srcFilepath)
	if out, err := Exec(cmdStr); err != nil {
		glog.Infof("Error[%s] while sorting files[%s] with output[%s]", err, srcFilepath, out)
		return err
	}
	return nil
}

func keysCmd(keys []int) (cmd string) {
	var keyCmds []string
	for _, key := range keys {
		if key < 0 {
			keyCmds = append(keyCmds, strconv.Itoa(-key)+"r")
		} else {
			keyCmds = append(keyCmds, strconv.Itoa(-key))
		}
	}

	cmd = strings.Join(keyCmds, ",")

	return cmd
}

//http://www.ibm.com/developerworks/library/l-tiptex6/
//http://stackoverflow.com/questions/5429840/eliminate-duplicate-lines-and-keep-the-last-one
//tac temp.txt | sort -k2,2 -r -u
//http://unix.stackexchange.com/questions/30173/how-to-remove-duplicate-lines-inside-a-text-file
func Unique() {

}
