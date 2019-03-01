package kvcache

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSimpleKeyValueCache(t *testing.T) {
	t.Run("new cache created", func(t *testing.T) {
		testCache := NewSimpleKVCache()
		assert.NotNil(t, testCache)
	})
}

func TestPut(t *testing.T) {
	t.Run("it can put and read", func(t *testing.T) {

		testCache := &SimpleKeyValueCache{map[string]string{}}
		require.NotNil(t, testCache)

		key := "testKey"
		value := "testValue"
		err := testCache.Put(key,value)

		assert.NoError(t,err)
		b, _ := testCache.Read(key)
		assert.Equal(t, b, value)
	})

	t.Run("second put test", func(t *testing.T) {
		testCache := &SimpleKeyValueCache{map[string]string{}}
		require.NotNil(t, testCache)
		key2 := "123"
		value2 := "Sooz"

		err2 := testCache.Put(key2, value2)
		assert.NoError(t, err2)

		a,_ := testCache.Read(key2)
		assert.Equal(t, a, value2)
	})

	//added to align with read error and tests
	t.Run(" put test for error working", func(t *testing.T) {
		testCache := &SimpleKeyValueCache{map[string]string{}}
		require.NotNil(t, testCache)
		key2 := ""
		value2 := ""

		err2 := testCache.Put(key2, value2)
		assert.Error(t,err2,"put failed: check key '' and value '' parameters")

		_,err := testCache.Read(key2)
		assert.ObjectsAreEqualValues(err, "read failed: key '' invalid")
	})
}


func TestRead(t *testing.T){
	t.Run("it can read", func(t *testing.T) {
		testCache := &SimpleKeyValueCache{map[string]string{}}
		require.NotNil(t, testCache)

		key := "name"
		value := "Scott"

		err := testCache.Put(key,value)

		assert.NoError(t, err)

		f, _ := testCache.Read(key)

		assert.Equal(t, f, value)
	})

	t.Run("read test for diff keys", func(t *testing.T) {
		testCache := &SimpleKeyValueCache{map[string]string{}}
		require.NotNil(t, testCache)

		key := "name"
		value := "Benelli"

		key2 := "nickname"
		value2 := "Benny"


		err := testCache.Put(key, value)
		assert.NoError(t, err)

		f, _ := testCache.Read(key)
		assert.Equal(t, f, value)

		err2 := testCache.Put(key2, value2)
		assert.NoError(t, err2)

		v, _ := testCache.Read(key2)
		assert.Equal(t,v,value2)

		//being sure that the Read is reading different values for different keys with different test
		assert.NotEqual(t, f,v)
	})

	t.Run("read test for error working", func(t *testing.T) {
		testCache := &SimpleKeyValueCache{map[string]string{}}
		require.NotNil(t, testCache)

		key := "name"
		value := "Scott"

		err := testCache.Put(key, value)
		assert.NoError(t, err)

		f, _ := testCache.Read(key)

		assert.Equal(t, f, value)

		_, err2 := testCache.Read("")

		//updated tests to reflect new Read method signature and used Objects are Equal values due to the indirect reference to the error message in the assertion
		assert.ObjectsAreEqualValues(err2, "read failed: key ' ' invalid")
	})
}

func TestUpdate(t *testing.T){
	t.Run("it can update", func(t *testing.T) {
		testCache := &SimpleKeyValueCache{map[string]string{}}
		require.NotNil(t, testCache)

		key := "name"
		value := "Benelli"

		put := testCache.Put(key,value)
		assert.NoError(t,put)

		key = "name"
		value = "Benny"
		err := testCache.Update(key, value)

		assert.Equal(t, err, nil)

		_, read := testCache.Read(key)
		assert.ObjectsAreEqualValues(read, value)
	})
	
	t.Run("update error works", func(t *testing.T) {
		testCache := &SimpleKeyValueCache{map[string]string{}}
		require.NotNil(t, testCache)

		key := "name"
		value := "Hero"
		err := testCache.Update(key, value)

		assert.ObjectsAreEqualValues(err, "update failed: key '%v' not in cache")

		_, read := testCache.Read(key)
		assert.ObjectsAreEqualValues(read, value)
	})

	t.Run("empty key Update error test", func(t *testing.T) {
		testCache := &SimpleKeyValueCache{map[string]string{}}
		require.NotNil(t, testCache)

		key := "name"
		value := "Benelli"

		put := testCache.Put(key,value)
		assert.NoError(t,put)

		key = ""
		value = "Benny"
		err := testCache.Update(key, value)

		assert.ObjectsAreEqualValues(err, "update failed: key '%v' not in cache")

		_, read := testCache.Read(key)
		assert.ObjectsAreEqualValues(read, value)
	})

}

func TestDelete(t *testing.T){
	t.Run("it deletes", func(t *testing.T){
		testCache := &SimpleKeyValueCache{map[string]string{}}
		require.NotNil(t, testCache)

		key := "name"
		value := "Benelli"

		put := testCache.Put(key,value)
		assert.NoError(t,put)

		err := testCache.Delete(key)

		//trying a different assert method based on the fact that a successful delete returns nil
		assert.Nil(t,err,"delete test successful")
	})

	t.Run("delete error test", func(t *testing.T) {
		testCache := &SimpleKeyValueCache{map[string]string{}}
		require.NotNil(t, testCache)

		key := "cat"

		err := testCache.Delete(key)

		assert.Error(t, err, "delete error works as expected")
	})

	t.Run("delete error test with empty key string", func(t *testing.T) {
		testCache := &SimpleKeyValueCache{map[string]string{}}
		require.NotNil(t, testCache)

		key := ""

		err := testCache.Delete(key)

		assert.Error(t, err, "delete error works as expected")
	})
}
