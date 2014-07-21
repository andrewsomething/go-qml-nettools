import QtQuick 2.2
import QtQuick.Controls 1.0
import QtQuick.Layouts 1.0

ApplicationWindow {
    id: window
    visible: true
    width: 600
    height: 400
    minimumWidth: 500
    minimumHeight: 300

    title: "Net tools"

    SplitView {
        anchors.fill: parent

        SplitView {
            orientation: Qt.Vertical
            Layout.fillWidth: true

            Rectangle {
                id: row1
                height: 50
                color: "#DEDEDE"
                Layout.minimumHeight: 1

                RowLayout {
                    anchors.fill: parent
                    spacing: 10

                    ComboBox {
                        id: selector
                        anchors.verticalCenter: parent.verticalCenter
                      //  anchors.margins: 10
                        currentIndex: 0
                        model: ListModel {
                            id: cbItems
                            ListElement { text: "Ping"; }
                            ListElement { text: "Traceroute"; }
                        }
                    }

                    TextField {
                        id: input
                        objectName: "input"
                        Layout.fillWidth: true
                        anchors.left: selector.right
                        anchors.verticalCenter: parent.verticalCenter
                     //   anchors.margins: 10
                    }
                    
                    Button {
                        objectName: "btn"
                        text: "Run"
                        anchors.left: input.right
                        anchors.verticalCenter: input.verticalCenter
                     //   anchors.margins: 10
                        onClicked: app.handleClick()
                    }
                }
            }

            Rectangle {
                id: row2
                color: "#272822"
                
                TextArea {
                    objectName: "message"
                    anchors.fill: parent
                    readOnly: true
                    font.family: "monospace"
                    textColor: "#94B630"
                    backgroundVisible: false
                    text: "Click that button up there"
                }
            }
        }
    }
}
