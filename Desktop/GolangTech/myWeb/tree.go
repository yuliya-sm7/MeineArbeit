package main

import (
	"io"
	"os"
	"strconv"
)

var flag = make([]bool, 0, 10)

func dirTree(out io.Writer, path string, printFiles bool) error {
	dir, err := os.Open(path)
	ls, _ := dir.Readdir(-1)
	if printFiles {
		for i := range ls {
			for j := 0; j < len(flag); j++ {
				if flag[j] == true {
					out.Write([]byte("│"))
				}
				out.Write([]byte("\t"))
			}
			if i == len(ls)-1 {
				out.Write([]byte("└"))
			} else {
				out.Write([]byte("├"))
			}
			if ls[i].IsDir() {
				out.Write([]byte("───" + ls[i].Name() + "\n"))
				if i == len(ls)-1 {
					flag = append(flag, false)
				} else {
					flag = append(flag, true)
				}
				dirTree(out, path+"/"+ls[i].Name(), printFiles)
			} else {
				out.Write([]byte("───" + ls[i].Name()))
				if ls[i].Size() == 0 {
					out.Write([]byte(" (empty)" + "\n"))
				} else {
					out.Write([]byte("(" + strconv.Itoa(int(ls[i].Size())) + "b)" + "\n"))
				}
			}
		}
	} else {
		for i := range ls {
			if ls[i].IsDir() {
				for j := 0; j < len(flag); j++ {
					if flag[j] == true {
						out.Write([]byte("│"))
					}
					out.Write([]byte("\t"))
				}
				if i == len(ls)-1 || !ls[i+1].IsDir() {
					out.Write([]byte("└"))

				} else {
					out.Write([]byte("├"))
				}
				if i == len(ls)-1 {
					flag = append(flag, false)
				} else {
					flag = append(flag, true)
				}
				out.Write([]byte("───" + ls[i].Name() + "\n"))
				dirTree(out, path+"/"+ls[i].Name(), printFiles)
			}
		}
	}
	if len(flag) >= 1 {
		flag = flag[0 : len(flag)-1]
	}
	return err
}

func find_rune_number(str, char string) []int { // функция принимает (строку и символ) и выдаёт масиив индексов символа
	slice1 := []int{}               // создать массив
	for i := 0; i < len(str); i++ { // цыкл поиска символов
		j := i + 1
		if str[i:j] == char {
			slice1 = append(slice1, i) // добавить в массив номер искомого символа
		}
	}
	arr := slice1
	return arr
}
