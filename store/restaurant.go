package store

import (
	"errors"
	"math/rand"

	"github.com/vence722/gcoll/maps"
)

var (
	Restaurants *restaurantStore = (*restaurantStore)(maps.NewTypedSyncMap[string, struct{}]())
)

type restaurantStore maps.TypedSyncMap[string, struct{}]

func (this *restaurantStore) Len() int {
	var l int
	mThis := (*maps.TypedSyncMap[string, struct{}])(this)
	mThis.Range(func(_ string, _ struct{}) bool {
		l += 1
		return true
	})
	return l
}

func (this *restaurantStore) Add(restaurant string) {
	mThis := (*maps.TypedSyncMap[string, struct{}])(this)
	mThis.Store(restaurant, struct{}{})
}

func (this *restaurantStore) ListAll() []string {
	var restaurants []string
	mThis := (*maps.TypedSyncMap[string, struct{}])(this)
	mThis.Range(func(restaurant string, _ struct{}) bool {
		restaurants = append(restaurants, restaurant)
		return true
	})
	return restaurants
}

func (this *restaurantStore) Plan(days int) ([]string, error) {
	if this.Len() < days {
		return nil, errors.New("not enough items")
	}
	allRestaurants := this.ListAll()
	rand.Shuffle(len(allRestaurants), func(i, j int) {
		allRestaurants[i], allRestaurants[j] = allRestaurants[j], allRestaurants[i]
	})
	return allRestaurants[:days], nil
}

func (this *restaurantStore) Clear() {
	mThis := (*maps.TypedSyncMap[string, struct{}])(this)
	mThis.Range(func(restaurant string, _ struct{}) bool {
		mThis.Delete(restaurant)
		return true
	})
}
