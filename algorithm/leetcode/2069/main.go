package main

func main() {

}

// var dirs = [4][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

type Robot struct {
	width, height, step, mode int
	isMove                    bool
}

func Constructor(width int, height int) Robot {
	return Robot{
		width:  width,
		height: height,
		mode:   width + height - 2,
	}
}

func (this *Robot) Step(num int) {
	if !this.isMove {
		this.isMove = true
	}
	this.step = (this.step + num) % (this.mode * 2)
}

func (this *Robot) GetPos() []int {
	x, y, _ := this.cal()
	return []int{x, y}
}

func (this *Robot) GetDir() string {
	_, _, dir := this.cal()
	return dir
}

func (this *Robot) cal() (int, int, string) {
	var x, y int
	dir := "East"
	if this.step == 0 {
		if this.isMove {
			dir = "South"
		}
		return x, y, dir
	}
	if this.step <= this.width-1 {
		x = this.step
		return x, y, dir
	}
	if this.step <= this.mode {
		dir = "North"
		x = this.width - 1
		y = this.step - this.width + 1
		return x, y, dir
	}
	if this.step <= this.mode+this.width-1 {
		dir = "West"
		x = 2*this.width + this.height - 3 - this.step
		y = this.height - 1
		return x, y, dir
	}
	dir = "South"
	x = 0
	// y = this.height - 1 - (this.step - 2 * (this.width - 1) - (this.height - 1))
	y = this.mode*2 - this.step
	return x, y, dir
}

/**
 * Your Robot object will be instantiated and called as such:
 * obj := Constructor(width, height);
 * obj.Step(num);
 * param_2 := obj.GetPos();
 * param_3 := obj.GetDir();
 */
