package object


type Bullet struct {
    BaseGameObject
}

func NewBullet(x, y, colliderSize int) *Bullet {
    return &Bullet{
        BaseGameObject: BaseGameObject{
            X:           x,
            Y:           y,
            Appearance:  '🔺',
            active:      true,
            ColliderSize: colliderSize,
        },
    }
}

func (b *Bullet) Update() {
    b.Y--
    if b.Y < 0 {
        b.SetActive(false)
    }
}

func (b *Bullet) OnCollision(other GameObject) {
    if _, ok := other.(*Enemy); ok {
        b.SetActive(false)
        other.SetActive(false)
    }
}