# OTE

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
```
