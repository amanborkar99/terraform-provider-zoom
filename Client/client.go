package client

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const URL string= "https://api.zoom.us/v2/users"

type Client struct{
	HTTPCLient *http.Client
	Token string
}


func NewClient( token string) (*Client, error) {
	c := Client{
		HTTPCLient: &http.Client{Timeout: 10* time.Second},
		Token: token,
	}

	if (token != ""){
		
		return &c, nil
		
		}
		
	return nil, errors.New("invalid token") 
}

func (c *Client) doRequestforCreate(req *http.Request) ([]byte, error){
	req.Header.Set("Authorization", c.Token)
    req.Header.Set("Content-Type","application/json")
	res, err := c.HTTPCLient.Do(req)

	if(err != nil){
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil{
		return nil, err
	}
	
	if res.StatusCode != 201 {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}
	
	return body, err

}	

func (c *Client) doRequestforRead(req *http.Request) ([]byte, error){
	req.Header.Set("Authorization", c.Token)
    req.Header.Set("Content-Type","application/json")
	res, err := c.HTTPCLient.Do(req)

	if(err != nil){
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil{
		return nil, err
	}
	
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}
	
	return body, err

}	
func (c *Client) doRequestforUpdate(req *http.Request) ([]byte, error){
	req.Header.Set("Authorization", c.Token)
    req.Header.Set("Content-Type","application/json")
	res, err := c.HTTPCLient.Do(req)

	if(err != nil){
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil{
		return nil, err
	}
	
	if res.StatusCode != 204 {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}
	
	return body, err

}

func (c *Client) doRequestforDelete(req *http.Request) ([]byte, error){
	req.Header.Set("Authorization", c.Token)
	res, err := c.HTTPCLient.Do(req)
	if err != nil{
		return nil, err
	}
    defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil{
		return nil, err
	}
	
	if res.StatusCode != 204 {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}
	
	return body, err
}






