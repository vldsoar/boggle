import QtQuick 2.7
import QtQuick.Controls 2.0
import QtQuick.Layouts 1.3
import QtQuick.Controls.Material 2.0


Pane {
    id: paneRoom
    width: parent.width
    height: parent.height
    anchors.verticalCenter: parent.verticalCenter
    anchors.horizontalCenter: parent.horizontalCenter

    property int usersConnected: user.data["joinedRoom"]["users"].length

    BusyIndicator {
        id: loaderRoom
        x: 266
        y: 178
        width: 100
        height: 100
        anchors.verticalCenter: parent.verticalCenter
        anchors.horizontalCenter: parent.horizontalCenter
    }

    Text {
        id: textLoader
        text: qsTr("Waiting for game to start")
        anchors.top: loaderRoom.bottom
        anchors.topMargin: 13
        anchors.horizontalCenter: parent.horizontalCenter
        horizontalAlignment: Text.AlignHCenter
        verticalAlignment: Text.AlignVCenter
        font.pixelSize: 20
    }

    Text {
        id: textPlayers
        text: "players: " + usersConnected + "/5"
        anchors.top: parent.top
        anchors.topMargin: 13
        verticalAlignment: Text.AlignVCenter
        font.pixelSize: 10
    }

    Button {
        id: btnStartGame
        text: "Start Game"
        width: 160
        background: Material.color(Material.Purple)
        highlighted: true
        enabled: false
        visible: false
        anchors.verticalCenter: loaderRoom.verticalCenter
        anchors.horizontalCenter: loaderRoom.horizontalCenter
        z: 10
    }

    Connections {
        target: GameController

        onPreparingGame: {
            paneRoom.state = "preparingGame"
        }

        onReadyGame: {
            var jsonReply = JSON.parse(reply)
            var data = user.data
            data["currentGame"] = jsonReply
            user.data = data
            stackView.push("qrc:/pages/Game.qml")
        }

    }

    states: [
        State {
            name: "preparingGame"
            PropertyChanges {
                target: textLoader
                text: "Preparing Game..."
            }
            PropertyChanges {
                target: loaderRoom
                running: true
            }
            PropertyChanges {
                target: btnStartGame
                enabled: false
                visible: false
                onClicked: null
            }

        },

        State {
            name: "waitingUsers"
            when: {
                if (user.data["joinedRoom"].initial && user.data["joinedRoom"]["users"].length < 2) {
                    return true
                }
                return false
            }
            PropertyChanges {
                target: textLoader
                text: "Waiting for at least one player"
            }
        },

        State {
            name: "readyPlayGame"
            when: usersConnected >= 2 && user.data["joinedRoom"].initial
            PropertyChanges {
                target: textLoader
                text: "Ready to Play"
            }
            PropertyChanges {
                target: loaderRoom
                running: false
            }
            PropertyChanges {
                target: btnStartGame
                enabled: true
                visible: true
                onClicked: {
                    GameController.startGame(user.data["joinedRoom"].name)
                }
            }
        }

    ]
}
