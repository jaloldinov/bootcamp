package controller

import (
	"app/models"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"time"
)

// 1. Order boyicha default holati time sort bolishi kerak. DESC
func (c *Controller) Task_1() []*models.Order {
	resp, _ := c.OrderGetList(&models.OrderGetListRequest{})
	orders := resp.Orders

	sort.Slice(orders, func(i, j int) bool {
		time1, _ := time.Parse("2006-01-02 15:04:05", orders[i].DateTime)
		time2, _ := time.Parse("2006-01-02 15:04:05", orders[j].DateTime)
		return time1.After(time2)
	})
	return orders
}

// 2. Order Date boyicha filter qoyish kerak
// M: 2022-02-01 dan 2022-02-22 gacha
// from_date: 2022-02-22 to_date:  2022-02-22
func (c *Controller) Task_2(req models.Task2request) []*models.Order {
	resp, _ := c.OrderGetList(&models.OrderGetListRequest{})
	orders := resp.Orders

	startDate, err := time.Parse("2006-01-02", req.FromDate)
	if err != nil {
		fmt.Println("Error parsing start date:", err)
		return make([]*models.Order, 0)
	}
	endDate, err := time.Parse("2006-01-02", req.ToDate)
	if err != nil {
		fmt.Println("Error parsing end date:", err)
		return make([]*models.Order, 0)
	}

	var result []*models.Order

	for _, order := range orders {
		createdAt, err := time.Parse("2006-01-02 15:04:05", order.DateTime)
		if err != nil {
			fmt.Println("Error parsing createdAt:", err)
			continue
		}

		if createdAt.After(startDate) && createdAt.Before(endDate.AddDate(0, 0, 1)) {
			result = append(result, order)
		}
	}
	return result
}

// 3. User history chiqish kerak. Ya'ni sotib olgan mahsulotlari korsatish kerak
//  1. Name: Qora futbolka Price: 98000 Count: 3 Total: 294000 Time: 2022-03-29 11:10:52
//  2. Name: Qora futbolka Price: 98000 Count: 3 Total: 294000 Time: 2022-03-29 11:10:52
//  3. Name: Qora futbolka Price: 98000 Count: 3 Total: 294000 Time: 2022-03-29 11:10:52
//  4. Name: Qora futbolka Price: 98000 Count: 3 Total: 294000 Time: 2022-03-29 11:10:52
func (c *Controller) Task_3(ID string) []models.Task3Model {
	productData, _ := c.ProductGetList(&models.ProductGetListRequest{})
	products := productData.Products
	userData, _ := c.UserGetList(&models.UserGetListRequest{})
	users := userData.Users
	cards, _ := readCards("data/shop_cart.json")

	userIdName := make(map[string]string)
	for _, u := range users {
		userIdName[u.Id] = u.Name
	}
	productsIdName := make(map[string]string)
	productIdPrice := make(map[string]int)
	for _, p := range products {
		productIdPrice[p.Id] = p.Price
		productsIdName[p.Id] = p.Name
	}

	var userSlice []models.Task3Model
	for _, card := range cards {
		if card.UserId == ID {
			userSlice = append(userSlice, models.Task3Model{
				Name:  productsIdName[card.ProductId],
				Price: productIdPrice[card.ProductId],
				Count: card.Count,
				Total: (card.Count * productIdPrice[card.ProductId]),
				Time:  card.Time,
			})
		}
	}
	return userSlice
}

// 4. User qancha pul mahsulot sotib olganligi haqida hisobot.
func (c *Controller) Task_4(id string) models.NameCount {
	productData, _ := c.ProductGetList(&models.ProductGetListRequest{})
	products := productData.Products
	userData, _ := c.UserGetList(&models.UserGetListRequest{})
	users := userData.Users
	cards, _ := readCards("data/shop_cart.json")

	userIdName := make(map[string]string)
	for _, u := range users {
		userIdName[u.Id] = u.Name
	}

	productIdPrice := make(map[string]int)
	for _, p := range products {
		productIdPrice[p.Id] = p.Price
	}

	var result models.NameCount
	for _, u := range cards {
		if u.UserId == id {
			result.Name = userIdName[u.UserId]
			result.Count += (productIdPrice[u.ProductId] * u.Count)
		}
	}
	return result
}

// 5. Productlarni Qancha sotilgan boyicha hisobot
func (c *Controller) Task_5() []models.NameCount {
	productData, _ := c.ProductGetList(&models.ProductGetListRequest{})
	products := productData.Products
	// shop card getlist method yo'q shu sababli o'zim datani oldim
	cards, _ := readCards("data/shop_cart.json")

	productIdName := make(map[string]string)
	for _, p := range products {
		productIdName[p.Id] = p.Name
	}

	prodIdTotalSell := make(map[string]int)
	for _, c := range cards {
		prodIdTotalSell[c.ProductId] += c.Count
	}
	var resp []models.NameCount

	for id, count := range prodIdTotalSell {
		resp = append(resp, models.NameCount{
			Name:  productIdName[id],
			Count: count,
		})
	}
	return resp
}

// 6. Top 10 ta sotilayotgan mahsulotlarni royxati.
func (c *Controller) Task_6() []models.NameCount {
	productData, _ := c.ProductGetList(&models.ProductGetListRequest{})
	products := productData.Products
	// shop card getlist method yo'q shu sababli o'zim datani oldim
	cards, _ := readCards("data/shop_cart.json")

	productIdName := make(map[string]string)
	for _, p := range products {
		productIdName[p.Id] = p.Name
	}
	prodIdTotalSell := make(map[string]int)
	for _, c := range cards {
		prodIdTotalSell[c.ProductId] += c.Count
	}

	var resp []models.NameCount
	for id, count := range prodIdTotalSell {
		resp = append(resp, models.NameCount{
			Name:  productIdName[id],
			Count: count,
		})
	}
	sort.Slice(resp, func(i, j int) bool {
		return resp[i].Count > resp[j].Count
	})
	return resp[:10]
}

// 7. TOP 10 ta Eng past sotilayotgan mahsulotlar royxati
func (c *Controller) Task_7() []models.NameCount {
	productData, _ := c.ProductGetList(&models.ProductGetListRequest{})
	products := productData.Products
	// shop card getlist method yo'q shu sababli o'zim datani oldim
	cards, _ := readCards("data/shop_cart.json")

	productIdName := make(map[string]string)
	for _, p := range products {
		productIdName[p.Id] = p.Name
	}

	prodIdTotalSell := make(map[string]int)
	for _, c := range cards {
		prodIdTotalSell[c.ProductId] += c.Count
	}
	var resp []models.NameCount
	for id, count := range prodIdTotalSell {
		resp = append(resp, models.NameCount{
			Name:  productIdName[id],
			Count: count,
		})
	}
	sort.Slice(resp, func(i, j int) bool {
		return resp[i].Count < resp[j].Count
	})
	return resp[:10]
}

//  8. Qaysi Sanada eng kop mahsulot sotilganligi boyicha jadval
//     -- Count Sort DESC
//  1. Name: Asus Sana: 2022-12-20 Count: 10
//  2. Name: HP Sana: 2022-12-20 Count: 8
//  3. Name: CONON Sana: 2022-12-20 Count: 3
//  4. Name: DELL Sana: 2022-12-20 Count: 1
func (c *Controller) Task_8() []models.Task8Model {
	productData, _ := c.ProductGetList(&models.ProductGetListRequest{})
	products := productData.Products
	cards, _ := readCards("data/shop_cart.json")

	productIdName := make(map[string]string)
	for _, p := range products {
		productIdName[p.Id] = p.Name
	}

	resp := make(map[string]map[string]int)
	for _, c := range cards {
		if _, ok := resp[c.Time[:10]]; !ok {
			resp[c.Time[:10]] = make(map[string]int)
		}
		resp[c.Time[:10]][c.ProductId] += c.Count
	}

	var sordedSlice []models.Task8Model
	for time, innerMap := range resp {
		for pID, count := range innerMap {
			sordedSlice = append(sordedSlice, models.Task8Model{
				Name:  productIdName[pID],
				Time:  time,
				Count: count,
			})
		}
	}
	sort.Slice(sordedSlice, func(i, j int) bool {
		return sordedSlice[i].Count > sordedSlice[j].Count
	})

	return sordedSlice
}

//  9. Qaysi category larda qancha mahsulot sotilgan boyicha jadval
//     Name: Electronika Count: 12
func (c *Controller) Task_9() map[string]int {
	categoryData, _ := c.CategoryGetList(&models.CategoryGetListRequest{})
	productData, _ := c.ProductGetList(&models.ProductGetListRequest{})
	// shop card getlist method yo'q shu sababli o'zim datani oldim
	cards, _ := readCards("data/shop_cart.json")
	category := categoryData.Categorys
	products := productData.Products

	prodIdCategoryId := make(map[string]string)
	for _, p := range products {
		prodIdCategoryId[p.Id] = p.CategoryID
	}
	categoryName := make(map[string]string)
	for _, c := range category {
		categoryName[c.Id] = c.Name
	}

	categoryCount := make(map[string]int)
	for _, c := range cards {
		categoryCount[categoryName[prodIdCategoryId[c.ProductId]]] += c.Count
	}
	return categoryCount
}

// 10. Qaysi User eng Active xaridor. Bitta ma'lumot chiqsa yetarli.
func (c *Controller) Task_10() []models.NameCount {
	userData, _ := c.UserGetList(&models.UserGetListRequest{})
	users := userData.Users
	// shop card getlist method yo'q shu sababli o'zim datani oldim
	cards, _ := readCards("data/shop_cart.json")

	userIdName := make(map[string]string)
	for _, u := range users {
		userIdName[u.Id] = u.Name
	}

	userIdCount := make(map[string]int)
	for _, c := range cards {
		userIdCount[userIdName[c.UserId]] += c.Count
	}
	var resp []models.NameCount
	for name, count := range userIdCount {
		resp = append(resp, models.NameCount{
			Name:  name,
			Count: count,
		})
	}
	sort.Slice(resp, func(i, j int) bool {
		return resp[i].Count > resp[j].Count
	})
	return resp[:1]
}

//  11. Agar User 9 dan kop mahuslot sotib olgan bolsa,
//     1 tasi tekinga beriladi va 9 ta uchun pul hisoblanadi.
//     1 tasi eng arzon mahsulotni pulini hisoblamaysiz.

// ORDERPAYMENT ICHIGA HAM YOZIB QOYGANMAN, BU SHUNCHAKI TEKSHIRGANI EDI
func (c *Controller) Task_11() *models.OrderGetList {
	resp, _ := c.OrderGetList(&models.OrderGetListRequest{})
	orders := resp.Orders

	for i := range orders {
		if orders[i].SumCount > 9 {
			minPrice := math.MaxInt64
			minPriceIndex := -1
			count := 0
			for j, item := range orders[i].OrderItems {
				if (item.TotalPrice / item.Count) < minPrice {
					minPrice = item.TotalPrice / item.Count
					minPriceIndex = j
					count = item.Count
				}
			}
			if minPriceIndex != -1 {
				orders[i].Sum -= (orders[i].OrderItems[minPriceIndex].TotalPrice / count)
			}
		}
	}
	return resp
}

// ============== READERS ===================
func readCards(data string) ([]models.ShopCard, error) {
	var cards []models.ShopCard

	d, err := os.ReadFile(data)
	if err != nil {
		log.Printf("Error while Read data: %+v", err)
		return nil, err
	}
	err = json.Unmarshal(d, &cards)
	if err != nil {
		log.Printf("Error while Unmarshal data: %+v", err)
		return nil, err
	}
	return cards, nil
}
