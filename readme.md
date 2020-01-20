# Go Git Last 5

A Golang command line tool that gets the last 5 updated repositories for a Github user.

To run:
```
% go run go-git-last-5.go username

OR

% go build go-git-last-5.go
% ./go-git-last-5 username
```

Example:
```
% ./go-git-last-5 golang
******************************************************************
Last 5 updated repositories for golang
******************************************************************
* Repo: go
  - Language: Go
  - URL: https://github.com/golang/go
* Repo: protobuf
  - Language: Go
  - URL: https://github.com/golang/protobuf
* Repo: exp
  - Language: Go
  - URL: https://github.com/golang/exp
* Repo: lint
  - Language: Go
  - URL: https://github.com/golang/lint
* Repo: tools
  - Language: Go
  - URL: https://github.com/golang/tools
```
