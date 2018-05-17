package main

import "testing"

func TestReadJSON(t *testing.T) {
	sourceFile := []byte(`{
	    "products": [
	        {
	            "categories": [
	                "Customer Service",
	                "Call Center"
	            ],
	            "twitter": "@freshdesk",
	            "title": "Freshdesk"
	        }
	    ]
	}`)

	productList := Softwareadvice{}
	err := productList.unmarshalJSON(sourceFile)
	if err != nil {
		t.Errorf("unmarshalJSON failed to return a Softwareadvice object: %s", err)
	}

	if len(productList.Products) != 1 {
		t.Errorf("Wrong number of product in object (%d): expected 1", len(productList.Products))
	}

	if productList.Products[0].Twitter != "@freshdesk" {
		t.Errorf("Invalid value of Twitter field (%s): expected @freshdesk", productList.Products[0].Twitter)
	}

	if productList.Products[0].Title != "Freshdesk" {
		t.Errorf("Invalid value of Twitter field (%s): expected Freshdesk", productList.Products[0].Title)
	}
}
