curl -O https://raw.githubusercontent.com/microsoft/TypeScript/master/lib/lib.dom.d.ts
go run .
rm lib.dom.d.ts
mkdir dist
cp vjs.js.v dist/vjs.js.v
rm vjs.js.v