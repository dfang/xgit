# README

让`git clone`更方便。 

在梯子临时不好使了，或者还没来得及装梯子的环境中，想临时克隆下github仓库可以试试这个  
或者经常需要克隆仓库，这个也很方便  

运行 `xgit clone <user>/<repo>` 就可以了, 对比正常的`git clone https://github.com/<user>/<repo>`, `git`换成`xgit`， `https://github.com`都可以不用输了  

如果需要指定克隆到哪里. `xgit clone <user>/<repo> /tmp/<repo>`

## install

三种方式任选一种

```
curl -sf http://goblin.run/github.com/dfang/xgit | sh

curl -sf https://raw.githubusercontent.com/dfang/xgit/master/scripts/install.sh | sh

go install github.com/dfang/xgit@latest
```

## run

三种格式都支持

```
xgit clone golang/go
xgit clone github.com/golang/go
xgit clone https://github.com/golang/go

```

by default, `--depth=1` is added to `xgit clone`, you have to run `git fetch --unshallow` to get all old commits if you want after clone is done.

you can `--no-depth=` to skip `--depth=1`.

eg. `xgit clone https://github.com/golang/go` or `xgit clone golang/go`


if you want to change url for remote after:

```
git remote -v
git remote set-url origin <ORIGIN GITHUB LINK>
git remote -v
```
