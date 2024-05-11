package systems

type SystemBase interface {
	Update()
	Init()
}

type System struct {
	Entities     []int64
	Requirements []string
}

func (s *System) FitsRequirements(entityComponents []string) bool {
	for _, requirement := range s.Requirements {
		contains := false
		for _, component := range entityComponents {
			if requirement == component {
				contains = true
				break
			}
		}
		if !contains {
			return false
		}
	}
	return true
}

func (s *System) HasEntity(entityID int64) bool {
	for _, entity := range s.Entities {
		if entity == entityID {
			return true
		}
	}
	return false
}

func (s *System) AddEntity(entityID int64) {
	if !s.HasEntity(entityID) {
		s.Entities = append(s.Entities, entityID)
	}
}

func (s *System) RemoveEntity(entityID int64) {
	for i, entity := range s.Entities {
		if entity == entityID {
			s.Entities[i], s.Entities[len(s.Entities)-1] = s.Entities[len(s.Entities)-1], s.Entities[i]
			s.Entities = s.Entities[:len(s.Entities)-1]
		}
	}
}

type MovementSystem struct {
	System
}

func (s *MovementSystem) Update() {

}

func (s *MovementSystem) Init() {

}

func NewMovementSystem() *MovementSystem {
	return &MovementSystem{}
}
