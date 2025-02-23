package user

import "regexp"

type User struct {
	ID       string
	ModelID  string
	Entity   string
	Email    string
	Password string
	Type     UserType
}

type UserEntity struct {
	ID         string
	UserID     string
	Entity     UserEntityName
	Name       string
	Email      string
	LocationID string
}

type UserEntityName int

const (
	Unknown = iota
	Operator
	Carrier
	Courier
)

func (u UserEntityName) String() string {
	switch u {
	case Operator:
		return "operator"
	case Carrier:
		return "carrier"
	case Courier:
		return "courier"
	}

	return ""
}

func (u UserEntityName) CollectionName() string {
	entityName := u.String()

	// Words ending in -y (consonant + y) → replace -y with -ies
	if matched, _ := regexp.MatchString("[^aeiou]y$", entityName); matched {
		return entityName[:len(entityName)-1] + "ies"
	}

	// Words ending in -s, -x, -z, -ch, -sh → add -es
	if matched, _ := regexp.MatchString("(s|x|z|ch|sh)$", entityName); matched {
		return entityName + "es"
	}

	// Words ending in -f or -fe → replace with -ves
	if matched, _ := regexp.MatchString("(fe|f)$", entityName); matched {
		return regexp.MustCompile("(fe|f)$").ReplaceAllString(entityName, "ves")
	}

	// Default: add -s
	return entityName + "s"
}

func StringToUserEntityName(entity string) UserEntityName {
	switch entity {
	case "operator":
		return Operator
	case "carrier":
		return Carrier
	case "courier":
		return Courier
	}

	return Unknown
}
