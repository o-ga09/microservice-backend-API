// Code generated by ogen, DO NOT EDIT.

package api

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

func (s *CommonErrorStatusCode) Error() string {
	return fmt.Sprintf("code %d: %+v", s.StatusCode, s.Response)
}

// Ref: #/components/schemas/CommonError
type CommonError struct {
	Code    OptInt32  `json:"code"`
	Message OptString `json:"message"`
}

// GetCode returns the value of Code.
func (s *CommonError) GetCode() OptInt32 {
	return s.Code
}

// GetMessage returns the value of Message.
func (s *CommonError) GetMessage() OptString {
	return s.Message
}

// SetCode sets the value of Code.
func (s *CommonError) SetCode(val OptInt32) {
	s.Code = val
}

// SetMessage sets the value of Message.
func (s *CommonError) SetMessage(val OptString) {
	s.Message = val
}

// CommonErrorStatusCode wraps CommonError with StatusCode.
type CommonErrorStatusCode struct {
	StatusCode int
	Response   CommonError
}

// GetStatusCode returns the value of StatusCode.
func (s *CommonErrorStatusCode) GetStatusCode() int {
	return s.StatusCode
}

// GetResponse returns the value of Response.
func (s *CommonErrorStatusCode) GetResponse() CommonError {
	return s.Response
}

// SetStatusCode sets the value of StatusCode.
func (s *CommonErrorStatusCode) SetStatusCode(val int) {
	s.StatusCode = val
}

// SetResponse sets the value of Response.
func (s *CommonErrorStatusCode) SetResponse(val CommonError) {
	s.Response = val
}

// Ref: #/components/schemas/DeleteNote
type DeleteNote struct {
	Deletednoteid OptUUID   `json:"deletednoteid"`
	Message       OptString `json:"message"`
}

// GetDeletednoteid returns the value of Deletednoteid.
func (s *DeleteNote) GetDeletednoteid() OptUUID {
	return s.Deletednoteid
}

// GetMessage returns the value of Message.
func (s *DeleteNote) GetMessage() OptString {
	return s.Message
}

// SetDeletednoteid sets the value of Deletednoteid.
func (s *DeleteNote) SetDeletednoteid(val OptUUID) {
	s.Deletednoteid = val
}

// SetMessage sets the value of Message.
func (s *DeleteNote) SetMessage(val OptString) {
	s.Message = val
}

func (*DeleteNote) deleteNoteRes() {}

// Ref: #/components/schemas/DeleteUser
type DeleteUser struct {
	Deleteduserid OptUUID   `json:"deleteduserid"`
	Message       OptString `json:"message"`
}

// GetDeleteduserid returns the value of Deleteduserid.
func (s *DeleteUser) GetDeleteduserid() OptUUID {
	return s.Deleteduserid
}

// GetMessage returns the value of Message.
func (s *DeleteUser) GetMessage() OptString {
	return s.Message
}

// SetDeleteduserid sets the value of Deleteduserid.
func (s *DeleteUser) SetDeleteduserid(val OptUUID) {
	s.Deleteduserid = val
}

// SetMessage sets the value of Message.
func (s *DeleteUser) SetMessage(val OptString) {
	s.Message = val
}

func (*DeleteUser) deleteUserRes() {}

// Ref: #/components/schemas/ErrorBadRequest
type ErrorBadRequest struct {
	Code    OptInt32  `json:"code"`
	Message OptString `json:"message"`
}

// GetCode returns the value of Code.
func (s *ErrorBadRequest) GetCode() OptInt32 {
	return s.Code
}

// GetMessage returns the value of Message.
func (s *ErrorBadRequest) GetMessage() OptString {
	return s.Message
}

// SetCode sets the value of Code.
func (s *ErrorBadRequest) SetCode(val OptInt32) {
	s.Code = val
}

// SetMessage sets the value of Message.
func (s *ErrorBadRequest) SetMessage(val OptString) {
	s.Message = val
}

func (*ErrorBadRequest) createNoteRes() {}
func (*ErrorBadRequest) createUserRes() {}
func (*ErrorBadRequest) deleteNoteRes() {}
func (*ErrorBadRequest) deleteUserRes() {}
func (*ErrorBadRequest) getNoteRes()    {}
func (*ErrorBadRequest) getNotesRes()   {}
func (*ErrorBadRequest) getUserRes()    {}
func (*ErrorBadRequest) getUsersRes()   {}
func (*ErrorBadRequest) updateNoteRes() {}
func (*ErrorBadRequest) updateUserRes() {}

// Ref: #/components/schemas/ErrorForbidden
type ErrorForbidden struct {
	Code    OptInt32  `json:"code"`
	Message OptString `json:"message"`
}

// GetCode returns the value of Code.
func (s *ErrorForbidden) GetCode() OptInt32 {
	return s.Code
}

// GetMessage returns the value of Message.
func (s *ErrorForbidden) GetMessage() OptString {
	return s.Message
}

// SetCode sets the value of Code.
func (s *ErrorForbidden) SetCode(val OptInt32) {
	s.Code = val
}

// SetMessage sets the value of Message.
func (s *ErrorForbidden) SetMessage(val OptString) {
	s.Message = val
}

func (*ErrorForbidden) createNoteRes() {}
func (*ErrorForbidden) createUserRes() {}
func (*ErrorForbidden) deleteNoteRes() {}
func (*ErrorForbidden) deleteUserRes() {}
func (*ErrorForbidden) getNoteRes()    {}
func (*ErrorForbidden) getNotesRes()   {}
func (*ErrorForbidden) getUserRes()    {}
func (*ErrorForbidden) getUsersRes()   {}
func (*ErrorForbidden) updateNoteRes() {}
func (*ErrorForbidden) updateUserRes() {}

// Ref: #/components/schemas/ErrorInternalServerError
type ErrorInternalServerError struct {
	Code    OptInt32  `json:"code"`
	Message OptString `json:"message"`
}

// GetCode returns the value of Code.
func (s *ErrorInternalServerError) GetCode() OptInt32 {
	return s.Code
}

// GetMessage returns the value of Message.
func (s *ErrorInternalServerError) GetMessage() OptString {
	return s.Message
}

// SetCode sets the value of Code.
func (s *ErrorInternalServerError) SetCode(val OptInt32) {
	s.Code = val
}

// SetMessage sets the value of Message.
func (s *ErrorInternalServerError) SetMessage(val OptString) {
	s.Message = val
}

func (*ErrorInternalServerError) createNoteRes() {}
func (*ErrorInternalServerError) createUserRes() {}
func (*ErrorInternalServerError) deleteNoteRes() {}
func (*ErrorInternalServerError) deleteUserRes() {}
func (*ErrorInternalServerError) getNoteRes()    {}
func (*ErrorInternalServerError) getNotesRes()   {}
func (*ErrorInternalServerError) getUserRes()    {}
func (*ErrorInternalServerError) getUsersRes()   {}
func (*ErrorInternalServerError) updateNoteRes() {}
func (*ErrorInternalServerError) updateUserRes() {}

// Ref: #/components/schemas/ErrorNotfound
type ErrorNotfound struct {
	Code    OptInt32  `json:"code"`
	Message OptString `json:"message"`
}

// GetCode returns the value of Code.
func (s *ErrorNotfound) GetCode() OptInt32 {
	return s.Code
}

// GetMessage returns the value of Message.
func (s *ErrorNotfound) GetMessage() OptString {
	return s.Message
}

// SetCode sets the value of Code.
func (s *ErrorNotfound) SetCode(val OptInt32) {
	s.Code = val
}

// SetMessage sets the value of Message.
func (s *ErrorNotfound) SetMessage(val OptString) {
	s.Message = val
}

func (*ErrorNotfound) getNoteRes()  {}
func (*ErrorNotfound) getNotesRes() {}
func (*ErrorNotfound) getUserRes()  {}
func (*ErrorNotfound) getUsersRes() {}

// Ref: #/components/schemas/Note
type Note struct {
	ID        OptUUID     `json:"id"`
	Title     OptString   `json:"title"`
	Content   OptString   `json:"content"`
	CreatedAt OptDateTime `json:"createdAt"`
	UpdatedAt OptDateTime `json:"updatedAt"`
}

// GetID returns the value of ID.
func (s *Note) GetID() OptUUID {
	return s.ID
}

// GetTitle returns the value of Title.
func (s *Note) GetTitle() OptString {
	return s.Title
}

// GetContent returns the value of Content.
func (s *Note) GetContent() OptString {
	return s.Content
}

// GetCreatedAt returns the value of CreatedAt.
func (s *Note) GetCreatedAt() OptDateTime {
	return s.CreatedAt
}

// GetUpdatedAt returns the value of UpdatedAt.
func (s *Note) GetUpdatedAt() OptDateTime {
	return s.UpdatedAt
}

// SetID sets the value of ID.
func (s *Note) SetID(val OptUUID) {
	s.ID = val
}

// SetTitle sets the value of Title.
func (s *Note) SetTitle(val OptString) {
	s.Title = val
}

// SetContent sets the value of Content.
func (s *Note) SetContent(val OptString) {
	s.Content = val
}

// SetCreatedAt sets the value of CreatedAt.
func (s *Note) SetCreatedAt(val OptDateTime) {
	s.CreatedAt = val
}

// SetUpdatedAt sets the value of UpdatedAt.
func (s *Note) SetUpdatedAt(val OptDateTime) {
	s.UpdatedAt = val
}

func (*Note) createNoteRes() {}
func (*Note) getNoteRes()    {}

// Ref: #/components/schemas/Notes
type Notes struct {
	TotalCount    OptInt32  `json:"totalCount"`
	Count         OptInt32  `json:"count"`
	NextPagetoken OptString `json:"nextPagetoken"`
	Notes         []Note    `json:"notes"`
}

// GetTotalCount returns the value of TotalCount.
func (s *Notes) GetTotalCount() OptInt32 {
	return s.TotalCount
}

// GetCount returns the value of Count.
func (s *Notes) GetCount() OptInt32 {
	return s.Count
}

// GetNextPagetoken returns the value of NextPagetoken.
func (s *Notes) GetNextPagetoken() OptString {
	return s.NextPagetoken
}

// GetNotes returns the value of Notes.
func (s *Notes) GetNotes() []Note {
	return s.Notes
}

// SetTotalCount sets the value of TotalCount.
func (s *Notes) SetTotalCount(val OptInt32) {
	s.TotalCount = val
}

// SetCount sets the value of Count.
func (s *Notes) SetCount(val OptInt32) {
	s.Count = val
}

// SetNextPagetoken sets the value of NextPagetoken.
func (s *Notes) SetNextPagetoken(val OptString) {
	s.NextPagetoken = val
}

// SetNotes sets the value of Notes.
func (s *Notes) SetNotes(val []Note) {
	s.Notes = val
}

func (*Notes) getNotesRes() {}

// NewOptDateTime returns new OptDateTime with value set to v.
func NewOptDateTime(v time.Time) OptDateTime {
	return OptDateTime{
		Value: v,
		Set:   true,
	}
}

// OptDateTime is optional time.Time.
type OptDateTime struct {
	Value time.Time
	Set   bool
}

// IsSet returns true if OptDateTime was set.
func (o OptDateTime) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptDateTime) Reset() {
	var v time.Time
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptDateTime) SetTo(v time.Time) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptDateTime) Get() (v time.Time, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptDateTime) Or(d time.Time) time.Time {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptInt32 returns new OptInt32 with value set to v.
func NewOptInt32(v int32) OptInt32 {
	return OptInt32{
		Value: v,
		Set:   true,
	}
}

// OptInt32 is optional int32.
type OptInt32 struct {
	Value int32
	Set   bool
}

// IsSet returns true if OptInt32 was set.
func (o OptInt32) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptInt32) Reset() {
	var v int32
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptInt32) SetTo(v int32) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptInt32) Get() (v int32, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptInt32) Or(d int32) int32 {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptString returns new OptString with value set to v.
func NewOptString(v string) OptString {
	return OptString{
		Value: v,
		Set:   true,
	}
}

// OptString is optional string.
type OptString struct {
	Value string
	Set   bool
}

// IsSet returns true if OptString was set.
func (o OptString) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptString) Reset() {
	var v string
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptString) SetTo(v string) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptString) Get() (v string, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptString) Or(d string) string {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptUUID returns new OptUUID with value set to v.
func NewOptUUID(v uuid.UUID) OptUUID {
	return OptUUID{
		Value: v,
		Set:   true,
	}
}

// OptUUID is optional uuid.UUID.
type OptUUID struct {
	Value uuid.UUID
	Set   bool
}

// IsSet returns true if OptUUID was set.
func (o OptUUID) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptUUID) Reset() {
	var v uuid.UUID
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptUUID) SetTo(v uuid.UUID) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptUUID) Get() (v uuid.UUID, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptUUID) Or(d uuid.UUID) uuid.UUID {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// Ref: #/components/schemas/UpdateNote
type UpdateNote struct {
	Title     OptString   `json:"title"`
	Content   OptString   `json:"content"`
	CreatedAt OptDateTime `json:"createdAt"`
	UpdatedAt OptDateTime `json:"updatedAt"`
}

// GetTitle returns the value of Title.
func (s *UpdateNote) GetTitle() OptString {
	return s.Title
}

// GetContent returns the value of Content.
func (s *UpdateNote) GetContent() OptString {
	return s.Content
}

// GetCreatedAt returns the value of CreatedAt.
func (s *UpdateNote) GetCreatedAt() OptDateTime {
	return s.CreatedAt
}

// GetUpdatedAt returns the value of UpdatedAt.
func (s *UpdateNote) GetUpdatedAt() OptDateTime {
	return s.UpdatedAt
}

// SetTitle sets the value of Title.
func (s *UpdateNote) SetTitle(val OptString) {
	s.Title = val
}

// SetContent sets the value of Content.
func (s *UpdateNote) SetContent(val OptString) {
	s.Content = val
}

// SetCreatedAt sets the value of CreatedAt.
func (s *UpdateNote) SetCreatedAt(val OptDateTime) {
	s.CreatedAt = val
}

// SetUpdatedAt sets the value of UpdatedAt.
func (s *UpdateNote) SetUpdatedAt(val OptDateTime) {
	s.UpdatedAt = val
}

func (*UpdateNote) updateNoteRes() {}

// Ref: #/components/schemas/UpdateUser
type UpdateUser struct {
	Name      OptString   `json:"name"`
	Email     OptString   `json:"email"`
	Password  OptString   `json:"password"`
	CreatedAt OptDateTime `json:"createdAt"`
	UpdatedAt OptDateTime `json:"updatedAt"`
}

// GetName returns the value of Name.
func (s *UpdateUser) GetName() OptString {
	return s.Name
}

// GetEmail returns the value of Email.
func (s *UpdateUser) GetEmail() OptString {
	return s.Email
}

// GetPassword returns the value of Password.
func (s *UpdateUser) GetPassword() OptString {
	return s.Password
}

// GetCreatedAt returns the value of CreatedAt.
func (s *UpdateUser) GetCreatedAt() OptDateTime {
	return s.CreatedAt
}

// GetUpdatedAt returns the value of UpdatedAt.
func (s *UpdateUser) GetUpdatedAt() OptDateTime {
	return s.UpdatedAt
}

// SetName sets the value of Name.
func (s *UpdateUser) SetName(val OptString) {
	s.Name = val
}

// SetEmail sets the value of Email.
func (s *UpdateUser) SetEmail(val OptString) {
	s.Email = val
}

// SetPassword sets the value of Password.
func (s *UpdateUser) SetPassword(val OptString) {
	s.Password = val
}

// SetCreatedAt sets the value of CreatedAt.
func (s *UpdateUser) SetCreatedAt(val OptDateTime) {
	s.CreatedAt = val
}

// SetUpdatedAt sets the value of UpdatedAt.
func (s *UpdateUser) SetUpdatedAt(val OptDateTime) {
	s.UpdatedAt = val
}

// Ref: #/components/schemas/User
type User struct {
	ID        OptUUID     `json:"id"`
	Name      OptString   `json:"name"`
	Email     OptString   `json:"email"`
	Password  OptString   `json:"password"`
	CreatedAt OptDateTime `json:"createdAt"`
	UpdatedAt OptDateTime `json:"updatedAt"`
}

// GetID returns the value of ID.
func (s *User) GetID() OptUUID {
	return s.ID
}

// GetName returns the value of Name.
func (s *User) GetName() OptString {
	return s.Name
}

// GetEmail returns the value of Email.
func (s *User) GetEmail() OptString {
	return s.Email
}

// GetPassword returns the value of Password.
func (s *User) GetPassword() OptString {
	return s.Password
}

// GetCreatedAt returns the value of CreatedAt.
func (s *User) GetCreatedAt() OptDateTime {
	return s.CreatedAt
}

// GetUpdatedAt returns the value of UpdatedAt.
func (s *User) GetUpdatedAt() OptDateTime {
	return s.UpdatedAt
}

// SetID sets the value of ID.
func (s *User) SetID(val OptUUID) {
	s.ID = val
}

// SetName sets the value of Name.
func (s *User) SetName(val OptString) {
	s.Name = val
}

// SetEmail sets the value of Email.
func (s *User) SetEmail(val OptString) {
	s.Email = val
}

// SetPassword sets the value of Password.
func (s *User) SetPassword(val OptString) {
	s.Password = val
}

// SetCreatedAt sets the value of CreatedAt.
func (s *User) SetCreatedAt(val OptDateTime) {
	s.CreatedAt = val
}

// SetUpdatedAt sets the value of UpdatedAt.
func (s *User) SetUpdatedAt(val OptDateTime) {
	s.UpdatedAt = val
}

func (*User) createUserRes() {}
func (*User) getUserRes()    {}
func (*User) updateUserRes() {}

// Ref: #/components/schemas/Users
type Users struct {
	TotalCount    OptInt32  `json:"totalCount"`
	Count         OptInt32  `json:"count"`
	NextPagetoken OptString `json:"nextPagetoken"`
	Users         []User    `json:"users"`
}

// GetTotalCount returns the value of TotalCount.
func (s *Users) GetTotalCount() OptInt32 {
	return s.TotalCount
}

// GetCount returns the value of Count.
func (s *Users) GetCount() OptInt32 {
	return s.Count
}

// GetNextPagetoken returns the value of NextPagetoken.
func (s *Users) GetNextPagetoken() OptString {
	return s.NextPagetoken
}

// GetUsers returns the value of Users.
func (s *Users) GetUsers() []User {
	return s.Users
}

// SetTotalCount sets the value of TotalCount.
func (s *Users) SetTotalCount(val OptInt32) {
	s.TotalCount = val
}

// SetCount sets the value of Count.
func (s *Users) SetCount(val OptInt32) {
	s.Count = val
}

// SetNextPagetoken sets the value of NextPagetoken.
func (s *Users) SetNextPagetoken(val OptString) {
	s.NextPagetoken = val
}

// SetUsers sets the value of Users.
func (s *Users) SetUsers(val []User) {
	s.Users = val
}

func (*Users) getUsersRes() {}
