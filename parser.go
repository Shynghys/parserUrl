// Входные данные:
// На стандартный вход программе подаётся список URL, по одному в каждой строке, которая терминируется символом новой строки. Обработка (выполнение запросов) должна быть распараллелена по числу присутствующих в системе вычислительных ядер.

// Выходные данные:
// Необходимо для каждого URL определить код ответа HTTP на запрос GET, размер документа и время между отправкой запроса и получением ответа. Печать результатов должен производиться на стандартный выход в формате CSV, пример: "http://example.com/;200;1560;120ms".

// Возникающие ошибки при обращении к сайту, нужно отображать на стандатный выход.

// По завершении программы либо через Ctrl+C либо при закрытии входа, необходимо вывести статистику работы программы в виде количества отработанных запросов в разрезе параллельного исполнения, в виде "<порядковый номер паралельного потока запросов>:<число запросов>".

package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {

	file, err := os.Open("read.txt")
	if err != nil {
		panic(err)
	}
	csvfile, err := os.Create("result.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := csv.NewWriter(csvfile)
	defer writer.Flush()

	scanner := bufio.NewScanner(file)
	slice := []string{"URL name", "Status code", "Content length", "time"}

	err = writer.Write(slice)
	if err != nil {
		panic(err)
	}

	for scanner.Scan() {
		slice = make([]string, 0)
		fmt.Print("--")
		res, err := http.Get(scanner.Text())
		if err != nil {
			// handle error
			panic(err)
		}

		defer res.Body.Close()

		// start := time.Now()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}
		// elapsed := time.Since(start)
		text := scanner.Text()
		url := text[8:len(text)]
		slice = append(slice, url)

		l := len(body)
		slice = append(slice, strconv.Itoa(res.StatusCode))

		// slice = append(slice, strconv.Itoa(elapsed)
		slice = append(slice, strconv.Itoa(l))
		// fmt.Println(res.StatusCode)
		// fmt.Println(l)
		// log.Println(elapsed)
		err = writer.Write(slice)
		if err != nil {
			panic(err)
		}
		slice = make([]string, 0)

	}
	log.Println("done")
	// conn, err := net.Dial("tcp", "example.com:80")
	// if err != nil {
	// 	panic(err)
	// }
	// defer conn.Close()
	// conn.Write([]byte("GET / HTTP/1.0\r\n\r\n"))

	// start := time.Now()
	// oneByte := make([]byte, 1)
	// _, err = conn.Read(oneByte)
	// if err != nil {
	// 	panic(err)
	// }
	// a := time.Since(start)

	// _, err = ioutil.ReadAll(conn)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	// b := time.Since(start)
	// log.Println(b - a)
}
