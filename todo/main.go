package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/Hephesto0X01/Go-tools/todo/utils"
)

var p = fmt.Println

func main() {
	var lines []string
	var final []string

	// 读取命令行文件路径参数
	defaultPath := "C:/Users/28998/Desktop/TODO.txt"
	filePath := flag.String("f", defaultPath, "指定要读取的文件路径，如果不指定则使用默认路径 "+defaultPath)
	briefPrint := flag.Bool("b", false, "只打印任务列表")
	flag.Parse()

	// 读取文件内容并打印
	file, err := os.Open(*filePath)
	utils.CheckErr(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}
		lines = append(lines, scanner.Text())
	}
	utils.CheckErr(scanner.Err())
	final = lines

	// 打印任务列表的简要信息
	if *briefPrint {
		final = utils.GetBriefLines(lines)
	}

	// 打印完整任务列表
	for index, line := range final {
		p(strconv.Itoa(index)+":", line)
	}
	p("total missions:", len(lines))

	// 读取用户输入的索引，并打开对应的url
	for {
		var vaildURLs []string
		fmt.Printf("input the index you wanna open: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()
		if input == "" {
			p("bye")
			break
		}
		if input == "all" {
			vaildURLs = lines
		} else {
			indexes := strings.Fields(input)
			for _, idx := range indexes {
				index, err := strconv.Atoi(idx)
				if err != nil || index < 0 || index >= len(lines) {
					p("invalid index:", idx)
					continue
				}
				vaildURLs = append(vaildURLs, lines[index])
			}
		}
		for _, url := range vaildURLs {
			p("opening url:", url)
			err := exec.Command("cmd", "/c", "start", "msedge", url).Run()
			if err != nil {
				p("open url failed:", err)
				utils.CheckErr(err)
			}
		}
	}
}
