package main

import (
	"fmt"
	"image/color"

	//"reflect"

	//"math"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type ButtonFunc func(*Game)

type GUI interface {
	draw(screen *ebiten.Image)
	update(params ...interface{})
	//getStructType(GUI) reflect.Type
}

//func getStructType(g GUI) reflect.Type {
//	return reflect.TypeOf(g)
//}

type Button struct {
	//CommonDraw
	image *ebiten.Image
	rect  Rectangle
	cmd   ButtonFunc
}

type VolButton struct {
	Button
	player audio.Player
	status string
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
	tfaType   string
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

type RNAPolymerase struct {
	image *ebiten.Image
	rect  Rectangle
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

type Parallax struct {
	image *ebiten.Image
	rect  Rectangle
	layer float64
}

type InfoPage struct {
	btn_image *ebiten.Image
	pg_image  *ebiten.Image
	rect      Rectangle
	status    string
	// Functions: when screen switches, is drawn in btn status. When mouseButtonJustPressed + btn status,
	// changes to pg status. When mouseButtonJustPressed + pg status, changes to button status.
	// update function sets image to status + "_image". Pg image should say "Click to exit."
	// if status = btn, if status = pg
}

type StillImage struct {
	image *ebiten.Image
	rect  Rectangle
}

func newStillImage(path string, rect Rectangle) StillImage {
	var still_image, _, err = ebitenutil.NewImageFromFile(loadFile(path))

	if err != nil {
		fmt.Println("Error parsing date:", err)
	}
	return StillImage{
		image: still_image,
		rect:  rect,
	}
}

func (s StillImage) draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(s.rect.pos.x), float64(s.rect.pos.y))
	scaleW := 0.5 * float64(screenWidth) / 1250
	scaleH := 0.5 * float64(screenHeight) / 750
	op.GeoM.Scale(scaleW, scaleH)
	screen.DrawImage(s.image, op)
}

func (s StillImage) update(params ...interface{}) {}

func newInfoPage(path1 string, path2 string, rect Rectangle, stat string) InfoPage {
	var btn_img, _, err1 = ebitenutil.NewImageFromFile(loadFile(path1))
	var pg_img, _, err2 = ebitenutil.NewImageFromFile(loadFile(path2))

	if err1 != nil {
		fmt.Println("Error parsing date:", err)
	}
	if err2 != nil {
		fmt.Println("Error parsing date:", err)
	}

	return InfoPage{
		btn_image: btn_img,
		pg_image:  pg_img,
		rect:      rect,
		status:    stat,
	}
}

func (i *InfoPage) update(params ...interface{}) {
	var x_c, y_c = ebiten.CursorPosition()
	var b_pos = newVector(x_c, y_c)
	if rect_point_collision(i.rect, b_pos) && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		if i.status == "btn" {
			i.status = "pg"
			i.rect = newRect(0, 0, screenWidth, screenHeight)
		} else {
			i.status = "btn"
			i.rect = newRect(850, 0, 165, 165)
		}
	}
}

func (i InfoPage) draw(screen *ebiten.Image, g *Game) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(i.rect.pos.x), float64(i.rect.pos.y))
	if i.status == "btn" {
		screen.DrawImage(i.btn_image, op)
	}
	if i.status == "pg" {
		screen.DrawImage(i.pg_image, op)
		Purple := color.RGBA{50, 0, 50, 250}
		g.defaultFont.drawFont(screen, info, 300, 200, color.RGBA(Purple))
	}
}

func newParallax(path string, rect Rectangle, layer float64) Parallax {
	var par_image, _, err = ebitenutil.NewImageFromFile(loadFile(path))

	if err != nil {
		fmt.Println("Error parsing date:", err)
	}
	return Parallax{
		image: par_image,
		rect:  rect,
		layer: layer,
	}
}

func (p *Parallax) update(params ...interface{}) {
	var x_c, y_c = ebiten.CursorPosition()
	var l = int(p.layer)
	switch scene {
	case "Main Menu":
		p.rect.pos.x = -5 * (x_c + 100) / (7 * l)
		p.rect.pos.y = -3 * (y_c + 100) / (5 * l)
	case "Signal Reception":
		p.rect.pos.x = -6 * (x_c + 100) / (7 * l)
		p.rect.pos.y = -2 * (y_c + 100) / (3 * l)
	case "Signal Transduction":
		p.rect.pos.x = -5 * (x_c + 80) / (7 * l)
		p.rect.pos.y = -3 * (y_c + 100) / (5 * l)
	}
	//p.rect.pos.x = (x_c - 625) / (2*l);
	//p.rect.pos.y = (y_c - 375) / (2*l);
}

func (p Parallax) draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(p.rect.pos.x), float64(p.rect.pos.y))
	scaleW := (p.layer + 0.5) / (p.layer * 2) * float64(screenWidth) / 1250
	scaleH := (p.layer + 0.5) / (p.layer * 2) * float64(screenHeight) / 750
	op.GeoM.Scale(scaleW, scaleH)
	//if scene == "Signal Reception" {
	//var l = float32(p.layer)
	//op.ColorScale.ScaleAlpha(float32(math.Pow(p.layer/4, 1.2)))
	//}
	screen.DrawImage(p.image, op)
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

func newButton(path string, rect Rectangle, cmd ButtonFunc) Button {
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

// NEED TO FIX VOL BUTTON CHANGE IMAGE
func (b Button) update(params ... interface{}) {
	if len(params) > 0 {
        g, ok := params[0].(*Game)
        if !ok {
            return
        }
		var x_c, y_c = ebiten.CursorPosition()
		var b_pos = newVector(x_c, y_c)
		if rect_point_collision(b.rect, b_pos) && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			b.cmd(g)
		}
	}
}

/*
func (b Button) draw(screen *ebiten.Image) {
	b.CommonDraw.draw(screen)
} */

func (b Button) draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(b.rect.pos.x), float64(b.rect.pos.y))
	scaleW := float64(screenWidth) / 1250
	scaleH := float64(screenHeight) / 750
	op.GeoM.Scale(scaleW, scaleH)
	b.rect.width *= int(scaleW)
	b.rect.height *= int(scaleH)
	screen.DrawImage(b.image, op)
}

func newVolButton(path string, rect Rectangle, cmd ButtonFunc, player audio.Player) VolButton {
	btn := newButton(path, rect, cmd)
	return VolButton{
		Button: btn,
		player: player,
		status: "ON",
	}
}

func (v VolButton) update(params ...interface{}) {
	v.Button.update(params...)
	//curr := int64(v.player.Position())
	if v.status == "ON" && !v.player.IsPlaying() {
		v.player.Rewind()
	}

}

func (v VolButton) draw(screen *ebiten.Image) {
	v.image = volButton.image
	v.status = volButton.status
	v.Button.draw(screen)
}

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

func (s *Signal) update(params ...interface{}) {
	var x_c, y_c = ebiten.CursorPosition()
	var b_pos = newVector(x_c, y_c)

	if !s.is_dragged {
		if rect_point_collision(s.rect, b_pos) && ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			s.is_dragged = true
		}
	} else if s.is_dragged {
		if s.rect.pos.y <= receptorA.rect.pos.y && b_pos.y <= receptorA.rect.pos.y {
			s.rect.pos.y = b_pos.y - (s.rect.height / 2)
		} else {
			if s.rect.pos.y > receptorA.rect.pos.y {
				s.rect.pos.y -= 1
			} else if s.rect.pos.y < receptorA.rect.pos.y {
				s.rect.pos.y += 1
			}
		}
		s.rect.pos = newVector(b_pos.x-s.rect.width/2, s.rect.pos.y)

	}
}

func (s Signal) draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	scaleW := (0.5) * float64(screenWidth) / 1250
	scaleH := (0.5) * float64(screenHeight) / 750
	op.GeoM.Scale(scaleW, scaleH)
	s.rect.width *= int(scaleW)
	s.rect.height *= int(scaleH)
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
	scaleW := 0.5 * float64(screenWidth) / 1250
	scaleH := 0.5 * float64(screenHeight) / 750
	op.GeoM.Scale(scaleW, scaleH)
	r.rect.width *= int(scaleW)
	r.rect.height *= int(scaleH)
	op.GeoM.Translate(float64(r.rect.pos.x), float64(r.rect.pos.y))
	screen.DrawImage(r.image, op)
}

func (r *Receptor) update(params ...interface{}) {
	var x_c, y_c = ebiten.CursorPosition()
	switch r.receptorType {
	case "receptorA":
		r.rect.pos.x = ((-5 * (x_c + 100) / (7 * 3)) + 100) * screenWidth / 1250
		r.rect.pos.y = ((-2 * (y_c + 100) / (5 * 3)) + 390) * screenHeight / 750
		//r.rect.pos.x = plasmaMembrane.rect.pos.x + 50
		//r.rect.pos.y = plasmaMembrane.rect.pos.y + 400
	case "receptorB":
		r.rect.pos.x = ((-5 * (x_c + 100) / (7 * 3)) + 500) * screenWidth / 1250
		r.rect.pos.y = ((-2 * (y_c + 100) / (5 * 3)) + 340) * screenHeight / 750
		//r.rect.pos.x = plasmaMembrane.rect.pos.x + 500
		//r.rect.pos.y = plasmaMembrane.rect.pos.y + 400
	case "receptorC":
		r.rect.pos.x = ((-5 * (x_c + 100) / (7 * 3)) + 950) * screenWidth / 1250
		r.rect.pos.y = ((-2 * (y_c + 100) / (5 * 3)) + 340) * screenHeight / 750
		//r.rect.pos.x = plasmaMembrane.rect.pos.x + 950
		//r.rect.pos.y = plasmaMembrane.rect.pos.y + 400
	case "receptorD":
		r.rect.pos.x = ((-5 * (x_c + 100) / (7 * 3)) + 1300) * screenWidth / 1250
		r.rect.pos.y = ((-2 * (y_c + 100) / (5 * 3)) + 390) * screenHeight / 750
		//r.rect.pos.x = plasmaMembrane.rect.pos.x + 1400
		//r.rect.pos.y = plasmaMembrane.rect.pos.y + 400
	}
	if aabb_collision(signal.rect, r.rect) {
		r.is_touching_signal = true
	} else {
		r.is_touching_signal = false
	}
}

func (r *Receptor) animate(newImage string) {
	var rec_image, _, err = ebitenutil.NewImageFromFile(loadFile(newImage))
	if err != nil {
		fmt.Println("Error parsing date:", err)
	}
	r.image = rec_image
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

func (k *Kinase) update(params ... interface{}) {
	if len(params) > 0 {
        rect, ok := params[0].(Rectangle)
        if !ok {
            return
        }
		var x_c, y_c = ebiten.CursorPosition()
		var b_pos = newVector(x_c, y_c)
		if strings.Contains(k.kinaseType, "temp_tk1") {
			if !k.is_moving {
				var x_c, y_c = ebiten.CursorPosition()
				switch k.kinaseType {
				case "temp_tk1A":
					k.rect.pos.x = (-5 * (x_c + 100) / (7 * 3)) + 75
					k.rect.pos.y = (-1 * (y_c + 100) / (3 * 3)) + 625
				case "temp_tk1B":
					k.rect.pos.x = (-5 * (x_c + 100) / (7 * 3)) + 475
					k.rect.pos.y = (-(y_c + 100) / (3 * 3)) + 575
				case "temp_tk1C":
					k.rect.pos.x = (-5 * (x_c + 100) / (7 * 3)) + 875
					k.rect.pos.y = (-(y_c + 100) / (3 * 3)) + 575
				case "temp_tk1D":
					k.rect.pos.x = (-5 * (x_c + 100) / (7 * 3)) + 1275
					k.rect.pos.y = (-1 * (y_c + 100) / (3 * 3)) + 625
				}
			} else if k.is_moving {
				if k.rect.pos.y <= screenHeight {
					k.descend()
				}
			}
		} else if !k.is_clicked_on && k.is_moving {
			if k.rect.pos.y <= 400*(screenHeight/750) && k.kinaseType == "tk2" {
				k.descend()
			} else if k.rect.pos.y <= 50*(screenHeight/750) && k.kinaseType == "tk1" {
				k.descend()
			} else {
				if ebiten.IsFullscreen() {k.rect.pos.x += k.delta * widthRatio
				} else {k.rect.pos.x += k.delta}
			}
		}
		if rect_point_collision(k.rect, b_pos) && ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && aabb_collision(k.rect, rect) {
			k.is_clicked_on = true
		}
	
		if k.rect.pos.x+k.rect.width >= screenWidth {
			k.delta = -3
		} else if k.rect.pos.x <= 0 {
			k.delta = 3
		}
	}	
}

func (k Kinase) draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	scaleW := (0.5) * float64(screenWidth) / 1250
	scaleH := (0.5) * float64(screenHeight) / 750
	op.GeoM.Scale(scaleW, scaleH)
	k.rect.width *= int(scaleW)
	k.rect.height *= int(scaleH)
	op.GeoM.Translate(float64(k.rect.pos.x), float64(k.rect.pos.y))
	screen.DrawImage(k.image, op)
}

func (k *Kinase) activate() {
	if k.kinaseType == "tk1" {
		k.animate("act_TK1.png")
	}
	if k.kinaseType == "tk2" {
		k.rect.pos.y -= 3 * (screenHeight / 750)
		k.animate("act_TK2.png")
	}
	if strings.Contains(k.kinaseType, "temp_tk1") && !k.is_moving {
		k.rect.pos.y -= 3 * (screenHeight / 750)
		k.animate("act_TK1.png")
	}
	k.is_moving = true
}

func (k *Kinase) descend() {
	if ebiten.IsFullscreen() {k.rect.pos.y += 2*heightRatio
	} else {k.rect.pos.y += 2}
}

func (k *Kinase) animate(newImage string) {
	var kin_image, _, err = ebitenutil.NewImageFromFile(loadFile(newImage))
	if err != nil {
		fmt.Println("Error parsing date:", err)
	}
	k.image = kin_image
}

func (t *TFA) activate() {
	if t.tfaType == "tfa1" {
		t.rect.pos.y -= 3 * (screenHeight / 750)
	}
	t.animate("act_TFA.png")
	t.is_active = true
}

func newTFA(path string, rect Rectangle, tfaType string) TFA {
	var tfa_image, _, err = ebitenutil.NewImageFromFile(loadFile(path))
	if err != nil {
		fmt.Println("Error parsing date:", err)
	}
	return TFA{
		image:     tfa_image,
		rect:      rect,
		is_active: false,
		tfaType:   tfaType,
	}
}

func (t *TFA) update(params ...interface{}) {
	if t.is_active {
		if t.rect.pos.y <= screenHeight && t.tfaType == "tfa1" {
			t.rect.pos.y += 2 * (screenHeight / 750)
		}
		if t.rect.pos.y <= 400*screenHeight/750 && t.tfaType == "tfa2" {
			t.rect.pos.y += 2 * (screenHeight / 750)
			t.rect.pos.x -= 1 * (screenWidth / 1250)
		}
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
	scaleW := (0.5) * float64(screenWidth) / 1250
	scaleH := (0.5) * float64(screenHeight) / 750
	if t.tfaType == "tfa2" {
		scaleW *= 1.25
		scaleH *= 1.25
	}
	t.rect.width *= int(scaleW)
	t.rect.height *= int(scaleH)
	op.GeoM.Scale(scaleW, scaleH)
	op.GeoM.Translate(float64(t.rect.pos.x), float64(t.rect.pos.y))
	screen.DrawImage(t.image, op)
}

func newRNAPolymerase(path string, rect Rectangle) RNAPolymerase {
	var RNAPolym_image, _, err = ebitenutil.NewImageFromFile(loadFile(path))
	if err != nil {
		fmt.Println("Error parsing date:", err)
	}
	return RNAPolymerase{
		image: RNAPolym_image,
		rect:  rect,
	}
}

func (r *RNAPolymerase) update(params ... interface{}) {
	if len(params) > 0 {
        tfaPosY, ok := params[0].(int)
        if !ok {
            return
        }
		if tfaPosY >= 300 {
			if r.rect.pos.x <= 25 {
				r.rect.pos.y += 1 * (screenHeight / 750)
				r.rect.pos.x += 2 * (screenWidth / 1250)
			}
		}
	}
}

func (r RNAPolymerase) draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(r.rect.pos.x), float64(r.rect.pos.y))
	scaleW := float64(screenWidth) / 1250
	scaleH := float64(screenHeight) / 750
	op.GeoM.Scale(scaleW, scaleH)
	r.rect.width *= int(scaleW)
	r.rect.height *= int(scaleH)
	screen.DrawImage(r.image, op)
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

func (transcr Transcript) draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(transcr.rect.pos.x), float64(transcr.rect.pos.y))
	scaleW := float64(screenWidth) / 1250
	scaleH := float64(screenHeight) / 750
	op.GeoM.Scale(scaleW, scaleH)
	transcr.rect.width *= int(scaleW)
	transcr.rect.height *= int(scaleH)
	screen.DrawImage(transcr.image, op)
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

func (temp Template) draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(temp.rect.pos.x), float64(temp.rect.pos.y))
	scaleW := float64(screenWidth) / 1250
	scaleH := float64(screenHeight) / 750
	op.GeoM.Scale(scaleW, scaleH)
	temp.rect.width *= int(scaleW)
	temp.rect.height *= int(scaleH)
	screen.DrawImage(temp.image, op)
}

func nextDNACodon(g *Game) {
	if currentFrag < 4 {
		currentFrag++
		dna[currentFrag].is_complete = false
		reset = true
	} else {
		ToCyto2(g)
		reset = false
	}
}

func nextMRNACodon(g *Game) {
	if mrna_ptr < 4 {
		mrna_ptr++
		mrna[mrna_ptr].is_complete = false
		reset = true
	} else {
		ToMenu(g)
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

func (c CodonChoice) update1(dnaFrag string) bool {
	var x_c, y_c = ebiten.CursorPosition()
	var b_pos = newVector(x_c, y_c)
	if rect_point_collision(c.rect, b_pos) && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		if c.bases == transcribe(dnaFrag) {
			return true
		}
	}
	return false
}

func (c CodonChoice) update2(mrnaFrag string) bool {
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
	scaleW := float64(screenWidth) / 1250
	scaleH := float64(screenHeight) / 750
	op.GeoM.Scale(scaleW, scaleH)
	c.rect.width *= int(scaleW)
	c.rect.height *= int(scaleH)
	screen.DrawImage(c.image, op)
}

func (ribo *Ribosome) update_movement() bool {
	if mrna[mrna_ptr].is_complete {
		if mrna_ptr == 4 && ribo.rect.pos.x < screenWidth+50 {
			ribo.rect.pos.x += 4 * (screenWidth / 1250)
			ribo.rect.pos.y += 2 * (screenHeight / 750)
			return false
		} else if ribo.rect.pos.x < (165 * (mrna_ptr + 1)) {
			ribo.rect.pos.x += 2 * (screenWidth / 1250)
			return false
		} else {
			return true
		}
	}

	//ribo.rect.pos.x += screenWidth / 7
	return false
}

func (ribo Ribosome) draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(ribo.rect.pos.x), float64(ribo.rect.pos.y))
	scaleW := float64(screenWidth) / 1250
	scaleH := float64(screenHeight) / 750
	op.GeoM.Scale(scaleW, scaleH)
	ribo.rect.width *= int(scaleW)
	ribo.rect.height *= int(scaleH)
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
