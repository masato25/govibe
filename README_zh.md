# govibe

我已經厭倦了一直當個點擊農夫！！ 我覺得讓機器人對付機器人!!

## 功能簡介
自動偵測螢幕上的 `allow_button.png` 圖案，並自動將滑鼠移動到圖案正中央並點擊。

## 安裝步驟

### 1. 安裝 Homebrew（如已安裝可略過）
https://brew.sh/

### 2. 安裝依賴套件
```sh
brew install opencv pkg-config
```

### 3. 安裝 Go 套件
```sh
go get github.com/go-vgo/robotgo

go get github.com/kbinani/screenshot

go get gocv.io/x/gocv
```

### 4. 權限設定（macOS）
- 前往「系統設定」→「隱私權與安全性」→「輔助使用」
- 勾選 Terminal 或 VS Code，允許控制滑鼠

## 使用方式

1. 將 `allow_button.png` 放在與 `main.go` 同一目錄
2. 執行：
```sh
go run main.go
```
3. 程式會自動偵測螢幕數量，並持續比對螢幕第 0 個顯示器
4. 偵測到圖案時，滑鼠會自動移動到圖案正中央並點擊

## 常見問題

### 1. 執行時顯示 `pkg-config` 或 `opencv4.pc` 找不到
- 請確認已安裝 opencv 和 pkg-config，並設定 PKG_CONFIG_PATH：
```sh
export PKG_CONFIG_PATH="/opt/homebrew/lib/pkgconfig:$PKG_CONFIG_PATH"
```

### 2. 沒有任何訊息或 panic
- 請確認 `allow_button.png` 檔案存在於目錄
- 檢查終端機是否有錯誤訊息

### 3. 滑鼠無法移動或點擊
- 請確認已設定「輔助使用」權限

### 4. 點擊位置不正確
- 程式已修正為點擊圖案正中央
- 多螢幕請調整 `screenshot.GetDisplayBounds(0)` 的參數

## 其他
- 若需指定其他螢幕，請修改 `main.go` 內 `GetDisplayBounds(0)` 的參數
- 若需調整比對相似度，請修改 `if maxVal > 0.8` 內的數值
