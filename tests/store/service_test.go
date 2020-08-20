package main

import (
	"context"
	"testing"

	"github.com/Pierrelx/verre-tech/store"
	storesvc "github.com/Pierrelx/verre-tech/store/implementation"
	assert "github.com/stretchr/testify/assert"
)

var fakeRepo, err = NewFakeService()

var sut store.Service = storesvc.NewService(fakeRepo, nil)

func ClearList() {
	fakeRepo.Stores = make([]store.Store, 0)
}

//#region GetAll
func TestGetAll_Empty(t *testing.T) {
	//Arrange
	var expectedCount = 0
	//Act
	var result, err = sut.GetAll(context.TODO())
	//Assert
	if err != nil {
		t.Error("An error occurred")
	}
	assert.Equal(t, expectedCount, len(result))
}

func TestGetAll_Populated(t *testing.T) {
	//Arrange
	var expectedCount = 1
	var firstStore = store.Store{
		Name: "magas1",
	}
	fakeRepo.Stores = append(fakeRepo.Stores, firstStore)
	defer ClearList()
	//Act
	var result, err = sut.GetAll(context.TODO())
	//Assert
	if err != nil {
		t.Error("An error occurred")
	}
	assert.Equal(t, expectedCount, len(result))
}

//#endregion

//#region GetStoreById
func TestGetStoreById_ValidID(t *testing.T) {
	//Arrange
	var newStore = store.Store{
		ID:   1,
		Name: "magas1",
	}
	fakeRepo.Stores = append(fakeRepo.Stores, newStore)
	defer ClearList()
	//Act
	var result, err = sut.GetByID(context.TODO(), newStore.ID)
	//Assert
	if err != nil {
		t.Error("An error occurred")
	}
	assert.Equal(t, newStore, result)
}

func TestGetStoreById_InvalidID(t *testing.T) {
	//Arrange
	var newStore = store.Store{
		ID:   1,
		Name: "magas1",
	}
	var emptyStore store.Store
	fakeRepo.Stores = append(fakeRepo.Stores, newStore)
	defer ClearList()
	//Act
	var result, err = sut.GetByID(context.TODO(), 99)
	//Assert
	if err != nil {
		t.Error("An error occurred")
	}
	assert.NotEqual(t, newStore, result) //does not returns first any store from list
	assert.Equal(t, emptyStore, result)  //returns empty store object
}

//#endregion

//#region Create
func TestCreate_ReturnsCreatedStoreID(t *testing.T) {
	//Arrange
	var newStore = store.Store{
		Name: "magas1",
	}
	var expectedID = int64(1)
	//Act
	var result, err = sut.Create(context.TODO(), newStore)
	defer ClearList()
	//Assert
	if err != nil {
		t.Error("An error occurred")
	}
	assert.Equal(t, expectedID, result)
}

func TestCreate_AddsStore(t *testing.T) {
	//Arrange
	var newStore = store.Store{
		Name: "magas1",
	}
	var storeCountBeforeCreate = len(fakeRepo.Stores)
	var expectedNewCount = 1
	//Act
	var _, err = sut.Create(context.TODO(), newStore)
	defer ClearList()
	var storeCountAfterCreate = len(fakeRepo.Stores)
	//Assert
	if err != nil {
		t.Error("An error occurred")
	}
	assert.Greater(t, storeCountAfterCreate, storeCountBeforeCreate)
	assert.Len(t, fakeRepo.Stores, expectedNewCount)
}

//#endregion

//#region UpdateStore
func TestUpdateStore_OverwriteFields(t *testing.T) {
	//Arrange
	var newStore = store.Store{
		ID:         1,
		Name:       "magas1",
		Type:       "type",
		Address:    "adress",
		PostalCode: "postalcode",
		City:       "city",
		County:     "county",
		Latitude:   1.1,
		Longitude:  2.2,
	}
	fakeRepo.Stores = append(fakeRepo.Stores, newStore)
	defer ClearList()
	var updatedStore = store.Store{
		ID:         newStore.ID,
		Name:       "2gasin",
		Type:       "updatedtype",
		Address:    "updatedadress",
		PostalCode: "updatedpostalcode",
		City:       "updatedcity",
		County:     "updatedcounty",
		Latitude:   3.3,
		Longitude:  4.4,
	}
	//Act
	var result, err = sut.UpdateStore(context.TODO(), updatedStore)
	//Assert
	if err != nil {
		t.Error("An error occurred")
	}
	assert.Equal(t, newStore.ID, result.ID)
	assert.Equal(t, updatedStore.Name, result.Name)
	assert.Equal(t, updatedStore.Type, result.Type)
	assert.Equal(t, updatedStore.Address, result.Address)
	assert.Equal(t, updatedStore.PostalCode, result.PostalCode)
	assert.Equal(t, updatedStore.City, result.City)
	assert.Equal(t, updatedStore.County, result.County)
	assert.Equal(t, updatedStore.Latitude, result.Latitude)
	assert.Equal(t, updatedStore.Longitude, result.Longitude)
}

func TestUpdateStore_UpdatesTelemetryFields(t *testing.T) {
	//Arrange
	var newStore = store.Store{
		ID:         1,
		Name:       "magas1",
		Type:       "type",
		Address:    "adress",
		PostalCode: "postalcode",
		City:       "city",
		County:     "county",
		Latitude:   1.1,
		Longitude:  2.2,
	}
	fakeRepo.Stores = append(fakeRepo.Stores, newStore)
	defer ClearList()
	var updatedStore = store.Store{
		ID:         newStore.ID,
		Name:       "2gasin",
		Type:       "updatedtype",
		Address:    "updatedadress",
		PostalCode: "updatedpostalcode",
		City:       "updatedcity",
		County:     "updatedcounty",
		Latitude:   3.3,
		Longitude:  4.4,
	}
	//Act
	var result, err = sut.UpdateStore(context.TODO(), updatedStore)
	//Assert
	if err != nil {
		t.Error("An error occurred")
	}
	assert.NotNil(t, result.UpdatedOn)
}

//#endregion

//#region DeleteStore
func TestDeleteStore_RemovesEntry(t *testing.T) {
	//Arrange
	var newStore = store.Store{
		ID:   1,
		Name: "magas1",
	}
	fakeRepo.Stores = append(fakeRepo.Stores, newStore)
	var storeCountBeforeDelete = len(fakeRepo.Stores)
	var expectedNewCount = storeCountBeforeDelete - 1
	//Act
	var err = sut.DeleteStore(context.TODO(), newStore.ID)
	defer ClearList()
	var storeCountAfterDelete = len(fakeRepo.Stores)
	//Assert
	if err != nil {
		t.Error("An error occurred")
	}
	assert.Less(t, storeCountAfterDelete, storeCountBeforeDelete)
	assert.Len(t, fakeRepo.Stores, expectedNewCount)
}

//#endregion
