package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// 生成一个长度为4且不重复的数字字符串
func generateAnswer() string {
	rand.Seed(time.Now().UnixNano())
	digits := rand.Perm(10)
	answer := ""
	for i := 0; i < 4; i++ {
		answer += fmt.Sprintf("%d", digits[i])
	}
	return answer
}

// 计算 xAxB
func checkGuess(answer, guess string) (nA, nB int) {
	used := make(map[byte]bool)

	for i := 0; i < 4; i++ {
		if guess[i] == answer[i] {
			nA++
			used[guess[i]] = true
		}
	}

	for i := 0; i < 4; i++ {
		if guess[i] != answer[i] {
			for j := 0; j < 4; j++ {
				if i != j && guess[i] == answer[j] && !used[guess[i]] {
					nB++
					used[guess[i]] = true
					break
				}
			}
		}
	}
	return
}

func main() {
	answer := generateAnswer()
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("猜数字游戏：请输入4位不重复数字，例如：1234")

	for {
		fmt.Print("你的猜测：")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if len(input) != 4 {
			fmt.Println("必须是4位数字！")
			continue
		}

		nA, nB := checkGuess(answer, input)
		fmt.Printf("结果: %dA%dB\n", nA, nB)

		if nA == 4 {
			fmt.Println("恭喜你猜对了！")
			break
		}
	}
}
