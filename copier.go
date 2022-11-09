package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

var (
	cniBinSrcDir = os.Getenv("CNI_BIN_SRC")
	cniBinDstDir = os.Getenv("CNI_BIN_DST")
	override     = os.Getenv("OVERRIDE")
)

func main() {
	var err error

	if cniBinSrcDir == "" || cniBinDstDir == "" {
		if len(os.Args) > 2 {
			cniBinSrcDir = os.Args[1]
			cniBinDstDir = os.Args[2]
		} else {
			panic("empty CNI_BIN_SRC or CNI_BIN_DST")
		}
	}

	cniBinSrcDir, err = filepath.Abs(cniBinSrcDir)
	if err != nil {
		panic(fmt.Errorf("abs CNI_BIN_SRC:%v", err))
	}
	cniBinDstDir, err = filepath.Abs(cniBinDstDir)
	if err != nil {
		panic(fmt.Errorf("abs CNI_BIN_DST:%v", err))
	}

	inputs, err := os.ReadDir(cniBinSrcDir)
	if err != nil {
		panic(fmt.Errorf("err list plugins in CNI_BIN_SRC:%v", err))
	}
	exists, err := os.ReadDir(cniBinDstDir)
	if err != nil {
		panic(fmt.Errorf("err list plugins in CNI_BIN_DST:%v", err))
	}

	m := map[string]bool{}

	for _, plugin := range exists {
		if plugin.IsDir() || !plugin.Type().IsRegular() {
			continue
		}
		m[plugin.Name()] = true
	}

	for _, plugin := range inputs {
		if plugin.IsDir() || !plugin.Type().IsRegular() {
			continue
		}
		if m[plugin.Name()] && override == "" {
			continue
		}
		cpCmd := exec.Command("cp", filepath.Join(cniBinSrcDir, plugin.Name()), filepath.Join(cniBinDstDir, plugin.Name()))
		if err = cpCmd.Run(); err != nil {
			fmt.Println("unable to copy", plugin.Name(), err)
		} else {
			fmt.Println("copied", plugin.Name())
		}
	}
}
