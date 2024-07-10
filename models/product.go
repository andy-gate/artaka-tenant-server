package models

type Product struct {
	User_id          string   `json:"user_id"`
	Tenant_name		 string	  `json:"tenant_name"`
	Name             string   `json:"name"`
	Units            string   `json:"units"`
	Quantity	     int   	  `json:"quantity"`
	Price			 int	  `json:"price"`
}

type QueryProduct struct {
	User_id			string	  `json:"user_id"`
}