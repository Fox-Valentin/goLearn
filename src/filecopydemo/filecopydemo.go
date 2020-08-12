package filecopydemo

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func fileExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}
func copyFileAction(force, showProgress bool, src, dist string) {
	if !force {
		if fileExist(dist) {
			fmt.Println("%s exists, override? y/n", dist)
			reader := bufio.NewReader(os.Stdin)
			data, _, _ := reader.ReadLine()
			if strings.TrimSpace(string(data)) == "y" {
				return
			}
		}
	}
	copyFile(src, dist)
	if showProgress {
		fmt.Printf("%s -> %s", src, dist)
	}
}
func copyFile(src, dist string) (w int64, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer srcFile.Close()

	distFile, err := os.Create(dist)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer distFile.Close()
	return io.Copy(srcFile, distFile)
}
func OsStind() {
	var showProgress, force bool
	flag.BoolVar(&showProgress, "v", false, "explain what is being done.")
	flag.BoolVar(&force, "f", false, "force coping when existing.")
	flag.Parse()
	if flag.NArg() < 2 {
		flag.Usage()
		return
	}
	copyFileAction(showProgress, force, flag.Arg(0), flag.Arg(1))
}
