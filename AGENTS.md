# AGENTS.md

## 目的

このリポジトリは、SWE実務選考インターンに向けた事前学習用のタスク管理APIです。

主目的は、GoによるAPI実装とDDDの基本設計を、実務に近い形で練習することです。

単に動くコードを書くことではなく、以下を重視します。

- GoでAPIを設計・実装できること
- DDDの責務分離を理解すること
- Entity / Value Object / Aggregate / Repository / Usecase の役割を説明できること
- 「このロジックをどこに置くか」を判断できること
- 実装後に設計意図・トレードオフ・改善余地を言語化できること

## 技術スタック

- Monorepo: Go の実装（`go.mod`、`cmd/`、`internal/`）は **`backend/`** に置く
- Language: Go
- API（Phase 02 以降）: net/http または軽量な HTTP ルーター。Phase 01 では HTTP 層は不要
- 永続化: Phase 01〜03 は in-memory repository（DB は Phase 04 以降で検討）
- Architecture: DDD / Clean Architecture寄り

## 設計方針

以下のレイヤー分離を意識する。

```txt
handler
  ↓
usecase
  ↓
domain
  ↓
repository
```

### domain

ドメインルールを表現する層。

- Entity
- Value Object
- Aggregate
- Domain Error

を置く。

ただのデータ構造にせず、業務ルールや不変条件を持たせる。

### usecase

アプリケーションの操作単位を表現する層。

例（詳細は `docs/phase-00_requirements.md`）：

- タスクを作成する
- タスク一覧を取得する
- タスクを 1 件取得する
- タスクのタイトルを変更する
- タスクを削除する

外部I/OやRepository呼び出しの流れを制御する。

### repository

永続化の抽象を表現する層。

- interfaceは使う側、つまりusecase側に近い場所に置く
- Phase 01〜03 は in-memory 実装でよい。Phase 02 で HTTP により並行アクセスがあり得る場合は、実装側に排他制御を入れる（要件書セクション 3・6）
- DB詳細をdomainに漏らさない

### handler

HTTPリクエスト/レスポンスを扱う層（**Phase 02 から主に扱う**。Phase 01 に必須ではない）。

- JSONのdecode/encode
- status codeの返却
- usecaseの呼び出し

を担当する。

ドメインロジックはhandlerに書かない。

## 実装方針

- [`docs/phase-00_requirements.md`](docs/phase-00_requirements.md) のフェーズ計画に沿って進める
- いきなり大きく作らない
- 1ユースケースずつ実装する
- Phase 01 の着手は UC-01「タスク作成」から（要件書と整合）
- 完璧なDDDより、説明可能な設計を優先する
- 過剰設計を避ける
- ただし責務分離は崩さない

## コード生成時のルール

AIエージェントは、いきなり完成コードを大量に出さないこと。

必ず以下の順で進める。

1.  変更方針を説明する
2.  どの層に何を追加するか説明する
3.  実装する
4.  設計意図を説明する
5.  改善余地を示す

## レビュー観点

コードレビューでは、以下を重点的に見る。

- このロジックは適切な層にあるか
- domainがただのstructになっていないか
- usecaseが責務を持ちすぎていないか
- handlerに業務ロジックが漏れていないか
- repositoryの抽象が不自然でないか
- Goらしいエラーハンドリングになっているか（Validation / NotFound など要件書セクション 5.7 と整合しているか）
- テストしやすい構造になっているか（Phase 03 で自動テストを主に扱うが、要件書どおり Clock 注入なども検討する）

## 禁止事項

- handlerにドメインロジックを書く
- domainが外部I/Oに依存する
- repository実装の都合をdomainに漏らす
- 最初から複雑なフレームワークやDBを入れる
- 設計理由を説明せずにコードだけ変更する
