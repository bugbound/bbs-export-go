package main

import (
	"fmt"
	"os"
    "net/http"
    "io/ioutil"
    "time"
    "encoding/json"
)


type DomainPagedRecords struct {
    Num_results int  
    Page int
    Objects []DomainRecord
    Total_pages int
}

type DomainRecord struct {
    Domain string
    Id int
}

type UrlPagedRecords struct {
    Num_results int  
    Page int
    Objects []UrlRecord
    Total_pages int
}

type UrlRecord struct {
    Url string
    Id int
}

func main() {
    if os.Args[1] == "domain" {
        firstPage := new(DomainPagedRecords) 
        link := "http://bbsstore-service:7002/api/dns_store"
        getJson("http://bbsstore-service:7002/api/dns_store?page=1", firstPage)
        totalPages := firstPage.Total_pages
        //fmt.Println(totalPages)
        
        for i := 1; i <= totalPages; i++ {
            //fmt.Println(i)
            concatenated := fmt.Sprintf("%s?page=%d", link, i)
            //fmt.Println(concatenated)
            
            jsonData := new(DomainPagedRecords)
            getJson(concatenated, jsonData)
            
            for currentIndex := range jsonData.Objects {
                fmt.Println(jsonData.Objects[currentIndex].Domain)
            }
        }
    }
    
    if os.Args[1] == "url" {
        firstPage := new(UrlPagedRecords) 
        link := "http://bbsstore-service:7002/api/url_store"
        getJson("http://bbsstore-service:7002/api/url_store?page=1", firstPage)
        totalPages := firstPage.Total_pages
        //fmt.Println(totalPages)
        
        for i := 1; i <= totalPages; i++ {
            //fmt.Println(i)
            concatenated := fmt.Sprintf("%s?page=%d", link, i)
            //fmt.Println(concatenated)
            
            jsonData := new(UrlPagedRecords)
            getJson(concatenated, jsonData)
            
            for currentIndex := range jsonData.Objects {
                fmt.Println(jsonData.Objects[currentIndex].Url)
            }
        }
    }
}


var myClient = &http.Client{Timeout: 10 * time.Second}

func getJson(url string, target interface{}) error {
    r, err := myClient.Get(url)
    if err != nil {
        return err
    }
    defer r.Body.Close()

    return json.NewDecoder(r.Body).Decode(target)
}
