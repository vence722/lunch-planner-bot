package store

import (
	"errors"
	"math/rand"
	"time"
)

var (
	Restaurants = RestaurantStroe{}
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type RestaurantStroe map[string]struct{}

func (this *RestaurantStroe) Add(restaurant string) {
	(*this)[restaurant] = struct{}{}
}

func (this *RestaurantStroe) ListAll() []string {
	var restaurants = make([]string, 0, len(*this))
	for restaurant := range *this {
		restaurants = append(restaurants, restaurant)
	}
	return restaurants
}

func (this *RestaurantStroe) Plan(days int) ([]string, error) {
	if len(*this) < days {
		return nil, errors.New("not enough items")
	}
	allRestaurants := this.ListAll()
	rand.Shuffle(len(allRestaurants), func(i, j int) {
		allRestaurants[i], allRestaurants[j] = allRestaurants[j], allRestaurants[i]
	})
	return allRestaurants[:days], nil
}
