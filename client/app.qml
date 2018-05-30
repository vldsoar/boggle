import QtQuick.Window 2.2
import QtQuick 2.7
import QtQuick.Controls 2.0
import QtQuick.Layouts 1.0
import QtQuick.Controls.Material 2.0
import User 1.0

ApplicationWindow {
    id: application
    visible: true
    width: 800
    maximumWidth: 800
    height: 600
    maximumHeight: 600
    title: qsTr("Boggle Game")
//    Material.theme: Material.Dark
    Material.accent: Material.Purple

    User {
      id: user
      state: 0
      data: ({})
    }

    StackView {
        id: stackView
        anchors.rightMargin: 0
        anchors.bottomMargin: 0
        anchors.leftMargin: 0
        anchors.topMargin: 0
        anchors.fill: parent

        initialItem: Pane {
             id: loginPane
             width: application.width / 2
             height: 220
             anchors.bottomMargin: 100

             anchors.verticalCenter: parent.verticalCenter
             anchors.horizontalCenter: parent.horizontalCenter
             anchors.bottom: parent.bottom

             ColumnLayout {
                 anchors.fill: parent
                 anchors.margins: 3
                 spacing: 3

                 Column {
                     id: column1
                     height: 101
                     Layout.alignment: Qt.AlignHCenter | Qt.AlignVCenter
                     Layout.fillWidth: true
                     Layout.fillHeight: false
//                     anchors.fill: parent
//                     anchors.margins: 3
                     anchors.top: loginPane.top
                     anchors.topMargin: 15
                     spacing: 3

                     Text {
                         text: "Boggle Game"
                         font.family: "Ubuntu, Tahoma"
                         font.pointSize: 20
                         font.bold: true
                         color: Material.color(Material.Purple)
                         anchors.horizontalCenter: parent.horizontalCenter
                         Layout.fillWidth: true
                         Layout.alignment: Qt.AlignHCenter | Qt.AlignVCenter
                         verticalAlignment: Text.AlignVCenter
                         horizontalAlignment: Text.AlignHCenter
                     }
                 }

                 TextField {
                     id: login
                     Layout.fillWidth: true
                     placeholderText: "Username"
//                     enabled: user.state == 1
                 }

                 Button {
                     id: proccessButton
                     Layout.fillWidth: true
                     highlighted: true
                     Material.background: Material.color(Material.Purple)
                 }

                 TextArea {
                    id: data
                    text: "Not logged in.\n\n"
                    readOnly: true
                    Layout.fillHeight: false
                    Layout.fillWidth: true

                    Connections {
                      target: GameController
                      onSessionAuthenticated: {
                        user.state = 1
                        var jsonObject = JSON.parse(reply)
                        user.data = jsonObject
                        data.text = "User '"+ login.text +"' Authenticated \n\n"

                        stackView.push("qrc:/pages/Lobby.qml")
                      }
                      onSessionAuthenticationError: {
                        user.state = 2
                      }
                    }
                 }

                 states: [
                     State {
                         name: "NotAuthenticated"
                         when: user.state == 0
                         PropertyChanges {
                             target: proccessButton
                             text: "Login"
                             onClicked: {

                                 if (login.text.length >= 4) {
                                     GameController.login(login.text)
                                 } else {
                                     data.text = "username: minimum 4 characters"
                                 }

                             }
                         }
                     },
                     State {
                         name: "AuthenticationFailure"
                         when: user.state == 2
                         PropertyChanges {
                             target: proccessButton
                             text: "Authentication failed, restart"
                             onClicked: {
                                if (login.text.length >= 4) {
                                    GameController.login(login.text)
                                } else {
                                    data.text = "username: minimum 4 characters"
                                }

                             }
                         }
                     },
                     State {
                         name: "Authenticated"
                         when: user.state == 1
                         PropertyChanges {
                             target: proccessButton
                             text: "Logout"
                             onClicked: {
                                 console.log("logout")
                             }
                         }
                     }
                 ]
             }



        }
        Connections {
            target: GameController
            onAppendUser: {
                console.log("**** APPEND USER ****")
                var data = user.data
                data["joinedRoom"]["users"].push(reply)
                user.data = data
                console.log(user.data["joinedRoom"]["users"].length)
                console.log("*************************")
            }

            onDeleteUser: {
                console.log("****DELETE USERS*****")

                var data = user.data

                data["joinedRoom"]["users"].pop(reply)

                user.data = data

            }

        }
    }
}
