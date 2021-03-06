// Code generated by jrpc. DO NOT EDIT.

package userdef

import (
	time "time"
)

// User is defined in the .def file
type User struct {
	Id        *uint32    `json:"id,omitempty"`
	Name      *string    `json:"name,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// GetId returns the de-referenced value of Id.
// The second return value states whether the field was set.
func (m *User) GetId() (val uint32, set bool) {
	if m.Id == nil {
		return
	}

	return *m.Id, true
}

// SetId sets the value of Id
func (m *User) SetId(v uint32) *User {
	m.Id = &v
	return m
}

// GetName returns the de-referenced value of Name.
// The second return value states whether the field was set.
func (m *User) GetName() (val string, set bool) {
	if m.Name == nil {
		return
	}

	return *m.Name, true
}

// SetName sets the value of Name
func (m *User) SetName(v string) *User {
	m.Name = &v
	return m
}

// GetCreatedAt returns the de-referenced value of CreatedAt.
// The second return value states whether the field was set.
func (m *User) GetCreatedAt() (val time.Time, set bool) {
	if m.CreatedAt == nil {
		return
	}

	return *m.CreatedAt, true
}

// SetCreatedAt sets the value of CreatedAt
func (m *User) SetCreatedAt(v time.Time) *User {
	m.CreatedAt = &v
	return m
}

// GetUpdatedAt returns the de-referenced value of UpdatedAt.
// The second return value states whether the field was set.
func (m *User) GetUpdatedAt() (val time.Time, set bool) {
	if m.UpdatedAt == nil {
		return
	}

	return *m.UpdatedAt, true
}

// SetUpdatedAt sets the value of UpdatedAt
func (m *User) SetUpdatedAt(v time.Time) *User {
	m.UpdatedAt = &v
	return m
}

// Validate returns an error if any of the fields have bad values
func (m *User) Validate() error {
	return nil
}

// GetUserRequest is defined in the .def file
type GetUserRequest struct {
	UserId *uint32 `json:"user_id,omitempty"`
}

// GetUserId returns the de-referenced value of UserId.
// The second return value states whether the field was set.
func (m *GetUserRequest) GetUserId() (val uint32, set bool) {
	if m.UserId == nil {
		return
	}

	return *m.UserId, true
}

// SetUserId sets the value of UserId
func (m *GetUserRequest) SetUserId(v uint32) *GetUserRequest {
	m.UserId = &v
	return m
}

// Validate returns an error if any of the fields have bad values
func (m *GetUserRequest) Validate() error {
	return nil
}

// GetUserResponse is defined in the .def file
type GetUserResponse struct {
	User *User `json:"user,omitempty"`
}

// GetUser returns the de-referenced value of User.
// The second return value states whether the field was set.
func (m *GetUserResponse) GetUser() (val User, set bool) {
	if m.User == nil {
		return
	}

	return *m.User, true
}

// SetUser sets the value of User
func (m *GetUserResponse) SetUser(v User) *GetUserResponse {
	m.User = &v
	return m
}

// Validate returns an error if any of the fields have bad values
func (m *GetUserResponse) Validate() error {
	if err := m.User.Validate(); err != nil {
		return err
	}

	return nil
}

// ListUsersRequest is defined in the .def file
type ListUsersRequest struct {
}

// Validate returns an error if any of the fields have bad values
func (m *ListUsersRequest) Validate() error {
	return nil
}

// ListUsersResponse is defined in the .def file
type ListUsersResponse struct {
	Users *[]User `json:"users,omitempty"`
}

// GetUsers returns the de-referenced value of Users.
// The second return value states whether the field was set.
func (m *ListUsersResponse) GetUsers() (val []User, set bool) {
	if m.Users == nil {
		return
	}

	return *m.Users, true
}

// SetUsers sets the value of Users
func (m *ListUsersResponse) SetUsers(v []User) *ListUsersResponse {
	m.Users = &v
	return m
}

// Validate returns an error if any of the fields have bad values
func (m *ListUsersResponse) Validate() error {
	if m.Users != nil {
		for _, r := range *m.Users {
			if err := r.Validate(); err != nil {
				return err
			}
		}
	}

	return nil
}
