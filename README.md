Open a folder locally with blank<br>
commands to execute to initialize go project which creates a `go.mod` file
```shell
go mod init github.com/GannaSaiRam/data_validation
```
Create repository in webpage and initialize here git and add remote to webpage created one
```shell
git init
git remote add origin https://github.com/GannaSaiRam/amazon_datavalidation.git
git branch -M master
git push -u origin master
```
after adding git init better to modify the url in .git/config file and excute while pushing code in.

execute `go_path.sh` file, which is going to create files according to project stcture and btw add `.gitignore` and have line 
```text
bin
```
So we are ready and when you are going to install anything `go get` have GOPATH and GOBIN set