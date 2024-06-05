package terminal

import (
	"VDController/docker"
	"VDController/kubernetes"
	"VDController/web"
	"bufio"
	"fmt"
	"os"
	"sync"
)

type Menu struct {
	name    string
	subMenu []*Menu
	action  func()
}

var exitFlag bool

func Terminal(wg *sync.WaitGroup) {
	parentMenu := &Menu{
		name:   "📤返回上级菜单",
		action: func() {},
	}
	exitMenu := &Menu{
		name: "🔌退出控制台",
		action: func() {
			exitFlag = true
			wg.Done()
		},
	}
	Menu := &Menu{
		name: "主菜单" + "\n======",
		subMenu: []*Menu{
			{
				name: "🐳Docker选项",
				subMenu: []*Menu{
					parentMenu,
					{
						name: "🔋查看正在运行的容器",
						action: func() {
							returnValue, err := docker.Dockerls()
							if err != nil {
								fmt.Println(err)
							}
							for _, ctr := range returnValue {
								fmt.Printf("%s %s %s\n", ctr.ID[:3], ctr.Image, ctr.Status)
							}
						},
					},
					exitMenu,
				},
			},
			{
				name: "⚓️Kubernetes",
				subMenu: []*Menu{
					parentMenu,
					{
						name: "列出所有 Pod",
						action: func() {
							kubernetes.GetAllPods()
						},
					},
					exitMenu,
				},
			},
			{
				name: "🌐Web选项",
				subMenu: []*Menu{
					parentMenu,
					{
						name: "启动网页端",
						action: func() {
							go web.StartWeb()
							fmt.Println("✅启动网页端成功")
						},
					},
					exitMenu,
				},
			},
			exitMenu,
		},
	}
	showMenu(Menu)
}

func showMenu(menu *Menu) {
	for {
		// 退出终端不再显示列表
		if exitFlag {
			break
		}
		// 显示菜单名
		fmt.Println(menu.name)
		// 显示菜单项
		for num, subMenu := range menu.subMenu {
			fmt.Printf("%d. %s\n", num+1, subMenu.name)
		}
		// 读取用户输入
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("请选择菜单项: ")
		// 读取用户输入直到碰到换行符
		input, _ := reader.ReadString('\n')
		fmt.Print("============\n")
		// 去除换行符
		var choice int
		fmt.Sscanf(input, "%d", &choice)
		// 处理输入值
		if choice > 0 && choice <= len(menu.subMenu) {
			selected := menu.subMenu[choice-1]
			// 执行菜单项动作
			if selected.action != nil {
				selected.action()
				if selected.name == "📤返回上级菜单" {
					break
				}
			} else if len(selected.subMenu) > 0 {
				showMenu(selected)
			}
		} else {
			fmt.Println("无效的菜单项，请重新输入！\n------------")
		}
	}
}
