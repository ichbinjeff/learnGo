package main

import "fmt"

const (
	WHITE = iota
	BLACK
	BLUE
	RED
	YELLOW
)

type Color byte

type Box struct {
	width, height, depth float64
	color                Color
}

type BoxList []Box

func (b Box) Volume() float64 {
	return b.width * b.height * b.depth
}

func (b *Box) setColor(c Color) {
	b.color = c
}

func (bl BoxList) BiggestColor() Color {
	var rst Color = Color(WHITE)
	max := 0.00
	for _, b := range bl {
		if bv := b.Volume(); bv > max {
			max = bv
			rst = b.color
		}
	}
	return rst
}

func (bl BoxList) paintItBlack() {
	for i, _ := range bl {
		bl[i].setColor(BLACK)
	}
}

func (c Color) String() string {
	s := []string{"WHITE", "BLACK", "BLUE", "RED", "YELLOW"}
	return s[c]
}

func main() {
	boxes := BoxList{
		Box{4, 4, 4, RED},
		Box{10, 10, 1, YELLOW},
		Box{1, 1, 20, BLACK},
		Box{10, 10, 1, BLUE},
		Box{10, 30, 1, WHITE},
		Box{20, 20, 20, RED},
	}

	fmt.Printf("we have %d boxes in our set\n", len(boxes))
	fmt.Printf("the color of first box is ", boxes[0].color.String())
	fmt.Printf("the volumn of the first box is", boxes[0].Volume())
	fmt.Printf("the color of the last box is", boxes[len(boxes)-1].color)
	fmt.Printf("Now Let's paint them all black")
	boxes.paintItBlack()
	fmt.Println("The color of the second one is", boxes[1].color.String())
	fmt.Println("the biggest one is", boxes.BiggestColor())
}
