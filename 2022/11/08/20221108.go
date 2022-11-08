// 20221108
// 协程 交替打印数字和字⺟

// 问题描述：
// 使⽤两个 goroutine 交替打印序列，⼀个 goroutine 打印数字， 另外⼀个  goroutine 打印字⺟
// 最终效果如下：
// 12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728

package day

import "fmt"

func F20221108() string {
	str := make(chan string)

	c := make(chan struct{})
	n := make(chan struct{})

	// 信道在试用前必须创建，发送和接受操作在另一端准备好之前都会阻塞
	// 注意协程开始与退出
	go func() {
		n <- struct{}{}
	}()

	go func() {
		for i := 0; i <= 26; i += 2 {
			<-n
			str <- fmt.Sprintf("%d%d", i+1, i+2)
			c <- struct{}{}
		}
		close(str)
	}()

	go func() {
		for i := int('A') - 1; i <= int('Z')-2; i += 2 {
			<-c
			str <- fmt.Sprintf("%c%c", i+1, i+2)
			n <- struct{}{}
		}
		<-c
	}()

	reStr := ""
	for {
		s, ok := <-str
		reStr += s
		if !ok {
			break
		}
	}
	return reStr
}
