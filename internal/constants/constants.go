package constants

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var (
	APPLICATION_NAME          = "Gokomodo Test App"
	LOGIN_EXPIRATION_DURATION = time.Duration(1) * time.Hour
	JWT_SIGNING_METHOD        = jwt.SigningMethodHS256
	JWT_SIGNATURE_KEY         = []byte("Gokomodo Test Signature Key")
	SELLER_TYPE               = "seller"
	BUYER_TYPE                = "buyer"
	PENDING_STATUS            = "pending"
	ACCEPTED_STATUS           = "accepted"
)
