import { api, www, router } from "../main.js"

export default function InstallApp() {
    function effect() {
        const button = document.getElementById("install")
        button.addEventListener("click", install)
    }

    async function install(e) {
        e.preventDefault()
        const form = document.querySelector("form")
        const response = await fetch(api + "/installapp", {
            method: "POST",
            body: new FormData(form)
        })
        if (response.ok) {
            const text = await response.text()
            console.log(text)
            history.pushState("", "", www + "/")
            router()
        }
        else {
            console.log("response ERROR")
        }
    }

    return {
        title: "InstallApp",
        html: "installapp.html",
        effect: effect
    }
}