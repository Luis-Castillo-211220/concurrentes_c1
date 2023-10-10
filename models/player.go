package models

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type Player struct {
	posX, posY, direction float32
	running               bool
	pel                   *canvas.Image
	image                 *canvas.Image // Campo para almacenar la imagen del jugador
}

func NewPlayer(posx float32, posy float32, img *canvas.Image) *Player {
	player := &Player{
		posX:      posx,
		posY:      posy,
		running:   true,
		direction: 0,
	}

	// Intenta cargar la imagen desde el archivo
	player.pel = canvas.NewImageFromFile("./assets/Soldier.png")
	if player.pel == nil {
		fmt.Println("Error: failed to load image from file")

		// Si la carga de la imagen falla, crea una imagen en blanco
		emptyImageResource := fyne.NewStaticResource("empty.png", []byte{})
		player.pel = canvas.NewImageFromResource(emptyImageResource)
	}

	// Asigna image a pel
	player.image = player.pel

	// Ajusta el tamaño de la imagen
	player.pel.Resize(fyne.NewSize(75, 75))

	// Mueve la imagen a la posición inicial del jugador
	player.pel.Move(fyne.NewPos(player.posX, player.posY))

	return player
}

func (p *Player) GoRight() {
	p.direction = 1
}

func (p *Player) GoLeft() {
	p.direction = -1
}

func (p *Player) Run() {
	for p.running {
		var incX float32 = 50
		incX *= p.direction

		if p.posX < 50 {
			p.posX = 50
		} else if p.posX > 650 {
			p.posX = 650
		}

		p.posX += incX
		p.pel.Move(fyne.NewPos(p.posX, p.posY))
		time.Sleep(100 * time.Millisecond)
	}
}

func (p *Player) SetRunning(pause bool) {
	p.running = pause
}

func (p *Player) GetRunning() bool {
	return p.running
}

func (p *Player) GetPosY() float32 {
	return p.posY
}

func (p *Player) GetPosX() float32 {
	return p.posX
}

func (p *Player) GetImage() *canvas.Image {
	return p.image
}
