## shell脚本常用操作

### 逐行读取文件
```
while read line
do
  echo $line
done < FileName
```

### 获取命令的输出
```
check_results=`rpm -qa | grep "zlib"`
echo "command(rpm -qa) results are: $check_results"
if [[ $check_results =~ "zlib" ]] 
then 
    echo "package zlib has already installed. "
else 
    echo "This is going to install package zlib"
 fi
```

### 变量自增
```
1. i=`expr $i + 1`;
2. let i+=1;
3. ((i++));
4. i=$[$i+1];
5. i=$(( $i + 1 ))
```