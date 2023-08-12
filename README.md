# README


make `git clone https://github.com/<author>/<repo>` faster or work in your server that don't have VPN or proxy support.


## install

```
curl -sf https://gobinaries.com/dfang/xgit | sh

curl -sf https://github.com/dfang/xgit/blob/master/scripts/install.sh | sh
```

## run

```
xgit clone https://github.com/golang/go
```
equals to

```
git clone --depth=1 https://ghproxy.com/https://github.com/golang/go
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
