# Problem Statement

We need to store a set of objects to a key/value store. You can pick the store of your choice like etcd, redis or any DB of your choice.

Assume that you have a set of objects in Go, that implement the following interface:

```
type Object interface {
// GetKind returns the type of the object.
GetKind() string

// GetID returns a unique UUID for the object.
GetID() string

// GetName returns the name of the object. Names are not unique.
GetName() string

// SetID sets the ID of the object.
SetID(string)

// SetName sets the name of the object.
SetName(string)
}
```

Using the data store of your choice you need to provide an implementation of the following interface:

```
type ObjectDB interface {
// Store will store the object in the data store. The object will have a
// name and kind, and the Store method should create a unique ID.
Store(ctx context.Context, object Object) error

	// GetObjectByID will retrieve the object with the provided ID. 
	GetObjectByID(ctx context.Context, id string) (Object, error)

	// GetObjectByName will retrieve the object with the given name. 
	GetObjectByName(ctx context.Context, name string) (Object, error)

	// ListObjects will return a list of all objects of the given kind. 
	ListObjects(ctx context.Context, kind string) ([]Object, error)

	// DeleteObject will delete the object. 
	DeleteObject(ctx context.Context, id string) error 
}
```


Provide a mechanism to test your implementation where multiple objects can be used against the data store. For example, assume the following two objects that implement the interface. Make sure that we can read/write both animals and persons in the data store.

```
type Person struct {
Name string `json:"name"`
ID string   `json:"id"`\
LastName string `json:”last_name”`
Birthday string `json:”birthday”`
BirthDate time.Time `json:”birthday”`
}

func (p *Person) GetKind() string {
return reflect.TypeOf(p).String()
}

func (p *Person) GetID() string {
return p.ID
}

func (p *Person) GetName() string {
return p.Name
}

func (p *Person) SetID(s string) {
p.ID = s
}

func (p *Person) SetName(s string) {
p.Name = s
}
```


```
type Animal struct {
Name string `json:"name"`
ID   string `json:"id"`
Type string `json:”type”`
OwnerID string `json:”owner_id”`
}

func (p *Animal) GetKind() string {
return reflect.TypeOf(p).String()
}

func (p *Animal) GetID() string {
return p.ID
}

func (p *Animal) GetName() string {
return p.Name
}

func (p *Animal) SetID(s string) {
p.ID = s
}

func (p *Animal) SetName(s string) {
p.Name = s
}
```
