# Redirect pages to Vim binaries

## How to update redirects

### Pre requirements

*   of course checkout this [vim-jp/redirects][1] repo
*   [go 1.16 or above][2]

### Update redirects

```
$ cd vim-jp/redirects
$ go run _scripts/vim_jp-redirects-update/main.go
$ git add .
$ git commit -m "awesome comments"
$ git push
```

[1]:https://github.com/vim-jp/redirects
[2]:https://golang.org/dl/
