# CLI tool for creating go project and package skeleton

## Creates new Golang

### Project:

- Name restriction: same case, at least 2 symbols, no special symbols.

#### Creates:

1. new directory ( 'newproj' in both cases )
2. go mod init (if long name provided, with long otherwise short)
3. creates first test file from template.
4. :
```bash
 inits git repo
 git add .
 git commit -m 'first init'
```

---

### Package:
- Same name restrictions (only short name).
#### Creates:

1. new directory
2. test file

- You should be able to run the test that should fail :)