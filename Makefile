# Goのツールチェーンをセットアップします
GOOS=js
GOARCH=wasm

# 出力ファイルの名前
OUT_FILE=./web/wasm/main.wasm

# ソースコードのディレクトリ
SOURCE_DIR=./cmd

# wasm_exec.jsファイルへのパス
WASM_EXEC=$(shell go env GOROOT)/misc/wasm/wasm_exec.js

# ターゲットのビルド
build:
	@echo "Building WASM..."
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(OUT_FILE) $(SOURCE_DIR)
	@echo "Copying wasm_exec.js..."
	cp $(WASM_EXEC) ./web/js/
	@echo "Build complete"

# ビルド成果物のクリーン
clean:
	@echo "Cleaning..."
	rm -f $(OUT_FILE)
	rm -f ./web/js/wasm_exec.js
	@echo "Clean complete"

# デフォルトターゲット
all: build
