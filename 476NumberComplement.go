package main

import "fmt"

func findComplement(num int) int {
	//给定一个正整数，对该数的二进制表示形式，从最高位的1开始向后按位取反。
	//如果我们能知道该数最高位的1所在的位置，就可以构造一个长度和该数据所占位置一样长的一个掩码mask，然后概述和mask进行异或即可。
	//例如：5的二进制是101，我们的构造的掩码为mask=111，两者异或则为010，即是所要的结果。
	mask := 1
	temp := num

	for {
		mask = mask << 1
		temp = temp >> 1
		if temp <= 0 {
			break
		}
	}
	return num ^ (mask - 1)

}

func main() {
	fmt.Println(findComplement(5))
	fmt.Println(findComplement(1))
}
