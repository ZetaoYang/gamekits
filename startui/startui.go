package startui

import(
 	"fmt"
	"time"
	"syscall"
)

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

func init() {
	//nargs 代表参数个数
	var nargs uintptr = 1

	//参数需要全部转成uinptr
	hCon, _, _ = syscall.Syscall(uintptr(getStdHandle), nargs, uintptr(STD_OUTPUT_HANDLE), 0, 0)
}
func SetConsoleTextAttribute(hConsoleOutput uintptr, wAttributes uint32) bool {
	var nargs uintptr = 2
	ret, _, _ := syscall.Syscall(setConsoleTextAttribute, nargs, hConsoleOutput, uintptr(wAttributes), 0)
	return ret != 0
}

func Startui(){
	
			fmt.Println("                        仿真游戏集合®  By golang")
			fmt.Println("                                  ---by Ahmed_Zetao_Yang")
			fmt.Println()
			fmt.Println()
			fmt.Println()
			fmt.Println()
			fmt.Println()
			fmt.Println()
			fmt.Println()
			fmt.Println()
			fmt.Println()
			fmt.Println()
			fmt.Println()
			fmt.Print("                            ")
			defer syscall.FreeLibrary(kernel32)			
			SetConsoleTextAttribute(hCon, 0x001d)
		     fmt.Println("正在进入游戏集")
			
			SetConsoleTextAttribute(hCon, 0x0007)//恢复默认值
			fmt.Print("    ✈")
			for i:=1;i<=100;i++ {
				time.Sleep(50 * time.Millisecond)
				SetConsoleTextAttribute(hCon, 0x00b0)
				fmt.Print(".")
				SetConsoleTextAttribute(hCon, 0x0007)//恢复默认值
				if i%2==0||i%3==2||i%5==0||i%4==0{
				fmt.Print("\b\b")
				}
				SetConsoleTextAttribute(hCon, 0x00e0)
				fmt.Printf("赶集中%d%%",i)
				SetConsoleTextAttribute(hCon, 0x0007)//恢复默认值
       	 		time.Sleep(80 * time.Millisecond)//✈✳☁❄☀⤵⤴®✖™✖☑✔
       			if i<10{
       			 fmt.Printf("\b\b\b\b\b\b\b")//\b\b\b\b\b=赶集中
				}else {
						fmt.Printf("\b\b\b\b\b\b\b\b")
					}
				}
				
			}
			
/*
在win32环境中，控制台的文字和背景可以通过动态链接库kernel32.dll的一个函数SetConsoleTextAttribute()这个函数实现。 这个函数接受一个标准输出的handle作为第一个参数，第二个参数是用来控制颜色的attribute。控制台的颜色分成16种不同的值。 每个都可以用一个16进制的数来表示。

分别是：

0 = Black
1 = Blue
2 = Green
3 = Aqua
4 = Red
5 = Purple
6 = Yellow
7 = White
8 = Gray
9 = Light Blue
A = Light Green
B = Light Aqua
C = Light Red
D = Light Purple
E = Light Yellow
F = Bright White
例如：0x004f:表示字是Red，背景是Bright White
即：32位的高位表示背景，低位表示文字颜色。
*/
			