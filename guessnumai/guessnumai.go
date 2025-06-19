package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// 判断字符串是否为4位不重复数字
func isValid(s string) bool {
	if len(s) != 4 {
		return false
	}
	seen := make(map[byte]bool)
	for i := 0; i < 4; i++ {
		if seen[s[i]] {
			return false
		}
		seen[s[i]] = true
	}
	return true
}

// 构造所有不重复的4位数字
func generateAllCandidates() []string {
	var result []string
	for i := 0; i <= 9999; i++ {
		s := fmt.Sprintf("%04d", i)
		if isValid(s) {
			result = append(result, s)
		}
	}
	return result
}

// 计算猜测与目标之间的 xAyB
func getAB(answer, guess string) (nA, nB int) {
	usedGuess := make([]bool, 4)
	usedAnswer := make([]bool, 4)

	// A（位置和数字都对）
	for i := 0; i < 4; i++ {
		if guess[i] == answer[i] {
			nA++
			usedGuess[i] = true
			usedAnswer[i] = true
		}
	}

	// B（数字对但位置错）
	for i := 0; i < 4; i++ {
		if usedGuess[i] {
			continue
		}
		for j := 0; j < 4; j++ {
			if !usedAnswer[j] && guess[i] == answer[j] {
				nB++
				usedAnswer[j] = true
				break
			}
		}
	}
	return
}

func main() {
	candidates := generateAllCandidates()
	reader := bufio.NewReader(os.Stdin)

	tries := 0

	fmt.Println("你想的数字是4位不重复的，请我来猜。每次你只需要告诉我返回多少AxB")

	for {
		if len(candidates) == 0 {
			fmt.Println("没有候选答案了，可能你输入的反馈有误。")
			break
		}

		// 简单策略：随机一个候选
		guess := candidates[rand.Intn(len(candidates))]
		tries++
		switch tries {
		case 1:
			guess = "1234" // 第一次猜测固定为1234
		case 2:
			guess = "5678" // 第二次猜测固定为5678
		}

		fmt.Printf("第 %d 次猜测: %s, 数据库剩余%d\n", tries, guess, len(candidates))
		fmt.Print("请输入结果 (如 0A0B): ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// 特殊指令：输入 4A0B 或 4A 表示猜中了
		if strings.HasPrefix(input, "4A") {
			fmt.Println("太好了，我猜中了！🎉")
			break
		}

		var nA, nB int
		_, err := fmt.Sscanf(input, "%dA%dB", &nA, &nB)
		if err != nil {
			fmt.Println("输入格式不正确，请输入类似 1A2B 的格式")
			continue
		}

		// 筛选剩余候选
		newCandidates := []string{}
		for _, c := range candidates {
			a, b := getAB(c, guess)
			if a == nA && b == nB {
				newCandidates = append(newCandidates, c)
			}
		}
		candidates = newCandidates
	}
}
