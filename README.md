# Picture-communicator
go get github.com/Masterminds/glide
go install github.com/Masterminds/glide  //windows版可能有bug需要修改源码
glide create  # 创建glide工程，生成glide.yaml
glide install # 生成glide.lock，并拷贝依赖包
work, work, work
glide get -u github.com/go-sql-driver/mysql
glide get -u github.com/jinzhu/gorm
glide get -u github.com/julienschmidt/httprouter
glide update  # 更新依赖包信息，更新glide.lock

#router  https://studygolang.com/articles/11411
httprouter
控制器是苗条的，模型是丰满的

控制器只包含路由和http报文的打包和解包

模型则包含逻辑处理和数据分析

path 和 path/filepath
bytes  strings strconv unicode

bytes:提供查询、比较、替换、截断、拆分、合并等，但针对和string有相同结构的[]byte类型，bytes.Buffer更有效
string:提供字符串查询、比较、替换、截断、拆分、合并等
strconv:提供string和其他类型的转换，一级双引号转义的转换
unicode:rune  提供字符分类，比如isDigit,isLetter,isUper等蕾西分类功能的函数
