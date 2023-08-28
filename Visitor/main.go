package main

func main() {
	square := &Square{side: 2}
	circle := &Circle{radius: 3}
	rectangle := &Rectangle{l: 2, b: 3}

	ac := &AreaCalculator{}

	square.accept(ac)
	circle.accept(ac)
	rectangle.accept(ac)

	md := &MiddleCoordinates{}
	square.accept(md)
	circle.accept(md)
	rectangle.accept(md)
}
