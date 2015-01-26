package schema

import (
	"github.com/elos/data"
)

type RelationshipMap map[data.Kind]map[data.Kind]LinkKind

func (s *RelationshipMap) Valid() bool {
	for outerKind, links := range *s {
		for innerKind, _ /*linkKind*/ := range links {
			innerLinks, ok := (*s)[innerKind]
			if !ok {
				return false
			}

			_ /*matchingLink*/, ok = innerLinks[outerKind]

			if !ok {
				return false
			}
		}
	}

	return true
}

type versionedRelationshipMap struct {
	*RelationshipMap
	version int
}

func NewSchema(sm *RelationshipMap, version int) (Schema, error) {
	s := &versionedRelationshipMap{
		RelationshipMap: sm,
		version:         version,
	}

	if !s.Valid() {
		return nil, InvalidSchemaError
	}

	return s, nil
}

func (s *versionedRelationshipMap) Version() int {
	return s.version
}
