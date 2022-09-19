package repository

import (
	"fmt"
	base_model "gokomodo_test/internal/app/model"
	"gokomodo_test/internal/constants"
	"strings"
)

var (
	insertProduct  = "INSERT INTO product(id, product_name, description, price, seller_id) VALUES(?, ?, ?, ?, ?)"
	getProductById = "SELECT p.id, p.product_name, p.description, p.price, s.id as seller_id, s.name as seller_name FROM product p LEFT JOIN seller s ON p.seller_id = s.id WHERE p.id = ?"
)

func (r *ProductDbRepository) BuildGetProductQuery(filter *base_model.FilterListModel, userType string, returnCountOnly bool) (string, []interface{}) {
	var (
		whereQuery   = ""
		args         = make([]interface{}, 0)
		whereQueries = make([]string, 0)
		whereExist   = false
	)

	if userType == constants.SELLER_TYPE {
		whereQueries = append(whereQueries, "s.id = ?")
		args = append(args, filter.UserId)
	}

	if filter.Search != "" {
		whereQueries = append(whereQueries, "LOWER(p.product_name) LIKE CONCAT('%', ?, '%')")
		args = append(args, filter.Search)
	}

	if len(whereQueries) > 0 {
		whereExist = true
		whereQuery = strings.Join(whereQueries, " AND ")
	}

	if returnCountOnly {
		if !whereExist {
			return "SELECT COUNT(p.id) FROM product p LEFT JOIN seller s ON p.seller_id = s.id", args
		}
		return fmt.Sprintf("SELECT COUNT(p.id) FROM product p LEFT JOIN seller s ON p.seller_id = s.id WHERE %s", whereQuery), args

	}

	if !whereExist {
		return fmt.Sprintf("SELECT p.id, p.product_name, p.description, p.price, s.id as seller_id, s.name FROM product p LEFT JOIN seller s ON p.seller_id = s.id LIMIT %d OFFSET %d", filter.Limit, filter.Offset), args

	}
	return fmt.Sprintf("SELECT p.id, p.product_name, p.description, p.price, s.id as seller_id, s.name FROM product p LEFT JOIN seller s ON p.seller_id = s.id WHERE %s LIMIT %d OFFSET %d", whereQuery, filter.Limit, filter.Offset), args
}
