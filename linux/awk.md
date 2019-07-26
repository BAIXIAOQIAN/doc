## awk常用操作
> 示例
cat logs/device-online-offline.log | grep CURRENTFLAG4LOGS | awk -F , '{print $11}'   | awk -F : '{print $2}' | awk -F \" '{print $2}' > device_id.log