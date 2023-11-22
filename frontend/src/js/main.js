// import * as bootstrap from "bootstrap"
// import { Tooltip, Toast, Popover } from 'bootstrap'

import AddServer from "./pages/addserver.js"

import "../styles/styles.scss"
import "../styles/styles.css"

export const www = "/www"
export const api = "http://127.0.0.1:2082/api"

const routes = {
    "/": AddServer,
}

const app = document.getElementById("app");

async function router() {
    let path = location.pathname.replace(www, "")
    let render = routes[path]
    if (render) {
        await setPage(render)
    } else {
        history.replaceState("", "", "/www/")
        router()
    }
}

async function routerpop() {
    let path = location.pathname.replace(www, "")
    let page = routes[path]
    if (page) {
        await setPage(page)
    } else {
        history.replaceState("", "", "/www/")
        router()
    }
}

async function setPage(render) {
    const page = render()
    document.title = page.title
    const response = await fetch(www + "/" + page.html)
    const html = await response.text()
    app.innerHTML = html
    page.effect()
}

// window.addEventListener("click", e => {
//     if (e.target) {
//         e.preventDefault()
//         history.pushState("", "", e.target.href)
//         console.log(history)
//         router()
//     }
// })

window.addEventListener("popstate", routerpop)
window.addEventListener("DOMContentLoaded", router)
