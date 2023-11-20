// import * as bootstrap from "bootstrap"
// import { Tooltip, Toast, Popover } from 'bootstrap';

import serverconnect from "./pages/serverconnect.js"

import "./scss/styles.scss"
import { setupCounter } from './counter.js'

const routes = {
    "/www/": { title: "Connect to server", render: serverconnect},
}

function router() {
    let page = routes[location.pathname]
    console.log(history)
    if (page) {
        document.title = page.title
        app.innerHTML = page.render()
    } else {
        history.replaceState("", "", "/www/")
        router()
    }
}

function routerpop() {
    let page = routes[location.pathname]
    console.log(history)

    if (page) {
        document.title = page.title
        app.innerHTML = page.render()
    } else {
        history.replaceState("", "", "/www/")
        router()
    }
}

function routerdomloaded() {
    let page = routes[location.pathname]
    console.log(history)

    if (page) {
        document.title = page.title
        app.innerHTML = page.render()
    } else {
        history.replaceState("", "", "/www/")
        router()
    }
}

window.addEventListener("click", e => {
    if (e.target) {
        e.preventDefault();
        history.pushState("", "", e.target.href);
        console.log(history)
        router();        
    }
})

window.addEventListener("popstate", routerpop);
window.addEventListener("DOMContentLoaded", routerdomloaded);

document.getElementById('app').innerHTML = `
  <div>
    <button id="button" class="btn btn-primary">Primary button</button>
  </div>
`

setupCounter(document.getElementById('button'))
