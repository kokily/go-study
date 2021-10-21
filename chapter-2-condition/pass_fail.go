// pass_fail은 성적 합격 여부를 알려줍니다
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Enter a grade: ")

	// 키보드로부터 텍스트를 읽어 오기 위한 버퍼 리더를 만듦
	reader := bufio.NewReader(os.Stdin)

	// 엔터키가 눌리기 직전까지 입력된 모든 문자를 리턴
	input, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(input)
}
