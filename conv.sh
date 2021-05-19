cd astgen
npm install
curl -O https://raw.githubusercontent.com/microsoft/TypeScript/master/lib/lib.dom.d.ts
node index.js
cp ast.json ../ast.json
rm ast.json
cd ../
go run .
cp vjs.js.v dist/vjs.js.v
rm vjs.js.v