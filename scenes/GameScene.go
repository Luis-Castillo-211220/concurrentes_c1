package scenes

import (
	"prueba/driver"
	"prueba/models"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

type GameScene struct {
	window fyne.Window
}

var p *models.Player
var w *models.Windows
var c *driver.CollisionDriver

func NewGameScene(window fyne.Window) *GameScene {
	scene := &GameScene{window: window}
	   // Crea la imagen del jugador
	   playerPeel := createPeel("./assets/Soldier.png", 50, 50, 350, 450)

	     // Inicia el jugador con la imagen creada
		 p = models.NewPlayer(350, 450, playerPeel)
	scene.Render()
	scene.StartGame()
	return scene
}

func (s *GameScene) Render() {
	backgroundImage := canvas.NewImageFromURI(storage.NewFileURI("./assets/WarFondo.png"))
	backgroundImage.Resize(fyne.NewSize(800, 600))
	backgroundImage.Move(fyne.NewPos(0, 0))
	// Models Render
	

	playerPeel := p.GetImage() //imagen del jugador desde el modelo

	windowsPeel := createPeel("./assets/Bomba.png", 100, 100, 100, 50)
	w = models.NewWindows(350, 600, windowsPeel, p)

	// CollisionDriver
	c = driver.NewCollisionDriver(p, w)

	// Buttons Render
	btnLeft := widget.NewButton("<", p.GoLeft)
	btnLeft.Resize(fyne.NewSize(50, 50))
	btnLeft.Move(fyne.NewPos(350, 550))

	btnRight := widget.NewButton(">", p.GoRight)
	btnRight.Resize(fyne.NewSize(50, 50))
	btnRight.Move(fyne.NewPos(400, 550))

	s.window.SetContent(container.NewWithoutLayout(backgroundImage, playerPeel, windowsPeel, btnLeft, btnRight))
}

func (s *GameScene) StartGame() {
	go p.Run()
	go w.Run()
	go c.Run()
	go s.checkGameOver()
}

func (s *GameScene) StopGame() {
	p.SetRunning(!p.GetRunning())
	w.SetRunning(!w.GetRunning())
}

func (s *GameScene) checkGameOver() {
	running := true
	for running {
		if c.GetGameOver() {
			running = false
			time.Sleep(1000 * time.Millisecond)
			NewGameOverScene(s.window)
		}
	}
}

func createPeel(fileURI string, sizeX, sizeY, posX, posY float32) *canvas.Image {
    //  se encarga de cargar una imagen a partir de la ubicación de un archivo en el sistema de archivos
    image := canvas.NewImageFromURI(storage.NewFileURI(fileURI))

    // Ajusta el tamaño de la imagen
    imageSize := fyne.NewSize(sizeX, sizeY)
    image.Resize(imageSize)

    // Mueve la imagen a la posición especificada
    imagePosition := fyne.NewPos(posX, posY)
    image.Move(imagePosition)

    return image
}