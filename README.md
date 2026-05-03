# タスク管理アプリ（学習用）

SWE 実務選考インターンに向けた事前学習用の **タスク管理** を題材にした Go 実装です。単に動作させることより、**DDD に寄せたレイヤ分離**と「どこにどのロジックを置くか」を説明できることを優先しています。

## リポジトリ構成

| パス                     | 説明                                                                                                  |
| ------------------------ | ----------------------------------------------------------------------------------------------------- |
| [`backend/`](backend/)   | Go モジュール（`go.mod`）、`cmd/`、`internal/`                                                        |
| [`docs/`](docs/)         | 要件・フェーズ計画・設計メモ（正は [`docs/phase-00_requirements.md`](docs/phase-00_requirements.md)） |
| [`AGENTS.md`](AGENTS.md) | 開発方針・レイヤ責務・禁止事項（AI／人間とも参照用）                                                  |

## 必要環境

- [Go](https://go.dev/dl/) **1.22** 以上（[`backend/go.mod`](backend/go.mod) に準拠）

## 実行方法（Phase 01 デモ）

Phase 01 は **HTTP API を含みません**。ユースケースの動作確認は、組み込みのデモプログラムで行います。

```bash
cd backend
go run ./cmd/phase01-demo
```

作成・一覧・取得・タイトル変更・削除・バリデーション／NotFound の例が順に標準出力に表示されます。

## アーキテクチャ（概要）

レイヤの流れは次のイメージです（Phase 02 から HTTP の **handler** を主に追加する想定）。

```txt
handler（Phase 02〜）
 ↓
usecase
 ↓
domain（Entity / Value Object / ドメインエラー）
 ↓
repository（interface は usecase 側、実装は `internal/repository/...`）
```

現在の実装では **in-memory のタスクリポジトリ** を利用しています（DB は Phase 04 以降で検討）。

## 実装スコープ（現状）

要件・フェーズの詳細は [`docs/phase-00_requirements.md`](docs/phase-00_requirements.md) を参照してください。

| フェーズ                      | 内容                                                                                                                                                               |
| ----------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| **Phase 01**（進行中の前提）  | UC-01〜05（作成・一覧・1件取得・タイトル変更・削除）、ドメインルール、`TaskService`、メモリ永続化。**自動テストは必須完了条件に含めない**（Phase 03 で主に扱う）。 |
| **Phase 02**（予定）          | HTTP、JSON、ステータスコード、並行アクセス時のリポジトリ排他など。                                                                                                 |
| **Phase 03**（予定）          | ユースケース単位の自動テストの充実。                                                                                                                               |
| **Phase 04 以降**（計画のみ） | DB、認証など。                                                                                                                                                     |

## 依存関係

[`backend/go.mod`](backend/go.mod) に記載。現状はタスク ID 生成などに [`github.com/google/uuid`](https://pkg.go.dev/github.com/google/uuid) を利用しています。
