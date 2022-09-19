package procon

type Dequeue struct {
	data  []int
	start int
	end   int
}

func NewDequeue(cap int) *Dequeue {
	if cap < 2 {
		cap = 2
	}

	return &Dequeue{data: make([]int, cap), start: cap / 2, end: (cap / 2) - 1}
}

func (d *Dequeue) refresh(isBack bool) {
	d.data = d.data[d.start : d.end+1]

	if isBack {
		d.start = 0
		d.end = len(d.data) - 1
		d.data = append(d.data, make([]int, 100)...)
	} else {
		d.data = append(make([]int, 100), d.data...)
		d.start = 100
		d.end = len(d.data) - 1
	}
}

func (d *Dequeue) Length() int {
	return d.end - d.start + 1
}

func (d *Dequeue) Push(x int) {
	d.end += 1
	if d.end >= len(d.data) {
		d.refresh(true)
	}
	d.data[d.end] = x
}

func (d *Dequeue) FrontPush(x int) {
	if d.start == 0 {
		d.refresh(false)
	}

	d.data[d.start-1] = x
	d.start -= 1
}

func (d *Dequeue) Access(i int) int {
	return d.data[i+d.start]
}

func (d *Dequeue) Pop() {
	d.end -= 1
}

func (d *Dequeue) FrontPop() {
	d.start += 1
}
