package main

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"os"
	"github.com/therecipe/qt/quickcontrols2"
	"github.com/therecipe/qt/qml"
	"io/ioutil"
	"fmt"
	"encoding/json"
)

const (
	maxDatagramSize = 8192
)

//type Client struct {
//	Recipients []string
//	AddrMulticast string
//	Me models.Peer
//	Dic map[string]int
//}

type User struct {
	core.QObject
	_ int 				`property:"state"`
	_ core.QJsonObject  `property:"data"`
}


var dic = make(map[string]int)

func main() {
	fmt.Println("-> Start application")
	loadDic()
	fmt.Println("-> Dictionary loaded")

	config := make(map[string]string)

	raw := loadFile("./config.json")

	json.Unmarshal(raw, &config)

	controller := NewGameController(nil)

	controller.init(config["serverAddr"])
	guiInterface(controller)
}


func guiInterface(controller *GameController) {
	core.QCoreApplication_SetApplicationName("Boggle Game")
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	gui.NewQGuiApplication(len(os.Args), os.Args)

	quickcontrols2.QQuickStyle_SetStyle("material")

	User_QmlRegisterType2("User", 1, 0, "User")

	var engine = qml.NewQQmlApplicationEngine(nil)
	engine.Load(core.NewQUrl3("qrc:/app.qml", 0))


	engine.RootContext().SetContextProperty("GameController", controller)


	gui.QGuiApplication_Exec()
}


func loadDic() {
	temp := make(map[string]int)

	raw := loadFile("./dic.json")

	json.Unmarshal(raw, &temp)

	for word, _ := range temp {
		dic[word] = len(word)
	}
}

func loadFile(path string) []byte {
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return raw
}