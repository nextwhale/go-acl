// Copyright 2024 Shaotschaw Teng(github.com/nextwhale). All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ginrouteacl

import (
	"slices"
)

// This ACL package provide a RBAC(Role Based Access Control) model for managing and controlling access
//
// The full structure is  Multiple ACRoles -> Multiple ACScopes -> Multiple permisions
// Note: The concept "Scope" here is like permissions group, it is similar to "Resource" in some other ACL module


// This is a quick method to create role with unique permissions, 
// which means each permission identifier is unique for sure in your system, such as route path.
// And there is no permission scope in your access structure so we can ignore it with this function
func NewRoleWithUniquePermissions(roleID string, name string, permissions []string) *ACRole {
	sc := &ACScope{
		ID:          roleID,
		Permissions: permissions,
	}
	r := &ACRole{
		ID:   roleID,
		Name: name,
	}
	r.AddScope(sc)
	return r
}

// ACL
type ACL struct {
	roles map[string]*ACRole
}

// adding one or more roles to ACL
func (a *ACL) AddRole(roles ...*ACRole) *ACL {
	if a.roles == nil {
	    a.roles = make(map[string]*ACRole)
	}
	for _, role := range roles {
		a.roles[role.ID] = role
	}
	return a
}

// remove one or more ACRole by IDs
func (a *ACL) RemRoleByID(IDs ...string) *ACL {
	for _, ID := range IDs {
		delete(a.roles, ID)
	}
	return a
}

// check whethher role has permission by scoped and permission
func (a *ACL) IsRoleAllowed(roleIDs []string, scopeID string, permission string) bool {
	for _, roleID := range roleIDs {
		if r, ok := a.roles[roleID]; ok {
			if r.IsAllowed(scopeID, permission) {
				return true
			}
		}
	}
	return false
}

// check whethher role has permission by unique permission
func (a *ACL) IsRoleAllowedUniquely(roleIDs []string, permission string) bool {
	for _, roleID := range roleIDs {
		if r, ok := a.roles[roleID]; ok {
			if r.IsAllowedUniquely(permission) {
				return true
			}
		}
	}
	return false
}

// ACRole is a model which has mutiple ACScopes
// It typically stands for for "Position in componay" or "User group"
type ACRole struct {
	ID     string
	Name   string
	Scopes map[string]*ACScope
}

// add one or more Scopes to ACRole
func (acr *ACRole) AddScope(Scopes ...*ACScope) *ACRole {
	if acr.Scopes == nil {
		acr.Scopes = make(map[string]*ACScope)
	}
	for _, v := range Scopes {
		acr.Scopes[v.ID] = v
	}
	return acr
}

// Remove one or more Scopes from user group
func (acr *ACRole) RemScopeByID(IDs ...string) *ACRole {
	for _, v := range IDs {
		delete(acr.Scopes, v)
	}
	return acr
}

// Check whether role has a scoped permission.
// The permissions may be not unique, thus scopeID should be provided
func (acr *ACRole) IsAllowed(scopeID, permission string) bool {
	if acr.Scopes[scopeID] != nil {
		if slices.Contains(acr.Scopes[scopeID].Permissions, permission) {
			return true
		}
	}
	return false
}

// Check whether role has a unique permission.
// Use this function while every permission is unique, even if scope exists
func (acr *ACRole) IsAllowedUniquely(permission string) bool {
	for _, v := range acr.Scopes {
		if slices.Contains(v.Permissions, permission) {
			return true
		}
	}
	return false
}

// ACScope is a combination of permissions, it is the minium unit to classify permissions in application
// It stands for "actions group"
type ACScope struct {
	ID          string
	Name        string
	Permissions []string
}

// add one or more permissions
func (r *ACScope) AddPermission(p ...string) *ACScope {
	r.Permissions = slices.Concat(r.Permissions, p)
	return r
}

// remove one or more permissions
func (r *ACScope) RemPermission(p ...string) *ACScope {
	for _, v := range p {
		if i := slices.Index(r.Permissions, v); i >= 0 {
			r.Permissions = slices.Delete(r.Permissions, i, i+1)
		}
	}
	return r
}
