package snake

import (
	"fmt"
	"math/rand"
	"os"
	"syscall"
	"time"
	"os/exec"

	//"myown/2048test/menu"
)

/*
#include <windows.h>
#include <conio.h>

// 使用了WinAPI来移动控制台的光标
void gotoxy(int x,int y)
{
    COORD c;
    c.X=x,c.Y=y;
    SetConsoleCursorPosition(GetStdHandle(STD_OUTPUT_HANDLE),c);
}

// 从键盘获取一次按键，但不显示到控制台
int direct()
{
    return _getch();
}
*/
import "C"// go中嵌入C语言的函数

// 表示光标的位置
type loct struct {
	i, j int
}

var (
	area = [20][21]byte{} // 记录了蛇、食物的信息
	food bool             // 当前是否有食物
	lead byte             // 当前蛇头移动方向
	head loct             // 当前蛇头位置
	tail loct             // 当前蛇尾位置
	size int              // 当前蛇身长度
)

// 随机生成一个位置，来放置食物
func place() loct {
	k := rand.Int() % 400
	return loct{k / 20, k % 20}
}

// 用来更新控制台的显示，在指定位置写字符，使用错误输出避免缓冲
func draw(p loct, c int) {

	C.gotoxy(C.int(p.i*2+4), C.int(p.j+1))//蛇的初始化位置
	fmt.Fprintf(os.Stderr, "%c", c)

}

func initialize() {

	// 初始化蛇的位置和方向、首尾；初始化随机数
	head, tail = loct{4, 4}, loct{4, 4}
	lead, size = 'R', 1
	area[4][4] = 'H'
	rand.Seed(int64(time.Now().Unix()))

	defer syscall.FreeLibrary(kernel32)
	SetConsoleTextAttribute(hCon, 0x0075)
	// 输出初始画面
	fmt.Println("  ^0^$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$^0^ ")
	fmt.Println("  ||                                       ||  ")
	fmt.Println("  ||                                       ||  ")
	fmt.Println("  ||                                       ||  ")
	fmt.Println("  ||                                       ||  ")
	fmt.Println("  ||                                       ||  ")
	fmt.Println("  ||                                       ||  ")
	fmt.Println("  ||                                       ||  ")
	fmt.Println("  ||                                       ||  ")
	fmt.Println("  ||                                       ||  ")
	fmt.Println("  ||                                       ||  ")
	fmt.Println("  ||                                       ||  ")
	fmt.Println("  ||                                       ||  ")
	fmt.Println("  ||                                       ||  ")
	fmt.Println("  ||                                       ||  ")
	fmt.Println("  ||                                       ||  ")
	fmt.Println("  ||                                       ||  ")
	fmt.Println("  ||                                       ||  ")
	fmt.Println("  ||                                       ||  ")
	fmt.Println("  ||                                       ||  ")
	fmt.Println("  ||                                       ||  ")
	fmt.Println("  ||                                       ||  ")
	fmt.Println("  ^0^$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$^0^ ")

	// 使用一个单独的go程来捕捉键盘的动作，因为是单独的，不怕阻塞
	go func() {
		for { // 函数只写入lead，外部只读取lead，无需设锁
			switch byte(C.direct()) {
			case 72:
				lead = 'U'
			case 75:
				lead = 'L'
			case 77:
				lead = 'R'
			case 80:
				lead = 'D'
			case 32:
				lead = 'P'
			}
		}
	}()
}

func init() {
	//nargs 代表参数个数
	var nargs uintptr = 1

	//参数需要全部转成uinptr
	hCon, _, _ = syscall.Syscall(uintptr(getStdHandle), nargs, uintptr(STD_OUTPUT_HANDLE), 0, 0)
}

const (
	//标准输出宏
	STD_OUTPUT_HANDLE = uint32(-11 & 0xFFFFFFFF)
)

var (
	err         error
	kernel32, _ = syscall.LoadLibrary("kernel32.dll")
	//设置console属性
	setConsoleTextAttribute, _ = syscall.GetProcAddress(kernel32, "SetConsoleTextAttribute")
	//获取标准输入输出的函数
	getStdHandle, _ = syscall.GetProcAddress(kernel32, "GetStdHandle")
	//标准输出
	hCon uintptr
)

func SetConsoleTextAttribute(hConsoleOutput uintptr, wAttributes uint32) bool {
	var nargs uintptr = 2
	ret, _, _ := syscall.Syscall(setConsoleTextAttribute, nargs, hConsoleOutput, uintptr(wAttributes), 0)
	return ret != 0
}

func Snake() {

	initialize()
	// 主程序
	for {

		// 程序更新周期，400毫秒
		time.Sleep(400 * time.Millisecond)

		// 暂停
		if lead == 'P' {
			continue
		}

		// 放置食物
		if !food {
			give := place()
			if area[give.i][give.j] == 0 { // 食物只能放在空闲位置
				area[give.i][give.j] = 'F'
				draw(give, '◎') // 绘制食物▇◎
				food = true
			}
		}

		// 我们在蛇头位置记录它移动的方向
		area[head.i][head.j] = lead

		// 根据lead来移动蛇头
		switch lead {
		case 'U':
			head.j--
		case 'L':
			head.i--
		case 'R':
			head.i++
		case 'D':
			head.j++
		}

		// 判断蛇头是否出界
		if head.i < 0 || head.i >= 20 || head.j <0|| head.j >= 21 {
			C.gotoxy(0, 23) // 让光标移动到画面下方
			break           // 跳出死循环
		}

		// 获取蛇头位置的原值，来判断是否撞车，或者吃到食物
		eat := area[head.i][head.j]

		if eat == 'F' { // 吃到食物
			food = false

			// 增加蛇的尺寸，并且不移动蛇尾
			size++
		} else if eat == 0 { // 普通移动

			draw(tail, ' ') // 擦除蛇尾

			// 注意我们记录了它移动的方向
			dir := area[tail.i][tail.j]

			// 我们需要擦除蛇尾的记录
			area[tail.i][tail.j] = 0

			// 移动蛇尾
			switch dir {
			case 'U':
				tail.j--
			case 'L':
				tail.i--
			case 'R':
				tail.i++
			case 'D':
				tail.j++
			}
		} else { // 撞车了
			C.gotoxy(0, 23)
			break
		}
		draw(head, '●') // 绘制蛇头◎
	}
	SetConsoleTextAttribute(hCon, 0x0007)
	// 收尾了
//	switch {
//	case size < 22:
//		fmt.Fprintf(os.Stderr, "Faild! You've eaten %d *\n", size-1)
//	case size < 42:
//		fmt.Fprintf(os.Stderr, "Try your best! You've eaten %d *\n", size-1)
//	default:
//		fmt.Fprintf(os.Stderr, "Congratulations! You've eaten %d *\n", size-1)
//	}

switch {
	case size < 22:
		fmt.Printf( "Faild! You've eaten %d *\n", size-1)
	case size < 42:
		fmt.Printf( "Try your best! You've eaten %d *\n", size-1)
	default:
		fmt.Printf("Congratulations! You've eaten %d *\n", size-1)
	}

	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
   
	fmt.Println("当前游戏结束,3秒后退出本程序....")
	time.Sleep(3 * time.Second)
	os.Exit(0)

}
