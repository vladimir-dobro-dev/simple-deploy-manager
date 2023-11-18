import { useState } from "react"
import {
  createBrowserRouter,
  RouterProvider,
} from "react-router-dom"

import "./bootstrap.min.css"

import Installer from "./routes/installer"
import AddServer from "./routes/addserver"
import "./App.css"

const router = createBrowserRouter([
  {
    path: "/www",
    element: <Installer />,
  },
  {
    path: "/www/installer",
    element: <Installer />,
  },
  {
    path: "/www/addserver",
    element: <div><AddServer /></div>,
  },
]);

function App() {
  return (
    <div className="p-3">
      <RouterProvider router={router} />
    </div>
  )
}

export default App
