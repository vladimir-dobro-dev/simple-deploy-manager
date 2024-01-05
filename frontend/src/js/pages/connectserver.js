import { api, www, router } from "../main.js"

export default function AddServer() {
    function effect() {
        const button = document.getElementById("connect")
        // button.innerHTML = i18next.t("ServerConnect.ButtonConnect")
        button.addEventListener("click", serverConnect)
    }

    async function serverConnect(e) {
        e.preventDefault()
        const connectServerForm = document.connectServerForm
        let connectData = {
            address: connectServerForm.address.value,
            user: connectServerForm.username.value,
            password: connectServerForm.password.value,
            port: connectServerForm.port.value
        }
        const response = await fetch(api + "/connectserver", {
            method: "POST",
            headers: { "Accept": "application/json", "Content-Type": "application/json" },
            body: JSON.stringify(connectData)
        })
        if (response.ok) {
            const text = await response.text()
            console.log(text)
            history.pushState("", "", www + "/applications")
            router()
        }
        else {
            console.log("response ERROR")
        }
    }

    return {
        title: "Connect to server",
        html: "connectserver.html",
        effect: effect
    } 
}
