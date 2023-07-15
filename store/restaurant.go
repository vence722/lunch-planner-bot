package store

import (
	"errors"
	"math/rand"
)

var (
	Restaurants = restaurantStroe{}
)

type restaurantStroe map[string]struct{}

func (this *restaurantStroe) Add(restaurant string) {
	(*this)[restaurant] = struct{}{}
}

func (this *restaurantStroe) ListAll() []string {
	var restaurants = make([]string, 0, len(*this))
	for restaurant := range *this {
		restaurants = append(restaurants, restaurant)
	}
	return restaurants
}

func (this *restaurantStroe) Plan(days int) ([]string, error) {
	if len(*this) < days {
		return nil, errors.New("not enough items")
	}
	allRestaurants := this.ListAll()
	rand.Shuffle(len(allRestaurants), func(i, j int) {
		allRestaurants[i], allRestaurants[j] = allRestaurants[j], allRestaurants[i]
	})
	return allRestaurants[:days], nil
}

func (this *restaurantStroe) Clear() {
	for restaurant := range *this {
		delete(*this, restaurant)
	}
}
