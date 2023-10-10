package scenes

import (
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/canvas"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/storage"
    "fyne.io/fyne/v2/widget"
)

type MenuScene struct {
    window fyne.Window
}

func NewMenuScene(window fyne.Window) *MenuScene {
    scene := &MenuScene{window: window}
    scene.Render()
    return scene
}

func (s *MenuScene) Render() {
    background := canvas.NewImageFromURI(storage.NewFileURI("./assets/Fondo.jpg"))
    button := widget.NewButton("Start Game", s.StartGame)

    s.window.SetContent(container.NewWithoutLayout(
        background,
        button,
    ))

    background.Resize(fyne.NewSize(800, 600))
    background.Move(fyne.NewPos(0, 0))

    button.Resize(fyne.NewSize(150, 30))
    button.Move(fyne.NewPos(325, 10))
}

func (s *MenuScene) StartGame() {
    NewGameScene(s.window)
}
