package main

import (
	"fmt"
	"github.com/andlabs/ui"
	"time"
	"os"
	"strings"
	"os/exec"
	"io/ioutil"
	"path/filepath"
	"runtime"

)


var mBinCfgFilePath string
var currentPath string

/*
type BinFile_t struct {
   path string
   to string
}

var BinFile [5]BinFile_t
*/

func getCurrDir() string {
	d, e := filepath.Abs(filepath.Dir(os.Args[0]))
	if e != nil {
		fmt.Println("err")
	}
	return strings.Replace(d, "\\", "/", -1)
}

func Show_ui() {
	currentPath = getCurrDir()
	
	err := ui.Main(func() {
		ckbox := ui.NewCheckbox("ckbox") // checkbox
		input := ui.NewEntry()
		button := ui.NewButton("Select bin")
		button_d := ui.NewButton("Download")
		greeting := ui.NewLabel("")
		box := ui.NewVerticalBox()
		box.Append(ui.NewLabel("Terminal:"), false)
		box.Append(input, false)
		box.Append(button, false)
		box.Append(button_d, false)
		box.Append(greeting, false)
		box.Append(ckbox, false)
		window := ui.NewWindow("Hello1", 500, 500, false)
		window.SetMargined(true)
		window.SetChild(box)

		button.OnClicked(func(*ui.Button) {
			/*Usb_write(input.Text())
			if input.Text() != "" {
				greeting.SetText(input.Text())
			} else {
				greeting.SetText("no cmd..")
			}*/
			mBinCfgFilePath = ui.OpenFile(window) // 打开文件路径.
		})
		
		button_d.OnClicked(func(*ui.Button) {
			exec_stlink_download("b")
			exec_stlink_download("ftm-m")
			exec_stlink_download("ftm-f")
		})
		
		window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true
		})

		window.Show()
		
		//go UpdateUI(greeting)
		//time.Sleep(1000 * 1000)
	})
	if err != nil {
		panic(err)
	}
}

func substring(source string, start int, end int) string {
	var substring = ""
	var pos = 0
	for _, c := range source {
	    if pos < start {
	        pos++
	        continue
	    }
	    if pos >= end {
	        break
	    }
	    pos++
	    substring += string(c)
	}
	return substring
}

func exec_shell_sh(action string) string {
	var script string
	var path string
	
	if runtime.GOOS == "linux" {
		script = "/" + "download.sh"
		path = strings.Replace(mBinCfgFilePath, "/map.txt", "",  -1)
	} else {
		script = "\\" + "download.bat"
		path = strings.Replace(mBinCfgFilePath, "\\map.txt", "",  -1)
	}
	
    cmd := exec.Command(currentPath + script, path, action)
	fmt.Println(cmd.Args)
    stdout, err := cmd.StdoutPipe()
	cmd.StdoutPipe() 
    cmd.Start()
    content, err := ioutil.ReadAll(stdout)
    if err != nil {
        fmt.Println(err)
		return "err"
    } 
    return string(content)
}

func exec_stlink_download(action string) {
	result := exec_shell_sh(action)
	fmt.Println(result)
}


func UpdateUI(label *ui.Label) {
	for {
		ret, rev := Usb_read()
		if ret < 0 {
			fmt.Printf("Usb_read err: %d\n", ret)
		}
		ui.QueueMain(func() {
			label.SetText(rev)
		})
		time.Sleep(1000 * 1000)
	}
}


func ShowUI() {
	err := ui.Main(func() {
        input := ui.NewEntry()
        input.SetText("this is input element")
        input.LibuiControl()
        spinbox := ui.NewSpinbox(50, 150)
        spinbox.SetValue(55)
        slider := ui.NewSlider(0, 100)
        slider.SetValue(10)
        processbar := ui.NewProgressBar()
        processbar.SetValue(50)
        combobox := ui.NewCombobox()
        combobox.Append("select one")
        combobox.Append("select two")
        combobox.Append("select three")
        combobox.SetSelected(2)
        checkbox1 := ui.NewCheckbox("GoLang")
        checkbox1.SetChecked(true)
        checkbox2 := ui.NewCheckbox("C++")
        checkbox3 := ui.NewCheckbox("Python")
        checkbox4 := ui.NewCheckbox("Other")
        checkbox_div := ui.NewHorizontalBox()
        checkbox_div.Append(checkbox1, true)
        checkbox_div.Append(checkbox2, true)
        checkbox_div.Append(checkbox3, true)
        checkbox_div.Append(checkbox4, true)
        radio := ui.NewRadioButtons()
        radio.Append("GoLang")
        radio.Append("C++")
        radio.Append("Python")
        radio.Append("Other")
        checkbox_div.SetPadded(true)
        Separator := ui.NewHorizontalSeparator()
        Separator_label_l := ui.NewLabel("left")
        Separator_label_r := ui.NewLabel("right")
        Separator_div := ui.NewHorizontalBox()
        Separator_div.Append(Separator_label_l, true)
        Separator_div.Append(Separator, false)
        Separator_div.Append(Separator_label_r, true)
        Separator_div.SetPadded(true)
        datetimepicker := ui.NewDateTimePicker()

        //-----------------Set a single child to a  new group.------------

        container1 := ui.NewGroup("input(输入框)")
        container1.SetChild(input)
        container2 := ui.NewGroup("spinbox(自设值范围框，只能通过箭头控制值，不能手动输入)")
        container2.SetChild(spinbox)
        container3 := ui.NewGroup("slider(滑片)")
        container3.SetChild(slider)
        container4 := ui.NewGroup("processbar（进度条）")
        container4.SetChild(processbar)
        container5 := ui.NewGroup("checkbox（多选框）")
        container5.SetChild(checkbox_div)
        container6 := ui.NewGroup("radio（单选框）")
        container6.SetChild(radio)
        container7 := ui.NewGroup("combobox（下拉框）")
        container7.SetChild(combobox)
        container8 := ui.NewGroup("Separator（分隔符）")
        container8.SetChild(Separator_div)
        container9 := ui.NewGroup("datetimepicker(时间选取器)")
        container9.SetChild(datetimepicker)

        //------垂直排列的容器---------
        div := ui.NewVerticalBox()
        //------水平排列的容器
        boxs_1 := ui.NewHorizontalBox()
        boxs_1.Append(container1, true)
        boxs_1.Append(container2, true)

        boxs_1.SetPadded(false)
        boxs_2 := ui.NewHorizontalBox()
        boxs_2.Append(container3, true)
        boxs_2.Append(container4, true)

        boxs_3 := ui.NewHorizontalBox()
        boxs_3.Append(container5, true)
        boxs_3.Append(container6, true)

        boxs_4 := ui.NewHorizontalBox()
        boxs_4.Append(container7, true)
        boxs_4.Append(container8, true)

        div.Append(boxs_1, true)
        div.Append(boxs_2, true)
        div.Append(boxs_3, true)
        div.Append(boxs_4, true)
        div.Append(container9, true)
        div.SetPadded(false)
        window := ui.NewWindow("test111", 500, 500, true)

        window.SetChild(div)
        window.SetMargined(true)

        window.OnClosing(func(*ui.Window) bool {
            //ui.Quit()
            return true
        })
        window.Show()
    })
    if err != nil {
        panic(err)
    }
}







