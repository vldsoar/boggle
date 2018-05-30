import QtQuick 2.7
import QtQuick.Controls 2.0
import QtQuick.Layouts 1.3
import QtQuick.Controls.Material 2.0

import "/js/ranking.js" as Ranking


Pane {
    id: pageResults
//    anchors.fill: parent
    font.family: "Ubuntu"

    Pane {
        id: paneResults
        width: application.width
        height: application.height
        visible: false

        ColumnLayout {
            id: columnWinner
            x: 258
            width: 616
            height: 155
            anchors.horizontalCenterOffset: 0
            anchors.horizontalCenter: parent.horizontalCenter

            Rectangle {
                id: rectImage
                width: 80
                height: 80
                Layout.fillWidth: false
                Layout.alignment: Qt.AlignHCenter | Qt.AlignVCenter

                Image {
                    id: imgTrophy
                    sourceSize.width: rectImage.width
                    sourceSize.height: rectImage.height
                    autoTransform: true
                    anchors.verticalCenter: parent.verticalCenter
                    anchors.horizontalCenter: parent.horizontalCenter
                    source: "/images/trophy.png"
                }
            }

            Text {
                id: textUserWinner
                text: ""
                verticalAlignment: Text.AlignVCenter
                horizontalAlignment: Text.AlignHCenter
                Layout.fillWidth: true
                font.pixelSize: 25
                font.family: "Ubuntu"
                font.bold: true

            }

            Text {
                id: name
                text: qsTr("Winner!")
                font.pointSize: 16
                font.bold: true
                Layout.alignment: Qt.AlignHCenter | Qt.AlignVCenter
                anchors.top: textUserWinner.bottom
                anchors.topMargin: 5
                font.family: "Ubuntu"
                color: Material.color(Material.Purple)
            }
        }

        ColumnLayout {
            id: columnHeader
            y: 161
            width: 616
            height: 67
            anchors.horizontalCenter: parent.horizontalCenter

            Row {
                id: rowHeaders
                Layout.fillWidth: true

                Text {
                    text: "#"
                    font.bold: true
                    anchors.verticalCenter: parent.verticalCenter
                    Layout.alignment: Text.AlignHCenter
                    font.pixelSize: 20
                    font.family: "Ubuntu"
                    color: Material.color(Material.Purple)

                }
                Rectangle {
                    width: 260
                    height: 40
                    anchors.verticalCenter: parent.verticalCenter

                    Text {
                        text: "User"
                        font.pixelSize: 20
                        font.family: "Ubuntu"
                        font.bold: true
                        width: parent.width
                        anchors.verticalCenter: parent.verticalCenter
                        horizontalAlignment: Text.AlignHCenter
                        color: Material.color(Material.Purple)
                    }
                }

                Text {
                    text: "Words"
                    font.pixelSize: 20
                    font.family: "Ubuntu"
                    font.bold: true
                    anchors.verticalCenter: parent.verticalCenter
                    horizontalAlignment: Text.AlignHCenter
                    width: 110
                    color: Material.color(Material.Purple)
                }

                Text {
                    text: "Score"
                    font.pixelSize: 20
                    font.family: "Ubuntu"
                    font.bold: true
                    anchors.verticalCenter: parent.verticalCenter
                    width: 150
                    horizontalAlignment: Text.AlignHCenter
                    color: Material.color(Material.Purple)
                }


                spacing: 20

            }
        }

        ListView {
            id: listViewUsers
            x: 228
            y: 220
            width: 616
            height: 276
            anchors.horizontalCenter: parent.horizontalCenter
            highlight: Rectangle { color: Material.color(Material.Grey, Material.Shade300); z: 1; opacity: .3}
            focus: true
            clip: true
            boundsBehavior: Flickable.StopAtBounds
            model: ListModel {
                id: listFinalUsers
            }
            delegate: Item {
                x: 5
                Layout.fillWidth: true
                width: parent.width
                height: 40

                Row {
                    id: rowListUser
                    Layout.fillWidth: true

                    Text {
                        text: (index + 1) + "º"
                        font.bold: true
                        anchors.verticalCenter: parent.verticalCenter
                        Layout.alignment: Text.AlignHCenter
                        font.pixelSize: 20
                        font.family: "Ubuntu"

                    }
                    Rectangle {
                        width: 40
                        height: 40
                        anchors.verticalCenter: parent.verticalCenter
                        //                            color: colorCode
                        Image {
                            id: userPosition
                            source: "/images/userposition.png"
                            sourceSize.height: 40
                            sourceSize.width: 40
                            autoTransform: true
                        }
                    }

                    Text {
                        width: application.width / 4
                        text: name
                        font.pixelSize: 20
                        font.family: "Ubuntu"
                        anchors.verticalCenter: parent.verticalCenter
                    }


                    Text {
                        text: totalWords
                        font.pixelSize: 20
                        font.family: "Ubuntu"
                        anchors.verticalCenter: parent.verticalCenter
                        horizontalAlignment: Text.AlignHCenter
                        width: 100
                    }

                    Text {
                        text: score
                        font.pixelSize: 20
                        font.family: "Ubuntu"
                        anchors.verticalCenter: parent.verticalCenter
                        width: 150
                        horizontalAlignment: Text.AlignHCenter
                    }
                    spacing: 20

                }
                MouseArea {
                    onClicked: listViewUsers.currentIndex = index
                    anchors.fill: parent

                }


            }

        }

    }

    Column {
        id: columnBusyResults
//        x: 244
//        y: 90
        width: 400
        height: 400
        visible: true
        anchors.verticalCenter: parent.verticalCenter
        anchors.horizontalCenter: paneResults.horizontalCenter

        BusyIndicator {
            id: busyLoadingResults
            width: 100
            height: 100
            z: 10
//            anchors.top: parent.top
//            anchors.topMargin: 120
            anchors.verticalCenter: parent.verticalCenter
            anchors.horizontalCenter: parent.horizontalCenter
        }

        Text {
            id: textPrepResults
            text: qsTr("Preparing Results")
            verticalAlignment: Text.AlignVCenter
            horizontalAlignment: Text.AlignHCenter
            font.bold: false
            anchors.horizontalCenter: parent.horizontalCenter
            anchors.top: busyLoadingResults.bottom
            anchors.topMargin: 20
            font.pixelSize: 20
        }
    }

    Component.onCompleted: GameController.finishGame()

    Connections {
        target: GameController
        onFinishedGame: {
            var jsonResults = JSON.parse(reply)

            console.log(JSON.stringify(jsonResults))

            var ranking = Ranking.generate(jsonResults)

            textUserWinner.text = ranking[0]["name"]



            ranking.forEach(function(player) {
                listFinalUsers.append(player)
            })

            pageResults.state = "finished"

        }
    }

    states: [
        State {
            name: "finished"
            PropertyChanges {
                target: columnBusyResults
                visible: false
            }

            PropertyChanges {
                target: paneResults
                visible: true
            }

        }
    ]



}


