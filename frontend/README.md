# FOIP frontend
### 起動方法
```bash
npm run start
```

### ビルド方法
```bash
npm run build
```

### 整形方法
```bash
npm run format
```
### create-react-appを使っていた人へ
起動の際に`vite`を利用するように変更しているため、*移行後の初回起動時には以下のコマンドを実行後に起動すること*。
```bash
rm -rf node_modules
npm i
```