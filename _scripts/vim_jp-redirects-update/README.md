# Tool to update redirect entries

    $ go run _scripts/vim_jp-redirects-update/main.go

or install compiled binary with `go build .` and use it.

## Environment variables

*   `GITHUB_USERNAME` and `GITHUB_TOKEN`

    Set these variables to increase rate limit of github's API.
    Token will be get from **Settings>Personal access token** on github.
