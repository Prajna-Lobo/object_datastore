package domain

type IObject interface {
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

type Object struct {
	Kind string `bson:"kind"`
	Id   string `bson:"id"`
	Name string `bson:"name"`
}
