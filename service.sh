#!bin/sh
curl -o- -L https://yarnpkg.com/install.sh | bash   # 安装 yarn
 
$HOME/.yarn/bin/yarn install                        # 安装 vue 项目依赖包

env GIN_MODE=release ./apps > apps.log &      # 在后台启动后端（已将go写的后端编译成二进制文件）
 
cd ./resource && $HOME/.yarn/bin/yarn serve --host 0.0.0.0           # 启动前端（注意必须指定 --host 0.0.0.0）
