package index

import (
	"fmt"
	"social-network/core/user"
	"social-network/view/element"
)

// Insert fügt einen neuen Benutzer in den Index ein.
func (idx *Index) Insert(u *user.User) {
	current := idx.root

	for !current.IsEmpty() {
		key := idx.key(u)
		currentKey := idx.key(current.User)

		if key < currentKey {
			current = current.Left
		} else {
			current = current.Right
		}
	}

	current.SetUser(u)
}

// Find sucht einen Benutzer im Index anhand eines Schlüssels.
func (idx *Index) Find(key string) *element.Element {

	current := idx.root

	for !current.IsEmpty() {

		if idx.key(current.User) == key {
			return current

		} else if key < idx.key(current.User) {
			current = current.Left
		} else {
			current = current.Right
		}

	}

	return nil
}

// List liefert alle Benutzer im Index in sortierter Reihenfolge.
func (idx *Index) List() []*user.User {
	return idx.root.List()
}

// Keys liefert eine Liste aller Schlüssel im Index in sortierter Reihenfolge.
func (idx *Index) Keys() []string {
	keys := []string{}

	if idx.root.IsEmpty() {
		return keys
	}
	for _, u := range idx.List() {
		keys = append(keys, idx.key(u))
	}

	return keys
}

// String gibt eine menschenlesbare Listen-Darstellung aller Schlüssel im Index zurück.
func (idx *Index) String() string {

	return fmt.Sprintf("%v", idx.Keys())
}
