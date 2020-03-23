# IssueTerminal
This terminal helps you manage issues of repositories locally.

## 如何使用?
**1.下载或编译得到IssueTerminal**  
   - 从 Release 中获取最新版的 IssueTerminal  
   - 从 /bin 目录下获取最新版的 IssueTerminal  
   - clone 这个仓库并执行 "go build /src/github/terminal/issueTerminal.go" 进行编译   
  
**2. 生成你的个人 github access token**
   - Settings -> Developer settings -> Personal access tokens -> Generate new token  
   
**3. 将 Token 保存到 "access_token.txt"**
   - 请删除文件中的默认行，并将你的 token 保存到其中，请勿输入其它字符或换行  
   
**4. 请确保 access_token.txt 和 issueTerminal.exe 在同一文件夹下**  

**5. 开始使用**  

## 命令说明
- -h 获取以下所有指令的说明
- -get [Username / Organization] [Repository]  
   获取指定 Repo 下的所有 issue
- -put  
   创建一个 issue 并上传到指定 Repo  
- -i  
   进入交互模式输入参数.  
- -closed [Username / Organization] [Repository] [issue id]
   关闭一个 issue  
- -update  
   修改一个 issue
- -github  
   获取 github api 参数说明



## How to use?
**1. Download or compile this terminal**
   - Download this terminal from release
   - Or clone this repository and get terminal from "/bin"
   - Or clone this repository and "go build /src/github/terminal/issueTerminal.go"  
  
**2. Generate your github access token**
   - Settings -> Developer settings -> Personal access tokens -> Generate new token  
   
**3. Save the token to file "access_token.txt"**
   - You need to delete the initial line, and replace it with your access token.  
   
**4. Make sure the "access_token.txt" and "issueTerminal.exe" is in the same directory**  

**5. Run the issue terminal**  

## Command
- -get [Username / Organization] [Repository]  
   Get issues of a repository.
- -put  
   Create an new issue and put it into a repository.  
- -i  
   Enter the interactive mode.  
- -closed [Username / Organization] [Repository] [issue id]
   Make an issue closed.  
- -update  
   Update an issue.
- -github  
   Show the explaination of github api parameters.
