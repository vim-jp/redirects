# Redirect pages to Vim binaries

## How to update redirects

### Pre requirements

*   of course checkout this [vim-jp/redirects][1] repo
*   [go 1.5.3 or above (1.6 is recommended)][2]
*   golang external packages

    ```
    $ go get -u github.com/koron/go-github
    $ go get -u gopkg.in/yaml.v2
    ```

### Update redirects

```
$ cd vim-jp/redirects
$ go run _scripts/check_release/main.go
$ git add .
$ git commit -m "awesome comments"
$ git push
```

[1]:https://github.com/vim-jp/redirects
[2]:https://golang.org/dl/
