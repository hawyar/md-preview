## Online Markdown Parser

Clone repo

```bash
git clone https://github.com/hawyar/md-preview
```

Generate JavaScript code

```bash
go generate
```

Run

```bash
npx serve public
```

Open [http://localhost:3000](http://localhost:3000)

### How?

Uses [markdown](https://github.com/gomarkdown/markdown) to parse the mardkown and [gopherjs](https://github.com/gopherjs/gopherjs) to then expose the function to the document's global scope.