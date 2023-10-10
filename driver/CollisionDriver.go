package driver

import (
	"prueba/models"
)

//maneja la logica de colision entre el jugador y las ventanas.
type CollisionDriver struct {
	gameOver bool
	player   *models.Player
	windows  *models.Windows
}

//crea una nueva instancia de CollisionDriver con el jugador y las ventanas especificados.
func NewCollisionDriver(player *models.Player, windows *models.Windows) *CollisionDriver {
	return &CollisionDriver{
		gameOver: false,
		player:   player,
		windows:  windows,
	}
}

// Run inicia la logica de colision en un bucle.
func (c *CollisionDriver) Run() {
	for !c.gameOver {
		if c.windows.GetPosY() >= 400 {
			playerX := c.player.GetPosX()
			windowX := c.windows.GetPosX()

			if windowX >= playerX-50 && windowX <= playerX+50 {
				c.windows.SetRunning(false)
				c.player.SetRunning(false)
				c.gameOver = true
			}
		}
	}
}

//devuelve true si el juego ha terminado debido a una colision.
func (c *CollisionDriver) GetGameOver() bool {
	return c.gameOver
}
