package main

type buyer_rep_t struct {
	Tags				[]string				`json:"tags"`
}

type status_t struct {
	SiteStatus			string					`json:"site_status"`
}

type ratings_t struct {
	Negative			int					`json:"negative"`
	Neutral				int					`json:"neutral"`
	Positive			int					`json:"positive"`
}

type transactions_t struct {
	Cancelled			int					`json:"cancelled"`
	Completed			int					`json:"completed"`
	Period				string					`json:"period"`
	Ratings				ratings_t				`json:"ratings"`
	Total				int					`json:"total"`
}

type seller_reputation_t struct {
	LevelId				string					`json:"level_id"`
	PowerSellerStatus		string					`json:"power_seller_status"`
	Transactions   			transactions_t				`json:"transactions"`
}

type adress_t struct {
	City				string					`json:"city"`
	State				string					`json:"state"`
}

<<<<<<< HEAD:rest/model.go
type MercadoLibreUser struct {
	Id 					int					`json:"id"`
	Nickname			string 				`json:"nickname"`
	RegistrationDate	string 				`json:"registration_date"`
	CountryId			string				`json:"country_id"`
	Adress				adress_t			`json:"adress"`
	UserType			string				`json:"user_type"`
	Logo				string				`json:"logo"`
=======
type MercadoLibreUser struct { 
	Id 				int					`json:"id"`
	Nickname			string 					`json:"nickname"`
	RegistrationDate		string 				`	json:"registration_date"`
	CountryId			string					`json:"country_id"`
	Adress				adress_t				`json:"adress"`
	UserType			string					`json:"user_type"`
	Logo				string					`json:"logo"`
>>>>>>> 9552ebcf85ed7472005b75b5a2e1641718706af0:rest/test.go
	Points				int					`json:"points"`
	SiteId				string					`json:"site_id"`
	Permalink			string					`json:"permalink"`
	Tags				[]string 				`json:"tags"`
	SellerReputation  		seller_reputation_t			`json:"seller_reputation"`
	BuyerReputation			buyer_rep_t 	    			`json:"buyer_reputation"`
	Status 				status_t				`json:"status"`
}


