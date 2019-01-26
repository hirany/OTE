# OTE

## 開発環境の構築 golangのインストールとアップデート

```bash
sudo apt install -y git

git config --global user.name "{ username }"
git config --global user.email "{ mail address }"

bash goenv.sh

exec $SHELL
source ~/.bashrc

bash goupdate.sh

source ~/.bashrc
mkdir -p $GOPATH/src/github.com/{ username }
cd $GOPATH/src/github.com/{ username }

bash goget.sh
```

## 実行方法

linux
```bash
go build
./OTE
```
windows
```powershell, cmd
go build
./OTE.exe
``