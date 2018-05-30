import QtQuick 2.7
import QtQuick.Controls 2.0
import QtQuick.Layouts 1.3
import QtQuick.Controls.Material 2.0

import "/js/functions.js" as Functions
import "/js/actionsButton.js" as ActionButton

Pane {
    id: pageGame
    width: application.width
    height: application.height
    anchors.fill: parent

    property var stackWord: []
    property var partialScore: 0

    Pane {
        id: paneGame
        x: 46
        y: 8
        width: 310
        height: 442
        anchors.horizontalCenter: parent.horizontalCenter

        Timer {
            id: timer
            interval: 1000; running: false; repeat: true
            onTriggered: {
                countDown.seconds--
                textTime.text = Functions.toDateTime(countDown.seconds)
                if (countDown.seconds == 0) {
                    timer.stop()
                    stackView.push("qrc:/pages/ResultsPage.qml")
                }
            }
        }

        GridView {
            id: grid
            width: 280; height: 280
            anchors.horizontalCenter: parent.horizontalCenter
            contentWidth: 0
            keyNavigationWraps: false
            interactive: false
            boundsBehavior: Flickable.DragOverBounds
            layoutDirection: Qt.LeftToRight
            cellWidth: 70; cellHeight: 70


            ListModel {
                id: diceModel
            }


            Component {
                id: dicesDelegate

                Rectangle {
                    id: wrapper
                    width: grid.cellWidth - 10
                    height: grid.cellHeight - 10
                    radius: 4
                    color: Material.color(Material.Purple)
                    Material.elevation: 6

                    Text {
                        id: contactInfo
                        text: face
                        font.pointSize: 20
                        anchors.horizontalCenter: parent.horizontalCenter
                        anchors.verticalCenter: parent.verticalCenter
                        color: wrapper.GridView.isCurrentItem ? "red" : "white"
                        font.capitalization: Font.AllUppercase
                    }

                    MouseArea {
                        id: mouseDice
                        anchors.fill: parent
                        onClicked: {
                            stackWord.push(index)
                            wrapper.state = "disabled"
                            inputWord.text = inputWord.text.concat(face)
                            grid.currentIndex = index
                        }
                        hoverEnabled: true
                    }

                    states: [
                        State {
                            name: "disabled"

                            PropertyChanges {
                                target: mouseDice
                                enabled: false
                            }

                            PropertyChanges {
                                target: wrapper
                                color: "gray"

                            }
                        }

                    ]
                }
            }

            model: diceModel
            delegate: dicesDelegate
            focus: true

            Component.onCompleted: { // append
                user.data["currentGame"]["board"].forEach(function(dice) {
                    diceModel.append({face: dice})
                })
            }
        }

        ColumnLayout {
            id: colButtons
            y: 216
            width: 283
            height: 70
            spacing: 10
            anchors.left: application.right
            anchors.horizontalCenter: parent.horizontalCenter
            anchors.top: grid.bottom
            anchors.topMargin: 20

            Text {
                id: textError
                text: qsTr("")
                Layout.alignment: Qt.AlignLeft | Qt.AlignTop
                Layout.fillWidth: true
                color: "red"
//                anchors.top: grid.bottom
//                anchors.topMargin: 5
            }


            TextInput {
                id: inputWord
                text: qsTr("")
                font.family: "Ubuntu"
                font.capitalization: Font.AllUppercase
                Layout.alignment: Qt.AlignLeft | Qt.AlignTop
                Layout.rowSpan: 1
                Layout.fillWidth: true
                font.pointSize: 20
                readOnly: true
                verticalAlignment: Text.AlignVCenter

            }

            Button {
                id: sendButton
                width: 66
                height: 40
                text: qsTr("Send")
                spacing: 3
                Layout.fillWidth: true
                anchors.top: inputWord.bottom
                anchors.topMargin: 10
                topPadding: 6
                highlighted: true
                Material.background: Material.color(Material.Purple)
                onClicked: {
                    var send = GameController.sendWord(inputWord.text.toLowerCase())
                    if (send) {
                        partialScore += inputWord.text.length
                        ActionButton.pushWord(inputWord.text, listViewWords)
                        ActionButton.resetInputWord(inputWord, grid, stackWord)
                        textError.text = ""

                    } else {
                        textError.text = "Invalid Word"
                    }

                }
            }

            RowLayout {
                id: rowActionButtons
                Layout.fillWidth: true
                anchors.top: sendButton.bottom
                anchors.topMargin: 10
                spacing: 5
                height: 50

                Button {
                    id: resetButton
                    text: qsTr("Reset")
                    anchors.left: parent.left
                    anchors.leftMargin: 0
                    Layout.fillWidth: true
                    onClicked: {
                        ActionButton.resetInputWord(inputWord, grid, stackWord)
                    }

                }

                Button {
                    id: backButton
                    text: qsTr("Back")
                    anchors.right: parent.right
                    anchors.rightMargin: 0
                    Layout.fillWidth: true
                    onClicked: {
                        if (inputWord.text.length > 0) {
                            stackWord.pop()

                            inputWord.text = ""

                            stackWord.forEach(function(l) {
                                inputWord.text += diceModel.get(l).face
                            })

                            grid.currentItem.state = ""

                            grid.currentIndex = stackWord[stackWord.length - 1] || 0

                        }
                    }
                }
            }


        }


    }

    GridLayout {
        id: gridContainerInfo
        width: 310
        height: 40
        y: 520
        anchors.horizontalCenterOffset: 0
//        anchors.top: grid.bottom
//        anchors.topMargin: 237
        anchors.horizontalCenter: parent.horizontalCenter


        RowLayout {
            id: rowScore
            width: 100
            height: 100
            Layout.fillHeight: true
            Layout.fillWidth: true

            Label {
                id: labelScore
                text: qsTr("Partial Score:")
                font.bold: true
                font.pixelSize: 16
            }

            Text {
                id: textScore
                text: partialScore
                verticalAlignment: Text.AlignVCenter
                font.pixelSize: 15
            }
        }

        RowLayout {
            id: rowTime
            width: 100
            height: 100
            Layout.fillHeight: true
            Layout.fillWidth: true

            Label {
                id: labelTime
                text: qsTr("Time: ")
                Layout.fillWidth: false
                font.pixelSize: 16
                font.bold: true
            }

            Item {
                id: countDown
                property int seconds: 120
                anchors.left: labelTime.right

                Text {
                    id: textTime
                    text: "02:00"
                    horizontalAlignment: Text.AlignRight
                    verticalAlignment: Text.AlignVCenter
                    font.pixelSize: 17
                }
            }


        }
    }

    ColumnLayout {
        id: rowPlayers
        width: 180
        height: application.height
        anchors.top: parent.top
        anchors.topMargin: 0
        anchors.right: parent.right
        anchors.rightMargin: 0

        Label {
            id: labelPlayers
            text: qsTr("Players")
            wrapMode: Text.WrapAnywhere
            horizontalAlignment: Text.AlignHCenter
            Layout.alignment: Qt.AlignHCenter | Qt.AlignVCenter
            Layout.fillHeight: true
            Layout.fillWidth: true
            font.pointSize: 13
            anchors.top: parent.top
            anchors.topMargin: 9
            verticalAlignment: Text.AlignVCenter
        }

        ListView {
            id: listViewPlayers
            width: 110
            height: 160
            Layout.fillWidth: true
            Layout.fillHeight: true
            orientation: ListView.Vertical
            boundsBehavior: Flickable.StopAtBounds
            anchors.horizontalCenter: application.horizontalCenter
            clip: true
            ScrollBar.vertical: ScrollBar {}

            ListModel {
                id: usersModel
            }

            model: usersModel

            delegate: Item {
                x: 5
                width: 80
                height: 40
                Row {
                    Rectangle {
                        width: 40
                        height: 40
                        Image {
                            id: iconPerson
                            source: "/images/black/person@2x.png"
                            anchors.fill: parent
                        }
                    }

                    Text {
                        anchors.verticalCenter: parent.verticalCenter
                        text: name
                    }

                    spacing: 10
                }
            }
            Component.onCompleted: {

                console.log(">>>> Completed List Users")
            }
        }
    }


    ColumnLayout {
        id: colWords
        width: 150
        height: application.height
        anchors.top: parent.top
        anchors.topMargin: 0
        anchors.left: parent.left
        anchors.leftMargin: 0

        Label {
            id: labelValidWords
            text: qsTr("Words")
            wrapMode: Text.WrapAnywhere
            horizontalAlignment: Text.AlignHCenter
            Layout.alignment: Qt.AlignHCenter | Qt.AlignVCenter
            Layout.fillHeight: true
            Layout.fillWidth: true
            font.pointSize: 13
            font.italic: false
            anchors.top: parent.top
            anchors.topMargin: 9
            verticalAlignment: Text.AlignVCenter
        }

        ListView {
            id: listViewWords
            width: parent.width
            height: 160
            highlightRangeMode: ListView.NoHighlightRange
            Layout.fillHeight: true
            Layout.fillWidth: true
            flickableDirection: Flickable.VerticalFlick
            boundsBehavior: Flickable.StopAtBounds
            interactive: true
            clip: true
            ScrollBar.vertical: ScrollBar {}
            delegate: Item {
                x: 5
                width: 80
                height: 40
                Row {

                    Rectangle {
                        width: 2
                        height: 40
                        color: Material.color(Functions.getMaterialColor())
                    }

                    Text {
                        text: word
                        anchors.verticalCenter: parent.verticalCenter
                        font.weight: Font.Light
                    }
                    spacing: 7
                }
            }
            model: ListModel {
                id: listModelWords
            }
        }


    }



    Component.onCompleted: {
        var users = user.data["joinedRoom"]["users"]
        users.forEach(function(u) {
            usersModel.append({name: u})
        });
        timer.start()
    }


    Connections {
        target: GameController
        onDeleteUser: {
            console.log("****DELETE USERS 2 *****")
//            var jsonUsers = JSON.parse(reply)
            usersModel.clear()
            var users = user.data["joinedRoom"]["users"]
            users.forEach(function(u) {
                usersModel.append({name: u})
            });

        }
    }


}