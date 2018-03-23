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


项目设计：
用户 users  (id name pwd msg) 
好友组:friend_groups (id name u_id)
图片组 albums (id name u_id)
好友: friends (u_id g_id f_id f_nickname ) nickname为自定义昵称，为空则去用户那里读取用户的真实信息
图片 pictures (id name img g_id u_id)



设计：
用户登录后，后台维护一个session，返回登出和获取个人信息、获取相册列表、获取好友组列表的url
查看个人信息  返回修改个人信息和注销的url
修改个人信息，返回查看个人信息的url
注销个人信息，返回登录的url

获取相册列表，返回查看、编辑和删除每个相册的url
查看相册信息，返回新增、修改相册信息、删除相册、获取所有图片的url
获取图片，返回编辑、删除、移动图片（获取相册列表）的url
移动图片、提交移动图片的url



0、认证：
    登录 post    /authentication   name pwd
    登出 delete  /authentication/id
1、用户
    用户注册账户        post    /users/register   name pwd
    用户修改个人信息    put     /users/id       name pwd
    用户注销个人账户    delete  /users/id
    用户查看个人信息    get     /users/id

2、好友组
    有个默认好友组无法修改
    新增好友组              post    /users/id/groups        name
    删除好友组（会删除好友） delete   /users/id/groups/gid       
    修改好友组名称          put      /users/id/groups/gid   name 

3、好友
    新增好友                post    /users/id/frinds   name gid
    删除好友                delete   /users/id/frinds/fid   
    修改好友昵称            put      /users/id/frinds/fid   name 
    移动好友所在的好友组     put      /users/id/frinds/fid   gid 

4、图片组
    有个默认图片组无法修改
    新增图片组              post      /users/id/albums   name   visible
    修改图片组信息          put       /users/id/albums   name   visible
    删除图片组（会删除图片） delete    /users/id/albums/aid

5、图片
    新增图片                post   /users/id/pictures   name img aid  visible
    删除图片                delete /users/id/pictures/pid 
    修改图片信息            put    /users/id/pictures/pid    name     visible
    移动图片所在的图片组     put    /users/id/pictures/pid    aid 


用户注册-》成功返回到登录url 

登录成功-》跳转到个人主页
  个人主页包含： 个人信息、用户组、相册、新增好友、新增相册
  
  用户组列表：
    查看、编辑、删除某个好友组
  相册列表：
    查看、编辑、删除某个相册

  查看好友组：
    好友列表
   
  好友列表：
    好友的查看、编辑、删除
  
  好友的查看：
    好友相册的查看
    
  好友相册：
    好友照片
  
  查看相册：
    照片列表

  照片列表：
    照片的查看、删除、编辑
