# git-changed-tree

`git diff --name-only` の出力をディレクトリツリー形式で表示するGo製のツールです。

## 機能

- 変更されたファイルを視覚的なディレクトリツリー形式で表示
- 未コミットの変更（git status）の含め/除外を切り替え可能
- 比較対象のベース（branch/commit）を指定可能
- マルチプラットフォーム対応 (Windows, macOS, Linux)

## 使い方

### 基本実行

```bash
# origin/main...HEAD との差分をツリー表示（未コミットの変更も含む）
./git-changed-tree

# 特定のブランチとの比較
./git-changed-tree -base master

# 未コミットの変更を除外
./git-changed-tree -status=false
```

### オプション

- `-base`: 比較対象のリファレンス (デフォルト: `origin/main...HEAD`)
- `-status`: 未コミットの変更を含めるかどうか (デフォルト: `true`)

## インストール

[Releases](https://github.com/m4549071758/git-changed-tree/releases) ページから、お使いの環境に合ったバイナリをダウンロードしてパスの通った場所に配置してください。

## ビルド

```bash
go build -o git-changed-tree
```
