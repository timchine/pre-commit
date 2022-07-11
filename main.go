package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	b, err := ioutil.ReadFile("version")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("old version", string(b))
	file, err := os.OpenFile("version", os.O_RDWR|os.O_TRUNC, 0777)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	bs := bytes.Split(b, []byte{'.'})
	lb := bs[len(bs)-1]
	i, err := strconv.Atoi(string(lb))
	if err != nil {
		fmt.Println(err)
		return
	}
	bs[len(bs)-1] = []byte(fmt.Sprintf("%d", i+1))
	tb := bytes.Join(bs, []byte{'.'})
	fmt.Println("new version", string(tb))
	file.Write(tb)
}

