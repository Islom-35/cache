package main

import (
	"fmt"
	"sync"
)

type Cache struct {
	data map[any]interface{}
	mu   sync.Mutex
}

func NewCache() *Cache {
	return &Cache{
		data: make(map[any]interface{}),
	}
}

func (c *Cache) Set(key,value interface{})string{
	c.mu.Lock()
 	defer c.mu.Unlock()
	c.data[key] = value	
	return "saved"
}

func (c *Cache) Get(key any)(interface{},string){
	c.mu.Lock()
	defer c.mu.Unlock()
	val,ok :=c.data[key]

	if ok == false{
		return nil,"this doesn't exists"
	}
	return val, "ok"
}

func (c *Cache) Delete(key any)(string){
	c.mu.Lock()
	defer c.mu.Unlock()
	_,ok :=c.data[key]

	if ok == false{
		return "this does n't exists"
	}
	return "deleted"
}


func main(){
	c := NewCache()
	fmt.Println(c.Set(2,1))
	fmt.Println(c.Get(2))
	fmt.Println(c.Delete(2))
}