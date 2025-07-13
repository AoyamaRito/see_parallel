# see_parallel

GoベースのCLIツールで、複数のファイルをGemini AIで並列に分析します。

## セットアップ

1. ビルド:
```bash
go build -o see_parallel
```

2. API キーの設定（以下のいずれか）:

方法1: コマンドで設定
```bash
./see_parallel api set "your-api-key"
```

方法2: 環境変数で設定
```bash
export GEMINI_API_KEY="your-api-key"
```

## 使い方

### キューにタスクを追加
```bash
# 単一ファイル分析
./see_parallel queue '["このファイルの主要な関数は？", "lib/auth.ts"]'

# 複数ファイル横断分析
./see_parallel queue '["認証の仕組みを説明して", "lib/auth.ts", "lib/jwt.ts", "middleware.ts"]'

# 深い分析（上位モデル使用）
./see_parallel queue '["セキュリティリスクを詳細に分析", "lib/auth.ts", "api/routes.ts", "deep"]'

# ワイルドカード使用
./see_parallel queue '["プロジェクト全体の構造は？", "**/*.ts", "**/*.tsx"]'
```

### キューの確認
```bash
./see_parallel queue list
```

### キューの実行
```bash
# デフォルト（並列数10）
./see_parallel queue run

# 並列数を指定
./see_parallel queue run --parallel 5
```

### キューのクリア
```bash
./see_parallel queue clear
```

### 文脈情報の設定
```bash
# プロジェクトの文脈を設定
./see_parallel context set "REACTのプロジェクトです"

# 文脈の確認
./see_parallel context get

# 文脈のクリア
./see_parallel context clear
```

## ファイル保存場所
- APIキー: `~/.see_parallel/config` (グローバル)
- キュー: `./see_parallel/queue.json` (プロジェクト毎)
- 文脈情報: `./see_parallel/context` (プロジェクト毎)

## モデル
- 通常: Gemini 2.0 Flash Lite
- "deep"指定時: Gemini 2.0 Flash