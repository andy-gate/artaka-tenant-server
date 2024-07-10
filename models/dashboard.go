package models

type Dashboard struct {
	Tenant_count		int		`json:"tenant_count"`
	Sales_count			int		`json:"sales_count"`
	Sales_total			int64	`json:"sales_total"`
	Inventory_count		int		`json:"inventory_count"`
}

type QueryDashboard struct {
	Referral_code		string	`json:"referral_code"`
}