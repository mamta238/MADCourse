package usercrudhandler

import (
	"encoding/json"
	"io/ioutil"
	logger "log"
	"net/http"
	
	"github.com/gorilla/mux"
	"dbrepository"
	customerrors "packages/errors"
	"packages/httphandlers"
	mthdroutr "packages/mthdrouter"
	"packages/resputl"
	"gopkg.in/mgo.v2/bson"
)

type UserCrudHandler struct {
	httphandlers.BaseHandler
	usersvc dbrepository.Repository
}

func NewUserCrudHandler(usersvc dbrepository.Repository) *UserCrudHandler {
	return &UserCrudHandler{usersvc: usersvc}
}

func (p *UserCrudHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := mthdroutr.RouteAPICall(p, r)
	response.RenderResponse(w)
}




//Get http method to get data

func (p *UserCrudHandler) Get(r *http.Request) resputl.SrvcRes {
	
	
	pathParam := mux.Vars(r)
	usID := pathParam["id"]
	query := r.URL.Query()

	if value,ok := query["name"]; ok{
			resp,err := p.usersvc.FindByName(value[0])
			if err != nil {
				return resputl.ReponseCustomError(err)
			}
			
			responseObj := transformobjListToResponse(resp)
			return resputl.Response200OK(responseObj)

	}	else if value,ok := query["postcode"]; ok{
			resp,err := p.usersvc.FindByTypeOfPostCode(value[0])
			if err != nil {
				return resputl.ReponseCustomError(err)
			}
			
			responseObj := transformobjListToResponse(resp)
			return resputl.Response200OK(responseObj)

	}	else if value,ok := query["searchTerm"] ; ok{
			resp,err := p.usersvc.Search(value[0])
			if err != nil {
				return resputl.ReponseCustomError(err)
			}
		
			responseObj := transformobjListToResponse(resp)
			return resputl.Response200OK(responseObj)

	}	else if value,ok := query["typeOfFood"] ; ok{
			resp,err := p.usersvc.FindByTypeOfFood(value[0])
			if err != nil {
				return resputl.ReponseCustomError(err)
			}
			
			responseObj := transformobjListToResponse(resp)
			return resputl.Response200OK(responseObj)
	}	else {
			if usID == "" {
			
			resp, err := p.usersvc.GetAll()

			if err != nil {
				return resputl.ReponseCustomError(err)
			}
			
			responseObj := transformobjListToResponse(resp)
			return resputl.Response200OK(responseObj)

				
		} else {
			obj, err := p.usersvc.Get(bson.ObjectIdHex(usID))

			if err != nil {
				return resputl.ProcessError(customerrors.NotFoundError("User Object Not found"), "")
		}

			restObj := RestGetRespDTO{
				Id 				: obj.Id,					         
				Name 			: obj.Name,         	
				Address 		: obj.Address,     
				AddressLine2 	: obj.AddressLine2,
				URL				: obj.URL,         
				Outcode 		: obj.Outcode,     
				Postcode 		: obj.Postcode,   
				Rating 			: obj.Rating,      
				TypeOfFood 		: obj.TypeOfFood ,
			}

		return resputl.Response200OK(restObj)
		}
	}
	
}

//Post method creates new temporary schedule
func (p *UserCrudHandler) Post(r *http.Request) resputl.SrvcRes {
	
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return resputl.ReponseCustomError(err)
	}
	/*e, err := ValidateRestCreateUpdateRequest(string(body))
	logger.Println("After")
	if e == false {
		return resputl.ProcessError(err, body)
		return resputl.SimpleBadRequest("Invalid Input Data")

	}*/
	logger.Printf("Received POST request to Create schedule %s ", string(body))
	var requestdata *RestCreateReqDTO
	err = json.Unmarshal(body, &requestdata)
	if err != nil {
		return resputl.SimpleBadRequest("Error unmarshalling Data")
	}

	f := dbrepository.Factory{}
	userObj := f.NewRest(requestdata.Name, requestdata.Address, requestdata.AddressLine2, requestdata.URL,requestdata.Outcode, requestdata.Postcode, requestdata.Rating,  requestdata.TypeOfFood)
	id, err := p.usersvc.Store(userObj)
	if err != nil {
		logger.Fatalf("Error while creating in DB: %v", err)
		return resputl.ProcessError(customerrors.UnprocessableEntityError("Error in writing to DB"), "")
	}

	return resputl.Response200OK(&RestCreateRespDTO{ID: id})
	
}


//Put method modifies temporary schedule contents
func (p *UserCrudHandler) Put(r *http.Request) resputl.SrvcRes {
	
	
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return resputl.ReponseCustomError(err)
	}
	/*
	e, err := ValidateRestCreateUpdateRequest(string(body))

	logger.Println("After put 1")
	if e == false {
		return resputl.ProcessError(err, body)
		return resputl.SimpleBadRequest("Invalid Input Data")

	}*/

	logger.Printf("Received PUT request to Create schedule %s ", string(body))
	var requestdata *RestCreateReqDTO
	err = json.Unmarshal(body, &requestdata)
		

	if err != nil {
		return resputl.SimpleBadRequest("Error unmarshalling Data")
	}

	f := dbrepository.Factory{}
	userObj := f.NewRest(requestdata.Name, requestdata.Address, requestdata.AddressLine2, requestdata.URL,requestdata.Outcode, requestdata.Postcode, requestdata.Rating,  requestdata.TypeOfFood)
	userObj.Id = requestdata.Id
	id, err := p.usersvc.Store(userObj)
	if err != nil {
		logger.Fatalf("Error while creating in DB: %v", err)
		return resputl.ProcessError(customerrors.UnprocessableEntityError("Error in writing to DB"), "")
	}

	return resputl.Response200OK(&RestCreateRespDTO{ID: id})
	
}
	


//Delete method removes temporary schedule from db
func (p *UserCrudHandler) Delete(r *http.Request) resputl.SrvcRes {
	
	pathParam := mux.Vars(r)
	id := pathParam["id"]  		//Deleting accordin to id
	
	
	if (id==""){
		return resputl.SimpleBadRequest("Requires Id to delete accordingly")
	}	else {
		
		err := p.usersvc.Delete(bson.ObjectIdHex(id))
		
		if err != nil{
		return resputl.SimpleBadRequest("Entry with requested id does not appear in Schedule")	
		}
	}
	
	
	
	return resputl.Response200OK("Record Deleted")
}
