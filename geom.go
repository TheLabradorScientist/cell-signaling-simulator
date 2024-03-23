package main

// Storing coordinates
type Vector struct {
	x int
	y int
}

// Storing a rectangle's data
type Rectangle struct {
	pos    Vector
	width  int
	height int
}

// Creating a vector
func newVector(x int, y int) Vector {
	return Vector{
		x: x,
		y: y,
	}
}

// Creating a rectangle
func newRect(x int, y int, width int, height int) Rectangle {
	return Rectangle{
		pos:    newVector(x, y),
		width:  width,
		height: height,
	}
}

// Check if a point and a rectangle collide
func rect_point_collision(rect Rectangle, point Vector) bool {
	if rect.pos.x <= point.x && rect.pos.x+rect.width >= point.x &&
		rect.pos.y <= point.y && rect.pos.y+rect.height >= point.y {
		return true
	}
	return false
}

// Check if two rectangle collides
func aabb_collision(rect1 Rectangle, rect2 Rectangle) bool {
	if rect1.pos.x < rect2.pos.x+rect2.width &&
		rect1.pos.x+rect1.width > rect2.pos.x &&
		rect1.pos.y < rect2.pos.y+rect2.height &&
		rect1.pos.y+rect1.height > rect2.pos.y {
		return true
	}
	return false
}
