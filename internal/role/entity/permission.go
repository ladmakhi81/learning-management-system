package roleentity

import (
	"database/sql/driver"
	"fmt"
	"strings"
)

type Permission string

const (
	// super admin access to all resources
	SUPER_ADMIN Permission = "SUPER_ADMIN"

	// course
	CREATE_COURSE           Permission = "CREATE_COURSE"
	EDIT_COURSE             Permission = "EDIT_COURSE"
	DELETE_COURSE           Permission = "DELETE_COURSE"
	READ_COURSE             Permission = "READ_COURSE"
	VERIFY_COURSE           Permission = "VERIFY_COURSE"
	READ_PARTICIPANT_COURSE Permission = "READ_PARTICIPANT_COURSE"
	VERIFY_VIDEO            Permission = "VERIFY_VIDEO"

	// comment
	READ_COMMENT   Permission = "READ_COMMENT"
	SEND_COMMENT   Permission = "SEND_COMMENT"
	REPORT_COMMENT Permission = "REPORT_COMMENT"

	// role
	CREATE_ROLE Permission = "CREATE_ROLE"
	EDIT_ROLE   Permission = "EDIT_ROLE"
	DELETE_ROLE Permission = "DELETE_ROLE"
	READ_ROLE   Permission = "READ_ROLE"

	// banner
	CREATE_BANNER Permission = "CREATE_BANNER"
	DELETE_BANNER Permission = "DELETE_BANNER"
	READ_BANNER   Permission = "READ_BANNER"
	VERIFY_BANNER Permission = "VERIFY_BANNER"

	// user
	VERIFY_TEACHER  Permission = "VERIFY_TEACHER"
	READ_USER       Permission = "READ_USER"
	CREATE_USER     Permission = "CREATE_USER"
	ASSIGN_ROLE     Permission = "ASSIGN_ROLE"
	CHANGE_PASSWORD Permission = "CHANGE_PASSWORD"
	BLOCK_USER      Permission = "BLOCK_USER"

	// transaction
	READ_TRANSACTION Permission = "READ_TRANSACTION"

	// payment
	READ_PAYMENT Permission = "READ_PAYMENT"

	// order
	READ_ORDER        Permission = "READ_ORDER"
	READ_ORDER_DETAIL Permission = "READ_ORDER_DETAIL"

	// question
	READ_QUESTION       Permission = "READ_QUESTION"
	EDIT_QUESTION       Permission = "EDIT_QUESTION"
	READ_QUESTON_ANSWER Permission = "READ_QUESTON_ANSWER"

	// coupon
	CREATE_DISCOUNT Permission = "CREATE_DISCOUNT"
	READ_DISCOUNT   Permission = "READ_DISCOUNT"
	DELETE_DISCOUNT Permission = "DELETE_DISCOUNT"
	EDIT_DISCOUNT   Permission = "DELETE_DISCOUNT"
	SET_COUPON_FEE  Permission = "SET_COUPON_FEE"

	// category
	CREATE_CATEGORY Permission = "CREATE_CATEGORY"
	DELETE_CATEGORY Permission = "DELETE_CATEGORY"
	UPDATE_CATEGORY Permission = "UPDATE_CATEGORY"
	READ_CATEGORY   Permission = "READ_CATEGORY"

	// carts
	READ_CARTS Permission = "READ_CARTS"

	// forum
	READ_FORUM         Permission = "READ_FORUM"
	READ_FORUM_MESSAGE Permission = "READ_FORUM_MESSAGE"
)

type Permissions []Permission

func (p *Permissions) Scan(value interface{}) error {
	switch v := value.(type) {
	case string:
		permissions := make(Permissions, 0)
		transformedPermissions := strings.ReplaceAll(
			strings.ReplaceAll(
				v,
				"{",
				"",
			),
			"}",
			"",
		)

		for _, item := range strings.Split(transformedPermissions, ",") {
			permissions = append(permissions, Permission(item))
		}
		*p = permissions
		return nil
	default:
		return fmt.Errorf("failed to scan Permissions, invalid type %T", v)
	}
}

func (p Permissions) Value() (driver.Value, error) {
	if len(p) == 0 {
		return nil, nil
	}
	collection := make([]string, 0)
	for _, item := range p {
		collection = append(collection, string(item))
	}
	return fmt.Sprintf("{%s}", strings.Join(collection, ",")), nil
}
