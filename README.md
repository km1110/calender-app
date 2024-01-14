# task-copilot

## 概要

タスク管理の Web アプリケーション

### 想定されるユーザー

- カレンダー、TODO 意外に日記の記録もしたい方
- 自身の達成度を可視化したい方

## Requirement

- node 18.14.2
- React 18.2.0
- go 1.20.5

## function

- 1 年間の相対的な達成度の確認
- タスクの管理
- カレンダーで予定管理
- 毎日の日記の記録

## point

### カレンダーへの情報の反映

- 表示されたカレンダー内の日付に該当するデータのみを取得するようクエリーパラメータを設定した。

## CORSの作成
- CORSを作成し、特定のオリジンからしかAPIの結果を取得できないように設定した。

## Doc

- [DB](/docs/db.md)

## test account
- id : test@test.com
- pw : test123
