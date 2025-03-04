package object

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"os/user"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// We can use this to define the builtin functions//

// could it be used to add plugin functions?

var FunctionBuiltins = map[string]*Builtin{
	"len": {
		Fn: func(args ...Object) Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d", len(args))
			}

			switch arg := args[0].(type) {
			case *Array:
				return &Integer{Value: int64(len(arg.Elements))}
			case *String:
				return &Integer{Value: int64(len(arg.Value))}
			default:
				return newError("arguments to len not supported, got=%s", args[0].Type())
			}
		},
	},
	"append": {
		Fn: func(args ...Object) Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2", len(args))
			}
			if args[0].Type() != ARRAY_OBJ {
				return newError("argument to push must be ARRAY, got %s", args[0].Type())
			}

			arr := args[0].(*Array)
			length := len(arr.Elements)

			newElements := make([]Object, length+1)
			copy(newElements, arr.Elements)
			newElements[length] = args[1]

			return &Array{Elements: newElements}
		},
	},
	"pause": {
		Fn: func(args ...Object) Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			switch args[0].Type() {
			case INTEGER_OBJ:
				time.Sleep(time.Second * time.Duration(args[0].(*Integer).Value))
			case FLOAT_OBJ:
				time.Sleep(time.Second * time.Duration(args[0].(*Float).Value))
			default:
				return newError("argument to pause must be INTEGER or FLOAT, got %s", args[0].Type())
			}

			return &Null{}
		},
	},
	"Start": {
		Fn: func(args ...Object) Object {
			DeleteScreen

			config.StartChannel <- true

			return &Null{}
		},
	},
	/*------------os------------*/
	"OStype": {
		Fn: func(args ...Object) Object {
			return &String{Value: runtime.GOOS}
		},
	},
	"OSusername": {
		Fn: func(args ...Object) Object {
			user, err := user.Current()
			if err != nil {
				return newError("Not find OS user, got=%s", err)
			}

			switch username := user.Username; runtime.GOOS {
			case "windows":
				switch slice := strings.Split(username, "\\"); len(slice) {
				case 1:
					return &String{Value: slice[0]}
				case 2:
					return &String{Value: slice[1]}
				default:
					return newError("username is error, isn't your username? got=%s", username)
				}
			default:
				return newError("OSusername support only windows, your OS = %s", runtime.GOOS)
			}
		},
	},
	"OSname": {
		Fn: func(args ...Object) Object {
			user, err := user.Current()
			if err != nil {
				return newError("Not find OS user, got=%s", err)
			}

			switch username := user.Name; runtime.GOOS {
			case "windows":
				return &String{Value: username}
			default:
				return newError("OSname support only windows, your OS = %s", runtime.GOOS)
			}
		},
	},
	"OSFileNames": {
		Fn: func(args ...Object) Object {
			if len(args) != 1 {
				return newError("OSFile func has 0 args, got=%d", len(args))
			}

			path, ok := args[0].(*String)
			if !ok {
				return newError("OSFileNames func args is not string")
			}

			files, err := ioutil.ReadDir(path.Value)
			if err != nil {
				return newError("Golang ioutil.ReadDir Error\n\n%v", err)
			}

			var arr []Object

			for _, file := range files {
				arr = append(arr, &String{Value: file.Name()})
			}

			return &Array{
				Elements: arr,
			}
		},
	},
	/*
			"OSbgChange": {
				Fn: func(args ...Object) Object {
					if runtime.GOOS != "windows" {
						return newError("OSbgChange func supports only windows")
					}

					if len(args) != 1 {
						return newError("OSbgChange func has 1 args, got=%d", len(args))
					}

					changePath := syscall.StringToUTF16Ptr(config.Path + args[0].(*String).Value)
					currentPath := make([]uint16, syscall.MAX_PATH)

					proc := syscall.NewLazyDLL("user32.dll").NewProc("SystemParametersInfoW")

					proc.Call(
						0x0073, // SPI_GETDESKWALLPAPER
						syscall.MAX_PATH,
						uintptr(unsafe.Pointer(&currentPath[0])),
						0,
					)

					proc.Call(
						20, // SPI_SETDESKWALLPAPER
						0,
						uintptr(unsafe.Pointer(changePath)),
						0x01,
					)

					var n int
					for n = 0; n < len(currentPath) && currentPath[n] != 0; n++ {
					}

					return &String{Value: string(utf16.Decode(currentPath[:n]))}
				},
			},
			"OSbgReturn": {
				Fn: func(args ...Object) Object {
					if runtime.GOOS != "windows" {
						return newError("OSbgReturn func supports only windows")
					}

					if len(args) != 1 {
						return newError("OSbgChange func has 1 args, got=%d", len(args))
					}

					proc := syscall.NewLazyDLL("user32.dll").NewProc("SystemParametersInfoW")

					changePath := syscall.StringToUTF16Ptr(args[0].(*String).Value)
					currentPath := make([]uint16, syscall.MAX_PATH)

					proc.Call(
						0x0073, // SPI_GETDESKWALLPAPER
						syscall.MAX_PATH,
						uintptr(unsafe.Pointer(&currentPath[0])),
						0,
					)

					var n int
					for n = 0; n < len(currentPath) && currentPath[n] != 0; n++ {
					}

					if string(utf16.Decode(currentPath[:n])) != args[0].(*String).Value {
						proc.Call(
							20, // SPI_SETDESKWALLPAPER
							0,
							uintptr(unsafe.Pointer(changePath)),
							0x01,
						)
					}

					return NULL
				},
			},
		"MessageBox": {
			Fn: func(args ...Object) Object {
				if runtime.GOOS != "windows" {
					return newError("MessageBox func supports only windows")
				}

				if len(args) != 2 {
					return newError("MessageBox func has 2 args, got=%d", len(args))
				}

				title, ok := args[0].(*String)
				if !ok {
					return newError("MessageBox func Title args is not string")
				}

				contents, ok := args[1].(*String)
				if !ok {
					return newError("MessageBox func Contents args is not string")
				}

				syscall.NewLazyDLL("user32.dll").NewProc("MessageBoxW").Call(
					0,
					uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(contents.Value))),
					uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(title.Value))),
					0,
				)

				return NULL
			},
		},
	*/
	/*-----------Data System----------*/
	"IsExistDataStore": {
		Fn: func(args ...Object) Object {
			if len(args) != 1 {
				return newError("ExistDataStore func has 1 args, got=%d", len(args))
			}

			storeName, ok := args[0].(*String)
			if !ok {
				return newError("ExistDataStore func storeName args is not string")
			}

			if _, err := os.Stat(config.Path + "data\\" + storeName.Value); os.IsNotExist(err) {
				return &Boolean{Value: false}
			}
			return &Boolean{Value: true}
		},
	},
	"MakeDataStore": {
		Fn: func(args ...Object) Object {
			if len(args) != 1 {
				return newError("MakeDataStore func has 1 args, got=%d", len(args))
			}

			storeName, ok := args[0].(*String)
			if !ok {
				return newError("MakeDataStore func storeName args is not string")
			}

			if _, err := ioutil.ReadDir(config.Path + "data"); err != nil {
				err := os.Mkdir(config.Path+"data", 0755)
				if err != nil {
					return newError("Golang os.Mkdir Error\n\n%v", err)
				}

				_, err = os.Create(config.Path + "data\\" + storeName.Value)
				if err != nil {
					return newError("Golang os.Create Error\n\n%v", err)
				}
			} else {
				_, err := os.Create(config.Path + "data\\" + storeName.Value)
				if err != nil {
					return newError("Golang os.Create Error\n\n%v", err)
				}
			}
			return &Null{}
		},
	},
	/*-----------util-----------*/
	"Character": {
		Fn: func(args ...Object) Object {
			if len(args) != 2 {
				return newError("Character func has 2 args, got=%d", len(args))
			}

			name, ok := args[0].(*String)
			if !ok {
				return newError("Character func name args is not string")
			}

			color, ok := args[1].(*String)
			if !ok {
				return newError("Character func color args is not string")
			}

			hex := make([]int64, 3)
			switch color.Value[:1] {
			case "#":
				hex[0], _ = strconv.ParseInt(color.Value[1:3], 16, 32)
				hex[1], _ = strconv.ParseInt(color.Value[3:5], 16, 32)
				hex[2], _ = strconv.ParseInt(color.Value[5:7], 16, 32)
			default:
				return newError("Color support only hex code")
			}

			return &Character{
				Name: &String{
					Value: name.Value,
				},
				Color: core.CreateColor(int(hex[0]), int(hex[1]), int(hex[2])),
			}
		},
	},
	"Link": {
		Fn: func(args ...Object) Object {
			if runtime.GOOS == "windows" {
				cmd := exec.Command("cmd", "/C", "start", "/max", args[0].(*String).Value)
				cmd.Run()
			}

			return &Null{}
		},
	},
	"SetVolume": {
		Fn: func(args ...Object) Object {
			if len(args) != 2 {
				return newError("SetVolume func has 1 args, got = %d", len(args))
			}

			channel, ok := args[0].(*String)
			if !ok {
				return newError("SetVolume func channel args is not string")
			}

			value, ok := args[1].(*Integer)
			if !ok {
				return newError("SetVolume func value args is not integer")
			}

			switch channel.Value {
			case "music":
				SetVolume(-1, int(value.Value))
			case "sound":
				core.SetVolume(0, int(value.Value))
			case "voice":
				core.SetVolume(1, int(value.Value))
			default:
				// TODO
				return newError("channel index over")
			}

			return NULL
		},
	},
	"print": { // 테스트 용으로 제작된 임시 출력 함수
		Fn: func(args ...Object) Object {
			for _, arg := range args {
				fmt.Println(arg.Inspect())
			}

			return NULL
		},
	},
}

func GetBuiltinByName(name string) *Builtin {
	for _, def := range FunctionBuiltins {
		if def.Name == name {
			return def.Builtin
		}
	}

	return nil
}
