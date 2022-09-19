package repository

var (
	getBuyer = `
		SELECT 
			id,
			email,
			name,
			address
		FROM buyer
		WHERE 
			email = ? AND
			password = ?
	`
	getSeller = `
		SELECT 
			id,
			email,
			name,
			address
		FROM seller
		WHERE 
			email = ? AND
			password = ?
	`
	getBuyerById = `
		SELECT 
			id,
			email,
			name,
			address
		FROM buyer
		WHERE 
			id = ?
	`
	getSellerById = `
		SELECT 
			id,
			email,
			name,
			address
		FROM seller
		WHERE 
			id = ?
	`
)
