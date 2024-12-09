package object

type GameObject interface {
	Update()
	CheckCollision(other GameObject) bool
	OnCollision(other GameObject)
	IsActive() bool
	SetActive(active bool)
	GetX() int
	GetY() int
	GetAppearance() rune
	GetColliderSize() int
}

type BaseGameObject struct {
	X            int
	Y            int
	Appearance   rune
	active       bool
	ColliderSize int
}

func (obj *BaseGameObject) IsActive() bool {
	return obj.active
}

func (obj *BaseGameObject) SetActive(active bool) {
	obj.active = active
}

func (obj *BaseGameObject) GetX() int {
	return obj.X
}

func (obj *BaseGameObject) GetY() int {
	return obj.Y
}

func (obj *BaseGameObject) GetAppearance() rune {
	return obj.Appearance
}

func (obj *BaseGameObject) GetColliderSize() int {
	return obj.ColliderSize
}

func (obj *BaseGameObject) CheckCollision(other GameObject) bool {
	if !obj.IsActive() || !other.IsActive() {
		return false
	}

	deltaX := abs(obj.X - other.GetX())
	deltaY := abs(obj.Y - other.GetY())
	totalColliderSize := obj.ColliderSize + other.GetColliderSize()

	return (deltaX <= totalColliderSize) && (deltaY <= totalColliderSize)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
