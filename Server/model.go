package server

type Page struct {
	Page_count   int    `json:"page_count"`
	Page_number  int    `json:"page_number"`
	Page_size    int    `json:"page_size"`
	Next_page_tk string `json:"next_page_token,omitempty"`
	User         []User `json:"users"`
}

type User struct {
	Id         string `json:"id"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Email      string `json:"email"`
	Type       int    `json:"type"`
	Pmi        int    `json:"pmi"`
	Verified   int    `json:"verified"`
	Status     string `json:"status"`
}

type Create struct {
	Action    string `json:"action"`
	User_info Uz     `json:"user_info"`
}

type Uz struct {
	Email      string `json:"email"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Type       int    `json:"type"`
}

type Resp struct {
	Id         string `json:"id"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Email      string `json:"email"`
	Type       int    `json:"type"`
}
