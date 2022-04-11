package dao

import (
	"bookstores2/model"
	"bookstores2/utils"
	"fmt"
)

//添加购物车
func AddCart(cart *model.Cart) error {
	sqlStr := "insert into carts(id ,total_amount,total_count,user_id) values(?,?,?,?)"

	stmt, err := utils.Db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("预编译出现异常，err:", err)
		return err
	}
	//图书总数，总价 通过结构体的方法获取
	_, errExec := stmt.Exec(cart.CartId, cart.GetTotalAmount(), cart.GetTotalCount(), cart.UserId)
	if errExec != nil {
		fmt.Println("执行出错,err:", errExec)
		return errExec
	}

	//添加购物项到相应的购物车中(重要)
	for _, v := range cart.CartItems {
		AddCartItem(v)
	}

	//执行成功！
	return nil
}

//根据用戶 ID查询得到购物车
func FindCartByUserId(userId int) (*model.Cart, error) {
	sqlStr := "select id, total_amount,total_count,user_id  from carts where user_id = ?  "

	cart := &model.Cart{}
	errNotFound := utils.Db.QueryRow(sqlStr, userId).Scan(&cart.CartId, &cart.TotalAmount, &cart.TotalCount, &cart.UserId)
	if errNotFound != nil {
		return nil, errNotFound
	}

	cartItems, _ := FindCartItemsByCartId(cart.CartId)

	// // //根据购物车的id获取购物车中所有的购物项
	// func FindCartItemsByCartId(cartId string) ([]*model.CartItem,error) {
	// 	sqlStr := "select id ,COUNT,amount,book_id,cart_id from cart_itmes where cart_id = ?"
	// 	var cartItems []*model.CartItem

	// 	rows ,err := utils.Db.Query(sqlStr,cartId)
	// 	if err != nil{
	// 		fmt.Println("查询所有的购物项出现异常，err",err)
	// 		return nil ,err
	// 	}
	// 	for rows.Next(){
	// 		cartItem := &model.CartItem{}
	// 		var bookId int
	// 		//此处直接扫描写入 &cartItem.Book.ID 会panic 因为此时book还未声明，没有空间地址
	// 		err := rows.Scan(&cartItem.CartItemId,&cartItem.Count, &cartItem.Amount,&bookId,&cartItem.CartId)
	// 		if err != nil {
	// 			fmt.Println("扫描写入出现异常，err:",err)
	// 			return nil ,err
	// 		}
	// 		book ,err := FindBookById(bookId)

	// 		// 			//根据图书ID查询一条记录
	// 		// func FindBookById(id int) (book *model.Book, err error) {
	// 		// 	sqlStr := "select id , title,author ,price ,sales ,stock ,img_path  from books where id = ?  "
	// 		// 	row := utils.Db.QueryRow(sqlStr, id)

	// 		// 	book = &model.Book{}
	// 		// 	errNotFound := row.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImagePath)
	// 		// 	if errNotFound != nil {
	// 		// 		return nil, errNotFound
	// 		// 	}
	// 		// 	return book, nil
	// 		// }
	// 		if err != nil {
	// 			fmt.Println("查询图书信息出现异常，err:",err)
	// 		}
	// 		cartItem.Book = book
	// 		cartItems = append(cartItems,cartItem)
	// 	}
	// 	return cartItems,nil
	// }

	for _, cartItem := range cart.CartItems {
		cartItems = append(cartItems, cartItem)
	}
	cart.CartItems = cartItems
	return cart, nil
}

//更新购物车 总金额，图书总数量
func UpdateCart(cart *model.Cart) error {
	//set 后跟多个值用 ,  而不是 and 隔开！！！！！
	sqlStr := "update carts set  total_amount = ? , total_count = ? where id = ?"

	stmt, err := utils.Db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("预编译出现异常，err:", err)
		return err
	}
	//图书总数，总价 通过结构体的方法获取
	_, errExec := stmt.Exec(cart.GetTotalAmount(), cart.GetTotalCount(), cart.CartId)
	if errExec != nil {
		fmt.Println("执行出错,err:", errExec)
		return errExec
	}
	//执行成功！
	return nil
}

//删除购物车（清空购物车）
//清空之前需要删除全部购物车内的购物项
func DeleteCartById(cartId string) error {
	sqlStr := "delete from carts where id = ?"

	stmt, err := utils.Db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("预编译 Prepare 出错，err :", err)
		return err
	}
	_, errExec := stmt.Exec(cartId)
	if errExec != nil {
		fmt.Println("执行出错,err:", errExec)
		return errExec
	}
	return nil
}
