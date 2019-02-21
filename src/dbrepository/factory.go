package dbrepository

import (
	"domain"
 )

type Factory struct {
}

func (f *Factory) NewRest(name string, address string, addressLine2 string, URL string,outcode string, postcode string, rating float32,  typeOfFood string)  * domain.Restaurant {
	return &domain.Restaurant{			         
		Name:         name,
		Address:      address,
		AddressLine2: addressLine2,
		URL:          URL,
		Outcode:      outcode,
		Postcode:     postcode,
		Rating :      rating,
		TypeOfFood:   typeOfFood,
		}
	}

