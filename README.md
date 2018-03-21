# Picture-communicator
go get github.com/Masterminds/glide
go install github.com/Masterminds/glide  //windows版可能有bug需要修改源码
glide create  # 创建glide工程，生成glide.yaml
glide install # 生成glide.lock，并拷贝依赖包
work, work, work
glide get -u github.com/jinzhu/gorm
glide get -u github.com/julienschmidt/httprouter
glide update  # 更新依赖包信息，更新glide.lock