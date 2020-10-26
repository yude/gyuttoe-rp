package main

import (
	"fmt"
	"os"
	"time"

	"github.com/hugolgst/rich-go/client"
	"github.com/joho/godotenv"
)

func main() {
	// .env ファイルを読み込む (Discord の App ID 用)
	err := godotenv.Load(fmt.Sprintf("../%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		panic(err)
	}

	// 教材の種類, 初期化
	lesson := "None"
	// HIKAKIN & SEIKIN - 今
	now := time.Now()
	// Discord の App ID を .env から読み込む
	DISCORD_APP_ID := os.Getenv("DISCORD_APP_ID")

	// Discord にログインする
	err = client.Login("DISCORD_APP_ID")
	if err != nil {
		panic(err)
	}

	// Rich Presence を表示する
	err = client.SetActivity(client.Activity{
		State:      "Playing",
		Details:    lesson,
		LargeImage: "e",
		LargeText:  "ぎゅっとe - eラーニングシステム",
		Timestamps: &client.Timestamps{
			Start: &now,
		},
	})

	if err != nil {
		panic(err)
	}
}
