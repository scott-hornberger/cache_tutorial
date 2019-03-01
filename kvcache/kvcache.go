package kvcache

import (
	"fmt"
)

// Cache Interface to be Implemented
type KeyValueCache interface{
	Put(key, value string) error
	Read(key string) (string,error)
	Update(key,value string) error
	Delete(key string) error
}

// Simple Implementation of the Cache Interface
type SimpleKeyValueCache struct{
	data map[string]string
}

// Constructor
func NewSimpleKVCache() *SimpleKeyValueCache{
	return &SimpleKeyValueCache{map[string]string{}}
}

func (c *SimpleKeyValueCache) Put(key,value string) error{
	if key =="" || value =="" {
		return fmt.Errorf("put failed: args must not be empty strings")
	}
	_, ok := c.data[key]
	if ok {
		return fmt.Errorf("put failed: key '%v' already exists in cache", key)
	}
	c.data[key] = value
	return nil
}

func (c *SimpleKeyValueCache) Read(key string) (string,error){
	f, keyExists := c.data[key]
	if !keyExists {
		return "",fmt.Errorf("read failed: key '%v' not in cache", key)
	}
	return f, nil
}

func (c *SimpleKeyValueCache) Update(key, value string) error{
	_, keyExists := c.data[key]
	if keyExists {
		c.data[key] = value
		return nil
	}
	return fmt.Errorf("update failed: key '%v' not in cache", key)
}

func (c *SimpleKeyValueCache) Delete(key string) error{
	_, keyExist := c.data[key]
	if keyExist {
		delete(c.data, key)
		return nil
	}
	return fmt.Errorf("delete failed: key '%v' not in cache",key)
}
