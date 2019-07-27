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

### 分隔字符串
```
var1=`echo "hua nong jing chao"|awk -F ' ' '{print $1}'`
echo $var1

var2=`echo "hua nong jing chao"|awk -F ' ' '{print $2}'`
echo $var2

var3=`echo "hua nong jing chao"|awk -F ' ' '{print $3}'`
echo $var3

var4=`echo "hua nong jing chao"|awk -F ' ' '{print $4}'`
echo $var4
```

### 连接字符串
```
$value1=home

$value2=${value1}"="

echo $value2

把要添加的字符串变量添加{}，并且需要把$放到外面。

这样输出的结果是：home=，也就是说连接成功。
```

### 去除字符串前后空格
```
[root@localhost ~]# echo ' A B C ' | awk '{gsub(/^\s+|\s+$/, "");print}'
^\s+            匹配行首一个或多个空格
\s+$            匹配行末一个或多个空格
^\s+|\s+$    同时匹配行首或者行末的空格
如果不用awk命令，也可以使用eval命令来达到相同的目的

[root@local ~]# echo "  A BC  "
   A  BC
[root@local ~]# eval echo "  A BC  "
A BC
或者

[root@linux ~]# echo ' A BC  ' | python -c "s=raw_input();print(s.strip())"
A BC
或者

[root@linux ~]# s=`echo " A BC  "`
[root@linux ~]# echo $s
A BC
或者

[root@linux ~]# echo ' A BC ' | sed -e 's/^[ ]*//g' | sed -e 's/[ ]*$//g'
A BC
或者

[root@linux ~]# echo " A BC  " | awk '$1=$1'
A BC
或者

[root@linux ~]# echo " A BC  " | sed -r 's/^[ \t]+(.*)[ \t]+$//g'
A BC
或者

[root@linux ~]# echo ' A BC  ' | awk '{sub(/^ */, "");sub(/ *$/, "")}1'

```