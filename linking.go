package schema

type LinkKind string

const MulLink LinkKind = "MANY"
const OneLink LinkKind = "ONE"

func PossibleLink(s *RelationshipMap, this Model, other Model) (bool, error) {
	thisKind := this.Kind()

	links, ok := (*s)[thisKind]

	if !ok {
		return false, UndefinedKindError
	}

	otherKind := other.Kind()

	_, linkPossible := links[otherKind]

	if !linkPossible {
		return false, UndefinedLinkError
	}

	return true, nil
}

func (s *RelationshipMap) LinkType(this Model, other Model) (LinkKind, error) {
	_, err := PossibleLink(s, this, other)
	if err != nil {
		return "", err
	}

	return (*s)[this.Kind()][other.Kind()], nil
}

func LinkWith(lk LinkKind, this Model, that Model) error {
	switch lk {
	case MulLink:
		this.LinkMul(that)
	case OneLink:
		this.LinkOne(that)
	default:
		return UndefinedLinkKindError
	}

	return nil
}

func UnlinkWith(ln LinkKind, this Model, that Model) error {
	switch ln {
	case MulLink:
		this.UnlinkMul(that)
	case OneLink:
		this.UnlinkOne(that)
	default:
		return UndefinedLinkKindError
	}

	return nil
}

func Compatible(this Model, that Model) bool {
	return this.DBType() == that.DBType()
}

func (s *RelationshipMap) Link(this Model, that Model) error {
	if !Compatible(this, that) {
		return IncompatibleModelsError
	}

	thisLinkType, err := s.LinkType(this, that)

	if err != nil {
		return err
	} else {
		if err = LinkWith(thisLinkType, this, that); err != nil {
			return err
		}
	}

	thatLinkType, err := s.LinkType(that, this)

	if err != nil {
		return err
	} else {
		if err = LinkWith(thatLinkType, that, this); err != nil {
			return err
		}
	}

	return nil
}

func (s *RelationshipMap) Unlink(this Model, that Model) error {
	if !Compatible(this, that) {
		return IncompatibleModelsError
	}

	thisLinkType, err := s.LinkType(this, that)

	if err != nil {
		return err
	} else {
		if err = UnlinkWith(thisLinkType, this, that); err != nil {
			return err
		}
	}

	thatLinkType, err := s.LinkType(that, this)

	if err != nil {
		return err
	} else {
		if err = UnlinkWith(thatLinkType, that, this); err != nil {
			return err
		}
	}

	return nil
}
