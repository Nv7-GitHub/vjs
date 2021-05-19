async function main() {
  var typescriptParser = require("typescript-parser");
  const parser = new typescriptParser.TypescriptParser();

  const parsed = await parser.parseFile('lib.dom.d.ts', '.');
  const json = JSON.stringify(parsed);

  var fs = require('fs');
  await fs.writeFile("ast.json", json, function(err, result) {
    if(err) console.log('error', err);
  });
}

main();