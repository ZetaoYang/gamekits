package main

import "myown/gamekits/go2048"
import "myown/gamekits/menu"
import "myown/gamekits/snake"
import ."myown/gamekits/startui"
import "fmt"
import "os"
import "os/exec"
import "strconv"
import "time"
import "github.com/shiyanlou/termbox-go"

const(
	
	def termbox.Attribute = termbox.ColorDefault//cmd默认的颜色
	green termbox.Attribute = termbox.ColorGreen
	cyan termbox.Attribute = termbox.ColorCyan
	red termbox.Attribute = termbox.ColorRed 
	blue termbox.Attribute = termbox.ColorBlue
	yellow termbox.Attribute = termbox.ColorYellow
	black termbox.Attribute = termbox.ColorBlack
	magenta termbox.Attribute = termbox.ColorMagenta
	white termbox.Attribute = termbox.ColorWhite
)

var (
	wcolor termbox.Attribute= red
	tcolor termbox.Attribute= yellow
	ncolor termbox.Attribute= magenta
	wbcolor termbox.Attribute= white
	lcolor termbox.Attribute= black
	nbcolor termbox.Attribute= white
)

func clearscreen() {

	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {

	var game1 go2048.G2048
	var game2 go2048.G4096
	var str1 string
	var s string
	clearscreen()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	Startui()
Game:	for {
			
		clearscreen()
  		menu.Kitmenu(true)
		var str string
		fmt.Scanln(&str)
		j, _ := strconv.Atoi(str) //字符串转换成整型

		switch j {

		case 6:
			clearscreen()
			os.Exit(0)
		case 5:
			
		case 4:
	
		case 3:


		case 2:
 
	
		case 1:
			goto G2
		default:
			fmt.Print("                     输入错误，请重试!")			
			var progress int=0
			for progress < 200 {
			time.Sleep(100 * time.Millisecond)
			fmt.Print(".")
			progress += 10
			}
			
		}
	}
	
G2:	for {
		
		
	Here:
		var flagm bool
		var flagdicm bool
		flagm = true
		clearscreen()
		s=""
	Default:
		s=""
		menu.Menu(flagm)
		fmt.Scanln(&s)
		i, _ := strconv.Atoi(s) //字符串转换成整型

		switch i {

		case 6:
			clearscreen()
			goto Game		
		case 5:
			game2.Run()
		case 4:
			game1.Run(wcolor,tcolor,ncolor,wbcolor,lcolor,nbcolor)
	//game1.Run(red,yellow,magenta,white,black,white)默认的颜色设置
	//game1.Run(提示字颜色,格子颜色,数字颜色,界面背景颜色,纵横细线颜色,数字背景颜色),
	//即,game1.Run(wcolor,tcolor,ncolor,wbcolor,lcolor,nbcolor)
		case 3:
		    clearscreen()
			flagdicm=true
 Dicmenu:   fmt.Println()
			menu.Dicmenu(flagdicm)
			str1=""
			fmt.Scanln(&str1)

			if str1 == "menu" {
				goto Here
			}else {
				clearscreen()
				flagdicm=false
				goto Dicmenu
			}
		case 2:
 Set2048m:  clearscreen()
			fmt.Println()	
			menu.Set2048()
			str1=""	
			fmt.Scanln(&str1)
			
			
			if str1 == "menu" {
				goto Here
			}else if str1 == "21" {

		        clearscreen()
				var flag1 bool=true
	Wcolmenu:	menu.Wcolmenu(flag1)
				str1=""							
				
				fmt.Scanln(&str1)
				if str1 == "menu" {	
					str1=""
					goto Set2048m	
				}else if str1=="211"{
				wcolor = yellow
				game1.Run(wcolor,yellow,magenta,white,black,white)
				str1=""
				}else if str1=="212"{
				wcolor = blue
				game1.Run(wcolor,yellow,magenta,white,black,white)
				str1=""
				}else if str1=="213"{
				wcolor = red
				game1.Run(wcolor,yellow,magenta,white,black,white)
				str1=""
				}else if str1=="214"{
				wcolor = green
				game1.Run(wcolor,yellow,magenta,white,black,white)
				str1=""
				}else if str1=="215"{
				wcolor = cyan
				game1.Run(wcolor,yellow,magenta,white,black,white)
				str1=""
				}else if str1=="216"{
				wcolor = black
				game1.Run(wcolor,yellow,magenta,white,black,white)
				str1=""
				}else if str1=="217"{
				wcolor = magenta
				game1.Run(wcolor,yellow,magenta,white,black,white)
				str1=""
				}else if str1=="218"{
				wcolor = white
				game1.Run(wcolor,yellow,magenta,white,black,white)
				str1=""
				}else if str1=="219"{
				wcolor = def
				game1.Run(wcolor,yellow,magenta,white,black,white)
				str1=""
				}else{
					clearscreen()
					flag1=false
					str1=""
					goto Wcolmenu
				}

			}else if str1 == "22"{
	  			clearscreen()
				var flag2 bool=true
	Tcolmenu:   menu.Tcolmenu(flag2)
				str1=""
				fmt.Scanln(&str1)
				if str1 == "menu" {
					str1=""
					goto Set2048m
				}else if str1=="221"{
				tcolor = yellow
				game1.Run(red,tcolor,magenta,white,black,white)
				str1=""
				}else if str1=="222"{
				tcolor = blue
				game1.Run(red,tcolor,magenta,white,black,white)
				str1=""
				}else if str1=="223"{
				tcolor = red
				game1.Run(red,tcolor,magenta,white,black,white)
				str1=""
				}else if str1=="224"{
				tcolor = green
				game1.Run(red,tcolor,magenta,white,black,white)
				str1=""
				}else if str1=="225"{
				tcolor = cyan
				game1.Run(red,tcolor,magenta,white,black,white)
				str1=""
				}else if str1=="226"{
				tcolor = black
				game1.Run(red,tcolor,magenta,white,black,white)
				str1=""
				}else if str1=="227"{
				tcolor = magenta
				game1.Run(red,tcolor,magenta,white,black,white)
				str1=""
				}else if str1=="228"{
				tcolor =white
				game1.Run(red,tcolor,magenta,white,black,white)
				str1=""
				}else if str1=="229"{
				tcolor = def
				game1.Run(red,tcolor,magenta,white,black,white)
				str1=""
				}else{
					clearscreen()
					flag2=false					
					str1=""
					goto Tcolmenu
				}
				
			}else if str1=="23"{
				clearscreen()
				var flag3 bool=true
	Ncolmenu:   menu.Ncolmenu(flag3)
				str1=""
				fmt.Scanln(&str1)
				if str1 == "menu" {
					str1=""
					goto Set2048m
				}else if str1=="231"{
				ncolor = yellow
				game1.Run(red,yellow,ncolor,white,black,white)
				str1=""
				}else if str1=="232"{
				ncolor = blue
				game1.Run(red,yellow,ncolor,white,black,white)
				str1=""
				}else if str1=="233"{
				ncolor = red
				game1.Run(red,yellow,ncolor,white,black,white)
				str1=""
				}else if str1=="234"{
				ncolor = green
				game1.Run(red,yellow,ncolor,white,black,white)
				str1=""
				}else if str1=="235"{
				ncolor = cyan
				game1.Run(red,yellow,ncolor,white,black,white)
				str1=""
				}else if str1=="236"{
				ncolor = black
				game1.Run(red,yellow,ncolor,white,black,white)
				str1=""
				}else if str1=="237"{
				ncolor = magenta
				game1.Run(red,yellow,ncolor,white,black,white)
				str1=""
				}else if str1=="238"{
				ncolor = white
				game1.Run(red,yellow,ncolor,white,black,white)
				str1=""
				}else if str1=="239"{
				ncolor = def
				game1.Run(red,yellow,ncolor,white,black,white)
				str1=""
				}else{
					clearscreen()
					flag3=false				
					str1=""
					goto Ncolmenu
				}
				
			}else if str1=="24"{
				clearscreen()
				var flag4 bool=true
	Wbcolmenu:  menu.Wbcolmenu(flag4)
				str1=""
				fmt.Scanln(&str1)
				if str1 == "menu" {
					str1=""
					goto Set2048m
				}else if str1=="241"{
				wbcolor = yellow
				game1.Run(red,yellow,magenta,wbcolor,black,white)
				str1=""
				}else if str1=="242"{
				wbcolor = blue
				game1.Run(red,yellow,magenta,wbcolor,black,white)
				str1=""
				}else if str1=="243"{
				wbcolor = red
				game1.Run(red,yellow,magenta,wbcolor,black,white)
				str1=""
				}else if str1=="244"{
				wbcolor = green
				game1.Run(red,yellow,magenta,wbcolor,black,white)
				str1=""
				}else if str1=="245"{
				wbcolor = cyan
				game1.Run(red,yellow,magenta,wbcolor,black,white)
				str1=""
				}else if str1=="246"{
				wbcolor = black
				game1.Run(red,yellow,magenta,wbcolor,black,white)
				str1=""
				}else if str1=="247"{
				wbcolor = magenta
				game1.Run(red,yellow,magenta,wbcolor,black,white)
				str1=""
				}else if str1=="248"{
				wbcolor = white
				game1.Run(red,yellow,magenta,wbcolor,black,white)
				str1=""
				}else if str1=="249"{
				wbcolor = def
				game1.Run(red,yellow,magenta,wbcolor,black,white)
				str1=""
				}else{
					clearscreen()
					flag4=false							
					str1=""
					goto Wbcolmenu
				}
			
			}else if str1=="25"{
				clearscreen()
				var flag5 bool=true
	Lcolmenu:   menu.Lcolmenu(flag5)
				str1=""
				fmt.Scanln(&str1)
				if str1 == "menu" {
					str1=""
					goto Set2048m
				}else if str1=="251"{
				lcolor = yellow
				game1.Run(red,yellow,magenta,wbcolor,lcolor,white)
				str1=""
				}else if str1=="252"{
				lcolor = blue
				game1.Run(red,yellow,magenta,wbcolor,lcolor,white)
				str1=""
				}else if str1=="253"{
				lcolor = red
				game1.Run(red,yellow,magenta,wbcolor,lcolor,white)
				str1=""
				}else if str1=="254"{
				lcolor = green
				game1.Run(red,yellow,magenta,wbcolor,lcolor,white)
				str1=""
				}else if str1=="255"{
				lcolor = cyan
				game1.Run(red,yellow,magenta,wbcolor,lcolor,white)
				str1=""
				}else if str1=="256"{
				lcolor = black
				game1.Run(red,yellow,magenta,wbcolor,lcolor,white)
				str1=""
				}else if str1=="257"{
				lcolor = magenta
				game1.Run(red,yellow,magenta,wbcolor,lcolor,white)
				str1=""
				}else if str1=="258"{
				lcolor = white
				game1.Run(red,yellow,magenta,wbcolor,lcolor,white)
				str1=""
				}else if str1=="259"{
				lcolor = def
				game1.Run(red,yellow,magenta,wbcolor,lcolor,white)
				str1=""
				}else{
					clearscreen()
					flag5=false					
					str1=""
					goto Lcolmenu
				}
					
			}else if str1=="26"{
				clearscreen()
				var flag6 bool=true
	Nbcolmenu:  menu.Nbcolmenu(flag6)
				str1=""
				fmt.Scanln(&str1)
				if str1 == "menu" {
					str1=""
					goto Set2048m
				}else if str1=="261"{
				nbcolor = yellow
				game1.Run(red,yellow,magenta,wbcolor,lcolor,nbcolor)
				str1=""
				}else if str1=="262"{
				nbcolor = blue
				game1.Run(red,yellow,magenta,wbcolor,lcolor,nbcolor)
				str1=""
				}else if str1=="263"{
				nbcolor = red
				game1.Run(red,yellow,magenta,wbcolor,lcolor,nbcolor)
				str1=""
				}else if str1=="264"{
				nbcolor = green
				game1.Run(red,yellow,magenta,wbcolor,lcolor,nbcolor)
				str1=""
				}else if str1=="265"{
				nbcolor = cyan
				game1.Run(red,yellow,magenta,wbcolor,lcolor,nbcolor)
				str1=""
				}else if str1=="266"{
				nbcolor = black
				game1.Run(red,yellow,magenta,wbcolor,lcolor,nbcolor)
				str1=""
				}else if str1=="267"{
				nbcolor = magenta
				game1.Run(red,yellow,magenta,wbcolor,lcolor,nbcolor)
				str1=""
				}else if str1=="268"{
				nbcolor = white
				game1.Run(red,yellow,magenta,wbcolor,lcolor,nbcolor)
				str1=""
				}else if str1=="269"{
				nbcolor = def
				game1.Run(red,yellow,magenta,wbcolor,lcolor,nbcolor)
				str1=""
				}else{
					clearscreen()
					flag6=false
					str1=""
					goto Nbcolmenu
				}
				
				
			}else if str1=="27"{
				wcolor,tcolor,ncolor,wbcolor,lcolor,nbcolor=red,yellow,magenta,white,black,white
				fmt.Println("                       颜色设置一键还原成功!")
				time.Sleep(800*time.Millisecond)
				//game1.Run(wcolor,tcolor,ncolor,wbcolor,lcolor,nbcolor)
			}else{
			fmt.Print("                                ")
			fmt.Println("输入错误!")
			fmt.Print("           2秒后将返回选项菜单。")
			fmt.Print("请稍后")
			var progress int=0
			for progress < 200 {
			time.Sleep(100 * time.Millisecond)
			fmt.Print(".")
			progress += 10
			}
			goto Set2048m
								
		}
		case 1:
			clearscreen()
			snake.Snake()
			s=""
		default:
			clearscreen()
			fmt.Printf("\a\a")//警告声			
			flagm = false
			goto Default

		}
	}

}