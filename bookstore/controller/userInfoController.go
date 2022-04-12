package controller

import (
	"bookstores2/dao"
	"bookstores2/model"
	"bookstores2/utils"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

//用户登录处理函数
func Login(w http.ResponseWriter, r *http.Request) {
	//ToLogin(w ,r)
	username := r.PostFormValue("username")
	success, uid, err := dao.Login(username,
		r.PostFormValue("password"))

	// 		//根据用户名查询、验证密码
	// func Login(username string ,password string) (b bool ,uid int,err error)  {
	// 	sqlStr := "select id,password  from users where username = ? "
	// 	pwd := " "
	// 	b = false

	// 	err  = utils.Db.QueryRow(sqlStr,username).Scan(&uid,&pwd)
	// 	if err != nil{
	// 		return b ,0,err
	// 	}else {
	// 		//MD5加密后比较
	// 		code := utils.Md5(password)
	// 		if code == pwd {  //不区分大小比较
	// 			b = true
	// 		}else{
	// 			return b,0,nil
	// 		}
	// 	}
	// 	return b ,uid,nil
	// }

	if err != nil {
		fmt.Println("登录处理出现错误，err：", err)
	}
	if success {
		//登录成功
		//需要创建session
		uuid := utils.CreateUUID()
		session := &model.Session{
			Session_id: uuid,
			User_id:    uid,
			Username:   username,
		}
		err := dao.AddSession(session)
		// // 		//添加一个session，
		// // //session_id使用MD5加密。
		// func AddSession(s *model.Session) error {
		// 	sqlStr := "insert into sessions(session_id , username , user_id) values (?,?,?)"
		// 	//Prepare创建一个准备好的状态用于之后的查询和命令。
		// 	//返回值可以同时执行多个查询和命令。
		// 	stmt, err := utils.Db.Prepare(sqlStr)
		// 	if err != nil {
		// 		fmt.Println("预编译 Prepare 出错，err :", err)
		// 		return err
		// 	}
		// 	//MD5 加密密码
		// 	s.Session_id = utils.Md5(s.Session_id) //肯定是一个

		// 	_, errExec := stmt.Exec(s.Session_id, s.Username, s.User_id)
		// 	if errExec != nil {
		// 		fmt.Println("执行出错,err:", errExec)
		// 		return errExec
		// 	}
		// 	//执行成功！
		// 	return nil
		// }

		if err != nil {
			fmt.Println("添加session 失败，err:", err)
		}
		//创建一个cookie，并将cookie写入到浏览器中
		cookie := &http.Cookie{
			HttpOnly: true,
			Name:     "user",
			Value:    uuid, //未设置过期时间、最长时间，则为会话 cookie，关闭浏览器后失效
		}
		http.SetCookie(w, cookie)

		t := template.Must(template.ParseFiles("views/pages/user/login_success.html"))
		err = t.Execute(w, username)
		if err != nil {
			fmt.Fprintln(w, "解析模板出现异常 ，err:", err)
		}
	} else {
		//登录失败处理
		t := template.Must(template.ParseFiles("views/pages/user/logining.html"))
		err := t.Execute(w, "登录失败，请检查输入的用户名和密码。")
		if err != nil {
			fmt.Fprintln(w, "解析模板出现异常 ，err:", err)
		}
	}
}

//用户注册处理函数
func Register(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	email := r.PostFormValue("email")

	var info = &model.UserInfo{
		Username: username,
		Password: password,
		Email:    email,
	}

	row, errFind := dao.FindUserByName(username)

	// // 	//根据用户名查询一条记录
	// func FindUserByName(username string) (info *model.UserInfo ,err error)  {
	// 	sqlStr := "select ID ,username ,password ,email from users where username = ?  "
	// 	row  := utils.Db.QueryRow(sqlStr,username)

	// 	info = &model.UserInfo{}
	// 	errNotFound :=row.Scan(&info.ID,&info.Username,&info.Password,&info.Email)
	// 	if errNotFound != nil{
	// 		 return nil ,errNotFound
	// 	}
	// 	return info ,nil
	// }

	//验证用户名是否已经存在  (后期改用 AJAX 处理)
	//if row != nil && errFind == nil {
	//	t := template.Must(template.ParseFiles("views/pages/user/regist.html"))
	//	errExe := t.Execute(w, "用户名已存在！请重新输入。")
	//	if errExe != nil {
	//		fmt.Fprintln(w, "解析模板出现异常 ，err:", errExe)
	//	}
	//}

	if row == nil && errFind != nil {
		//进行添加用户操作
		errAdd := dao.AddUser(info)
		// // 		//新增一个用户
		// func AddUser(info *model.UserInfo) error {
		// 	//事务性，预编译
		// 	sqlStr := "insert into users(username,password,email) values (?,?,?);"

		// 	//Prepare创建一个准备好的状态用于之后的查询和命令。
		// 	//返回值可以同时执行多个查询和命令。
		// 	stmt,err := utils.Db.Prepare(sqlStr)
		// 	if err!= nil {
		// 		fmt.Println("预编译 Prepare 出错，err :",err)
		// 		return err
		// 	}
		// 	//MD5 加密密码
		// 	info.Password =  utils.Md5(info.Password)
		// 	_ ,errExec := stmt.Exec(info.Username,info.Password,info.Email)
		// 	if errExec != nil{
		// 		fmt.Println("执行出错,err:",errExec)
		// 		return errExec
		// 	}
		// 	//执行成功！
		// 	return nil
		// }

		if errAdd != nil {
			fmt.Println("注册处理出现错误，err：", errAdd)
		}
		//解析页面模板
		t := template.Must(template.ParseFiles("views/pages/user/logining.html"))
		errExe := t.Execute(w, "")
		if errExe != nil {
			fmt.Fprintln(w, "解析模板出现异常 ，err:", errExe)
		}
	}
}

//通过AJAX 验证用户名是否重复
func FindUserByName(w http.ResponseWriter, r *http.Request) {
	row, err := dao.FindUserByName(r.PostFormValue("username"))

	if err == nil && row != nil {
		w.Write([]byte("用户名已存在！请重新输入。"))
	} else {
		w.Write([]byte("<font style='color:blue'>用户名可用。</font>"))
	}
}

//注销当前用户，销毁Cookie
func Logout(w http.ResponseWriter, r *http.Request) {
	//获取Cookie 查看是否已经登录
	cookie, _ := r.Cookie("user")
	if cookie == nil {
		fmt.Println("获取Cookie失败，Cookie可能不存在，还未登录。")
	} else {
		//根据session_id获取Cookie
		session, errFindSession := dao.DeleteSessionById(cookie.Value)
		// // 		//根据session_id删除一个session，
		// // //MD5加密session_id删除数据库中的相应的session记录。
		// func DeleteSessionById(session_id string) (int64, error) {
		// 	sqlStr := "delete from sessions where session_id =?"
		// 	stmt, err := utils.Db.Prepare(sqlStr)
		// 	if err != nil {
		// 		fmt.Println("预编译 Prepare 出错，err :", err)
		// 		return 0, err
		// 	}
		// 	//MD5 加密密码
		// 	session_id = utils.Md5(session_id)

		// 	res, errExec := stmt.Exec(session_id)
		// 	if errExec != nil {
		// 		fmt.Println("执行出错,err:", errExec)
		// 		return 0, errExec
		// 	}
		// 	  // RowsAffected返回被update、insert或delete命令影响的行数。
		// 	// 不是所有的数据库都支持该功能。
		// 	affect, errRes := res.RowsAffected()
		// 	if errRes != nil {
		// 		fmt.Println("取出受影响行数时出现异常，err:", errRes)
		// 		return 0, errRes
		// 	}
		// 	return affect, nil
		// }

		if session == 0 && errFindSession != nil {
			fmt.Println("数据库中没查找到该session相关记录，err", errFindSession)
		} else {
			//设置为 -1 ，浏览器中的cookie立即销毁
			cookie.MaxAge = -1
			http.SetCookie(w, cookie)
			fmt.Println("刪除相关登录session成功，重定向到主页。")
			http.Redirect(w, r, "/main", 302)
		}
	}
}

//登录前预处理，检查是否登录
func ToLogin(w http.ResponseWriter, r *http.Request) {
	//获取Cookie 查看是否已经登录
	isLogin, _, _ := dao.IsLogin(r)
	if isLogin {
		//已经登录，重定向到主页
		http.Redirect(w, r, "/main", 302)
	} else {
		t := template.Must(template.ParseFiles("views/pages/user/logining.html"))
		t.Execute(w, "")
	}
}

//首页查找所有图书
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	//mapValue := r.URL.Query()
	//pageNo := mapValue.Get("PageNo")
	// ==
	pageNo := r.FormValue("PageNo") //没有此参数pageNo会为空
	if pageNo == "" {
		pageNo = "1"
	}
	indexPage, _ := strconv.ParseInt(pageNo, 10, 64)
	pages, err := dao.GetPageBooks(indexPage)
	if err != nil {
		fmt.Println("分页查询全部图书出现异常,err：", err)
	}

	isLogin, session, errFindSession := dao.IsLogin(r) //判断是否有登录
	// // 	//根据session_id查找一个session，用于验证用户是否登录,返回session（包含用户名等信息）
	// // //MD5加密session_id查找数据库中的相应的session记录。
	// func IsLogin(r *http.Request) (bool, *model.Session, error) {
	// 	//获取Cookie 查看是否已经登录
	// 	cookie, _ := r.Cookie("user")
	// 	if cookie == nil {
	// 		fmt.Println("获取Cookie失败，Cookie可能不存在，还未登录。")
	// 		return false, nil, nil
	// 	} else {
	// 		sqlStr := "select session_id,user_id,username from sessions where session_id =?"

	// 		//MD5 加密密码
	// 		session_id := utils.Md5(cookie.Value)

	// 		res := &model.Session{}
	// 		err := utils.Db.QueryRow(sqlStr, session_id).Scan(&res.Session_id, &res.User_id, &res.Username)
	// 		if err != nil {
	// 			fmt.Println("数据库中没查找到该session相关记录，err:", err)
	// 			return false, nil, err
	// 		} else {
	// 			return true, res, nil
	// 		}
	// 	}
	// }

	if !isLogin || session == nil {
		fmt.Println("数据库中没查找到该session相关记录，err", errFindSession)
		pages.IsLogin = false
	} else {
		pages.IsLogin = true
		pages.Username = session.Username
	}
	t := template.Must(template.ParseFiles("views/index.html"))
	t.Execute(w, pages)
}

//跳转主页面
func AnyHandler(w http.ResponseWriter, r *http.Request) {
	//重定向
	// http.Redirect(w, r, "/user/index", 302)
	http.Redirect(w, r, "/main", 302)
}

// stmt
// 声明;全自动裱纸机;语句对象
