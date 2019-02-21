package usercrudhandler

import "domain"
import "fmt"


func transformobjListToResponse(resp []*domain.Restaurant) RestGetListRespDTO {
	responseObj := RestGetListRespDTO{}
	fmt.Println("In transform")
	for _, obj := range resp {
		userObj := RestGetRespDTO{
			Id : obj.Id,					         
			Name : obj.Name,         	
			Address : obj.Address,     
			AddressLine2 : obj.AddressLine2,
			URL :	obj.URL,         
			Outcode : obj.Outcode,     
			Postcode : obj.Postcode,   
			Rating : obj.Rating,      
			TypeOfFood : obj.TypeOfFood  ,
		}
		responseObj.Rests = append(responseObj.Rests, userObj)
	}
	responseObj.Count = len(responseObj.Rests)

	return responseObj
}
