package main

import (
	"fmt"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type SceneSwapFunc func(*Game)

type Button struct {
	//CommonDraw
	image *ebiten.Image
	rect  Rectangle
	cmd   SceneSwapFunc
}

type Signal struct {
	image      *ebiten.Image
	rect       Rectangle
	is_dragged bool
	signalType string
}

type Receptor struct {
	image              *ebiten.Image
	rect               Rectangle
	is_touching_signal bool
	receptorType       string
}

type Kinase struct {
	image         *ebiten.Image
	rect          Rectangle
	is_moving     bool
	is_clicked_on bool
	delta         int
	kinaseType    string
}

type TFA struct {
	image     *ebiten.Image
	rect      Rectangle
	is_active bool
}

type Transcript struct {
	image *ebiten.Image
	rect  Rectangle
	codon string
}

type Template struct {
	image       *ebiten.Image
	rect        Rectangle
	codon       string
	fragment    int
	is_complete bool
}

type Nucleobase struct {
	image    *ebiten.Image
	rect     Rectangle
	baseType string
}

type CodonChoice struct {
	image *ebiten.Image
	rect  Rectangle
	bases string
	// codonType string // Correct vs Incorrect
}

type Ribosome struct {
	image *ebiten.Image
	rect  Rectangle
}

/*
type Drawable interface {
	draw()
}

type CommonDraw struct{
	Drawable
	rect   Rectangle
	image  *ebiten.Image
}

func (cd CommonDraw) draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(cd.rect.pos.x), float64(cd.rect.pos.y))
	screen.DrawImage(cd.image, op)
} */

func newSignal(path string, rect Rectangle) Signal {
	var sig_image, _, err = ebitenutil.NewImageFromFile(loadFile(path))

	if err != nil {
		fmt.Println("Error parsing date:", err)
	}
	return Signal{
		image:      sig_image,
		rect:       rect,
		is_dragged: false,
	}
}

func newKinase(path string, rect Rectangle, ktype string) Kinase {
	var kin_image, _, err = ebitenutil.NewImageFromFile(loadFile(path))
	if err != nil {
		fmt.Println("Error parsing date:", err)
	}
	return Kinase{
		image:         kin_image,
		rect:          rect,
		is_moving:     false,
		is_clicked_on: false,
		delta:         3,
		kinaseType:    ktype,
	}
}

func newTFA(path string, rect Rectangle) TFA {
	var tfa_image, _, err = ebitenutil.NewImageFromFile(loadFile(path))
	if err != nil {
		fmt.Println("Error parsing date:", err)
	}
	return TFA{
		image:     tfa_image,
		rect:      rect,
		is_active: false,
	}
}

func newButton(path string, rect Rectangle, cmd SceneSwapFunc) Button {
	var btn_image, _, err = ebitenutil.NewImageFromFile(loadFile(path))

	if err != nil {
		fmt.Println("Error parsing date:", err)
	}
	return Button{
		image: btn_image,
		rect:  rect,
		cmd:   cmd,
	}
}

func (b Button) on_click(g *Game) {
	var x_c, y_c = ebiten.CursorPosition()
	var b_pos = newVector(x_c, y_c)
	if rect_point_collision(b.rect, b_pos) && ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		b.cmd(g)
	}
}

/* 
func (b Button) draw(screen *ebiten.Image) {
	b.CommonDraw.draw(screen)
} */


func (b Button) draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(b.rect.pos.x), float64(b.rect.pos.y))
	screen.DrawImage(b.image, op)
}


func (s *Signal) on_click(g *Game) {
	var x_c, y_c = ebiten.CursorPosition()
	var b_pos = newVector(x_c, y_c)

	if !s.is_dragged {
		if rect_point_collision(s.rect, b_pos) && ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			s.is_dragged = true
		}
	}

	if s.is_dragged {
		s.rect.pos = newVector(b_pos.x-s.rect.width/2, b_pos.y-s.rect.height/2)
	}
}

func (s Signal) draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(s.rect.pos.x), float64(s.rect.pos.y))
	screen.DrawImage(s.image, op)
}

func newReceptor(path string, rect Rectangle, rtype string) Receptor {
	var rec_image, _, err = ebitenutil.NewImageFromFile(loadFile(path))

	if err != nil {
		fmt.Println("Error parsing date:", err)
	}
	return Receptor{
		image:              rec_image,
		rect:               rect,
		is_touching_signal: false,
		receptorType:       rtype,
	}
}

func (r Receptor) draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(r.rect.pos.x), float64(r.rect.pos.y))
	screen.DrawImage(r.image, op)
}

func (r *Receptor) update() {
	if aabb_collision(signal.rect, r.rect) {
		r.is_touching_signal = true
	} else {
		r.is_touching_signal = false
	}
}

func (k *Kinase) update(rect Rectangle) {
	var x_c, y_c = ebiten.CursorPosition()
	var b_pos = newVector(x_c, y_c)
	if !k.is_clicked_on && k.is_moving {
		if k.rect.pos.y <= 425 && k.kinaseType == "tk2" {
			k.descend()
		} else {
			k.rect.pos.x += k.delta
		}
	}

	if rect_point_collision(k.rect, b_pos) &&
		ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) &&
		aabb_collision(k.rect, rect) {
		k.is_clicked_on = true
	}

	if k.rect.pos.x+k.rect.width >= screenWidth {
		k.delta = -3
	} else if k.rect.pos.x <= 0 {
		k.delta = 3
	}

}

func (k Kinase) draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(k.rect.pos.x), float64(k.rect.pos.y))
	screen.DrawImage(k.image, op)
}

func (k *Kinase) activate() {
	if k.kinaseType == "tk1" {k.animate("act_TK1.png")}
	if k.kinaseType == "tk2" {k.animate("act_TK2.png")}
	k.is_moving = true
}

func (k *Kinase) descend() {
	k.rect.pos.y += 2
}

func (k *Kinase) animate(newImage string) {
	var kin_image, _, err = ebitenutil.NewImageFromFile(loadFile(newImage))
	if err != nil {
		fmt.Println("Error parsing date:", err)
	}
	k.image = kin_image
}

func (t *TFA) activate() {
	t.animate("act_TFA.png")
	tfa.is_active = true
}

func (t *TFA) update() {
	if t.is_active && t.rect.pos.y <= 750 {
		t.rect.pos.y += 2
	}
}

func (t *TFA) animate(newImage string) {
	var tfa_image, _, err = ebitenutil.NewImageFromFile(loadFile(newImage))
	if err != nil {
		fmt.Println("Error parsing date:", err)
	}
	t.image = tfa_image
}

func (t TFA) draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(t.rect.pos.x), float64(t.rect.pos.y))
	screen.DrawImage(t.image, op)
}

func newTranscript(path string, rect Rectangle, codon string) Transcript {
	var transcript_image, _, err = ebitenutil.NewImageFromFile(loadFile(path))
	if err != nil {
		fmt.Println("Error parsing date:", err)
	}
	return Transcript{
		image: transcript_image,
		rect:  rect,
		codon: codon,
	}
}

func (transcript Transcript) draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(transcript.rect.pos.x), float64(transcript.rect.pos.y))
	screen.DrawImage(transcript.image, op)
}

func newTemplate(path string, rect Rectangle, codon string, fragment int) Template {
	var template_image, _, err = ebitenutil.NewImageFromFile(loadFile(path))
	if err != nil {
		fmt.Println("Error parsing date:", err)
	}
	return Template{
		image:       template_image,
		rect:        rect,
		codon:       codon,
		fragment:    fragment,
		is_complete: false,
	}
}

func (template Template) draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(template.rect.pos.x), float64(template.rect.pos.y))
	screen.DrawImage(template.image, op)
}

func nextDNACodon(g *Game) {
	if currentFrag < 4 {
		currentFrag++
		dna[currentFrag].is_complete = false
		reset = true
	} else {
		NucleusToCyto2(g)
		reset = false
	}
}

func nextMRNACodon(g *Game) {
	if mrna_ptr < 4 {
		mrna_ptr++
		mrna[mrna_ptr].is_complete = false
		reset = true
		ribosome.update_movement()
	} else {
		os.Exit(3)
		reset = false
	}
}

func newRibosome(path string, rect Rectangle) Ribosome {
	var ribo_image, _, err = ebitenutil.NewImageFromFile(loadFile(path))
	if err != nil {
		fmt.Println("Error parsing date:", err)
	}
	return Ribosome{
		image: ribo_image,
		rect:  rect,
	}
}

func newCodonChoice(path string, rect Rectangle, bases string) CodonChoice {
	var cdn_image, _, err = ebitenutil.NewImageFromFile(loadFile(path))

	if err != nil {
		fmt.Println("Error parsing date:", err)
	}
	return CodonChoice{
		image: cdn_image,
		rect:  rect,
		bases: bases,
	}
}

func (c CodonChoice) update1(g *Game, dnaFrag string) bool {
	var x_c, y_c = ebiten.CursorPosition()
	var b_pos = newVector(x_c, y_c)
	if rect_point_collision(c.rect, b_pos) && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		if c.bases == transcribe(dnaFrag) {
			return true
		}
	}
	return false
}

func (c CodonChoice) update2(g *Game, mrnaFrag string) bool {
	var x_c, y_c = ebiten.CursorPosition()
	var b_pos = newVector(x_c, y_c)
	if rect_point_collision(c.rect, b_pos) && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		if c.bases == translate(mrnaFrag) {
			return true
		}
	}
	return false
}

func (c CodonChoice) draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(c.rect.pos.x), float64(c.rect.pos.y))
	screen.DrawImage(c.image, op)
}

func (ribo *Ribosome) update_movement() {
	ribo.rect.pos.x += screenWidth / 6
}

func (ribo Ribosome) draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(ribo.rect.pos.x), float64(ribo.rect.pos.y))
	screen.DrawImage(ribo.image, op)
}

func newNucelobase(path string, rect Rectangle, btype string) Nucleobase {
	var base_image, _, err = ebitenutil.NewImageFromFile((loadFile(path)))
	if err != nil {
		fmt.Println("Error parsing date:", err)
	}
	return Nucleobase{
		image:    base_image,
		rect:     rect,
		baseType: btype,
	}
}
