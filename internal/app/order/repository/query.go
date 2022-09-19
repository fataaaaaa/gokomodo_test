package repository

import (
	"fmt"
	base_model "gokomodo_test/internal/app/model"
	"gokomodo_test/internal/constants"
	"strings"
)

var (
	createOrder  = "INSERT INTO `order`(id, buyer_id, seller_id, delivery_source_address, delivery_destination_address, item_id, quantity, price, total_price, status) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	approveOrder = "UPDATE `order` SET status = ? WHERE id = ? AND seller_id = ?"
)

func (r *OrderDbRepository) BuildGetOrderQuery(filter *base_model.FilterListModel, userType string, returnCountOnly bool) (string, []interface{}) {
	var (
		whereQuery   = ""
		args         = make([]interface{}, 0)
		whereQueries = make([]string, 0)
	)

	if userType == constants.SELLER_TYPE {
		whereQueries = append(whereQueries, "o.seller_id = ?")
		args = append(args, filter.UserId)
	}

	if userType == constants.BUYER_TYPE {
		whereQueries = append(whereQueries, "o.buyer_id = ?")
		args = append(args, filter.UserId)
	}

	whereQuery = strings.Join(whereQueries, " AND ")

	if returnCountOnly {
		return fmt.Sprintf("SELECT COUNT(o.id) FROM `order` o LEFT JOIN seller s ON o.seller_id = s.id LEFT JOIN buyer b ON o.buyer_id = b.id WHERE %s", whereQuery), args

	}

	return fmt.Sprintf("SELECT o.id, o.buyer_id, b.name as buyer_name, o.seller_id, s.name as seller_name, o.delivery_source_address, o.delivery_destination_address, o.item_id, p.product_name, o.quantity, o.price, o.total_price, o.status FROM `order` o LEFT JOIN seller s ON o.seller_id = s.id LEFT JOIN buyer b ON o.buyer_id = b.id LEFT JOIN product p ON o.item_id = p.id WHERE %s LIMIT %d OFFSET %d", whereQuery, filter.Limit, filter.Offset), args
}
