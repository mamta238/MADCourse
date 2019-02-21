package usercrudhandler

import  "gopkg.in/mgo.v2/bson"

type RestCreateReqDTO struct {
	Id	         bson.ObjectId  `json:"id" bson:"_id"`
	Name         string  `json:"name" bson:"name"`
	Address      string  `json:"address" bson:"address"`
	AddressLine2 string  `json:"address_line_2" bson:"addressLine2"`
	URL          string  `json:"url" bson:"url"`
	Outcode      string  `json:"outcode" bson:"outcode"`
	Postcode     string  `json:"postcode" bson:"postcode"`
	Rating       float32 `json:"rating" bson:"rating"`
	TypeOfFood   string  `json:"type_of_food" bson:"typeOfFood"`
}

type UserUpdateDTO struct {
	ID		  bson.ObjectId `json:"id"`		
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int    `json:"age"`
}


