package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
)

var inputfile *string = flag.String("i", "inputfile", "the name of the input file, ABS path")
var outputfile *string = flag.String("o", "outputfile", "the name of the outputfile , ABS path")
var option *string = flag.String("O", "option", "option")

func main() {
	flag.Parse()
	if inputfile != nil {
		values, err := ReadValues(*inputfile)
		if err != nil {
			fmt.Println("Error: ", err)
		} else {
			fmt.Println("the values are", values)
		}
		values, err = Popsort(values)
		if err != nil {
			fmt.Println("Pop sort error", err)
		}
		WriteValues(values, *outputfile)
	}
}
func ReadValues(inputfile string) (values []int, err error) {
	file, err := os.Open(inputfile)
	if err != nil {
		fmt.Println("Can not open filename: ", inputfile)
		return
	}
	defer file.Close()
	br := bufio.NewReader(file)
	values = make([]int, 0) //在ReadValues处以返回参数定义
	for {
		line, isPrefix, err1 := br.ReadLine()
		if err1 != nil {
			if err1 != io.EOF {
				err = err1
			}
			break
		}
		if isPrefix {
			fmt.Println("The line is too long...")
			return
		}
		str := string(line)
		value, err1 := strconv.Atoi(str)
		if err1 != nil {
			err = err1
			return
		}
		values = append(values, value)
	}
	return
}

func WriteValues(values []int, outputfile string) error {
	file, err := os.Create(outputfile)
	if err != nil {
		fmt.Println("can not create file :", outputfile)
		return err
	}
	defer file.Close()
	for _, value := range values {
		str := strconv.Itoa(value)
		file.WriteString(str + "\r\n")
	}
	return nil
}
func Popsort(values []int) (values1 []int, err error) {
	for i := 0; i < len(values); i++ {
		for j := i + 1; j < len(values); j++ {
			if values[i] > values[j] {
				temp := values[i]
				values[i] = values[j]
				values[j] = temp
			}
		}
	}
	return values, nil
}
