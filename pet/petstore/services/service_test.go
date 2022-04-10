package services

import (
	"context"
	"errors"
	"petstore/models"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

// TestService_ServicePost is used to test the post method in the service layer
func TestService_ServicePost(t *testing.T) {
	id := uuid.New()
	pet := models.Pet{ID: id, Name: "bruno", Age: 2, Species: "dog", Status: "Available"}
	testcases := []struct {
		desc   string
		input  models.Pet
		output models.Pet
		err    error
	}{
		{"success case", pet, pet, nil},
	}
	ctrl := gomock.NewController(t)
	MockStore := NewMockStore(ctrl)
	s := New(MockStore)
	ctx := context.Background()

	for _, tc := range testcases {
		MockStore.EXPECT().Post(ctx, pet).Return(pet, nil)

		_, err := s.ServicePost(ctx, tc.input)
		if err != nil {
			return
		}

		assert.Equal(t, tc.err, err)
	}
}

func Test_PostStoreErr(t *testing.T) {
	id := uuid.New()
	pet := models.Pet{ID: id, Name: "bruno", Age: 2, Species: "dog", Status: "Available"}
	testcases := []struct {
		desc   string
		input  models.Pet
		output models.Pet
		err    error
	}{
		{"success case", pet, models.Pet{}, errors.New("error from store layer")},
	}
	ctrl := gomock.NewController(t)
	MockStore := NewMockStore(ctrl)
	s := New(MockStore)
	ctx := context.Background()

	for _, tc := range testcases {
		MockStore.EXPECT().Post(ctx, pet).Return(pet, tc.err)

		_, err := s.ServicePost(ctx, tc.input)
		if err != nil {
			return
		}

		assert.Equal(t, tc.err, err)
	}
}

func TestPostInvalidIP(t *testing.T) {
	id := uuid.New()
	pet := models.Pet{ID: id, Name: "bruno", Age: 12, Species: "Cat", Status: "Available"}
	pet1 := models.Pet{ID: id, Name: "bruno", Age: 12, Species: "Dog", Status: "Available"}
	pet2 := models.Pet{ID: id, Name: "bruno", Age: 12, Species: "Bird", Status: "Available"}
	testcases := []struct {
		desc   string
		input  models.Pet
		output models.Pet
		err    error
	}{
		{"invalid age for cat", pet, models.Pet{}, errors.New("invalid age for Cat")},
		{"invalid age for Dog", pet1, models.Pet{}, errors.New("invalid age for Dog")},
		{"invalid age for Bird", pet2, models.Pet{}, errors.New("invalid age for Bird")},
	}
	ctrl := gomock.NewController(t)
	MockStore := NewMockStore(ctrl)
	s := New(MockStore)
	ctx := context.Background()
	for _, tc := range testcases {
		_, err := s.ServicePost(ctx, tc.input)
		assert.Equal(t, tc.err, err)
	}

}

func TestService_ServicePostStatusErr(t *testing.T) {
	id := uuid.New()
	pet := models.Pet{ID: id, Name: "bruno", Age: 2, Species: "dog", Status: "NotAvailable"}
	testcases := []struct {
		desc   string
		input  models.Pet
		output models.Pet
		err    error
	}{
		{"success case", pet, models.Pet{}, errors.New("invalid status")},
	}

	ctrl := gomock.NewController(t)
	MockStore := NewMockStore(ctrl)
	s := New(MockStore)
	ctx := context.Background()

	for _, tc := range testcases {
		_, err := s.ServicePost(ctx, tc.input)
		if err != nil {
			return
		}

		assert.Equal(t, tc.err, err)
	}
}

// TestService_ServiceGetByID is used to test the GetByID in service layer
func TestService_ServiceGetByID(t *testing.T) {
	id := uuid.New()
	pet := models.Pet{ID: id, Name: "bruno", Age: 2, Species: "dog", Status: "Available"}
	testcases := []struct {
		desc   string
		input  uuid.UUID
		output models.Pet
		err    error
	}{
		{"success", id, pet, nil},
		{"nil id", uuid.Nil, models.Pet{}, errors.New("error from store layer")},
	}
	ctrl := gomock.NewController(t)
	MockStore := NewMockStore(ctrl)
	s := New(MockStore)
	ctx := context.Background()

	for _, tc := range testcases {
		MockStore.EXPECT().GetByID(ctx, tc.input).Return(tc.output, tc.err)
		_, err := s.ServiceGetByID(ctx, tc.input)
		assert.Equal(t, tc.err, err)
	}
}

func TestService_Delete(t *testing.T) {
	id := uuid.New()
	testcases := []struct {
		desc  string
		input uuid.UUID
		err   error
	}{
		{"success case", id, nil},
		{"success case", uuid.Nil, errors.New("error from store layer")},
	}
	ctrl := gomock.NewController(t)
	MockStore := NewMockStore(ctrl)
	s := New(MockStore)
	ctx := context.Background()
	for _, tc := range testcases {
		MockStore.EXPECT().Delete(ctx, tc.input).Return(tc.err)
		err := s.Delete(ctx, tc.input)
		assert.Equal(t, tc.err, err)
	}
}
