package main

import (
	"bookstores2/controller"
	"net/http"
)

//主方法初始化，处理静态资源
func main() {
	http.Handle("/static/",
		http.StripPrefix("/static/", http.FileServer(http.Dir("views/static/"))))

	http.Handle("/pages/",
		http.StripPrefix("/pages/", http.FileServer(http.Dir("views/pages/"))))

	//处理请求
	//用户相关
	http.HandleFunc("/main", controller.IndexHandler)
	http.HandleFunc("/toLogin", controller.ToLogin)
	http.HandleFunc("/login", controller.Login)
	http.HandleFunc("/logout", controller.Logout)
	http.HandleFunc("/register", controller.Register)
	http.HandleFunc("/FindUserByName", controller.FindUserByName)

	//图书相关
	http.HandleFunc("/getPageBooks", controller.GetPageBooks)
	http.HandleFunc("/toUpdateBookPage", controller.ToUpdateBookPage)
	http.HandleFunc("/deleteBook", controller.DeleteBookById)
	http.HandleFunc("/updateOraddBook", controller.AddOrUpdateBook)
	http.HandleFunc("/queryPrice", controller.QueryPrice)

	//购物车相关
	http.HandleFunc("/AddBook2Cart", controller.AddBook2Cart)
	http.HandleFunc("/getCartInfo", controller.GetCartInfo)
	http.HandleFunc("/deleteCart", controller.DeleteCart)
	http.HandleFunc("/deleteCartItem", controller.DeleteCartItem)
	http.HandleFunc("/updateCartItem", controller.UpdateCartItem)

	//订单相关（结账，发货，收货）
	http.HandleFunc("/checkout", controller.Checkout)
	http.HandleFunc("/getMyOrder", controller.GetMyOrder)
	http.HandleFunc("/getOrders", controller.GetAllOrder)
	http.HandleFunc("/getOrderInfo", controller.GetOrderInfo)
	http.HandleFunc("/sendOrder", controller.SendOrder)
	http.HandleFunc("/takeOrder", controller.TakeOrder)

	////获取SSL 证书和 RSA 私钥
	// 使用https协议时一般需要两个文件，cert.pem和key.pem
	// cert.pem文件是ssl证书，而key.pem是私钥
	// 可以使用Go标准库中的crypto包群来生成证书与私钥
	//utils.GetTLS("utils/pem/cert.pem","utils/pem/key.pem")

	//捕获没有的路由
	http.HandleFunc("/", controller.AnyHandler)

	//设置服务器路径,使用默认多路服务器
	http.ListenAndServe(":8081", nil)
}
