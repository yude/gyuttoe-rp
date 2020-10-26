# [yudete](https://github.com/yudete) / gyuttoe-rp
eラーニングシステム「ぎゅっとe」Discord Rich Presence (半自動)  
## For English:
Discord Rich Presence front of e-learning system called "Gyutto-e". (Semi-automatic)

## Goal
* 解答などといったクリティカルなデータを公開しない: 不正行為の手助けとならないツール
* API が無いので手動（アホ）

## To Do
* 閉じるための処理が、バグを利用しているような状況なので何とかする。  
(ウィンドウを閉じた瞬間に RPC の表示が止まるわけではなく、5 秒程度のラグがある。)  

## Requirements
* Go
* [hugolgst/rich-go](https://github.com/hugolgst/rich-go)
```
go get github.com/hugolgst/rich-go
```
* [Fyne - Cross platform GUI in Go based on Material Design](https://github.com/fyne-io/fyne)
```
go get fyne.io/fyne
go get fyne.io/fyne/app
go get fyne.io/fyne/widget
```

### For packaging
* For decrease file size of executable:  
[UPX - the Ultimate Packer for eXecutables](https://upx.github.io/)

## License
This project is licensed under the [MIT License](http://opensource.org/licenses/MIT).
