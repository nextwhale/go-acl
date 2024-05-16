A Role-Based-Access-Control (RBAC) package for general access verification.

## Features
There is a concept "scope"(ACScope) , which is similar to "resource" in some other ACL library. It is necessary when permission identifiers are not unique, and is ignorable when permission identifiers are unique.

## install
```bash
go get -u github.com/nextwhale/go-acl@latest
```
## Usage 1
While permission identifiers are not unique in different scopes, scope must be provided.

#### To initialize ACL
```go
// Example: A merchant business system
var bizACL *ga.ACL

func ACL() {
    if bizACL != nil {
        return bizACL
    }
	// While permission identifiers are not unique, ACScope must be provided
	scope1 := &ga.ACScope{
		ID:          "order_editting",
		Name:        "Order permissions",
		Permissions: []string{"add", "edit", "delete", "close"},
	}
	scope2 := &ga.ACScope{
		ID:          "video_editting",
		Name:        "Video permissions",
		Permissions: []string{"add", "edit", "delete", "audit"},
	}
	roleEditor := &ga.ACRole{
		ID:   "editors",
		Name: "Editors Group",
	}
	roleAssistant := &ga.ACRole{
		ID:"assistants",
		Name: "Assistants Group",
	}

	roleEditor.AddScope(scope1, scope2)
	roleAssistant.AddScope(scope1)

	bizACL := &ga.ACL{}
	bizACL.AddRole(roleEditor, roleAssistant)

    return bizACL
}
```

### To verify
```go
if ACL().IsRoleAllowed([]string{"editors"}, "order_editting", "delete") {
    fmt.Println("Editors have access to delete order!")
}
```


## Usage 2
While permissions are unique in different scopes, scope is not necessary.

### To initialize ACL
```go
// Example: A routes access system
var adminACL *ga.ACL

func ACL() {
    if adminACL != nil {
        return adminACL
    }
    roleAdmin := ga.NewRoleWithUniquePermissions("1", "Administrators Group", []string{"/admin/admin/list", "/admin/admin/edit/:id", "/admin/admin/del/:id"})
    roleEditor := ga.NewRoleWithUniquePermissions("2", "Editors Group", []string{"/admin/article/list", "/admin/article/edit/:id", "/admin/article/del/:id"})

	adminACl := &ga.ACL{}
	adminACl.AddRole(roleAdmin, roleEditor)

    return adminACL
}
```

### To verify permissions
```go
if ACL().IsRoleAllowedUniquely([]string{"1","2"}, "/admin/article/del/:id") {
    fmt.Println("You have access to delete article")
}
```

## Note
if you encounter any issue, feel free to post it. 
And I strongly encourage contributing to this project.

## License
Distributed under MIT License, please see license file in code for more details.
