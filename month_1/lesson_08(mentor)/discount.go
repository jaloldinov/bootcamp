package main

import (
	"errors"
	"fmt"
	"time"
)

type Discount struct {
	Id              int    //unique id
	Name            string //
	Type            string // fixed  || percent
	DiscountType    string // discount || allowance
	Value           int
	DiscountObject  string   //order || product
	Sources         []string // website,bot,app,admin-panel
	PaymentTypes    []string // cash, payme, click,uzum
	ClientsOrderNum int      //1,2,....
	Priority        int
	Active          bool
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type Discounts struct {
	Data []Discount
}

func main() {

	discounts := Discounts{Data: make([]Discount, 0)}

	for i := 0; i < 10; i++ {
		discounts.Create(Discount{
			Name:           "first discount",
			Type:           "percent",
			DiscountType:   "order",
			Value:          15,
			DiscountObject: "order",
			Sources:        []string{"bot"},
			PaymentTypes:   []string{"cash", "click"},
			Priority:       1,
			Active:         true,
		})
	}

	// fmt.Printf("%+v\n", discounts)
	fmt.Println(discounts.GetAll(10, 1))
	// resp, err := discounts.GetById(2)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// } else {
	// 	fmt.Println(resp)
	// }

	// fmt.Println(discounts.DeleteById(2))
}

func (d *Discounts) Create(newDiscount Discount) string {
	newDiscount.Id = len(d.Data)
	newDiscount.CreatedAt = time.Now()
	d.Data = append(d.Data, newDiscount)
	return "successfully added"
}
func (d *Discounts) Update(changedDiscount Discount) string {
	for i, v := range d.Data {
		if v.Id == changedDiscount.Id {
			d.Data[i].Name = changedDiscount.Name
			d.Data[i].Type = changedDiscount.Type
			d.Data[i].DiscountType = changedDiscount.DiscountType
			d.Data[i].Value = changedDiscount.Value
			d.Data[i].DiscountObject = changedDiscount.DiscountObject
			d.Data[i].Sources = changedDiscount.Sources
			d.Data[i].PaymentTypes = changedDiscount.PaymentTypes
			d.Data[i].ClientsOrderNum = changedDiscount.ClientsOrderNum
			d.Data[i].Priority = changedDiscount.Priority
			d.Data[i].Active = changedDiscount.Active
			d.Data[i].UpdatedAt = time.Now()
		}
	}
	return "successfully added"
}

func (d *Discounts) GetAll(limit, page int) []Discount {
	start := (page - 1) * limit
	end := start + limit
	if end > len(d.Data) {
		end = len(d.Data)
	}
	return d.Data[start:end]
}

func (d *Discounts) GetById(id int) (Discount, error) {
	for _, v := range d.Data {

		if v.Id == id {
			return v, nil
		} else {
			return Discount{}, errors.New("not found")
		}
	}
	return Discount{}, errors.New("not found")
}

func (d *Discounts) DeleteById(id int) string {
	for i, v := range d.Data {
		if v.Id == id {
			d.Data = append(d.Data[:i], d.Data[i+1:]...)
			return "successfully deleted"
		}
	}
	return fmt.Sprintf("discount with ID %d not found", id)
}
