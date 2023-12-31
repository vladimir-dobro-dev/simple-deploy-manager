// import * as bootstrap from "bootstrap"
// import { Tooltip, Toast, Popover } from 'bootstrap'

import AddServer from "./pages/connectserver.js"
import InstallApp from "./pages/installapp.js"

import "../styles/styles.scss"
import "../styles/styles.css"

export const www = "/www"
export const api = "http://127.0.0.1:2082/api"

const routes = {
    "/": AddServer,
    "/applications": InstallApp,
}

const app = document.getElementById("app");

export async function router() {
    let path = location.pathname.replace(www, "")
    let page = routes[path]
    if (page) {
        await render(page)
    } else {
        history.replaceState("", "", www + "/")
        router()
    }
}

async function render(page) {
    let pageData = page()
    document.title = pageData.title
    const response = await fetch(www + "/" + pageData.html)
    const html = await response.text()
    app.innerHTML = html
    pageData.effect()
}

window.addEventListener("popstate", router)
window.addEventListener("DOMContentLoaded", router)
