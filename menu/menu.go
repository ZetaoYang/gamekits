package menu

import (
	"fmt"
	"syscall"
    "time"
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

func Kitmenu(flag bool) {
	defer syscall.FreeLibrary(kernel32)
	SetConsoleTextAttribute(hCon, 0x001b)

	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	if flag==true{
	fmt.Println()}else{
		SetConsoleTextAttribute(hCon, 0x000c)
		
		fmt.Print("                               ")		
		var j int     //以下为闪烁文字的代码
        for j=1;j<=10;j++{//闪烁5次
            switch j{
        case 1:
        fmt.Printf("请输入正确的选项!")
        time.Sleep(200 * time.Millisecond)
        
        case 2:

        fmt.Printf("\r ")
        fmt.Printf("                          ")
        time.Sleep(200 * time.Millisecond)
         
        case 3:
        fmt.Printf("\r ")
        fmt.Printf("                              请输入正确的选项!")
        time.Sleep(200 * time.Millisecond)

        case 4:

        fmt.Printf("\r ")
        fmt.Printf("                                                    ")
        time.Sleep(200 * time.Millisecond)

        case 5:
        fmt.Printf("\r ")
        fmt.Printf("                              请输入正确的选项!")
        time.Sleep(200 * time.Millisecond)

        case 6:

        fmt.Printf("\r ")
        fmt.Printf("                                                    ")
        time.Sleep(200 * time.Millisecond)

        case 7:
        fmt.Printf("\r ")
        fmt.Printf("                              请输入正确的选项!")
        time.Sleep(200 * time.Millisecond)

        case 8:

        fmt.Printf("\r ")
        fmt.Printf("                                                    ")
        time.Sleep(200 * time.Millisecond)

        case 9:
        fmt.Printf("\r ")
        fmt.Printf("                              请输入正确的选项!")
        time.Sleep(200* time.Millisecond)

        case 10:

        fmt.Printf("\r ")
        fmt.Printf("                                                    ")
        time.Sleep(200 * time.Millisecond)

        default:

            }
    	}
			//fmt.Println()
			fmt.Printf("\r ")
			fmt.Println("                              请输入正确的选项!")
	}
	
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x00e0)
	fmt.Println("----------游戏集菜单----------")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x00e0)
	fmt.Println("**   (a).挑战2048：    1    **")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x00e0)
	fmt.Println("**   (b).挑战    ：    2    **")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x00e0)
	fmt.Println("**   (c).挑战    ：    3    **")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x00e0)
	fmt.Println("**   (d).挑战    ：    4    **")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x00e0)
	fmt.Println("**   (e).挑战    ：    5    **")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x00e0)
	fmt.Println("**   (f).退出游戏菜单：6    **")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x00e0)
	fmt.Println("------------------------------")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x00e0)
	fmt.Println(" 完成相应功能 ，请输入对应数字")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x00e0)
	fmt.Printf(" 您选择的是.................:")

	//恢复默认值
	SetConsoleTextAttribute(hCon, 0x0007)
}

func Menu(flag bool) {
	defer syscall.FreeLibrary(kernel32)
	SetConsoleTextAttribute(hCon, 0x001b)

	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	if flag==true{
	fmt.Println()}else{
		SetConsoleTextAttribute(hCon, 0x000c)
		
		fmt.Print("                               ")		
		var j int     //以下为闪烁文字的代码
        for j=1;j<=10;j++{//闪烁5次
            switch j{
        case 1:
        fmt.Printf("请输入正确的选项!")
        time.Sleep(200 * time.Millisecond)
        
        case 2:

        fmt.Printf("\r ")
        fmt.Printf("                          ")
        time.Sleep(200 * time.Millisecond)
         
        case 3:
        fmt.Printf("\r ")
        fmt.Printf("                              请输入正确的选项!")
        time.Sleep(200 * time.Millisecond)

        case 4:

        fmt.Printf("\r ")
        fmt.Printf("                                                    ")
        time.Sleep(200 * time.Millisecond)

        case 5:
        fmt.Printf("\r ")
        fmt.Printf("                              请输入正确的选项!")
        time.Sleep(200 * time.Millisecond)

        case 6:

        fmt.Printf("\r ")
        fmt.Printf("                                                    ")
        time.Sleep(200 * time.Millisecond)

        case 7:
        fmt.Printf("\r ")
        fmt.Printf("                              请输入正确的选项!")
        time.Sleep(200 * time.Millisecond)

        case 8:

        fmt.Printf("\r ")
        fmt.Printf("                                                    ")
        time.Sleep(200 * time.Millisecond)

        case 9:
        fmt.Printf("\r ")
        fmt.Printf("                              请输入正确的选项!")
        time.Sleep(200* time.Millisecond)

        case 10:

        fmt.Printf("\r ")
        fmt.Printf("                                                    ")
        time.Sleep(200 * time.Millisecond)

        default:

            }
    	}
			//fmt.Println()
			fmt.Printf("\r ")
			fmt.Println("                              请输入正确的选项!")
	}
	
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x001b)
	fmt.Println("---------游戏选项菜单---------")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x001b)
	fmt.Println("**   (a).挑战贪食蛇：    1  **")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x001b)
	fmt.Println("**   (b).设置2048：      2  **")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x001b)
	fmt.Println("**   (c).2048玩法说明：  3  **")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x001b)
	fmt.Println("**   (d).开始2048吧：    4  **")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x001b)
	fmt.Println("**   (e).挑战4096：      5  **")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x001b)
	fmt.Println("**   (f).退出选项菜单：  6  **")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x001b)
	fmt.Println("------------------------------")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x001b)
	fmt.Println(" 完成相应功能 ，请输入对应数字")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x001b)
	fmt.Printf(" 您选择的是.................:")

	//恢复默认值
	SetConsoleTextAttribute(hCon, 0x0007)
}

func Dicmenu(flag bool) {
	defer syscall.FreeLibrary(kernel32)
	SetConsoleTextAttribute(hCon, 0x008c)

	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	
	if flag==true{
	fmt.Println()}else{
		SetConsoleTextAttribute(hCon, 0x000c)
		fmt.Print("                            ")		
		var j int     //以下为闪烁文字的代码
        for j=1;j<=10;j++{ //闪烁5次
            switch j{
        case 1:
		fmt.Printf("\a\a")//警告声
        fmt.Printf("输入错误,请重试!")
		
        time.Sleep(200 * time.Millisecond)
        
        case 2:

        fmt.Printf("\r ")
        fmt.Printf("                       ")
        time.Sleep(200 * time.Millisecond)
         
        case 3:
        fmt.Printf("\r ")
        fmt.Printf("                           输入错误,请重试!")
        time.Sleep(200 * time.Millisecond)

        case 4:

        fmt.Printf("\r ")
        fmt.Printf("                                                 ")
        time.Sleep(200 * time.Millisecond)

        case 5:
        fmt.Printf("\r ")
        fmt.Printf("                           输入错误,请重试!")
        time.Sleep(200 * time.Millisecond)

        case 6:

        fmt.Printf("\r ")
        fmt.Printf("                                                 ")
        time.Sleep(200 * time.Millisecond)

        case 7:
        fmt.Printf("\r ")
        fmt.Printf("                           输入错误,请重试!")
        time.Sleep(200 * time.Millisecond)

        case 8:

        fmt.Printf("\r ")
        fmt.Printf("                                                 ")
        time.Sleep(200 * time.Millisecond)

        case 9:
        fmt.Printf("\r ")
        fmt.Printf("                           输入错误,请重试!")
        time.Sleep(200* time.Millisecond)

        case 10:

        fmt.Printf("\r ")
        fmt.Printf("                                                 ")
        time.Sleep(200 * time.Millisecond)

        default:

            }
    	}
			
			fmt.Printf("\r ")
			fmt.Println("                           输入错误,请重试!")
	}
	
	
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                         ")
	SetConsoleTextAttribute(hCon, 0x0075)
	fmt.Println("-----2048游戏说明-----")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                         ")
	SetConsoleTextAttribute(hCon, 0x0075)
	fmt.Println("    按Enter键重玩     ")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                         ")
	SetConsoleTextAttribute(hCon, 0x0075)
	fmt.Println("    按方向键移动      ")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                         ")
	SetConsoleTextAttribute(hCon, 0x0075)
	fmt.Println("    按Esc键退出游戏   ")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                         ")
	SetConsoleTextAttribute(hCon, 0x0075)
	fmt.Println("       Enjoy !!!      ")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                         ")
	SetConsoleTextAttribute(hCon, 0x0075)
	fmt.Println(" <输入menu返回主菜单> ")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                         ")
	SetConsoleTextAttribute(hCon, 0x0075)
	fmt.Println("----------------------")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                         ")
	SetConsoleTextAttribute(hCon, 0x0075)
	fmt.Printf(" 请返回主菜单.......:")
	//恢复默认值
	SetConsoleTextAttribute(hCon, 0x0007)
}

func Set2048(){
	
	defer syscall.FreeLibrary(kernel32)
	SetConsoleTextAttribute(hCon, 0x003d)
	
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x0072)
	fmt.Println("----------请选择设置内容-----------")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x0072)
	fmt.Println("*     提示字颜色设置：     21     *")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x0072)
	fmt.Println("*     格子颜色设置：       22     *")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x0072)
	fmt.Println("*     数字颜色设置：       23     *")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x0072)
	fmt.Println("*     界面背景颜色设置：   24     *")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x0072)
	fmt.Println("*     纵横细线颜色设置：   25     *")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x0072)
	fmt.Println("*     数字背景颜色设置：   26     *")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x0072)
	fmt.Println("*     一键还原颜色设置：   27     *")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x0072)
	fmt.Println("       <输入menu返回主菜单>        ")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x0072)
	fmt.Println("-----------------------------------")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x0072)
	fmt.Printf("  请输入所选项或返回主菜单.......:")
	
	SetConsoleTextAttribute(hCon, 0x0007)//恢复默认值
}

func Wcolmenu(flag bool) {
	defer syscall.FreeLibrary(kernel32)
	SetConsoleTextAttribute(hCon, 0x003d)
	if flag==true{
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	}else{
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
		SetConsoleTextAttribute(hCon, 0x000c)
		fmt.Print("                               ")
		fmt.Println("输入错误，请重试!")
		SetConsoleTextAttribute(hCon, 0x0007)
	}
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x003d)
	fmt.Println("-------请选择提示字颜色-------")
	SetConsoleTextAttribute(hCon, 0x0007)//
	fmt.Print("                       ")//
	SetConsoleTextAttribute(hCon, 0x003d)//
	fmt.Println("*     Yellow：     211       *")//
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x003d)
	fmt.Println("*     Blue：       212       *")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x003d)
	fmt.Println("*     Red：        213       *")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x003d)
	fmt.Println("*     Green：      214       *")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x003d)
	fmt.Println("*     Cyan：       215       *")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x003d)
	fmt.Println("*     Black：      216       *")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x003d)
	fmt.Println("*     Magenta：    217       *")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x003d)
	fmt.Println("*     White：      218       *")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x003d)
	fmt.Println("*     Default：    219       *")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x003d)
	fmt.Println("    <输入menu返回上一菜单>    ")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x003d)
	fmt.Println("------------------------------")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x003d)
	fmt.Printf("请输入所选项或返回上一菜单..:")
	
	SetConsoleTextAttribute(hCon, 0x0007)//恢复默认值
}

func Tcolmenu(flag bool) {
	defer syscall.FreeLibrary(kernel32)
	SetConsoleTextAttribute(hCon, 0x003d)
	if flag==true{
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	}else{
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
		SetConsoleTextAttribute(hCon, 0x000c)
		fmt.Print("                               ")
		fmt.Println("输入错误，请重试!")
		SetConsoleTextAttribute(hCon, 0x0007)
	}
	
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x003d)
	fmt.Println("--------请选择格子颜色--------")
	SetConsoleTextAttribute(hCon, 0x0007)//
	fmt.Print("                       ")//
	SetConsoleTextAttribute(hCon, 0x003d)//
	fmt.Println("*     Yellow：     221       *")//
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x003d)
	fmt.Println("*     Blue：       222       *")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x003d)
	fmt.Println("*     Red：        223       *")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x003d)
	fmt.Println("*     Green：      224       *")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x003d)
	fmt.Println("*     Cyan：       225       *")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x003d)
	fmt.Println("*     Black：      226       *")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x003d)
	fmt.Println("*     Magenta：    227       *")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x003d)
	fmt.Println("*     White：      228       *")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x003d)
	fmt.Println("*     Default：    229       *")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x003d)
	fmt.Println("    <输入menu返回上一菜单>    ")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x003d)
	fmt.Println("------------------------------")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x003d)
	fmt.Printf("请输入所选项或返回上一菜单..:")
	
	SetConsoleTextAttribute(hCon, 0x0007)//恢复默认值
}

func Ncolmenu(flag bool) {
	defer syscall.FreeLibrary(kernel32)
	SetConsoleTextAttribute(hCon, 0x003d)
	if flag==true{
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	}else{
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
		SetConsoleTextAttribute(hCon, 0x000c)
		fmt.Print("                               ")
		fmt.Println("输入错误，请重试!")
		SetConsoleTextAttribute(hCon, 0x0007)
	}
	
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x003d)
	fmt.Println("--------请选择数字颜色--------")
	SetConsoleTextAttribute(hCon, 0x0007)//
	fmt.Print("                       ")//
	SetConsoleTextAttribute(hCon, 0x003d)//
	fmt.Println("*     Yellow：     231       *")//
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x003d)
	fmt.Println("*     Blue：       232       *")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x003d)
	fmt.Println("*     Red：        233       *")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x003d)
	fmt.Println("*     Green：      234       *")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x003d)
	fmt.Println("*     Cyan：       235       *")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x003d)
	fmt.Println("*     Black：      236       *")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x003d)
	fmt.Println("*     Magenta：    237       *")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x003d)
	fmt.Println("*     White：      238       *")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x003d)
	fmt.Println("*     Default：    239       *")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x003d)
	fmt.Println("    <输入menu返回上一菜单>    ")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x003d)
	fmt.Println("------------------------------")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x003d)
	fmt.Printf("请输入所选项或返回上一菜单..:")
	
	SetConsoleTextAttribute(hCon, 0x0007)//恢复默认值
}

func Wbcolmenu(flag bool) {
	defer syscall.FreeLibrary(kernel32)
	SetConsoleTextAttribute(hCon, 0x003d)
	if flag==true{
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	}else{
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
		SetConsoleTextAttribute(hCon, 0x000c)
		fmt.Print("                               ")
		fmt.Println("输入错误，请重试!")
		SetConsoleTextAttribute(hCon, 0x0007)
	}
	
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x003d)
	fmt.Println("------请选择界面背景颜色------")
	SetConsoleTextAttribute(hCon, 0x0007)//
	fmt.Print("                       ")//
	SetConsoleTextAttribute(hCon, 0x003d)//
	fmt.Println("*     Yellow：     241       *")//
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x003d)
	fmt.Println("*     Blue：       242       *")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x003d)
	fmt.Println("*     Red：        243       *")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x003d)
	fmt.Println("*     Green：      244       *")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x003d)
	fmt.Println("*     Cyan：       245       *")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x003d)
	fmt.Println("*     Black：      246       *")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x003d)
	fmt.Println("*     Magenta：    247       *")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x003d)
	fmt.Println("*     White：      248       *")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x003d)
	fmt.Println("*     Default：    249       *")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x003d)
	fmt.Println("    <输入menu返回上一菜单>    ")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x003d)
	fmt.Println("------------------------------")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x003d)
	fmt.Printf("请输入所选项或返回上一菜单..:")
	
	SetConsoleTextAttribute(hCon, 0x0007)//恢复默认值
}

func Lcolmenu(flag bool) {
	defer syscall.FreeLibrary(kernel32)
	SetConsoleTextAttribute(hCon, 0x003d)
	if flag==true{
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	}else{
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
		SetConsoleTextAttribute(hCon, 0x000c)
		fmt.Print("                               ")
		fmt.Println("输入错误，请重试!")
		SetConsoleTextAttribute(hCon, 0x0007)
	}
	
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x003d)
	fmt.Println("------请选择纵横细线颜色------")
	SetConsoleTextAttribute(hCon, 0x0007)//
	fmt.Print("                       ")//
	SetConsoleTextAttribute(hCon, 0x003d)//
	fmt.Println("*     Yellow：     251       *")//
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x003d)
	fmt.Println("*     Blue：       252       *")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x003d)
	fmt.Println("*     Red：        253       *")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x003d)
	fmt.Println("*     Green：      254       *")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x003d)
	fmt.Println("*     Cyan：       255       *")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x003d)
	fmt.Println("*     Black：      256       *")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x003d)
	fmt.Println("*     Magenta：    257       *")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x003d)
	fmt.Println("*     White：      258       *")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x003d)
	fmt.Println("*     Default：    259       *")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x003d)
	fmt.Println("    <输入menu返回上一菜单>    ")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x003d)
	fmt.Println("------------------------------")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x003d)
	fmt.Printf("请输入所选项或返回上一菜单..:")
	
	SetConsoleTextAttribute(hCon, 0x0007)//恢复默认值
}

func Nbcolmenu(flag bool) {
	defer syscall.FreeLibrary(kernel32)
	SetConsoleTextAttribute(hCon, 0x003d)
	if flag==true{
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	}else{
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
		SetConsoleTextAttribute(hCon, 0x000c)
		fmt.Print("                               ")
		fmt.Println("输入错误，请重试!")
		SetConsoleTextAttribute(hCon, 0x0007)
	}
	
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x003d)
	fmt.Println("------请选择数字背景颜色------")
	SetConsoleTextAttribute(hCon, 0x0007)//
	fmt.Print("                       ")//
	SetConsoleTextAttribute(hCon, 0x003d)//
	fmt.Println("*     Yellow：     261       *")//
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x003d)
	fmt.Println("*     Blue：       262       *")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x003d)
	fmt.Println("*     Red：        263       *")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x003d)
	fmt.Println("*     Green：      264       *")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x003d)
	fmt.Println("*     Cyan：       265       *")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x003d)
	fmt.Println("*     Black：      266       *")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x003d)
	fmt.Println("*     Magenta：    267       *")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x003d)
	fmt.Println("*     White：      268       *")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x003d)
	fmt.Println("*     Default：    269       *")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x003d)
	fmt.Println("    <输入menu返回上一菜单>    ")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x003d)
	fmt.Println("------------------------------")
	SetConsoleTextAttribute(hCon, 0x0007)
	fmt.Print("                       ")
	SetConsoleTextAttribute(hCon, 0x003d)
	fmt.Printf("请输入所选项或返回上一菜单..:")
	
	SetConsoleTextAttribute(hCon, 0x0007)//恢复默认值
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
