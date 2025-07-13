package cli

import (
	"fmt"

	"see_parallel/internal/config"
)

func SetAPIKey(apiKey string) {
	if apiKey == "" {
		fmt.Println("エラー: APIキーが空です")
		return
	}

	if err := config.SetAPIKey(apiKey); err != nil {
		fmt.Printf("エラー: APIキーを保存できません: %v\n", err)
		return
	}

	fmt.Println("APIキーを設定しました")
}