package models

type Tenant struct {
	User_id          string   `json:"user_id"`
	Outlet_id		 string	  `json:"outlet_id"`
	Nama       		 string   `json:"nama"`
	Address			 string	  `json:"address"`
	Referral_code    string   `json:"referral_code"`
}

type QueryTenant struct {
	Referral_code    string   `json:"referral_code"`
}

