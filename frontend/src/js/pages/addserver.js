import i18next from "../i18n"

import { api } from "../main"

export default function AddServer() {
    function effect() {
        const button = document.getElementById("button-connectserver")
        button.innerHTML = i18next.t("ServerConnect.ButtonConnect")
        button.addEventListener("click", serverConnectClick)
    }

    async function serverConnectClick(e) {
        e.preventDefault()
        const connectForm = document.addserver
        let connectData = {
            address: connectForm.address.value,
            username: connectForm.username.value,
            password: connectForm.password.value,
            port: connectForm.port.value
        }
        const response = await fetch(api + "/addserver", {
            method: "POST",
            headers: { "Accept": "application/json", "Content-Type": "application/json" },
            body: JSON.stringify(connectData)
        })
        if (response.ok === true) {
            const text = await response.text()
            console.log(text)
        }
        else {
            console.log("response ERROR")
        }
    }

    return {
        title: "Add server",
        html: "addserver.html",
        effect: effect
    } 
}
