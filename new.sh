##!/usr/bin/env bash
# 创建 Go 每日一题

dir="$(date +%Y-%m-%d)"
date="$(date +%Y%m%d)"

sort=$1
exercism=$2

content="// $date
// $sort $exercism

package day

func F${date}() {

}"

content_test="package day

import \"testing\"

func Test_F${date}(t *testing.T) {

}"

dir_path="."
(
    IFS="-"
    for e in $dir; do
        dir_path="${dir_path}/${e}"
    done
    mkdir -p ${dir_path} && \
    echo -e "${content}" >${dir_path}/${date}.go
    echo -e "${content_test}" >${dir_path}/${date}_test.go
    echo -e "| ${sort}| [${exercism}](${dir_path}/${date}.go) |" >> README.md
    code ${dir_path}/${date}.go
)

# 排序
reserve=$(head -n 7 README.md)
lists=$(grep '\[*\]' README.md  | sort)

echo  "${reserve}" > README.md
echo "${lists}" >> README.md
