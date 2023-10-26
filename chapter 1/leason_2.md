# Setup Golang

1. pertama download golang
```
https://go.dev/doc/install
```
2. setup path variable
```
export PATH=$PATH:/usr/local/go/

source ~/.zshrc
```
3. check where golang is installed
```
which go
```
4. setup working space
```
go env
```
```
export GOPATH=$HOME/your-workspace"
```

5. uninstall go 
```
which go

sudo rm -rf /usr/local/go
sudo rm /etc/paths.d/go
```