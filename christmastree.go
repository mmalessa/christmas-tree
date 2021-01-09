package main

import (
	"time"
	"os"
	"math/rand"
	"fmt"

	"github.com/jgarff/rpi_ws281x/golang/ws2811"
)

const (
	pin = 18
	count = 250
	tick = 23
	brightness = 120
	baseColor = uint32(0x003000)
)

// GRB
var     rainbow_tail = []uint32 {
            0x004000,  // red
            0x204000,  // orange
            0x404000,  // yellow
            0x400000,  // green
            0x400040,  // lightblue
            0x000040,  // blue
            0x002020,  // purple
            0x004020,  // pink
	    0x001000,  // dark red
        }

var white_tail = []uint32 {
        0x606060,
	0x606060,
	0x404050,
	0x303040,
	0x202030,
	0x101020,
	0x000010,
}

var vertical_tail = []uint32 {
    0x002000,
    0x004000,
    0x004000,
    0x006000,
    0x006000,
    0x005000,
    0x004000,
    0x003000,
    0x000010,
    0x000030,
    0x000030,
    0x000050,
    0x000050,
    0x000040,
    0x000030,
    0x000020,
    0x100000,
    0x300000,
    0x300000,
    0x500000,
    0x500000,
    0x400000,
    0x300000,
    0x200000,
}

var carousel_tail = []uint32 {
            0x004000,  // red
            0x000040,  // blue
            0x400000,  // green
}

var test_tail = []uint32 {
            0x004000,  // red
            0x400000,  // green
            0x000040,  // blue
            0x004000,  // red
            0x400000,  // green
            0x000040,  // blue
            0x004000,  // red
            0x400000,  // green
            0x000040,  // blue
            0x004000,  // red
            0x400000,  // green
            0x000040,  // blue
}

var tree_matrix =[][]int {
    {249, 249, 249, 249, 249, 249, 249, 249, 249, 249, 249, 249, 249, 249, 249, 249, 249, 249, 249, 249, 249, 249, 249, 249, 249, 249, 249, 249, 249, 249, 249, 249, 249, 249, 249, 249, 249, 249, 249, 249, 249, 249, 249, 249, 249, 249, 249},
    {248, 248, 248, 248, 248, 248, 248, 247, 247, 247, 247, 247, 247, 247, 246, 246, 246, 246, 246, 246, 245, 245, 245, 245, 245, 245, 244, 244, 244, 244, 244, 244, 244, 243, 243, 243, 243, 243, 243, 243, 242, 242, 242, 242, 242, 242, 242},
    {241, 241, 241, 241, 240, 240, 240, 240, 239, 239, 239, 239, 238, 238, 238, 237, 237, 237, 236, 236, 236, 235, 235, 235, 234, 234, 234, 233, 233, 233, 232, 232, 232, 231, 231, 231, 230, 230, 230, 229, 229, 229, 229, 228, 228, 228, 228},
    {227, 227, 227, 226, 226, 226, 225, 225, 224, 224, 223, 223, 222, 222, 221, 221, 220, 220, 219, 219, 218, 218, 217, 217, 216, 216, 215, 215, 214, 214, 213, 213, 212, 212, 211, 211, 210, 210, 209, 209, 208, 208, 207, 207, 206, 206, 206},
    {205, 205, 205, 204, 204, 203, 203, 202, 202, 201, 201, 200, 200, 199, 199, 198, 198, 197, 197, 196, 196, 195, 195, 194, 194, 193, 193, 192, 192, 191, 192, 190, 189, 189, 188, 188, 187, 187, 186, 186, 185, 185, 184, 184, 183, 183, 183},
    {182, 182, 181, 181, 180, 180, 179, 179, 178, 178, 177, 176, 176, 175, 174, 174, 173, 172, 172, 171, 170, 170, 169, 168, 168, 167, 166, 166, 165, 164, 164, 163, 162, 162, 161, 160, 160, 159, 158, 158, 157, 156, 156, 155, 155, 154, 154},
    {153, 153, 152, 151, 151, 150, 149, 148, 148, 147, 146, 146, 145, 144, 144, 143, 142, 141, 141, 140, 139, 139, 138, 137, 137, 136, 135, 135, 134, 134, 133, 132, 132, 131, 130, 129, 129, 128, 127, 127, 126, 125, 124, 124, 123, 122, 122},
    {121, 120, 120, 119, 118, 117, 116, 116, 115, 114, 113, 113, 112, 111, 110, 109, 109, 108, 107, 106, 106, 105, 104, 103, 102, 102, 101, 100, 100, 99,  98,  97,  96,  96,  95,  94,  93,  93,  92,  91,  90,  90,  89,  88,  87,  86,  86},
    {85,  85,  84,  83,  82,  82,  81,  80,  79,  79,  78,  77,  76,  76,  75,  74,  73,  72,  72,  71,  70,  69,  68,  67,  66,  66,  65,  64,  63,  62,  61,  61,  60,  59,  58,  57,  57,  56,  55,  54,  53,  53,  52,  51,  50,  48,  47},
    {46,  45,  44,  43,  42,  41,  40,  39,  38,  37,  36,  35,  34,  33,  32,  31,  30,  29,  28,  27,  26,  25,  24,  23,  22,  21,  20,  19,  18,  17,  16,  15,  14,  13,  12,  11,  10,  9,   8,   7,   6,   5,   4,   3,   2,   1,   0},
}

func main() {
	defer ws2811.Fini()
	err := ws2811.Init(pin, count, brightness)
	if err != nil {
		fmt.Println(err)
	} else {
            err = fillColor(baseColor)
            if err != nil {
                fmt.Println("Error during wipe " + err.Error())
                os.Exit(-1)
            }
            for {
                //err = waterfall(vertical_tail[8:15])
		err = waterfall(vertical_tail[0:7])
		err = fallDown(rainbow_tail)
                //time.Sleep(4000 * time.Millisecond)
                //time.Sleep(4000 * time.Millisecond)
                //err = waterfall(vertical_tail[16:23])
                //time.Sleep(4000 * time.Millisecond)
	        if err != nil {
		    fmt.Println("Error during wipe " + err.Error())
                    os.Exit(-1)
	        }
		err = carousel()
                time.Sleep(1000 * time.Millisecond)
	    }
	}
}

func safeSetLed(id int, color uint32) error {
    if (id >=0 && id <= count) {
        ws2811.SetLed(id, color)
    }
    return nil;
}

func fallDown(tail []uint32) error {
    var taillen = len(tail)
    for i := count; i >= (0 - taillen) + 1 ; i-- {
        for t :=0; t < taillen; t++ {
           safeSetLed(i + t, tail[t])
        }
	//safeSetLed(i + taillen + 1, baseColor)
	err := ws2811.Render()
	if err != nil {
	    ws2811.Clear()
	    return err
	}
	time.Sleep(tick * time.Millisecond)
    }
    return nil
}

func randomBlick(tail []uint32) error {
    var taillen = len(tail)
    var randomLed = rand.Intn(count -1)

    for i := randomLed; i >= (randomLed - taillen); i-- {
        for t := 0; t < taillen; t++ {
            safeSetLed(i + t, tail[t])
	}
        time.Sleep(tick * time.Millisecond)
	err := ws2811.Render()
	if err != nil {
	    ws2811.Clear()
	    return err
	}
    }
    return nil
}

func colorWipe(color uint32) error {
	for i := count; i > 0; i-- {
		ws2811.SetLed(i - 1, uint32(color))
		err := ws2811.Render()
		if err != nil {
			ws2811.Clear()
			return err
		}
		time.Sleep(1 * time.Millisecond)
	}
	return nil
}

func waterfall(tail []uint32) error {
    var rownum = len(tree_matrix)
    var tailnum = len(tail)
    for l :=0; l < (rownum + tailnum); l++ {
        for t := 0; t < tailnum; t++ {
	    var y = l - t;
	    if (y >=0 && y < rownum) {
                var colnumber = len(tree_matrix[y])
                for x :=0; x < colnumber; x++ {
                    safeSetLed(tree_matrix[y][x], tail[t])
                }
            }
	}
        err := ws2811.Render()
        if err != nil {
            ws2811.Clear()
            return err
        }
        time.Sleep(2 * tick * time.Millisecond)
    }
    return nil
}

func carousel() error {
    var colnum = len(tree_matrix[0])
    var rownum = len(tree_matrix)
    var tailnum = len(carousel_tail)

    for c :=0; c < tailnum; c++ {
        for x := 0; x < colnum; x++ {
            for y := 0; y < rownum; y++ {
                safeSetLed(tree_matrix[y][x], carousel_tail[c])
            }
            err := ws2811.Render()
            if err != nil {
                ws2811.Clear()
                return err
            }
            time.Sleep(2 * tick * time.Millisecond)
	}
    }

    return nil
}

func fillColor(color uint32) error {
	for i := count; i > 0; i-- {
		ws2811.SetLed(i - 1, uint32(color))
	}
	err := ws2811.Render()
	if err != nil {
		ws2811.Clear()
		return err
	}
	return nil
}
