package object

type Enemy struct {
	BaseGameObject
	maxY int
}

func NewEnemy(x, y, maxY, colliderSize int) *Enemy {
	return &Enemy{
		BaseGameObject: BaseGameObject{
			X:            x,
			Y:            y,
			Appearance:   '👽',
			active:       true,
			ColliderSize: colliderSize,
		},
		maxY: maxY,
	}
}

func (e *Enemy) Update() {
	e.Y++
	if e.Y >= e.maxY {
		e.SetActive(false)
	}
}

func (e *Enemy) OnCollision(other GameObject) {
	if _, ok := other.(*Bullet); !ok {
		return
	}
	if _, ok := other.(*Player); !ok {
		return
	}
	e.SetActive(false)
	other.SetActive(false)
}
