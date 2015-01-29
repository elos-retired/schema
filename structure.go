package schema

import (
	"encoding/json"
	"sync"

	"github.com/elos/data"
)

type LinkKind string

const (
	MulLink LinkKind = "MANY"
	OneLink LinkKind = "ONE"
)

type RelationshipMap map[data.Kind]map[data.Kind]LinkKind

func (s *RelationshipMap) valid() bool {
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

type ModelConstructor func() Model

type versionedRelationshipMap struct {
	*RelationshipMap
	registered map[data.Kind]ModelConstructor
	version    int
	data.DB

	sync.Mutex
}

func NewSchema(sm *RelationshipMap, version int) (Schema, error) {
	s := &versionedRelationshipMap{
		RelationshipMap: sm,
		registered:      make(map[data.Kind]ModelConstructor),
		version:         version,
	}

	if !s.valid() {
		return nil, ErrInvalidSchema
	}

	return s, nil
}

func (s *versionedRelationshipMap) Version() int {
	return s.version
}

func (s *versionedRelationshipMap) Register(k data.Kind, c ModelConstructor) {
	s.Lock()
	defer s.Unlock()

	s.registered[k] = c
}

func (s *versionedRelationshipMap) ModelFor(kind data.Kind) (Model, error) {
	s.Lock()
	defer s.Unlock()
	c, ok := s.registered[kind]

	if !ok {
		return nil, ErrUndefinedKind
	}

	return c(), nil
}

func (s *versionedRelationshipMap) Unmarshal(k data.Kind, attrs data.AttrMap) (Model, error) {
	bytes, _ := json.Marshal(attrs)
	m, err := s.ModelFor(k)
	if err != nil {
		return m, err
	}

	return m, json.Unmarshal(bytes, m)
}
