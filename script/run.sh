#!/bin/bash

# 检查Go是否已安装
if command -v go >/dev/null 2>&1; then
    echo "Go is installed, continue."
    # 继续执行您的命令或脚本
    go mod tidy
else
    echo "Cannot find go in PATH, please install it first."
    exit 1
fi