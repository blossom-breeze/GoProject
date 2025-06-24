package main

import (
	"fmt"
	"sync"
)

func main() {
	letterChan := make(chan struct{})
	numberChan := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(2)

	// 打印字母的协程
	go func() {
		defer wg.Done()
		for c := 'A'; c <= 'Z'; c++ {
			<-letterChan
			fmt.Printf("%c", c)
			numberChan <- struct{}{}
		}
	}()

	// 打印数字的协程
	go func() {
		defer wg.Done()
		for i := 1; i <= 26; i++ {
			<-numberChan
			fmt.Print(i)
			if i < 26 {
				letterChan <- struct{}{}
			}
		}
	}()

	// 启动打印
	letterChan <- struct{}{}

	wg.Wait()
	fmt.Println() // 换行
}
