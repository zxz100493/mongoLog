#!/bin/sh
# wget https://registry.npmmirror.com/-/binary/node/latest-v18.x/node-v18.6.0-linux-x64.tar.gz

# tar -xvf node-v18.6.0-linux-x64.tar.gz

# mv node-v18.6.0-linux-x64 nodejs

# echo "# nodejs env" >> /etc/profile
# echo "export PATH=`pwd`/nodejs/bin:\$PATH" >> /etc/profile

# bash /etc/profile

env GIN_MODE=release ./apps > apps.log &      # 在后台启动后端（已将go写的后端编译成二进制文件）

yarn --version
yarn config set registry https://registry.npm.taobao.org -g 
yarn config set sass_binary_site http://cdn.npm.taobao.org/dist/node-sass -g

cd ./resource && yarn install && yarn add element-plus && yarn serve --host 0.0.0.0           # 启动前端（注意必须指定 --host 0.0.0.0）
