# foip_copy

## Description
FOIP開発用リポジトリ（就活用コピー）

### Directory
- backend/
  - chat/
  - core/
- frontend/

### How to run
初めて起動する人向け
```bash
cd frontend
npm i
npm run build
cd ../
docker compose up
```

`docker compose`が立ち上がったら[localhost](http://localhost)をブラウザで開いてみよう

---
#### for developer
サービス全体の起動
```
$ docker compose up
```
フロントの開発の場合
```
$ cd frontend
$ npm run build/dev/watch
```

## Other
- [Commit Rules](./docs/commit_rules.md)
- [Branch Rules](./docs/branch_rules.md)
- [Issue Rules](./docs/issue_rules.md)
