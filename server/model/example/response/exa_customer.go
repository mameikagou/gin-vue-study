package response

import "github.com/mameikagou/gin-vue-study/server/model/example"

type ExaCustomerResponse struct {
	Customer example.ExaCustomer `json:"customer"`
}
