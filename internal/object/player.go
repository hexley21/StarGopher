package object

import (
	"sync"

	"github.com/eiannone/keyboard"
)

type Player struct {
	BaseGameObject
	maxX     int
	key      keyboard.Key
	mu       sync.Mutex
	addGameObject func(obj GameObject)
}

func NewPlayer(x, y, maxX, colliderSize int, addGameObject func(obj GameObject)) *Player {
	p := &Player{
		BaseGameObject: BaseGameObject{
			X:            x,
			Y:            y,
			Appearance:   '🦫',
			active:       true,
			ColliderSize: colliderSize,
		},
        addGameObject: addGameObject,
		maxX: maxX,
	}

	go p.listenForKeyPress()
	return p
}

func (p *Player) listenForKeyPress() {
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer keyboard.Close()

	for p.IsActive() {
		_, key, err := keyboard.GetKey()
		if err != nil {
			continue
		}

		p.mu.Lock()
		p.key = key
		p.mu.Unlock()

		if !p.IsActive() {
			break
		}
	}
}

func (p *Player) Update() {
	p.mu.Lock()
	key := p.key
	p.key = 0
	p.mu.Unlock()

	switch key {
	case keyboard.KeyArrowLeft:
		if p.X > 0 {
			p.X--
		}
	case keyboard.KeyArrowRight:
		if p.X < p.maxX-1 {
			p.X++
		}
	case keyboard.KeySpace:
        p.addGameObject(NewBullet(p.X, p.Y-1, 1))
	}
}

func (p *Player) OnCollision(other GameObject) {
	if _, ok := other.(*Enemy); ok {
		p.SetActive(false)
		other.SetActive(false)
	}
}
