package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	sl "terraform-provider-zoom/Server"
)

func (c *Client) CreateUser(user sl.Create) (*sl.Resp, error) {
	rb, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", "https://api.zoom.us/v2/users", strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}
	body, err := c.doRequestforCreate(req)
	if err != nil {
		return nil, err
	}

	r := sl.Resp{}
	err = json.Unmarshal(body, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

func (c *Client) GetUser(userID string) (*sl.Resp, error) {
	req, err := http.NewRequest("GET", "https://api.zoom.us/v2/users/"+userID, nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequestforRead(req)
	if err != nil {
		return nil, err
	}
	user := sl.Resp{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (c *Client) UpdateUser(id string, user sl.Uz) (error){
   rb, err := json.Marshal(user)
   if err != nil{
	   return err
   }

   req, err := http.NewRequest("PATCH", "https://api.zoom.us/v2/users/"+id, strings.NewReader(string(rb)))
   if err!= nil{
	   return  err
   }
   _, err = c.doRequestforUpdate(req)
   if err != nil{
	   return err
   }
//    r := sl.Uz{}
//    err = json.Unmarshal(body, &r)
//    if err != nil{
// 	   return nil, err
//    }
   
   return nil
}

func (c *Client) DeleteUser(id string, st string) (error){
   var req *http.Request
   var err error 
	if st == "pending"{
		req, err = http.NewRequest("DELETE","https://api.zoom.us/v2/users/"+id+"?action=disassociate", nil)
	}else{
		req, err = http.NewRequest("DELETE","https://api.zoom.us/v2/users/"+id+"?action=delete", nil)
	}
	//req, err := http.NewRequest("DELETE","https://api.zoom.us/v2/users/"+id+"?action=", nil) 
    if err != nil{
		return err
	}
 client := &http.Client{Timeout: 10 * time.Second}
 req.Header.Set("Authorization", c.Token)
 re, err := client.Do(req)
 fmt.Println(re.StatusCode)
 if err != nil {
	 return fmt.Errorf("status: %v", err)
 }
	return nil

}

