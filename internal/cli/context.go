package cli

import (
	"fmt"

	"see_parallel/internal/config"
)

func SetContext(context string) {
	if context == "" {
		fmt.Println("エラー: 文脈情報が空です")
		return
	}

	if err := config.SetContext(context); err != nil {
		fmt.Printf("エラー: 文脈情報を保存できません: %v\n", err)
		return
	}

	fmt.Println("文脈情報を設定しました")
	fmt.Printf("設定内容: %s\n", context)
}

func GetContext() {
	context := config.GetContext()
	if context == "" {
		fmt.Println("文脈情報は設定されていません")
		return
	}

	fmt.Printf("現在の文脈情報: %s\n", context)
}

func ClearContext() {
	if err := config.ClearContext(); err != nil {
		fmt.Printf("エラー: 文脈情報をクリアできません: %v\n", err)
		return
	}

	fmt.Println("文脈情報をクリアしました")
}