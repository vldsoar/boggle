import QtQuick 2.7
import QtQuick.Controls 2.0
import QtQuick.Layouts 1.3
import QtQuick.Controls.Material 2.0

import "/js/functions.js" as Functions

Pane {
    id: paneLobby
    width: parent.width
    height: parent.height
    anchors.horizontalCenter: parent.horizontalCenter
    anchors.verticalCenter: parent.verticalCenter

    ColumnLayout {
        id: columnLeftLobby
        width: 386
        height: 456
        anchors.top: parent.top
        anchors.topMargin: 0
        anchors.left: parent.left
        anchors.leftMargin: 0

        Text {
            id: welcomeText
            text: "Hello, " + user.data.username
            Layout.fillWidth: true
            anchors.top: parent.top
            anchors.topMargin: 50
//                anchors.bottom: columnCreateRoom.top
//                anchors.bottomMargin: 50
            anchors.horizontalCenter: parent.horizontalCenter
            verticalAlignment: Text.AlignVCenter
            horizontalAlignment: Text.AlignHCenter
            font.pixelSize: 20
        }

        Column {
            id: columnCreateRoom
            width: 200
            height: 258
            Layout.alignment: Qt.AlignHCenter | Qt.AlignVCenter
            spacing: 10
            Layout.fillHeight: false
            Layout.fillWidth: true
            anchors.verticalCenter: paneLobby.verticalCenter
            anchors.top: welcomeText.bottom
            anchors.topMargin: 120

            TextField {
                id: textFieldCreateRoom
                width: 300
                text: qsTr("")
                placeholderText: "Room Name"
                selectionColor: "#4f0080"
                anchors.horizontalCenter: parent.horizontalCenter
                font.pixelSize: 16
            }


            Button {
                id: btnCreateRoom
                x: 0
                y: 30
                width: 300
                text: qsTr("Create Room")
                highlighted: true
                anchors.horizontalCenter: parent.horizontalCenter
                background: Material.color(Material.Purple)
                onClicked: {
                    GameController.createRoom(textFieldCreateRoom.text)
                }
            }

            Text {
                id: textErrorCreateRoom
                text: ""
                anchors.horizontalCenter: parent.horizontalCenter
                font.pixelSize: 12
                color: "red"
            }


        }

    }

    ColumnLayout {
        id: columnRightLobby
        x: 386
        width: 180
        height: parent.height
        anchors.top: parent.top
        anchors.topMargin: 0
        anchors.right: parent.right
        anchors.rightMargin: 0

        Label {
            id: labelRooms
            text: qsTr("Rooms")
            font.pointSize: 15
            horizontalAlignment: Text.AlignHCenter
            verticalAlignment: Text.AlignVCenter
            Layout.fillWidth: true
            Layout.fillHeight: true
        }

        ListView {
            id: listViewRooms
            width: 180
            height: 160
            clip: true
            boundsBehavior: Flickable.StopAtBounds
            Layout.fillHeight: true
            Layout.fillWidth: true
            highlight: Rectangle { color: Material.color(Material.Grey, Material.Shade300); }
            focus: true
            model: ListModel {
                id: listModelRooms
            }
            delegate: Item {
                x: 5
                Layout.fillWidth: true
                width: 180
                height: 40
                Row {
                    id: row1
                    Rectangle {
                        width: 2
                        height: 40
                        color: Material.color(Functions.getMaterialColor())
                    }

                    Text {
                        width: 120
                        text: name
                        font.bold: true
                        anchors.verticalCenter: parent.verticalCenter
                        MouseArea {
                            anchors.fill: parent
                            onClicked: listViewRooms.currentIndex = index
                        }
                    }

                    Rectangle {
                        id: rect
                        width: 40
                        height: 40
                        z: 10
                        Image {
                            width: 30
                            height: 30
                            anchors.fill: parent
                            id: iconDoor
                            source: "/images/black/sign_in.png"
                        }
                        MouseArea {
                            anchors.fill: rect
                            onClicked: GameController.joinRoom(name)
                        }
                    }
                    spacing: 10
                }

            }
        }



    }

    Component.onDestruction: listModelRooms.clear()


    Connections {
        target: GameController
        onJoinRoomError: {
            textErrorCreateRoom.text = reply
        }
        onCreateRoomError: {
            console.log(reply)
        }
        onJoinedRoom: {
            console.log(reply)
            var jsonObject = JSON.parse(reply)
            user.data = jsonObject

            console.log(JSON.stringify(user.data))
            stackView.push("qrc:/pages/JoinedRoom.qml")

        }
        onUpdateRooms: {
            console.log("****UPDATE ROOMS*****")
            var jsonRooms = JSON.parse(reply)
            var data = user.data

            data["rooms"] = jsonRooms

            user.data = data

            console.log(user.data)

            listModelRooms.clear()
            jsonRooms.forEach(function(r) {
                listModelRooms.append({name: r})
            })
        }
    }
}