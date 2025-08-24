package main

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/sirupsen/logrus"
)

func main() {

	fmt.Println("################ Strings Package in Go ################")
	text := "This is Golang Advanced Course From KodeKloud"
	result := strings.Contains(text, "Golang")
	fmt.Println(result)
	result1 := strings.ReplaceAll(text, "Golang", "python")
	fmt.Println(result1)
	result2 := strings.Count(text, "Golang")
	fmt.Println(result2)

	fmt.Println("################ I/O Package in Go ###################")
	r := strings.NewReader("Learning Golang")
	buf := make([]byte, 4)
	for {
		n, err := r.Read(buf)
		fmt.Println(string(buf[:n]), err)
		if err != nil {
			fmt.Println("Breaking")
			break
		}
	}

	r2 := strings.NewReader("Learning python")
	_, err := io.Copy(os.Stdout, r2)
	if err != nil {
		// log.Fatal(err)
		fmt.Println(err)
	}
	fmt.Println("################ Error handling Package in Go ###################")
	check_error(process(2))
	fmt.Println("################ Crypto Package in Go ###################")
	var pass string
	pass = "KodeKloud"
	fmt.Println(encodingwithmd5(pass))
	fmt.Println("################ file handling Package in Go ###################")
	// path := filepath.Join("D:/", "atiia_practice", "mail.md")
	// fmt.Println(path)
	// fmt.Println(filepath.Dir(path))
	// fmt.Println(filepath.Base(path))
	// fmt.Println(filepath.Ext(path))
	// fmt.Println(filepath.IsAbs(path))
	// fmt.Println("################ file Info Package in Go ###################")
	// fileInfo, err4 := os.Stat(path)
	// if err4 != nil {
	// 	fmt.Println(err4)
	// }
	// fmt.Println(fileInfo.Name())
	// fmt.Println(fileInfo.Size())
	// fmt.Println(fileInfo.Mode())
	// fmt.Println(fileInfo.IsDir())
	// fmt.Println("################ Read file in Go ###################")
	// data, err5 := os.ReadFile(path)
	// if err5 != nil {
	// 	fmt.Println(err5)
	// }
	// fmt.Println(string(data))

	// file, err6 := os.Open(path)
	// if err6 != nil {
	// 	fmt.Println(err6)
	// } else {
	// 	buffer := make([]byte, 5)
	// 	for {
	// 		n, err6 := file.Read(buffer)
	// 		if err6 != nil {
	// 			fmt.Println(err6)
	// 			break
	// 		}
	// 		fmt.Println(string(buffer[:n]))
	// 	}
	// }
	fmt.Println("################ Write file in Go ###################")
	// path2 := filepath.Join("D:/", "atiia_practice", "test.txt")
	// file, err7 := os.OpenFile(path2, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0660)
	// if err7 != nil {
	// 	fmt.Println(err7)
	// }
	// _, err8 := file.WriteString("Testing")
	// if err8 != nil {
	// 	fmt.Println(err8)
	// }
	fmt.Println("################ Logging in Go ###################")
	log.Println("Hello")
	path3 := filepath.Join("/media", "aattia", "DATA", "projects", "Go-practice", "test.md")
	fmt.Println(path3)
	file, err9 := os.OpenFile(path3, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err9 != nil {
		fmt.Println(err9)
	}
	log.SetOutput(file)
	log.Println("write log to a file")
	log.Println("write log2 to a file")

	logrus.Println("Log using Logrus")
	logrus.SetOutput(file)
	// logrus.Fatal("This is fatal log")
	// logrus.Panic("This is panic log")
	logrus.Error("This is error log")
	// logrus.Warn("This is Warn log")
	// logrus.Info("This is Info log")
	fmt.Println("################ Sorting in Go ###################")
	nums := []int{1, 7, 5}
	sort.Ints(nums)
	fmt.Println(nums)
	letters := []string{"b", "s", "a"}
	sort.Strings(letters)
	fmt.Println(letters)
	fmt.Println("################ Testing in Go ###################")

}

func process(i int) error {
	if i%2 == 0 {
		return errors.New("Custome Error")
	}
	return nil
}

func check_error(e error) {
	if e != nil {
		fmt.Println(e)
	}
	fmt.Println("Done")

}

func encodingwithmd5(str string) string {
	hashcode := md5.Sum([]byte(str))
	return hex.EncodeToString(hashcode[:])
}
func checkeven(i int) bool {
	if i%2 == 0 {
		return true
	}
	return false
}
