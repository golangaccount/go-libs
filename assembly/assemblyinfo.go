package assembly

import (
	"runtime"
)

var (
	AssemblyTitle         string //标题
	AssemblyDescription   string //描述
	AssemblyConfiguration string //配置
	AssemblyCompany       string //公司
	AssemblyProduct       string //产品名称
	AssemblyCopyright     string //产权
	AssemblyTrademark     string //商标
	AssemblyCulture       string //
	AssemblyVersion       string //版本
	AssemblyFileVersion   string //文件版本
	AssemblyBuildTime     string //编译时间
	AssemblyGitRepository string //git的repository    cmd:
	AssemblyGitBranch     string //git的分支     cmd:git symbolic-ref --short -q HEAD
	AssemblyGitHead       string //git的提交head cmd:git rev-parse HEAD
	GoVersion             string //go 版本
)

func init() {
	GoVersion = runtime.Version()
}

//build 命令
//window下
//go build -ldflags "-X git.lichengsoft.com/lichengsoft/go-libs/assembly.GitBranch=$(git symbolic-ref --short -q HEAD) -X git.lichengsoft.com/lichengsoft/go-libs/assembly.GitHead=$(git rev-parse HEAD) -X git.lichengsoft.com/lichengsoft/go-libs/assembly.BuildTime=$(date -format yyyyMMddHHmmss)"
//linux下
//go build -ldflags "-X git.lichengsoft.com/lichengsoft/go-libs/assembly.GitBranch=$(git symbolic-ref --short -q HEAD) -X git.lichengsoft.com/lichengsoft/go-libs/assembly.GitHead=$(git rev-parse HEAD) -X git.lichengsoft.com/lichengsoft/go-libs/assembly.BuildTime=$(date +%Y%m%d%H%M%S)"
