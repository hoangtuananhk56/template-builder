package user

import (
	"db/mgo"
	"gopkg.in/mgo.v2/bson"
	"template-builder/tbs/o/model"
)

var TableUser = model.NewTable("users", "usr")

func NewUserID() string {
	return TableUser.Next()
}

func AllUserOnBranch(branchid string, role string) ([]User, error) {
	var res = make([]User, 0)
	var query = bson.M{"branch_id": branchid}
	if role != "" {
		query["role"] = role
	}
	return res, TableUser.C().Find(query).All(&res)
}

func (b *User) Create() error {
	if err := b.ensureUniqueUsername(); err != nil {
		return err
	}
	var p = password(b.Password)
	// replace
	if err := p.HashTo(&b.Password); err != nil {
		return err
	}
	return TableUser.Create(b)
}

func MarkDelete(id string) error {
	return TableUser.MarkDelete(id)
}

func (v *User) Update(newValue *User) error {
	var values = map[string]interface{}{
		"firstname": newValue.Firstname,
	}

	// if newValue.Username != v.Username {
	// 	if err := newValue.ensureUniqueUsername(); err != nil {
	// 		return err
	// 	}
	// 	values["username"] = newValue.Username
	// }

	if len(newValue.Password) > 0 {
		if newValue.Password != v.Password {
			var p = password(newValue.Password)
			if err := p.HashTo(&newValue.Password); err != nil {
				return err
			}
		}
		values["password"] = newValue.Password
	}

	if newValue.GetPhone() != v.GetPhone() && newValue.GetPhone() != "" {
		values["phone"] = newValue.GetPhone()
	}
	if newValue.GetOrigin() != v.GetOrigin() && newValue.GetOrigin() != "" {
		values["origin"] = newValue.GetOrigin()
	}
	if newValue.GetFirstname() != v.GetFirstname() && newValue.GetFirstname() != "" {
		values["firstname"] = newValue.GetFirstname()
	}
	if newValue.GetLastname() != v.GetLastname() && newValue.GetLastname() != "" {
		values["lastname"] = newValue.GetLastname()
	}
	if newValue.GetCompany() != v.GetCompany() && newValue.GetCompany() != "" {
		values["company"] = newValue.GetCompany()
	}

	values["supcription_id"] = newValue.SupcriptionID

	values["role"] = newValue.Role
	values["expried"] = newValue.Expried

	return TableUser.UnsafeUpdateByID(v.ID, values)
}

var _ = TableUser.EnsureIndex(mgo.Index{
	Key:        []string{"ctime"},
	Background: true,
})
