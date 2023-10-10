package scenes

import (
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/canvas"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/storage"
    "fyne.io/fyne/v2/widget"
)

type GameOverScene struct {
    window fyne.Window
}

func NewGameOverScene(window fyne.Window) *GameOverScene {
    scene := &GameOverScene{window: window}
    scene.Render()
    return scene
}

func (s *GameOverScene) Render() {
    btnGoMenu := widget.NewButton("Go Menu", s.goMenu)
    btnRestart := widget.NewButton("Restart Game", s.restart)

    gameOverImage := canvas.NewImageFromURI(storage.NewFileURI("./assets/Fin.jpg"))
    backgroundImage := canvas.NewImageFromURI(storage.NewFileURI("./assets/WarFondo.png"))

    btnGoMenu.Resize(fyne.NewSize(150, 30))
    btnRestart.Resize(fyne.NewSize(150, 30))

    gameOverImage.Resize(fyne.NewSize(200, 150))
    backgroundImage.Resize(fyne.NewSize(800, 600))

    btnGoMenu.Move(fyne.NewPos(325, 50))
    btnRestart.Move(fyne.NewPos(325, 10))

    gameOverImage.Move(fyne.NewPos(300, 225))
    backgroundImage.Move(fyne.NewPos(0, 0))

    s.window.SetContent(container.NewWithoutLayout(backgroundImage, gameOverImage, btnGoMenu, btnRestart))
}

func (s *GameOverScene) goMenu() {
    NewMenuScene(s.window)
}

func (s *GameOverScene) restart() {
    NewGameScene(s.window)
}
