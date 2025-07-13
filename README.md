# see_parallel

GoベースのCLIツールで、複数のファイルをGemini AIで並列に分析します。プロジェクトごとに文脈情報を設定でき、大量のファイル分析を効率的に実行できます。

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

## 実践的な使用例

### 1. プロジェクト全体の把握
```bash
# プロジェクトの文脈を設定
see_parallel context set "Next.js 14とTypeScriptを使用したWebアプリケーション"

# 複数の質問をキューに追加
see_parallel queue '["プロジェクトの全体構造を説明して", "**/*.ts", "**/*.tsx", "package.json"]'
see_parallel queue '["使用されている主要なライブラリとその用途は？", "package.json", "tsconfig.json"]'
see_parallel queue '["ルーティング構造を説明して", "app/**/*.tsx", "app/**/route.ts"]'

# 並列で分析実行
see_parallel queue run --parallel 3
```

### 2. コードレビュー
```bash
# セキュリティレビュー
see_parallel queue '["セキュリティ上の懸念点を指摘して", "app/api/**/*.ts", "lib/auth/**/*.ts", "deep"]'

# パフォーマンスレビュー
see_parallel queue '["パフォーマンスの改善点を提案して", "components/**/*.tsx", "hooks/**/*.ts", "deep"]'

see_parallel queue run --parallel 2
```

### 3. バッチ分析
```bash
# 複数のコンポーネントを一括分析
for component in Header Footer Navigation Card Button; do
  see_parallel queue "[\"$componentコンポーネントの実装を説明して\", \"components/$component.tsx\"]"
done

see_parallel queue run --parallel 5
```

### 4. ドキュメント生成
```bash
# API仕様の抽出
see_parallel context set "APIドキュメントを生成するため、エンドポイント、パラメータ、レスポンスを詳しく説明"
see_parallel queue '["APIエンドポイントの仕様をまとめて", "app/api/**/route.ts", "deep"]'
see_parallel queue run
```

## プロジェクト別管理

異なるプロジェクトで独立した設定を維持できます：

```bash
# プロジェクトA
cd ~/projects/react-app
see_parallel context set "React 18のSPAプロジェクト"
see_parallel queue '["コンポーネントの構造は？", "src/**/*.jsx"]'

# プロジェクトB
cd ~/projects/backend-api
see_parallel context set "Express.jsのREST APIサーバー"
see_parallel queue '["APIの認証方式は？", "routes/**/*.js", "middleware/**/*.js"]'
```

## ヒント

1. **文脈情報を活用**: プロジェクト固有の情報を設定することで、より正確な分析結果を得られます
2. **並列数の調整**: APIレート制限に注意しながら、並列数を調整して効率化
3. **深い分析の使い分け**: 重要な分析には`deep`オプションを使用
4. **ワイルドカード活用**: `**/*.ts`のようなパターンで効率的にファイルを指定

## モデル
- 通常: Gemini 2.0 Flash Lite (高速・軽量)
- "deep"指定時: Gemini 2.0 Flash (詳細分析)