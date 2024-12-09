package engine

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/hexley21/star-gopher/internal/object"
	"github.com/hexley21/star-gopher/internal/renderer"
)

type GameEngine struct {
	renderer      *renderer.Renderer
	gameObjects   []object.GameObject
	player        *object.Player
	random        *rand.Rand
	width, height int
	isRunning     bool
}

func NewGameEngine(width, height int) *GameEngine {
	engine := &GameEngine{
		renderer:  renderer.NewRenderer(width, height),
		random:    rand.New(rand.NewSource(time.Now().UnixNano())),
		width:     width,
		height:    height,
		isRunning: true,
	}
	engine.InitializeGame()
	return engine
}

func (ge *GameEngine) InitializeGame() {
	ge.player = object.NewPlayer(ge.width/2, ge.height-1, ge.width, 1, ge.AddGameObject)
	ge.AddGameObject(ge.player)
}

func (ge *GameEngine) AddGameObject(gameObject object.GameObject) {
	ge.gameObjects = append(ge.gameObjects, gameObject)
}

func (ge *GameEngine) Run() {
	for ge.isRunning {

		ge.SpawnEnemies()
		ge.Update()
		ge.Render()

		time.Sleep(80 * time.Millisecond)

		if !ge.player.IsActive() {
			renderer.ClearScreen()
			fmt.Println("Game Over!")
			time.Sleep(2 * time.Second)
			ge.isRunning = false
		}
	}
}

func (ge *GameEngine) SpawnEnemies() {
	if ge.random.Intn(10) < 2 {
		ge.AddGameObject(object.NewEnemy(ge.random.Intn(ge.width), 0, ge.height, 0))
	}
}

func (ge *GameEngine) Update() {
	for _, obj := range ge.gameObjects {
		obj.Update()
	}

	ge.CheckCollisions()

	activeObjects := []object.GameObject{}
	for _, obj := range ge.gameObjects {
		if obj.IsActive() {
			activeObjects = append(activeObjects, obj)
		}
	}
	ge.gameObjects = activeObjects
}

func (ge *GameEngine) CheckCollisions() {
	for i := 0; i < len(ge.gameObjects); i++ {
		objA := ge.gameObjects[i]
		if !objA.IsActive() {
			continue
		}

		for j := i + 1; j < len(ge.gameObjects); j++ {
			objB := ge.gameObjects[j]
			if !objB.IsActive() {
				continue
			}

			if objA.CheckCollision(objB) {
				objA.OnCollision(objB)
				objB.OnCollision(objA)
			}
		}
	}
}

func (ge *GameEngine) Render() {
	ge.renderer.Render(ge.gameObjects)
}
